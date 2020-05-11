package main

import (
	"context"
	"contrib.go.opencensus.io/exporter/jaeger"
	"contrib.go.opencensus.io/exporter/prometheus"
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
	"github.com/micro/go-plugins/wrapper/trace/opencensus/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
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

	if err := view.Register(opencensus.DefaultServerViews...); err != nil {
		log.Fatal(err)
	}
	exporter1, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Fatal(err)
	}
	view.RegisterExporter(exporter1)

	go func() {
		metricsEndPoint := fmt.Sprintf(":%d", setting.ServerSetting.MetricsPort)

		http.Handle("/opencensus/metrics", exporter1)
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("[info] start metrics server of prometheus listening %s", metricsEndPoint)
		http.ListenAndServe(metricsEndPoint, nil)
	}()

	exporter, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: setting.ServerSetting.JaegerAgentAddr,
		Process: jaeger.Process{
			ServiceName: "saveitem",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

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
		micro.WrapClient(opencensus.NewClientWrapper()),
		micro.WrapHandler(opencensus.NewHandlerWrapper()),
		micro.WrapSubscriber(opencensus.NewSubscriberWrapper()),
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
