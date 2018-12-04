package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"gopkg.in/mgo.v2/bson"

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
	router.HandleFunc("/post/{postId}/comment", a.validate(postRouter.commentHandler)).Methods("POST")
	router.HandleFunc("/post/{postId}/like", a.validate(postRouter.likeHandler)).Methods("GET")
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
	pi, err2 := pr.postService.Post(&post, pr.userService)
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

	var param []string
	param = append(param, Username.(string), UserID.(string), LastLoggedIn.(string))
	check := pr.userService.CheckToken(param)
	if check == false {
		resp.Message = "Invalid Token"
		Json(w, http.StatusUnauthorized, resp)
		return
	}

	data, emptyFields, err := decodeData(r, []string{"Limit", "IDs"})
	if err != nil {
		fmt.Println(err)
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
	var IDs []bson.ObjectId
	for _, v := range data.IDs {
		IDs = append(IDs, bson.ObjectIdHex(v))
	}
	posts, err := pr.postService.GetPosts(data.Limit, IDs)
	resp.Message = "Operation successful"
	resp.Data = posts
	Json(w, http.StatusOK, resp)
	return
}

func (pr *postRouter) commentHandler(w http.ResponseWriter, r *http.Request) {
	var resp root.ResponseSlice

	vars := mux.Vars(r)

	Username := context.Get(r, "Username")
	UserID := context.Get(r, "ID")
	LastLoggedIn := context.Get(r, "LastLoggedIn")
	var param []string

	param = append(param, Username.(string), UserID.(string), LastLoggedIn.(string))
	check := pr.userService.CheckToken(param)
	if check == false {
		resp.Message = "Invalid Token"
		Json(w, http.StatusUnauthorized, resp)
		return
	}
	data, emptyFields, err := decodeData(r, []string{"Comment"})
	if err != nil {
		fmt.Println(err)
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
	postID := vars["postId"]
	post := pr.postService.AddComment(postID, root.Comments{ID: bson.NewObjectId(), UserID: bson.ObjectIdHex(UserID.(string)), Comment: data.Comment})
	if post == nil {
		resp.Message = "Post Not Found"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Operation successful"
	resp.Data = post
	Json(w, http.StatusOK, resp)
	return
}

func (pr *postRouter) likeHandler(w http.ResponseWriter, r *http.Request) {
	var resp root.ResponseSlice

	vars := mux.Vars(r)

	Username := context.Get(r, "Username")
	UserID := context.Get(r, "ID")
	LastLoggedIn := context.Get(r, "LastLoggedIn")
	var param []string

	param = append(param, Username.(string), UserID.(string), LastLoggedIn.(string))
	check := pr.userService.CheckToken(param)
	if check == false {
		resp.Message = "Invalid Token"
		Json(w, http.StatusUnauthorized, resp)
		return
	}
	postID := vars["postId"]
	post := pr.postService.UpdateLike(postID, root.Likes{ID: bson.NewObjectId(), UserID: bson.ObjectIdHex(UserID.(string))})
	if post == nil {
		resp.Message = "Post Not Found"
		Json(w, http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Operation successful"
	resp.Data = post
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
			if (fieldValue.Type().String() == "string" && fieldValue.Len() == 0) || (fieldValue.Type().String() != "string" && fieldValue.Type().String() != "int" && fieldValue.IsNil()) {
				fmt.Println("EMPTY->", check)

				emptyFields = append(emptyFields, string(check))
			}
		} else {
			fmt.Println("EMPTY->", check)

			emptyFields = append(emptyFields, string(check))
		}
	}
	return p, emptyFields, nil
}
