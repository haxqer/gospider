package main

import (
	"context"
	"fmt"
	item "git.trac.cn/nv/spider/services/item/proto"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {

	registryReg := etcd.NewRegistry(registry.Addrs("172.31.0.134:2379"))
	microSelector := selector.NewSelector(
		selector.Registry(registryReg),
		selector.SetStrategy(selector.RoundRobin),
	)

	client.DefaultClient = grpc.NewClient(
		client.Selector(microSelector),
		client.Registry(registryReg),
	)

	req := client.NewRequest("api.trac.cn.saveitem", "Save.SaveItem", &item.Item{
		ChannelId:   2,
		DramaId:     334346,
		DramaTitle:  "电影有鸡汤 2020",
		EpisodeId:   7653401,
		Title1:      "18",
		Title2:      "洼田正孝好看到飞起",
		Title3:      "《初恋》中二又热血！洼田正孝好看到飞起",
		Title4:      "05:07",
		EpisodeUrl:  "https://www.mgtv.com/b/1/21199.html",
		Duration:    888,
		ContentType: "0",
		Image:       "https://0img.hitv.com/preview/sp_images/2020/3/4/dianying/334346/7653401/20200304184507992.jpg_220x125.jpg",
		IsIntact:    "1",
		IsNew:       "0",
		IsVip:       "0",
		PlayCounter: 518000,
		Ts:          "2020-03-04 18:40:50.0",
		NextId:      "7616861",
		SrcClipId:   "1111",
	})

	rsp := &item.SaveResponse{}
	// Call service
	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}




	//newService := grpc.NewService(
	//	service.Registry(registryReg),
	//	)
	//newService.Init()
	//
	//// use the generated client stub
	//cl := item.NewSaveService("api.trac.cn.saveitem", newService.Client())
	//
	//// Set arbitrary headers in context
	//ctx := metadata.NewContext(context.Background(), map[string]string{
	//	"X-User-Id": "john",
	//	"X-From-Id": "script",
	//})
	//
	//rsp, err := cl.SaveItem(ctx, &item.Item{
	//	ChannelId:   2,
	//	DramaId:     334346,
	//	DramaTitle:  "电影有鸡汤 2020",
	//	EpisodeId:   7653401,
	//	Title1:      "18",
	//	Title2:      "洼田正孝好看到飞起",
	//	Title3:      "《初恋》中二又热血！洼田正孝好看到飞起",
	//	Title4:      "05:07",
	//	EpisodeUrl:  "https://www.mgtv.com/b/1/21199.html",
	//	Duration:    888,
	//	ContentType: "0",
	//	Image:       "https://0img.hitv.com/preview/sp_images/2020/3/4/dianying/334346/7653401/20200304184507992.jpg_220x125.jpg",
	//	IsIntact:    "1",
	//	IsNew:       "0",
	//	IsVip:       "0",
	//	PlayCounter: 518000,
	//	Ts:          "2020-03-04 18:40:50.0",
	//	NextId:      "7616861",
	//	SrcClipId:   "334346",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	fmt.Println(rsp.Code)
}