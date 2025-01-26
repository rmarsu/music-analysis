package repository

import (
	"music-service/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MusicRepo struct {
	db *pgxpool.Conn
}

func NewMusicRepo(db *pgxpool.Conn) *MusicRepo {
	return &MusicRepo{
		db: db,
	}
}

func (r *MusicRepo) Create(music *models.Music) (int64, error) {
	panic("not implemented")
}

func (r *MusicRepo) Delete(id int64) error {
	panic("not implemented")
}

func (r *MusicRepo) GetByGenre(genre string) ([]models.Music, error) {
	panic("not implemented")
}
