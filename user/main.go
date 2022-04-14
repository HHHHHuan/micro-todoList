package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"user/conf"
	"user/core"
	"user/services"
)
/*
* @Author: hh
* @Date:   2022/4/13 15:57
 */

func main() {
	conf.Init()
	// etcd注册
	etcdReg:= etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	microService:=micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_=services.RegisterUserServiceHandler(microService.Server(),new(core.UserService))
	_=microService.Run()
}
