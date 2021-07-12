package main

import (
	"flag"
	"fmt"
	httptranport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	_ "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"go-kit-demo/user_api/service"
	"go-kit-demo/user_api/util"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	name := flag.String("name", "", "service_name")
	port := flag.Int("p", 0, "port")
	flag.Parse()
	if *name == "" {
		logrus.Fatal("need service name")
	}
	if *port == 0 {
		logrus.Fatal("need service name")
	}
	util.SetServiceNameAndPort(*name,*port)

	user := service.UserService{}
	endp := service.GenUserEndpoint(&user)

	serverHandler := httptranport.NewServer(endp, service.DecodeUserRequest, service.EncodeUserResponse)

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
		err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
		if err != nil {
			logrus.Error(err)
			errChan <- err
		}
	}()

	go func() {
		sigC := make(chan os.Signal)
		signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sigC)
	}()
	getErr := <-errChan

	util.DeRegister()
	logrus.Error(getErr)
}
