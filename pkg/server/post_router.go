package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"Go-Social/pkg"
)

type postRouter struct {
	postService root.PostService
	auth        *authHelper
	userService root.UserService
}

func NewPostRouter(u root.UserService, p root.PostService, router *mux.Router, a *authHelper) *mux.Router {
	postRouter := postRouter{p, a, u}
	router.HandleFunc("/post", a.validate(postRouter.postHandler)).Methods("PUT")
	router.HandleFunc("/homepage", a.validate(postRouter.homepageHandler)).Methods("POST")
	return router
}

func (pr *postRouter) postHandler(w http.ResponseWriter, r *http.Request) {
	var resp root.ResponseSlice

	Username := context.Get(r, "Username")
	UserID := context.Get(r, "ID")
	LastLoggedIn := context.Get(r, "LastLoggedIn")
	post, emptyFields, err := decodeData(r, []string{"Text"})
	if err != nil {
		fmt.Println(err)
		resp.Message = "Error Occured"
		resp.Err = err
		Json(w, http.StatusInternalServerError, resp)
		return
	}

	var param []string
	param = append(param, Username.(string), UserID.(string), LastLoggedIn.(string))
	check := pr.userService.CheckToken(param)
	if check == false {
		resp.Message = "Invalid Token"
		Json(w, http.StatusUnauthorized, resp)
		return
	}

	if len(emptyFields) != 0 {
		resp.Message = "Bad Request."
		resp.Data = emptyFields
		Json(w, http.StatusBadRequest, resp)
		return
	}
	post.OwnerID = UserID.(string)
	pi, err2 := pr.postService.Post(&post)
	if err2 != nil {
		resp.Message = "Error Occured"
		resp.Err = err2
		Json(w, http.StatusInternalServerError, resp)
		return
	}

	resp.Message = "Operation successful"
	resp.Data = pi
	Json(w, http.StatusOK, resp)
	return
}

func (pr *postRouter) homepageHandler(w http.ResponseWriter, r *http.Request) {
	var resp root.ResponseSlice
	Username := context.Get(r, "Username")
	UserID := context.Get(r, "ID")
	LastLoggedIn := context.Get(r, "LastLoggedIn")

	data, emptyFields, err := decodeData(r, []string{"Limit", "IDs"})
	if err != nil {
		fmt.Println(err)
		resp.Message = "Error Occured"
		resp.Err = err
		Json(w, http.StatusInternalServerError, resp)
		return
	}
	var param []string
	param = append(param, Username.(string), UserID.(string), LastLoggedIn.(string))
	check := pr.userService.CheckToken(param)
	if check == false {
		resp.Message = "Invalid Token"
		Json(w, http.StatusUnauthorized, resp)
		return
	}
	if len(emptyFields) != 0 {
		resp.Message = "Bad Request."
		resp.Data = emptyFields
		Json(w, http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Operation successful"
	resp.Data = data
	Json(w, http.StatusOK, resp)
	return
}

func decodeData(r *http.Request, checks []string) (root.Post, []string, error) {
	var p root.Post
	if r.Body == nil {
		return p, []string{}, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	emptyFields := []string{}
	if err != nil {
		return p, []string{}, err
	}
	for _, check := range checks {
		if check == "" {
			emptyFields = append(emptyFields, check)
			continue
		}
		temp := reflect.Indirect(reflect.ValueOf(&p))
		fieldValue := temp.FieldByName(string(check))
		if fieldValue != temp.FieldByName("") {
			if (fieldValue.Type().String() == "string" && fieldValue.Len() == 0) || (fieldValue.Type().String() != "string" && fieldValue.IsNil()) {
				fmt.Println("EMPTY->", check)

				emptyFields = append(emptyFields, check)
			}
		} else {
			fmt.Println("EMPTY->", check)

			emptyFields = append(emptyFields, check)
		}
	}
	return p, emptyFields, nil
}
