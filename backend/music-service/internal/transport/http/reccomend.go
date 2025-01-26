package handler

import (
	"mime/multipart"
	music_service "music-service/grpc"
	"music-service/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	classification_server = "localhost:4000"
)

func (h *Handler) getRecommendations(c *gin.Context) {
	logger.Infof("Someone trying to get recommendations... %s", classification_server)
	conn, err := grpc.Dial(classification_server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf(err.Error())
		newErrorResponse(c, http.StatusServiceUnavailable, ErrServiceUnavailable)
		return
	}
	defer conn.Close()

	client := music_service.NewClassificationServiceClient(conn)
	

	file, err := c.FormFile("song")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, ErrNoAudio)
		return
	}

	file_data, err := ReadFile(file)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, ErrRead)
		return
	}
	maxSizeOption := grpc.MaxCallRecvMsgSize(1024 * 1024 * 200)
	genre, err := client.UploadAudio(c.Request.Context(), &music_service.UploadAudioRequest{
		Filename: file.Filename,
		FileData: file_data,
	}, maxSizeOption)

	logger.Infof("%v", genre)
	if err != nil {
		logger.Errorf(err.Error())
		st, ok := status.FromError(err)
		if ok {
			sc := st.Code()
			switch sc {
			case codes.InvalidArgument:
				newErrorResponse(c, http.StatusBadRequest, ErrInvalidArgument)
			default:
				newErrorResponse(c, http.StatusInternalServerError, ErrInternal)
			}
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"genre": genre.GetGenre().String(),
	})
}

func ReadFile(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	fileData := make([]byte, file.Size)
	if _, err := src.Read(fileData); err != nil {
		return nil, err
	}
	return fileData, nil
}
