package parser

import (
	"git.trac.cn/nv/spider/engine"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseEpisode(t *testing.T) {
	type args struct {
		contents []byte
	}

	testCase01Contents, err := ioutil.ReadFile("episode_test_data.txt")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name string
		args args
		want engine.ParseResult
	}{
		{
			name: "testCase01",
			args: args{contents: testCase01Contents},
			want: engine.ParseResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseEpisode(tt.args.contents, "", "")
			//for _, m := range got.Requests {
			//	fmt.Printf("%s", m.Url)
			//}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseEpisode() = %v, want %v", got, tt.want)
			}
		})
	}
}
