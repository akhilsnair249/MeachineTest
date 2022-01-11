package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"klok/OnlineBilling/models"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const (
	sessionKeyHeader = "Session-Key"
)

func SetupRoutes(connection *gorm.DB, mainRouter *mux.Router) {

	var articleRepo Article = CreateArticle(connection)
	var articleBll Article = CreateArticleg(articleRepo)

	artController := CreateArticle(articleBll)

	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/articles", artController.ShowAllArticle).Methods(http.MethodGet)
	apiRouter.HandleFunc("/articles", APISessionAuthMiddlewareFunc(artController.AddArticle)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/articles/{id}", APISessionAuthMiddlewareFunc(artController.AproveArticle)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/articles/{id}", APISessionAuthMiddlewareFunc(artController.DeclineArticle)).Methods(http.MethodPost)

}

func APISessionAuthMiddlewareFunc(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok, tokenData, err := verifySession(w, r)
		isError := false

		if err == nil && tokenData != nil && ok {

			user, err1 := verifier.userFullDetailsRepo.ShowWithCode(r.Context(), tokenData.UserCode)
			if err1 == nil && user.Code == tokenData.UserCode {
				sessionContext := context.WithValue(r.Context(), "TokenData", *tokenData)
				h.ServeHTTP(w, r.WithContext(sessionContext))
			} else {
				isError = true
			}

		} else {
			isError = true
		}

		if isError {

			message := "Access denied"

			if err != nil {
				message = err.Error()
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.CreateJsonError(message, http.StatusUnauthorized))
		}

	})
}

func verifySession(w http.ResponseWriter, r *http.Request) (bool, *TokenData, error) {

	sessKey := r.Header.Get(sessionKeyHeader)
	queryKey := r.URL.Query().Get("Token")
	if len(sessKey) == 0 && len(queryKey) == 0 {
		return false, nil, errors.New("Session invalid , Please login again : No token specified")
	}

	if len(sessKey) == 0 {
		sessKey = queryKey
	}

	pasedToken, err := jwt.ParseWithClaims(sessKey, &TokenData{}, func(token *jwt.Token) (interface{}, error) {
		tokenData, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		if tokenData.Name != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("Unexpected method %s", tokenData.Name)
		}

		var signingKey = []byte("4667awkfh233j3h5h345kjh345h34kj5hk")

		return signingKey, nil
	})

	if err != nil {

		return false, nil, errors.New("Session invalid , Please login again  : " + err.Error())
	}

	claims, ok := pasedToken.Claims.(*TokenData)

	if !ok {
		return false, nil, errors.New("Session invalid , Please login again  : Token cast failed")
	}

	validationError := claims.Valid()

	if validationError != nil {
		return false, nil, errors.New("Session invalid , Please login again  :  " + err.Error())
	}

	return true, claims, nil

}
