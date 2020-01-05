package repo

import (
	"github.com/alallema/picture_dictionnary.git/core/service"
	"reflect"
	"testing"
)

func TestGetPicture(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want service.Picture
	}{
		{
			name: "No picture",
			args: args{
				id: "noexist",
			},
			want: service.Picture{},
		},
		{
			name: "simple picture",
			args: args{
				id: "1111111",
			},
			want: service.Picture{
				Id:          "1111111",
				Title:       "img_001.jpg",
				Format:      "jpg",
				Source:      "instagram:Lala",
				PicturePath: "/path/img_001.jpg",
				PictureURL:  "https://path/img_001.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPicture(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPicture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLabelByPicture(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				id: "noexist",
			},
			want: []string{},
		},
		{
			name: "First picture",
			args: args{
				id: "1111111",
			},
			want: []string{
				"label3",
				"label1",
				"label2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLabelByPicture(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLabelByPicture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGettLocalizedObjectByPicture(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				id: "noexist",
			},
			want: []string{},
		},
		{
			name: "First picture",
			args: args{
				id: "1111111",
			},
			want: []string{
				"object3",
				"object1",
				"object2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GettLocalizedObjectByPicture(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GettLocalizedObjectByPicture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLabelDescription(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Don't exist",
			args: args{
				id: "none",
			},
			want: "Not found",
		},
		{
			name: "label 1",
			args: args{
				id: "label1",
			},
			want: "first test label",
		},
		{
			name: "label 2",
			args: args{
				id: "label2",
			},
			want: "second test label",
		},
		{
			name: "label 3",
			args: args{
				id: "label3",
			},
			want: "third test label",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLabelDescription(tt.args.id); got != tt.want {
				t.Errorf("GetLabelDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetObjectDescription(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty",
			args: args{
				id: "noexist",
			},
			want: "",
		},
		{
			name: "object 1",
			args: args{
				id: "object1",
			},
			want: "first test object",
		},
		{
			name: "object 2",
			args: args{
				id: "object2",
			},
			want: "second test object",
		},
		{
			name: "object 3",
			args: args{
				id: "object3",
			},
			want: "third test object",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetObjectDescription(tt.args.id); got != tt.want {
				t.Errorf("GetObjectDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPictureByLabel(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				id: "noexist",
			},
			want: []string{},
		},
		{
			name: "label 1",
			args: args{
				id: "label1",
			},
			want: []string{
				"1111111",
				"2222222",
				"3333333",
			},
		},
		{
			name: "object 1",
			args: args{
				id: "object1",
			},
			want: []string{
				"1111111",
				"2222222",
				"3333333",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPictureByLabel(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPictureByLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPictureUrlByLabel(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				id: "noexist",
			},
			want: []string{},
		},
		{
			name: "Url 1",
			args: args{
				id: "label1",
			},
			want: []string{
				"https://path/img_003.jpg",
				"https://path/img_001.jpg",
				"https://path/img_002.jpg",
			},
		},
		{
			name: "Url 1",
			args: args{
				id: "object1",
			},
			want: []string{
				"https://path/img_003.jpg",
				"https://path/img_001.jpg",
				"https://path/img_002.jpg",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPictureUrlByLabel(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPictureUrlByLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}
