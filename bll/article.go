package main

import (
	"context"
	"strconv"
)

type Article struct {
	repo Article
}

func CreateArticle(repo Article) Article {
	return Article{repo}
}

func (bll Article) ShowAllArticle(ctx context.Context) ([]Article, error) {

	allArticle, err := bll.repo.Show(ctx)

	if err != nil {
		return nil, DBError{Messsage: "Unable to get all article details!", ActualError: err, ErrorFrom: nil}
	}

	return allArticle, nil
}

func (bll Article) AddArticle(ctx context.Context, data Article) error {

	allArticle, err := bll.repo.Create(ctx, data)

	if err != nil {
		return nil, DBError{Messsage: "Unable to add article details!", ActualError: err, ErrorFrom: nil}
	}

	return allArticle, nil
}

func (bll Article) AproveArticle(ctx context.Context, id string) error {

	reqCode64, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return DBError{Messsage: "Invalid ID", ActualError: err, ErrorFrom: nil}
	}

	article, err := bll.repo.Show(ctx)

	if err != nil {
		return DBError{Messsage: "Unable to get article details!", ActualError: err, ErrorFrom: nil}
	}

	err = bll.repo.Aprove(ctx, &article)

	if err != nil {
		return DBError{Messsage: "Unable to approve article details!", ActualError: err, ErrorFrom: nil}
	}

	return nil
}

func (bll Article) DeclineArticle(ctx context.Context, id string) error {

	reqCode64, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return DBError{Messsage: "Invalid ID", ActualError: err, ErrorFrom: nil}
	}

	article, err := bll.repo.Show(ctx)

	if err != nil {
		return DBError{Messsage: "Unable to get article details!", ActualError: err, ErrorFrom: nil}
	}

	err = bll.repo.Decline(ctx, &article)

	if err != nil {
		return DBError{Messsage: "Unable to decline article details!", ActualError: err, ErrorFrom: nil}
	}

	return nil
}
