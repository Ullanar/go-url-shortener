package app

import (
	"gorm.io/gorm"
	"url-shortener/internal/config"
	"url-shortener/internal/database"
	"url-shortener/internal/repository"
)

type appProvider struct {
	config     *config.Config
	db         *gorm.DB
	repository repository.Repository
}

func New() *appProvider {
	return &appProvider{}
}

func (p *appProvider) Config() *config.Config {
	if p.config == nil {
		p.config = config.MustLoad()
	}
	return p.config
}

func (p *appProvider) Database() *gorm.DB {
	if p.db == nil {
		p.db = database.New(p.Config().Database)
	}
	return p.db
}

func (p *appProvider) Repository() repository.Repository {
	if p.repository == nil {
		repo := repository.New(p.Database())
		p.repository = repo
	}
	return p.repository
}
