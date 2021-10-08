package controller

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpcproject/project/resources"
)

func extactToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	//var authorization string
	if ok {
		fmt.Println("meta data =%v", md["authorization"])
		a, okay := md["authorization"]
		if okay {
			bearer := a[0]
			tokn := bearer[6:]
			fmt.Println("%v", tokn, okay)
			return tokn, nil
		}
	}
	resources.Log.Warn("error token not found %v")

	return "", status.Errorf(
		codes.PermissionDenied, fmt.Sprintf("token not found"))
}

func tokenwork(ctx context.Context, id int) (bool, error) {

	token, err := extactToken(ctx)
	if err != nil {
		return false, err
	}

	uid, err := Verifytoken(token)
	if err != nil {
		resources.Log.Warn("error token not invalid %v", err)

		return false, status.Errorf(
			codes.Unauthenticated, err.Error())
	}
	if uid != id && id != 0 {
		resources.Log.Warn("error token not invalid/fake id did not match %v", err)
		return false, status.Errorf(
			codes.Unauthenticated, err.Error())
	}
	return true, nil
}
