package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

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
		"Username":  user.Username,
		"UpdatedAt": user.UpdatedAt,
	})

	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		fmt.Println("Error : token not generated")
		fmt.Println(err)
		return ""
	}
	fmt.Println(tokenString, "tokenString")
	return tokenString
}

func (a *authHelper) validate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("Authorization")
		if tokenString == "" {
			Error(res, http.StatusUnauthorized, "No authorization token")
			return
		}
		token, err1 := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return a.secret, nil
		})
		if err1 != nil {
			Error(res, http.StatusUnauthorized, "Invalid Token")
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
		}
		Error(res, http.StatusUnauthorized, "Unauthorized")
		return
	})
}
