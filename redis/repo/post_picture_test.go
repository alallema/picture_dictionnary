package repo

import (
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"testing"
)

func TestPostPicture(t *testing.T) {
	type args struct {
		picture core.Picture
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "simple picture",
			args: args{
				core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostPicture(tt.args.picture)
		})
	}
}

func TestPostLabelByPicture(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LabelData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "First picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LabelData{
					Id:          "label1",
					Language:    "en",
					Description: "first test label",
					Score:       0.55,
					Confidence:  0.55,
					Topicality:  0.55,
				},
			},
		},
		{
			name: "Same picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LabelData{
					Id:          "label2",
					Language:    "en",
					Description: "second test label",
					Score:       0.77,
					Confidence:  0.77,
					Topicality:  0.77,
				},
			},
		},
		{
			name: "Same same picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LabelData{
					Id:          "label3",
					Language:    "en",
					Description: "third test label",
					Score:       0.44,
					Confidence:  0.44,
					Topicality:  0.44,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostLabelByPicture(tt.args.picture, tt.args.data)
		})
	}
}

func TestPostLocalizedObjectByPicture(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LocalizedObjectData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "First picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
			},
		},
		{
			name: "Same picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object2",
					Language: "en",
					Name:     "second test object",
					Score:    0.77,
				},
			},
		},
		{
			name: "Same same picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object3",
					Language: "en",
					Name:     "third test object",
					Score:    0.44,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostLocalizedObjectByPicture(tt.args.picture, tt.args.data)
		})
	}
}

func TestPostCreateLabel(t *testing.T) {
	type args struct {
		data core.LabelData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "label 1",
			args: args{
				data: core.LabelData{
					Id:          "label1",
					Language:    "en",
					Description: "first test label",
					Score:       0.55,
					Confidence:  0.55,
					Topicality:  0.55,
				},
			},
		},
		{
			name: "label 2",
			args: args{
				data: core.LabelData{
					Id:          "label2",
					Language:    "en",
					Description: "second test label",
					Score:       0.77,
					Confidence:  0.77,
					Topicality:  0.77,
				},
			},
		},
		{
			name: "label 3",
			args: args{
				data: core.LabelData{
					Id:          "label3",
					Language:    "en",
					Description: "third test label",
					Score:       0.44,
					Confidence:  0.44,
					Topicality:  0.44,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostCreateLabel(tt.args.data)
		})
	}
}

func TestPostCreateObject(t *testing.T) {
	type args struct {
		data core.LocalizedObjectData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "object 1",
			args: args{
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
			},
		},
		{
			name: "object 2",
			args: args{
				data: core.LocalizedObjectData{
					Id:       "object2",
					Language: "en",
					Name:     "second test object",
					Score:    0.77,
				},
			},
		},
		{
			name: "object 3",
			args: args{
				data: core.LocalizedObjectData{
					Id:       "object3",
					Language: "en",
					Name:     "third test object",
					Score:    0.44,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostCreateObject(tt.args.data)
		})
	}
}

func TestPostPicturebyLabel(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LabelData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "First picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LabelData{
					Id:          "label1",
					Language:    "en",
					Description: "first test label",
					Score:       0.55,
					Confidence:  0.55,
					Topicality:  0.55,
				},
			},
		},
		{
			name: "Second picture",
			args: args{
				picture: core.Picture{
					Id:          "2222222",
					Title:       "img_002.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_002.jpg",
					PictureURL:  "https://path/img_002.jpg",
				},
				data: core.LabelData{
					Id:          "label1",
					Language:    "en",
					Description: "first test label",
					Score:       0.55,
					Confidence:  0.55,
					Topicality:  0.55,
				},
			},
		},
		{
			name: "Third picture",
			args: args{
				picture: core.Picture{
					Id:          "3333333",
					Title:       "img_003.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_003.jpg",
					PictureURL:  "https://path/img_003.jpg",
				},
				data: core.LabelData{
					Id:          "label1",
					Language:    "en",
					Description: "first test label",
					Score:       0.55,
					Confidence:  0.55,
					Topicality:  0.55,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostPicturebyLabel(tt.args.picture, tt.args.data)
		})
	}
}

func TestPostPicturebyObject(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LocalizedObjectData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "First picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
			},
		},
		{
			name: "Second picture",
			args: args{
				picture: core.Picture{
					Id:          "2222222",
					Title:       "img_002.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_002.jpg",
					PictureURL:  "https://path/img_002.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
			},
		},
		{
			name: "Third picture",
			args: args{
				picture: core.Picture{
					Id:          "3333333",
					Title:       "img_003.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_003.jpg",
					PictureURL:  "https://path/img_003.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostPicturebyObject(tt.args.picture, tt.args.data)
		})
	}
}

func TestPostImageUrlbyLabel(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LabelData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "First picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LabelData{
					Id:          "label1",
					Language:    "en",
					Description: "first test label",
					Score:       0.55,
					Confidence:  0.55,
					Topicality:  0.55,
				},
			},
		},
		{
			name: "Second picture",
			args: args{
				picture: core.Picture{
					Id:          "2222222",
					Title:       "img_002.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_002.jpg",
					PictureURL:  "https://path/img_002.jpg",
				},
				data: core.LabelData{
					Id:          "label1",
					Language:    "en",
					Description: "first test label",
					Score:       0.55,
					Confidence:  0.55,
					Topicality:  0.55,
				},
			},
		},
		{
			name: "Third picture",
			args: args{
				picture: core.Picture{
					Id:          "3333333",
					Title:       "img_003.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_003.jpg",
					PictureURL:  "https://path/img_003.jpg",
				},
				data: core.LabelData{
					Id:          "label1",
					Language:    "en",
					Description: "first test label",
					Score:       0.55,
					Confidence:  0.55,
					Topicality:  0.55,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostImageUrlbyLabel(tt.args.picture, tt.args.data)
		})
	}
}

func TestPostImageUrlbyObject(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LocalizedObjectData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "First picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
			},
		},
		{
			name: "Second picture",
			args: args{
				picture: core.Picture{
					Id:          "2222222",
					Title:       "img_002.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_002.jpg",
					PictureURL:  "https://path/img_002.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
			},
		},
		{
			name: "Third picture",
			args: args{
				picture: core.Picture{
					Id:          "3333333",
					Title:       "img_003.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_003.jpg",
					PictureURL:  "https://path/img_003.jpg",
				},
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostImageUrlbyObject(tt.args.picture, tt.args.data)
		})
	}
}
