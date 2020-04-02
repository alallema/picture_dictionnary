package repo

import (
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPostPicture(t *testing.T) {
	type args struct {
		picture core.Picture
		id      string
	}
	tests := []struct {
		name string
		args args
		want core.Picture
	}{
		{
			name: "No picture",
			args: args{
				picture: core.Picture{},
				id:      "noexist",
			},
			want: core.Picture{},
		},
		{
			name: "simple picture",
			args: args{
				picture: core.Picture{
					Id:          "1111111",
					Title:       "img_001.jpg",
					Format:      "jpg",
					Source:      "instagram:Lala",
					PicturePath: "/path/img_001.jpg",
					PictureURL:  "https://path/img_001.jpg",
				},
				id: "1111111",
			},
			want: core.Picture{
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
			PostPicture(tt.args.picture)
			got := GetPicture(tt.args.id)
			if got != tt.want {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		DelHSet("picture:" + tt.args.picture.Id)
	}
}

func TestPostLabelByPicture(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LabelData
		id      string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{
			name: "Empty",
			args: args{
				picture: core.Picture{},
				data:    core.LabelData{},
				id:      "noexist",
			},
			want:    []string{},
			wantErr: nil,
		},
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
				id: "1111111",
			},
			want: []string{
				"label1",
			},
			wantErr: nil,
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
				id: "1111111",
			},
			want: []string{
				"label1",
				"label2",
			},
			wantErr: nil,
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
				id: "1111111",
			},
			want: []string{
				"label3",
				"label1",
				"label2",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostLabelByPicture(tt.args.picture, tt.args.data)
			got := GetLabelByPicture(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	DelAll("pictureIdLabel:")
	for _, tt := range tests {
		DelAll("pictureIdLabel:" + tt.args.id)
	}
}

func TestPostLocalizedObjectByPicture(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LocalizedObjectData
		id      string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{
			name: "Empty",
			args: args{
				picture: core.Picture{},
				data:    core.LocalizedObjectData{},
				id:      "noexist",
			},
			want:    []string{},
			wantErr: nil,
		},
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
				id: "1111111",
			},
			want: []string{
				"object1",
			},
			wantErr: nil,
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
				id: "1111111",
			},
			want: []string{
				"object1",
				"object2",
			},
			wantErr: nil,
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
				id: "1111111",
			},
			want: []string{
				"object3",
				"object1",
				"object2",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostLocalizedObjectByPicture(tt.args.picture, tt.args.data)
			got := GettLocalizedObjectByPicture(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	DelAll("pictureIdObject:")
	for _, tt := range tests {
		DelAll("pictureIdObject:" + tt.args.id)
	}
}

func TestPostCreateLabel(t *testing.T) {
	type args struct {
		data core.LabelData
		id   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Don't exist",
			args: args{
				data: core.LabelData{},
				id:   "none",
			},
			want:    "Not found",
			wantErr: nil,
		},
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
				id: "label1",
			},
			want:    "first test label",
			wantErr: nil,
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
				id: "label2",
			},
			want:    "second test label",
			wantErr: nil,
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
				id: "label3",
			},
			want:    "third test label",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostCreateLabel(tt.args.data)
			got := GetLabelDescription(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	DelHSet("labelDescr:")
	DelMembersInSet("labelId", "")
	for _, tt := range tests {
		DelHSet("labelDescr:" + tt.args.id)
		DelMembersInSet("labelId", tt.args.id)
	}
}

func TestPostCreateObject(t *testing.T) {
	type args struct {
		data core.LocalizedObjectData
		id   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Empty",
			args: args{
				data: core.LocalizedObjectData{},
				id:   "noexist",
			},
			want:    "",
			wantErr: nil,
		},
		{
			name: "object 1",
			args: args{
				data: core.LocalizedObjectData{
					Id:       "object1",
					Language: "en",
					Name:     "first test object",
					Score:    0.55,
				},
				id: "object1",
			},
			want:    "first test object",
			wantErr: nil,
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
				id: "object2",
			},
			want:    "second test object",
			wantErr: nil,
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
				id: "object3",
			},
			want:    "third test object",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostCreateObject(tt.args.data)
			got := GetObjectDescription(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	DelHSet("objectDescr:")
	DelMembersInSet("objectId", "")
	for _, tt := range tests {
		DelHSet("objectDescr:" + tt.args.id)
		DelMembersInSet("objectId", tt.args.id)
	}
}

func TestPostPicturebyLabel(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LabelData
		id      string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{
			name: "Empty",
			args: args{
				picture: core.Picture{},
				data:    core.LabelData{},
				id:      "noexist",
			},
			want:    []string{},
			wantErr: nil,
		},
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
				id: "label1",
			},
			want: []string{
				"1111111",
			},
			wantErr: nil,
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
				id: "label1",
			},
			want: []string{
				"1111111",
				"2222222",
			},
			wantErr: nil,
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
				id: "label1",
			},
			want: []string{
				"1111111",
				"2222222",
				"3333333",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostPicturebyLabel(tt.args.picture, tt.args.data)
			got := GetPictureByLabel(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	DelAll("Id:")
	for _, tt := range tests {
		DelAll("Id:" + tt.args.id)
	}
}

func TestPostPicturebyObject(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LocalizedObjectData
		id      string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{
			name: "Empty",
			args: args{
				picture: core.Picture{},
				data:    core.LocalizedObjectData{},
				id:      "noexist",
			},
			want:    []string{},
			wantErr: nil,
		},
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
				id: "object1",
			},
			want: []string{
				"1111111",
			},
			wantErr: nil,
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
				id: "object1",
			},
			want: []string{
				"1111111",
				"2222222",
			},
			wantErr: nil,
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
				id: "object1",
			},
			want: []string{
				"1111111",
				"2222222",
				"3333333",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostPicturebyObject(tt.args.picture, tt.args.data)
			got := GetPictureByLabel(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	DelAll("Id:")
	for _, tt := range tests {
		DelAll("Id:" + tt.args.id)
	}
}

func TestPostImageUrlbyLabel(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LabelData
		id      string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
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
				id: "label1",
			},
			want: []string{
				"https://path/img_001.jpg",
			},
			wantErr: nil,
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
				id: "label1",
			},
			want: []string{
				"https://path/img_001.jpg",
				"https://path/img_002.jpg",
			},
			wantErr: nil,
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
				id: "label1",
			},
			want: []string{
				"https://path/img_003.jpg",
				"https://path/img_001.jpg",
				"https://path/img_002.jpg",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostImageUrlbyLabel(tt.args.picture, tt.args.data)
			got := GetPictureUrlByLabel(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		DelAll("URLId:" + tt.args.id)
	}
}

func TestPostImageUrlbyObject(t *testing.T) {
	type args struct {
		picture core.Picture
		data    core.LocalizedObjectData
		id      string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
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
				id: "object1",
			},
			want: []string{
				"https://path/img_001.jpg",
			},
			wantErr: nil,
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
				id: "object1",
			},
			want: []string{
				"https://path/img_001.jpg",
				"https://path/img_002.jpg",
			},
			wantErr: nil,
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
				id: "object1",
			},
			want: []string{
				"https://path/img_003.jpg",
				"https://path/img_001.jpg",
				"https://path/img_002.jpg",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostImageUrlbyObject(tt.args.picture, tt.args.data)
			got := GetPictureUrlByLabel(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		DelAll("URLId:" + tt.args.id)
	}
}
