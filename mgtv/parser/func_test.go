package parser

import (
	"reflect"
	"testing"
	"time"
)

func TestDurationUnmarshalText(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "testCase01",
			args:    args{s: "45:12"},
			want:    2712 * time.Second,
			wantErr: false,
		},
		{
			name:    "testCase02",
			args:    args{s: "145:12"},
			want:    8712 * time.Second,
			wantErr: false,
		},
		{
			name:    "testCase03",
			args:    args{s: "-1:12"},
			want:    0 * time.Second,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DurationUnmarshalText(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DurationUnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DurationUnmarshalText() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMgtvTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "testCase01",
			args:    args{s: "2019-12-23 11:28:54.0"},
			want:    time.Date(2019, time.December, 23, 11, 28, 54, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "testCase02",
			args:    args{s: "2019-12-23 11:28:54"},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "testCase03",
			args:    args{s: "2020-04-02 09:50:42.0"},
			want:    time.Date(2020, time.April, 2, 9, 50, 42, 0, time.UTC),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMgtvTime(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMgtvTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMgtvTime() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestParseMgtvPlayCounter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "testCase01",
			args:    args{s: "1.9万"},
			want:    19000,
			wantErr: false,
		},
		{
			name:    "testCase02",
			args:    args{s: "1896"},
			want:    1896,
			wantErr: false,
		},
		{
			name:    "testCase03",
			args:    args{s: "6484.1万"},
			want:    64841000,
			wantErr: false,
		},
		{
			name:    "testCase04",
			args:    args{s: "4.1亿"},
			want:    409999999,
			wantErr: false,
		},
		{
			name:    "testCase04",
			args:    args{s: "4.1 亿"},
			want:    409999999,
			wantErr: false,
		},
		{
			name:    "testCase04",
			args:    args{s: "192.1 亿 "},
			want:    19210000000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMgtvPlayCounter(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMgtvPlayCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMgtvPlayCounter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
