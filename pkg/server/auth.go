package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"

	"Go-Social/pkg"
)

type authHelper struct {
	secret string
}

type claims struct {
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type contextKey string

func (c contextKey) String() string {
	return "mypackage context key " + string(c)
}

var (
	contextKeyAuthtoken = contextKey("auth-token")
)

func (a *authHelper) newToken(user root.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":           user.ID,
		"Username":     user.Username,
		"LastLoggedIn": user.LastLoggedIn,
	})

	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		fmt.Println("Error : token not generated")
		fmt.Println(err)
		return ""
	}
	return tokenString
}

func (a *authHelper) validate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var resp root.ResponseSlice

		tokenString := req.Header.Get("Authorization")
		if tokenString == "" {
			resp.Message = "No authorization token"
			Json(res, http.StatusUnauthorized, resp)
			return
		}
		token, err1 := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(a.secret), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			context.Set(req, "Username", claims["Username"])
			context.Set(req, "ID", claims["ID"])
			context.Set(req, "LastLoggedIn", claims["LastLoggedIn"])
			next(res, req)
		} else {
			resp.Message = "Invalid token"
			resp.Err = err1
			Json(res, http.StatusBadRequest, resp)
		}
	})
}
