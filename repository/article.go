package main

import (
	"context"
)

type Article interface {
	Show(ctx context.Context, identifier string) (Article, error)

	ShowAll(ctx context.Context) ([]Article, error)

	Aprove(ctx context.Context, req Article) error

	Decline(ctx context.Context, req Article) error

	Create(ctx context.Context, newModel Article) error
}
