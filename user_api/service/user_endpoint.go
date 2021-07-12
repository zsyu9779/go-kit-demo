package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-kit-demo/user_api/util"
	"strconv"
)

type UserRequest struct {
	Uid int `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndpoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error){
		r := request.(UserRequest)
		result := userService.GetName(r.Uid)+strconv.Itoa(util.ServicePort)
		return UserResponse{Result: result},nil
	}
}