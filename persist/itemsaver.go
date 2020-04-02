package persist

import (
	"fmt"
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/pkg/logging"
)

func ItemSaver() (chan engine.Item, error) {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		var storedID = make(map[int]bool)
		for {
			item := <-out

			if storedID[item.EpisodeId] {
				continue
			}
			storedID[item.EpisodeId] = true
			itemCount++

			if itemCount%10000 == 0 {
				logging.Info(fmt.Sprintf("Item Saver: got item #%d: %v", itemCount, item))
			}
			err := Save(&item)
			if err != nil {
				logging.Error(fmt.Sprintf("Item Saver: error saving item %v: %v", item, err))
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

//func Save(client *elastic.Client, item engine.Item) error {
//	if item.ID == "" {
//		return errors.New("must supply ID")
//	}
//	_, err := client.Index().
//		Index("mgtv_episode").
//		Id(item.ID).
//		BodyJson(item).Do(context.Background())
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
