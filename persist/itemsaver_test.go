package persist

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/model"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		URL: "https://www.mgtv.com/b/1/2111.html",
		ID:  "7653401",
		Payload: model.Episode{
			ChannelID:   "3",
			DramaID:     "334346",
			DramaTitle:  "电影有鸡汤 2020",
			EpisodeID:   "7653401",
			Title1:      "18",
			Title2:      "洼田正孝好看到飞起",
			Title3:      "《初恋》中二又热血！洼田正孝好看到飞起",
			Title4:      "05:07",
			EpisodeURL:  "https://www.mgtv.com/b/1/2111.html",
			Duration:    "05:07",
			ContentType: "0",
			Image:       "https://0img.hitv.com/preview/sp_images/2020/3/4/dianying/334346/7653401/20200304184507992.jpg_220x125.jpg",
			IsIntact:    "1",
			IsNew:       "0",
			IsVIP:       "0",
			PlayCounter: "4.4万",
			TS:          "2020-03-04 18:40:50.0",
			NextID:      "7616861",
			SrcClipID:   "334346",
		},
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	//const index= "mgtv_episode"

	Save(client, expected)
}
