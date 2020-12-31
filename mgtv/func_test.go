package mgtv

import "testing"

func TestGenJsonp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "testCase-01", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := genJsonp(); got != tt.want {
			//	t.Errorf("genJsonp() = %v, want %v", got, tt.want)
			//}
			got := GenJsonp()
			println(got)
		})
	}
}
