package main

import (
	"fmt"
	httptranport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	_ "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	. "go-kit-demo/service"
	"go-kit-demo/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	user := UserService{}
	endp := GenUserEndpoint(&user)

	serverHandler := httptranport.NewServer(endp, DecodeUserRequest, EncodeUserResponse)

	r := mux.NewRouter()
	//r.Handle(`/user/{uid:\d+}`,serverHandler)
	r.Methods("GET").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	r.Methods("GET").Path(`/health`).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "application/json")
		writer.Write([]byte(`{"status":"ok"}`))
	})

	errChan := make(chan error)
	go func() {
		util.Register()
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			logrus.Error(err)
			errChan <- err
		}
	}()

	go func() {
		sigC := make(chan os.Signal)
		signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
		errChan<-fmt.Errorf("%s",<-sigC)
	}()
	getErr := <-errChan
	util.DeRegister()
	logrus.Error(getErr)
}

