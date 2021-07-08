package util

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
	"go-kit-demo/product/service"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"io"
	"net/url"
	"os"
)

func ConsulClient() {
	{
		//第一步，创建client
		config := api.DefaultConfig()
		config.Address = "127.0.0.1:8500"//注册中心地址
		apiClient, _ := api.NewClient(config)
		client := consul.NewClient(apiClient)

		var logger log.Logger
		{
			logger = log.NewLogfmtLogger(os.Stdout)
		}
		{
			tags := []string{"primary"}
			//可实时查询服务实例的状态信息
			instancer := consul.NewInstancer(client, logger, "userservice", tags, true)

			{
				factory := func(serviceUrl string) (endpoint.Endpoint, io.Closer, error) {
					target, _ := url.Parse("http://" + serviceUrl)
					return httptransport.NewClient("GET", target, service.GetUserInfoRequest, service.GetUserInfoResponse).Endpoint(), nil, nil
				}
				endpointer := sd.NewEndpointer(instancer, factory, logger)
				endpoints, _ :=endpointer.Endpoints()
				fmt.Println("has",len(endpoints),"services")
				getUserInfo := endpoints[0]
				res, _ := getUserInfo(context.Background(),service.UserRequest{
					Uid: 102,
				})
				userInfo := res.(service.UserResponse)
				fmt.Printf("%+v\n",userInfo)

			}
		}

	}

}
