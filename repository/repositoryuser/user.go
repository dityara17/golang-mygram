package repositoryuser

import "gorm.io/gorm"

type RepositoryUser interface {
	Create(data entity.User) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryUser {
	return &repository{db: db}
}

func (r *repository) Create(data entity.User) (entity.User, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.User{}, err
	}

	return data, nil
}
