package parser

import (
	"git.trac.cn/nv/spider/engine"
	"io/ioutil"
	"testing"
)

func TestParseChannelList(t *testing.T) {
	type args struct {
		contents []byte
	}

	testCase01Contents, err := ioutil.ReadFile("channel_test_data.html")
	if err != nil {
		panic(err)
	}
	var parserFunc = ParseChannelList

	var tests = []struct {
		name        string
		args        args
		want        engine.ParseResult
		requestSize int
		itemSize    int
	}{
		{
			name: "testCase01",
			args: args{contents: testCase01Contents},
			want: engine.ParseResult{
				Requests: []engine.Request{
					{
						Url:        "https://list.mgtv.com/-------------.html?channelId=1",
						ParserFunc: parserFunc,
					},
				},
				Items: nil,
			},
			requestSize: 42,
			itemSize:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseChannelList(tt.args.contents)
			if len(got.Requests) != tt.requestSize {
				t.Errorf("result should have %d "+"requests; but had %d",
					tt.requestSize, len(got.Requests))
			}
			if len(got.Items) != tt.itemSize {
				t.Errorf("result should have %d "+"item; but had %d",
					tt.itemSize, len(got.Items))
			}
			//if got := ParseChannelList(tt.args.contents); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ParseChannelList() = %v, want %v", got, tt.want)
			//}
		})
	}
}
