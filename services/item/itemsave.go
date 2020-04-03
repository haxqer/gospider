package main

import (
	"context"
	"fmt"
	item "git.trac.cn/nv/spider/services/item/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport/grpc"

	"log"
)

type Save struct{}

func (s *Save) SaveItem(ctx context.Context, req *item.Item, rsp *item.SaveResponse) error {
	log.Print("Received Save.SaveItem request")
	fmt.Printf("%+v \n", req)
	rsp.Code = 123
	return nil
}

func main() {
	registryReg := etcd.NewRegistry(registry.Addrs("172.31.0.134:2379"))
	//broker := kafka.NewBroker()
	transport := grpc.NewTransport()

	newService := micro.NewService(
		micro.Name("api.trac.cn.saveitem"),
		micro.Transport(transport),
		//micro.Broker(broker),
		micro.Registry(registryReg),
		//micro.Address(":19999"),
	)

	newService.Init()

	item.RegisterSaveHandler(newService.Server(), new(Save))

	// Run server
	if err := newService.Run(); err != nil {
		log.Fatal(err)
	}
}

