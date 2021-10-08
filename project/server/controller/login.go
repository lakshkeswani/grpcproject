package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcproject/project/server/model"
	_ "strconv"
	"time"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenrateToken(credentials model.Credentials, ID int) (string, error) {

	expirationTime := time.Now().Add(time.Hour * 5)
	claims := &Claims{
		Id:       int(ID),
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", status.Errorf(
			codes.PermissionDenied, fmt.Sprintf("worong username or password"))
	}

	return tokenString, nil

}
