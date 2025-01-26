package handler

import (
	"music-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

const (
	ErrServiceUnavailable = "сервис сейчас не доступен. мы работаем над этим!"
	ErrInternal = "у нас что то сломалось! мы уже работаем над этим."
	ErrInvalidArgument = "хм , похоже какие то данные сломаны! попробуйте позже или повторите попытку с другими данными."
	ErrNoAudio = "отсутсвует аудио-файл."
	ErrRead = "ошибка обработки файла."
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
