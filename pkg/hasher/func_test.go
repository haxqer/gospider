package hasher

import "testing"

func TestGetMD5Hash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:    "testCase01",
			args:    args{text: "123"},
			want:    "202cb962ac59075b964b07152d234b70",
		},
		{
			name:    "testCase02",
			args:    args{text: "http://www.mgtv.com/b/291325/2979284.html"},
			want:    "d0e2109a8ff867ee55f83f8765273ff9",
		},
		{
			name:    "testCase03",
			args:    args{text: "abc"},
			want:    "900150983cd24fb0d6963f7d28e17f72",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMD5Hash(tt.args.text); got != tt.want {
				t.Errorf("GetMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}