package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"Go-Social/pkg"
)

type postRouter struct {
	postService root.PostService
	auth        *authHelper
}

func NewPostRouter(p root.PostService, router *mux.Router, a *authHelper) *mux.Router {
	postRouter := postRouter{p, a}
	router.HandleFunc("/post", a.validate(postRouter.postHandler)).Methods("PUT")
	return router
}

func (pr *postRouter) postHandler(w http.ResponseWriter, r *http.Request) {

}
