package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"

	"Go-Social/pkg"
	"Go-Social/pkg/helpers"
)

type userRouter struct {
	userService root.UserService
	auth        *authHelper
}

func NewUserRouter(u root.UserService, router *mux.Router, a *authHelper) *mux.Router {
	userRouter := userRouter{u, a}
	router.HandleFunc("/signup", userRouter.signUpHandler).Methods("PUT")
	router.HandleFunc("/resendMail", userRouter.mailHandler).Methods("POST")
	router.HandleFunc("/verify/{secret}", userRouter.verifyAccountHandler).Methods("GET")
	return router
}

func (ur *userRouter) signUpHandler(w http.ResponseWriter, r *http.Request) {

	var resp root.ResponseSlice

	user, emptyFields, err := decodeUser(r)
	fmt.Println(emptyFields, "emptyFields")
	if err != nil {
		resp.Message = "Error Occured"
		resp.Err = err
		Json(w, http.StatusInternalServerError, resp)
		return
	}
	//change later
	if len(emptyFields) == 0 {
		resp.Message = "Bad Request."
		resp.Data = emptyFields
		Json(w, http.StatusBadRequest, resp)
		return
	}
	check1 := ur.userService.CheckUserName(user.Username)
	if check1 == false {
		resp.Message = "Username should be unique"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	check2 := ur.userService.CheckEmail(user.Email)
	if check2 == false {
		resp.Message = "Email should be unique"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	random := helper.GenerateRandomString()
	link := "http://localhost:1377/user/verify/" + random
	user.VerificationSecret = random
	err2 := ur.userService.CreateUser(&user)
	if err2 != nil {
		resp.Message = "Error Occured"
		resp.Err = err2
		Json(w, http.StatusInternalServerError, resp)
		return
	}
	//sending mail
	c := make(chan string)
	body := "Wecome abroad, " + user.FirstName
	body = body + "\r\n Click on the link below to activate your account "
	body = body + link
	go helper.SendMail(c, user.Email, body)
	resp.Message = "User Signed Up"
	// resp.Data = user
	Json(w, http.StatusOK, resp)
	return
}

func (ur *userRouter) mailHandler(w http.ResponseWriter, r *http.Request) {

}

func (ur *userRouter) verifyAccountHandler(w http.ResponseWriter, r *http.Request) {

	var resp root.ResponseSlice

	vars := mux.Vars(r)
	secret := vars["secret"]
	err := ur.userService.HandleSecret(secret)
	if err != nil {
		resp.Message = "Account already activated or Link Expired"
		resp.Err = err
		Json(w, http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Successfully Verified"
	resp.Data = []string{}
	Json(w, http.StatusOK, resp)
	return
}

func decodeUser(r *http.Request) (root.User, []string, error) {
	var u root.User
	if r.Body == nil {
		return u, []string{}, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	checks := []string{
		"Username",
		"FirstName",
		"LastName",
		"Email",
		"Password",
	}
	emptyFields := []string{}
	for _, check := range checks {
		r := reflect.ValueOf(u)
		f := reflect.Indirect(r).FieldByName(check)
		fmt.Println(f)
		fieldValue := reflect.Indirect(reflect.ValueOf(u)).FieldByName(check)
		if (fieldValue.Type().String() == "string" && fieldValue.Len() == 0) || (fieldValue.Type().String() != "string" && fieldValue.IsNil()) {
			fmt.Println("fieldValue")
			fmt.Println(reflect.Indirect(reflect.ValueOf(u)).FieldByName(check))
			fmt.Println(reflect.Indirect(reflect.ValueOf(u)).FieldByName(check))
			emptyFields = append(emptyFields, check)
		}
	}
	err := decoder.Decode(&u)
	return u, emptyFields, err
}
