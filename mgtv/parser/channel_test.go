package parser

import (
	"git.trac.cn/nv/spider/engine"
	"io/ioutil"
	"testing"
)

func TestParseChannel(t *testing.T) {
	type args struct {
		contents []byte
	}

	testCase01Contents, err := ioutil.ReadFile("channel_test_data.html")
	if err != nil {
		panic(err)
	}
	//var parserFunc = engine.NilParser

	tests := []struct {
		name        string
		args        args
		want        engine.ParseResult
		requestSize int
		itemSize    int
	}{
		{
			name:        "testCase01",
			args:        args{contents: testCase01Contents},
			want:        engine.ParseResult{},
			requestSize: 60,
			itemSize:    60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseChannel(tt.args.contents)
			if len(got.Requests) != tt.requestSize {
				t.Errorf("result should have %d "+"requests; but had %d",
					tt.requestSize, len(got.Requests))
			}
			if len(got.Items) != tt.itemSize {
				t.Errorf("result should have %d "+"item; but had %d",
					tt.itemSize, len(got.Items))
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ParseChannel() = %v, want %v", got, tt.want)
			//}
		})
	}
}
