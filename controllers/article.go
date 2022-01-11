package controllers

import (
	"context"
	"klok/OnlineBilling/controllers"
	"net/http"
)

type Article struct {
	bll Article
}

func CreateArticle(bll Article) Article {
	return Article{bll}
}

func (con Article) ShowAllArticle(w http.ResponseWriter, r *http.Request) {

	res, err := con.bll.ShowAllArticle(r.Context())
	if err != nil {
		con.RespondJsonFailure(w, err)
		return
	}
	con.RespondSuccess(w, res)
}

func (bll Article) AddArticle(ctx context.Context) ([]Article, error) {

	res, err := con.bll.AddArticle(ctx)

	if err != nil {
		con.RespondJsonFailure(w, err)
		return
	}
	con.RespondSuccess(w, res)
}

func (bll Article) AproveArticle(ctx context.Context) ([]Article, error) {
	idString := controllers.GetID(r)

	res, err := con.bll.AproveArticle(ctx)

	if err != nil {
		con.RespondJsonFailure(w, err)
		return
	}
	con.RespondSuccess(w, res)
}
func (bll Article) DeclineArticle(ctx context.Context) ([]Article, error) {

	idString := controllers.GetID(r)

	res, err := con.bll.DeclineArticle(ctx)

	if err != nil {
		con.RespondJsonFailure(w, err)
		return
	}
	con.RespondSuccess(w, res)
}
