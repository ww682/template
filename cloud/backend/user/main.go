package main

import (
	"fmt"
	"log"
	"user/config"
	//"user/handler"
	//pb "user/proto"

	"github.com/go-micro/plugins/v4/registry/consul"

	"go-micro.dev/v4"
)

func main() {
	cfg, err := config.LoadConfig("conf/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	_, err = config.InitMysql(&cfg.Mysql)
	if err != nil {
		return
	}

	reg := consul.NewRegistry()
	// 创建一个服务
	service := micro.NewService(
		micro.Registry(reg),
		micro.Version("v1.0"),
		micro.Name("user"),
	)
	// 初始化服务
	service.Init()

	// 注册一个gRPC服务处理器
	//pb.RegisterUserHandler(service.Server(), handler.New())

	// Run service
	if err := service.Run(); err != nil {
		fmt.Println("Error running service:", err)
	}
}
