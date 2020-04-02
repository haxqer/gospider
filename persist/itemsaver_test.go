package persist

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/pkg/logging"
	"git.trac.cn/nv/spider/pkg/setting"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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
		SrcClipId:   "334346",
	}

	setting.Setup()
	logging.Setup()
	model.Setup()

	err := Save(&expected)
	if err != nil {
		t.Errorf("err : %+v", err)
	}
}
