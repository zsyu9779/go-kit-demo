package util

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"net"
)

var consulClient *api.Client

func init() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := api.NewClient(config)
	if err != nil {
		logrus.Fatal(err)
	}
	consulClient = client
}
func Register() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	check := api.AgentServiceCheck{
		Interval: "5s",
		HTTP:     fmt.Sprintf("http://%s:8080/health", getClientIp()),
	}
	reg := api.AgentServiceRegistration{
		Kind:    "",
		ID:      "userservice",
		Name:    "userservice",
		Tags:    []string{"primary"},
		Port:    8080,
		Address: getClientIp(),
		Check:   &check,
	}
	err := consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		logrus.Fatal(err)
	}
}
func DeRegister() {
	consulClient.Agent().ServiceDeregister("userservice")
}
func getClientIp() string {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}

	return ""
}
