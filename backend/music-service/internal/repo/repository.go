package repository

import (
	"music-service/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Music Music
}

type Music interface {
	Create(music *models.Music) (int64, error)
	Delete(id int64) (error)
	GetByGenre(genre string) ([]models.Music, error)
}

func New(db *pgxpool.Conn) *Repository {
	return &Repository{
		Music: NewMusicRepo(db),
	}
}