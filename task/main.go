package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"task/conf"
	"task/core"
	"task/service"
)

func main() {
	conf.Init()
	// etcd注册
	etcdReg:= etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到微服务实例
	microService:=micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_=service.RegisterTaskServiceHandler(microService.Server(),new(core.TaskService))
	_=microService.Run()
}
