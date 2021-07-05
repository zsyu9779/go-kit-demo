package main

import (
	"context"
	"fmt"
	httptranport "github.com/go-kit/kit/transport/http"
	"net/url"
	"product/service"
)

func main() {
	target, _ := url.Parse("http://localhost:8080")
	client := httptranport.NewClient("GET", target, service.GetUserInfoRequest,service.GetUserInfoResponse)
	getUserInfo := client.Endpoint()
	res, _ := getUserInfo(context.Background(),service.UserRequest{
		Uid: 102,
	})
	userInfo := res.(service.UserResponse)
	fmt.Printf("%+v\n",userInfo)
}
