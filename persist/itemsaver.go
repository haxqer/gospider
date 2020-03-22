package persist

import (
	"context"
	"errors"
	"git.trac.cn/nv/spider/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		var storedID = make(map[string]bool)
		for {
			item := <-out

			if storedID[item.ID] {
				continue
			}
			storedID[item.ID] = true
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item) error {
	if item.ID == "" {
		return errors.New("must supply ID")
	}
	_, err := client.Index().
		Index("mgtv_episode").
		Id(item.ID).
		BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
