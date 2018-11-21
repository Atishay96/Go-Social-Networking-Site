package server

import (
	"net/http"

	"Go-Social/pkg"

	"github.com/gorilla/mux"
)

type userRouter struct {
	userService root.UserService
	auth        *authHelper
}

func NewUserRouter(u root.UserService, router *mux.Router, a *authHelper) *mux.Router {
	userRouter := userRouter{u, a}
	router.HandleFunc("/profile", a.validate(userRouter.profileHandler)).Methods("GET")
	return router
}

func (ur *userRouter) profileHandler(w http.ResponseWriter, r *http.Request) {
	claim, ok := r.Context().Value(contextKeyAuthtoken).(claims)
	if !ok {
		Error(w, http.StatusBadRequest, "no context")
		return
	}
	username := claim.Username

	user := "Hello, " + username

	Json(w, http.StatusOK, user)
}
