package persist

import (
	"context"
	"fmt"
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/pkg/logging"
	"git.trac.cn/nv/spider/pkg/setting"
	itemsave "git.trac.cn/nv/spider/services/itemsave/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/patrickmn/go-cache"
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
		micro.Name("greeter.client"),
		micro.Selector(microSelector),
		micro.Transport(microTransport),
	)
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
		fmt.Println("call err: ", err, rsp)
		return err
	}

	return nil
}
