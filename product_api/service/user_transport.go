package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetUserInfoRequest(ctx context.Context, request *http.Request, r interface{}) error {
	userRequest := r.(UserRequest)
	request.URL.Path+="/user/"+fmt.Sprintf("%d", userRequest.Uid)
	return nil
}
func GetUserInfoResponse(ctx context.Context, response *http.Response) (interface{},error){
	if response.StatusCode>400{
		return nil,errors.New("no data")
	}
	var userResponse UserResponse
	err := json.NewDecoder(response.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, err
}