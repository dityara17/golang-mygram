package repositorysocialmedia

import (
	"github.com/arfan21/golang-mygram/entity"
	"gorm.io/gorm"
)

type RepositorySocialMedia interface {
	Create(data entity.SocialMedia) (entity.SocialMedia, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositorySocialMedia {
	return &repository{db: db}
}

func (r *repository) Create(data entity.SocialMedia) (entity.SocialMedia, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}
	return data, nil
}
