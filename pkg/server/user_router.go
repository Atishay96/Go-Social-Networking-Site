package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
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
	router.HandleFunc("/login", userRouter.loginHandler).Methods("POST")
	router.HandleFunc("/profile", a.validate(userRouter.loggedInUserHandler)).Methods("GET")
	router.HandleFunc("/profile/{friendId}", a.validate(userRouter.ProfileHandler)).Methods("GET")
	return router
}

func (ur *userRouter) signUpHandler(w http.ResponseWriter, r *http.Request) {

	var resp root.ResponseSlice

	checks := []string{
		"username",
		"FirstName",
		"LastName",
		"Email",
		"Password",
	}
	user, emptyFields, err := decodeUser(r, checks)

	if err != nil {
		resp.Message = "Error Occured"
		resp.Err = err
		Json(w, http.StatusInternalServerError, resp)
		return
	}
	//change later
	if len(emptyFields) != 0 {
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
	var resp root.ResponseSlice

	user, emptyFields, err := decodeUser(r, []string{"Email"})

	if err != nil {
		resp.Message = "Error Occured"
		resp.Err = err
		Json(w, http.StatusInternalServerError, resp)
		return
	}
	//change later
	if len(emptyFields) != 0 {
		resp.Message = "Bad Request."
		resp.Data = emptyFields
		Json(w, http.StatusBadRequest, resp)
		return
	}
	check2 := ur.userService.CheckEmail(user.Email)
	if check2 != false {
		resp.Message = "Email is not registered"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	random := helper.GenerateRandomString()
	link := "http://localhost:1377/user/verify/" + random
	user.VerificationSecret = random
	err2 := ur.userService.UpdateUser([]string{"VerificationSecret"}, user.VerificationSecret, user.Email)
	if err2 != nil {
		resp.Message = "Account already activated"
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
	resp.Message = "Mail sent"
	// resp.Data = user
	Json(w, http.StatusOK, resp)
	return
}

func (ur *userRouter) verifyAccountHandler(w http.ResponseWriter, r *http.Request) {

	var resp root.ResponseSlice

	vars := mux.Vars(r)
	secret := vars["secret"]
	user, err := ur.userService.HandleSecret(secret)
	if err != nil {
		resp.Message = "Account already activated or Link Expired"
		resp.Err = err
		Json(w, http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Successfully Verified"
	token := ur.auth.newToken(user)
	resp.Data = map[string]string{"AuthToken": token}
	Json(w, http.StatusOK, resp)
	return
}

func (ur *userRouter) loginHandler(w http.ResponseWriter, r *http.Request) {
	var resp root.ResponseSlice

	user, emptyFields, err := decodeUser(r, []string{"Email", "Password"})
	if err != nil {
		resp.Message = "Error Occured"
		resp.Err = err
		Json(w, http.StatusInternalServerError, resp)
		return
	}
	//change later
	if len(emptyFields) != 0 {
		resp.Message = "Bad Request."
		resp.Data = emptyFields
		Json(w, http.StatusBadRequest, resp)
		return
	}
	check2 := ur.userService.CheckEmail(user.Email)
	if check2 != false {
		resp.Message = "Email is not registered"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	check3, userData := ur.userService.CheckStatus(user.Email)
	if check3 == false {
		resp.Message = "Account is not verified or is blocked. Please contact ADMIN!"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Successfully LoggedIn"
	token := ur.auth.newToken(userData)
	resp.Data = map[string]string{"AuthToken": token}
	Json(w, http.StatusOK, resp)
	return
}

func (ur *userRouter) loggedInUserHandler(w http.ResponseWriter, r *http.Request) {

	var resp root.ResponseSlice

	Username := context.Get(r, "Username")
	ID := context.Get(r, "ID")
	UpdatedAt := context.Get(r, "UpdatedAt")

	var param []string
	param = append(param, Username.(string), ID.(string), UpdatedAt.(string))
	user := ur.userService.GetUserByParams(param)
	if user == nil {
		resp.Message = "Data unavailable"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Operation successful"
	resp.Data = user
	Json(w, http.StatusOK, resp)
	return
}

func (ur *userRouter) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	var resp root.ResponseSlice

	vars := mux.Vars(r)
	UserId := context.Get(r, "ID")
	ID := vars["friendId"]
	if UserId == ID {
		resp.Message = "Wrong API call"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	user := ur.userService.GetOtherUserByParams(ID)
	if user == nil {
		resp.Message = "Data unavailable"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Operation successful"
	resp.Data = user
	Json(w, http.StatusOK, resp)
	return
}

func decodeUser(r *http.Request, checks []string) (root.User, []string, error) {
	var u root.User
	if r.Body == nil {
		return u, []string{}, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	emptyFields := []string{}
	for _, check := range checks {
		fmt.Println(check)
		// fieldValue := reflect.Indirect(reflect.ValueOf(&u)).FieldByName(string(check))
		// fmt.Println(reflect.Indirect(reflect.ValueOf(u)))
		// if (fieldValue.Type().String() == "string" && fieldValue.Len() == 0) || (fieldValue.Type().String() != "string" && fieldValue.IsNil()) {
		// 	emptyFields = append(emptyFields, check)
		// }
	}
	err := decoder.Decode(&u)
	return u, emptyFields, err
}
