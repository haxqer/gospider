package persist

import (
	"context"
	"fmt"
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/pkg/logging"
	"git.trac.cn/nv/spider/pkg/setting"
	itemsave "git.trac.cn/nv/spider/services/itemsave/proto"
	hystrixGo "github.com/afex/hystrix-go/hystrix"
	metricCollector "github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/afex/hystrix-go/plugins"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	"github.com/patrickmn/go-cache"
	"net"
	"net/http"
	"strconv"
	"time"
)

func ItemSaver() (chan engine.Item, error) {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		storedID := cache.New(setting.ServerSetting.SaveItemExpire, setting.ServerSetting.SaveItemExpire + 3*time.Minute)
		for {
			item := <-out

			episodeStr := strconv.Itoa(int(item.DramaId)) + ":" + strconv.Itoa(int(item.EpisodeId))
			if _, found := storedID.Get(episodeStr); found {
				continue
			}
			storedID.SetDefault(episodeStr, true)

			err := call(&item)
			if err != nil {
				logging.Error(fmt.Sprintf("Item Saver: error saving itemsave %v: %v", item, err))
				continue
			}
			itemCount++

			if itemCount%10000 == 0 {
				logging.Info(fmt.Sprintf("Item Saver: got itemsave #%d: %v", itemCount, item))
			}
		}
	}()
	return out, nil
}

func Save(mgtv *model.Mgtv) error {
	err := model.InsertOnDuplicate(mgtv)
	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

var itemSaveClient itemsave.SaveService
//var itemPub micro.Event

func Setup() {
	microRegistry := etcd.NewRegistry(registry.Addrs(setting.ServerSetting.RegistryAddr))
	microSelector := selector.NewSelector(
		selector.Registry(microRegistry),
		selector.SetStrategy(selector.RoundRobin),
	)
	microTransport := grpc.NewTransport()
	microService := micro.NewService(
		micro.Name("saveitem.client"),
		micro.Selector(microSelector),
		micro.Transport(microTransport),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	microService.Init()

	hystrixGo.DefaultMaxConcurrent = 1
	hystrixGo.DefaultTimeout = 1
	hystrixGo.DefaultSleepWindow = 1000
	hystrixStreamHandler := hystrixGo.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	c, err := plugins.InitializeStatsdCollector(&plugins.StatsdCollectorConfig{
		StatsdAddr: "localhost:8125",
		Prefix:     "myapp.hystrix",
	})
	if err != nil {
		panic(fmt.Sprintf("could not initialize statsd client: %v", err))
	}

	metricCollector.Registry.Register(c.NewStatsdCollector)


	itemSaveClient = itemsave.NewSaveService("api.trac.cn.saveitem", microService.Client())
	//itemPub = micro.NewEvent("trac.saveitem", microService.Client())
}

func call(mgtv * model.Mgtv) error {
	rsp, err := itemSaveClient.SaveItem(context.TODO(), &itemsave.Item{
		ChannelId:   mgtv.ChannelId,
		DramaId:     mgtv.DramaId,
		DramaTitle:  mgtv.DramaTitle,
		EpisodeId:   mgtv.EpisodeId,
		Title1:      mgtv.Title1,
		Title2:      mgtv.Title2,
		Title3:      mgtv.Title3,
		Title4:      mgtv.Title4,
		EpisodeUrl:  mgtv.EpisodeUrl,
		Duration:    mgtv.Duration,
		ContentType: mgtv.ContentType,
		Image:       mgtv.Image,
		IsIntact:    mgtv.IsIntact,
		IsNew:       mgtv.IsNew,
		IsVip:       mgtv.IsVip,
		PlayCounter: mgtv.PlayCounter,
		Ts:          mgtv.Ts,
		NextId:      mgtv.NextId,
		SrcClipId:   mgtv.SrcClipId,
	})
	if err != nil {
		//if errors.Cause(err) == hystrixGo.ErrCircuitOpen {
		if hystrixErr, ok := err.(hystrixGo.CircuitError); ok {
			fmt.Println(hystrixErr)
			return err
		}
		fmt.Println("grpc call err: ", err, rsp)
		return err
	}

	return nil
}
