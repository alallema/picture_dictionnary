package repo

import (
	"testing"

	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/google/go-cmp/cmp"
)

func TestPostVideo(t *testing.T) {
	type args struct {
		video core.Picture
		id    string
	}
	tests := []struct {
		name string
		args args
		want core.Picture
	}{
		{
			name: "No video",
			args: args{
				video: core.Picture{},
				id:    "noexist",
			},
			want: core.Picture{},
		},
		{
			name: "simple video",
			args: args{
				video: core.Picture{
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
			PostVideo(tt.args.video)
			got := GetVideo(tt.args.id)
			if got != tt.want {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		DelHSet("video:" + tt.args.video.Id)
	}
}

func TestPostCreateVideoLabel(t *testing.T) {
	type args struct {
		data core.VideoLabelData
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
				data: core.VideoLabelData{},
				id:   "none",
			},
			want:    "Not found",
			wantErr: nil,
		},
		{
			name: "label 1",
			args: args{
				data: core.VideoLabelData{
					Entity: core.Entity{
						Id:          "label1",
						Language:    "en",
						Description: "first test label",
					},
					CategoryEntities: []core.Entity{
						core.Entity{
							Id:          "category1",
							Language:    "en",
							Description: "first category",
						},
					},
					Segments: []core.LabelSegment{
						core.LabelSegment{
							Confidence: 0.55,
						},
					},
				},
				id: "label1",
			},
			want:    "first test label",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostCreateVideoLabel(tt.args.data)
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

func TestPostCreateVideoCategory(t *testing.T) {
	type args struct {
		categories []core.Entity
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "test first category",
			args: args{
				categories: []core.Entity{
					core.Entity{
						Id:          "category1",
						Language:    "en",
						Description: "first category",
					},
				},
			},
			want:    "first category",
			wantErr: nil,
		},
		{
			name: "test two categories",
			args: args{
				categories: []core.Entity{
					core.Entity{
						Id:          "category2",
						Language:    "en",
						Description: "second category",
					},
				},
			},
			want:    "second category",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostCreateVideoCategory(tt.args.categories)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := GetCategoryDescription(tt.args.categories[0].Id)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		DelHSet("categoryDescr:" + tt.args.categories[0].Id)
		DelMembersInSet("categoryId", tt.args.categories[0].Id)
	}
}

func TestPostLabelbyCategory(t *testing.T) {
	type args struct {
		label      core.VideoLabelData
		categories []core.Entity
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "label 1",
			args: args{
				label: core.VideoLabelData{
					Entity: core.Entity{
						Id:          "label1",
						Language:    "en",
						Description: "first test label",
					},
					CategoryEntities: []core.Entity{
						core.Entity{
							Id:          "category1",
							Language:    "en",
							Description: "first category",
						},
					},
					Segments: []core.LabelSegment{
						core.LabelSegment{
							Confidence: 0.55,
						},
					},
				},
				categories: []core.Entity{
					core.Entity{
						Id:          "category1",
						Language:    "en",
						Description: "first category",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostLabelbyCategory(tt.args.label, tt.args.categories)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	for _, tt := range tests {
		DelAll("categoryId:" + tt.args.categories[0].Id)
	}
}

func TestPostSegbyVideobyLabel(t *testing.T) {
	type args struct {
		segments []core.LabelSegment
		data     core.VideoLabelData
		video    core.Picture
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "simple video",
			args: args{
				segments: []core.LabelSegment{
					core.LabelSegment{
						Confidence: 0.55,
					},
				},
				data: core.VideoLabelData{
					Entity: core.Entity{
						Id: "label1",
					},
					CategoryEntities: []core.Entity{
						core.Entity{
							Id:          "category1",
							Language:    "en",
							Description: "first category",
						},
					},
					Segments: []core.LabelSegment{
						core.LabelSegment{
							Confidence: 0.55,
						},
					},
				},
				video: core.Picture{
					Id: "1111111",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostSegbyVideobyLabel(tt.args.segments, tt.args.data, tt.args.video)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	for _, tt := range tests {
		DelAll("seg:" + tt.args.video.Id + ":" + tt.args.data.Entity.Id)
	}
}

func TestPostVideobyLabel(t *testing.T) {
	type args struct {
		video core.Picture
		data  core.VideoLabelData
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{
			name: "simple video",
			args: args{
				video: core.Picture{
					Id: "1111111",
				},
				data: core.VideoLabelData{
					Entity: core.Entity{
						Id:          "label1",
						Language:    "en",
						Description: "first test label",
					},
					CategoryEntities: []core.Entity{
						core.Entity{
							Id:          "category1",
							Language:    "en",
							Description: "first category",
						},
					},
					Segments: []core.LabelSegment{
						core.LabelSegment{
							Confidence: 0.55,
						},
					},
				},
			},
			want: []string{
				"1111111",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostVideobyLabel(tt.args.video, tt.args.data)
			got := GetVideoByLabel(tt.args.data.Entity.Id)
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
		DelAll("VidId:" + tt.args.data.Entity.Id)
	}
}
