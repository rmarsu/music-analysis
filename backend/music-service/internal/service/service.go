package service

import (
	repository "music-service/internal/repo"
	"music-service/models"
)

type Services struct {
	// TODO implement auth
	Music Music
}

type Music interface {
	GetByGenre(genre string) ([]models.Music, error)
	Create(music *models.Music) (int64, error)
	Delete(id int64) error
}

type Deps struct {
	repo *repository.Repository
}

func New(d *Deps) *Services {
	return &Services{
		Music: NewMusicService(d.repo),
	}
}