package service

import (
	repository "music-service/internal/repo"
	"music-service/models"
	"music-service/pkg/logger"
)

type MusicService struct {
	repo *repository.Repository
}

const (
	NoValue = 0
)

func NewMusicService(repo *repository.Repository) *MusicService {
	return &MusicService{
		repo: repo,
	}
}

func (s *MusicService) GetByGenre(genre string) ([]models.Music, error) {
	music , err := s.repo.Music.GetByGenre(genre)
	if err != nil {
		logger.Warnf("No music for genre: %s", genre)
		return nil, err
	}
	return music, nil
}

func (s *MusicService) Create(music *models.Music) (int64, error) {
	id , err := s.repo.Music.Create(music)
	if err != nil {
		logger.Warnf("Cannot creare music %v. Reason : %s", music, err)
		return NoValue, err
	}
	return id, nil
}

func (s *MusicService) Delete(id int64) error {
	if err := s.repo.Music.Delete(id); err != nil {
		logger.Warnf("Cannot delete music with id %d. Reason: %s", id, err)
		return err
	}
	return nil
}
