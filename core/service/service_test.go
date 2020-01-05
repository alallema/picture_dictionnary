package service

import "testing"

func TestCheckFormat(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "No format",
			args: args{
				name: "file.txt",
			},
			want: 0,
		},
		{
			name: "Test picture format jpeg",
			args: args{
				name: "picture.jpeg",
			},
			want: 1,
		},
		{
			name: "Test picture format png",
			args: args{
				name: "picture.png",
			},
			want: 1,
		},
		{
			name: "Test video format mpeg",
			args: args{
				name: "picture.mpeg",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFormat(tt.args.name); got != tt.want {
				t.Errorf("CheckFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
