package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcproject/project/resources"
)

func Verifytoken(tokenStr string) (int, error) {
	ok := resources.Localsession.ValidateUser(tokenStr)
	if !ok {
		resources.Log.Warn(" user already logged out trying use invalid token")

		return 0, status.Errorf(
			codes.Unauthenticated, fmt.Sprintf("invaild tokens please login again session expired"))
	}
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, status.Errorf(
				codes.Unauthenticated, fmt.Sprintf("invaild tokens "))
		}
		return 0, status.Errorf(
			codes.Unauthenticated, fmt.Sprintf("unknown error "))
	}

	if !tkn.Valid {
		return 0, status.Errorf(
			codes.Unauthenticated, fmt.Sprintf("token expired"))
	}
	//w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
	return (*claims).Id, nil
}
