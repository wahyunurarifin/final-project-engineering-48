package api

import (
	"fmt"
	"net/http"

	"server/repository"
)

type API struct {
	usersRepo    repository.UserRepository
	articlesRepo repository.ArticleRepository
	mux          *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, articlesRepo repository.ArticleRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, articlesRepo, mux,
	}

	// Handler untuk users (dengan middleware)
	mux.Handle("/api/register", api.POST(http.HandlerFunc(api.register)))
	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/logout", api.POST(http.HandlerFunc(api.logout)))

	// Handler untuk guest (tanpa middleware)
	mux.Handle("/api/article/category", api.GET(http.HandlerFunc(api.getArticleCategory)))
	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
