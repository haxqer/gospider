package main

import (
	"context"
	"fmt"
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/pkg/logging"
	"git.trac.cn/nv/spider/pkg/setting"
	item "git.trac.cn/nv/spider/services/itemsave/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport/grpc"
	limiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

type Save struct{}

func (s *Save) SaveItem(ctx context.Context, req *item.Item, rsp *item.SaveResponse) error {
	err := insertOrDuplicate(&model.Mgtv{
		EpisodeId:   req.EpisodeId,
		ChannelId:   req.ChannelId,
		DramaId:     req.DramaId,
		DramaTitle:  req.DramaTitle,
		Title1:      req.Title1,
		Title2:      req.Title2,
		Title3:      req.Title3,
		Title4:      req.Title4,
		EpisodeUrl:  req.EpisodeUrl,
		Duration:    req.Duration,
		ContentType: req.ContentType,
		Image:       req.Image,
		IsIntact:    req.IsIntact,
		IsNew:       req.IsNew,
		IsVip:       req.IsVip,
		PlayCounter: req.PlayCounter,
		Ts:          req.Ts,
		NextId:      req.NextId,
		SrcClipId:   req.SrcClipId,
	})
	if err != nil {
		rsp.Code = 500
		logging.Error(err)
		return err
	}
	rsp.Code = 200
	return nil
}

func insertOrDuplicate(item *model.Mgtv) error {
	err := model.PersistDB.Save(item).Error
	if err != nil {
		return err
	}
	return nil
}

func init() {
	setting.Setup()
	logging.Setup()
	model.Setup()
}

func main() {
	go func() {
		metricsEndPoint := fmt.Sprintf(":%d", setting.ServerSetting.MetricsPort)
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("[info] start metrics server of prometheus listening %s", metricsEndPoint)
		http.ListenAndServe(metricsEndPoint, nil)
	}()

	const QPS = 1000
	registryReg := etcd.NewRegistry(registry.Addrs(setting.ServerSetting.RegistryAddr))
	//broker := brokerRedis.NewBroker()
	transport := grpc.NewTransport()

	newService := micro.NewService(
		micro.Name("api.trac.cn.saveitem"),
		micro.Transport(transport),
		//micro.Broker(broker),
		micro.Registry(registryReg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		micro.WrapHandler(limiter.NewHandlerWrapper(QPS)),
		//micro.Address(":19999"),
	)

	newService.Init()

	item.RegisterSaveHandler(newService.Server(), new(Save))

	// Run server
	if err := newService.Run(); err != nil {
		log.Fatal(err)
	}
}
