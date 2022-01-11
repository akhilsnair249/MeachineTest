package main

import (
	"context"

	"gorm.io/gorm"
)

type Article struct {
	connection *gorm.DB
}

func CreateArticle(connection *gorm.DB) Article {
	return Article{connection}
}

func (repo Article) Show(ctx context.Context, identifier string) (Article, error) {
	var res Article
	err := repo.connection.WithContext(ctx).First(&res, identifier).Error
	return res, err

}

func (repo Article) ShowAll(ctx context.Context) ([]Article, error) {
	var allres []Article
	err := repo.connection.Find(&allres).Error

	if err != nil {
		return allres, err
	}

	return allres, nil
}

func (repo Article) Aprove(ctx context.Context, req *Article) error {
	return repo.connection.WithContext(ctx).Save(req).Error

}

func (repo Article) Decline(ctx context.Context, req *Article) error {
	return repo.connection.WithContext(ctx).Save(req).Error
}

func (repo Article) Create(ctx context.Context, newModel Article) error {
	return repo.connection.WithContext(ctx).Create(newModel).Error
}
