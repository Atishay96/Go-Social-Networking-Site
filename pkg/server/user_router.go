package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"Go-Social/pkg"
)

type userRouter struct {
	userService root.UserService
	auth        *authHelper
}

func NewUserRouter(u root.UserService, router *mux.Router, a *authHelper) *mux.Router {
	userRouter := userRouter{u, a}
	// router.HandleFunc("/profile", a.validate(userRouter.profileHandler)).Methods("GET")
	router.HandleFunc("/signup", userRouter.signUpHandler).Methods("POST")
	return router
}

func (ur *userRouter) signUpHandler(w http.ResponseWriter, r *http.Request) {

	var resp root.ResponseSlice

	user, emptyFields, err := decodeUser(r)
	if err != nil {
		resp.Message = "Error Occured"
		resp.Err = err
		Json(w, http.StatusInternalServerError, resp)
		return
	}
	if len(emptyFields) != 0 {
		resp.Message = "Bad Request."
		resp.Data = emptyFields
		Json(w, http.StatusBadRequest, resp)
		return
	}
	err1 := ur.userService.CreateUser(&user)
	if err1 != nil {
		resp.Message = "Error Occured"
		resp.Err = err1
		Json(w, http.StatusInternalServerError, resp)
		return
	}
	resp.Message = "User Signed Up"
	resp.Data = user
	Json(w, http.StatusOK, resp)
	return
}

func decodeUser(r *http.Request) (root.User, []string, error) {
	var u root.User
	if r.Body == nil {
		return u, []string{}, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	// checks := []string{
	// 	"UserName",
	// 	"FirstName",
	// 	"LastName",
	// 	"Email",
	// "Password",
	// }
	emptyFields := []string{}
	// for _, check := range checks {
	// 	r := reflect.ValueOf(u)
	// 	f := reflect.Indirect(r).FieldByName(check)
	// 	fmt.Println(f)
	// 	// if f == nil {
	// 	// 	emptyFields = append(emptyFields, check)
	// 	// }
	// }
	err := decoder.Decode(&u)
	return u, emptyFields, err
}

// func (ur *userRouter) profileHandler(w http.ResponseWriter, r *http.Request) {
// 	claim, ok := r.Context().Value(contextKeyAuthtoken).(claims)
// 	if !ok {
// 		Error(w, http.StatusBadRequest, "no context")
// 		return
// 	}
// 	username := claim.Username

// 	user := "Hello, " + username

// 	Json(w, http.StatusOK, user)
// }
