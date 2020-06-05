package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	httpapi "github.com/spadesk1991/go-micro-grpc/client/api/http"
	"github.com/spadesk1991/go-micro-grpc/client/api/http/middleware"
	"github.com/spadesk1991/go-micro-grpc/server/Services"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs(":8500"))
	router := gin.Default()
	srv := web.NewService(
		web.Name("httpGETProdService"),
		web.Address(":8301"),
		web.Registry(reg),
		web.Handler(router),
	)
	prodGrpcServer := Services.NewProdService("prodGRPCService", nil)
	router.Use(middleware.AddGRPCServer(prodGrpcServer))
	httpapi.InitRouter(router)
	srv.Init()
	srv.Run()
}
