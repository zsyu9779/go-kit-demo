package main

import (
	httptranport "github.com/go-kit/kit/transport/http"
	."go-kit-demo/service"
	"net/http"
)

func main()  {
	user :=UserService{}
	endp :=GenUserEndpoint(&user)

	serverHandler :=httptranport.NewServer(endp,DecodeUserRequest,EncodeUserResponse)

	http.ListenAndServe(":8080",serverHandler)
}