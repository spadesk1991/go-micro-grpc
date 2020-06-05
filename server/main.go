package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/spadesk1991/go-micro-grpc/server/Services"
	"github.com/spadesk1991/go-micro-grpc/server/ServicesImpl"
)

func main() {
	reg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	s := micro.NewService(
		micro.Name("prodGRPCService"),
		micro.Address(":8300"),
		micro.Registry(reg),
	)
	Services.RegisterProdServiceHandler(s.Server(), new(ServicesImpl.ProdService))
	s.Init()
	s.Run()
}
