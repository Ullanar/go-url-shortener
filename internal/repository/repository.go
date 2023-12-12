package repository

import (
	"gorm.io/gorm"
	"url-shortener/internal/models"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateAlias(alias, dest string) error
	GetDestinationByAlias(alias string) (string, error)
}

func New(db *gorm.DB) *repository {
	_ = db.AutoMigrate(models.Link{})
	return &repository{db}
}

func (r *repository) CreateAlias(alias, dest string) error {
	result := r.db.Create(&models.Link{Alias: alias, Dest: dest})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) GetDestinationByAlias(alias string) (string, error) {
	var link models.Link
	result := r.db.Where(&models.Link{Alias: alias}).First(&link)
	if result.Error != nil {
		return "", result.Error
	}
	return link.Dest, nil
}
