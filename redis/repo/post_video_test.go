package repo

import (
	"testing"

	core "github.com/alallema/picture_dictionnary.git/core/service"
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
			got := GetPicture(tt.args.id)
			if got != tt.want {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
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
						// Segment: core.Segment{
						// 	StartTimeOffset: ptypes.Duration(time.Duration(rand.Int31n(1000))),
						// 	EndTimeOffset:   ptypes.Duration(time.Duration(rand.Int31n(1000))),
						// },
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
		// {
		// 	name: "label 2",
		// 	args: args{
		// 		data: core.VideoLabelData{
		// 			Id:          "label2",
		// 			Language:    "en",
		// 			Description: "second test label",
		// 			Score:       0.77,
		// 			Confidence:  0.77,
		// 			Topicality:  0.77,
		// 		},
		// 		id: "label2",
		// 	},
		// 	want:    "second test label",
		// 	wantErr: nil,
		// },
		// {
		// 	name: "label 3",
		// 	args: args{
		// 		data: core.VideoLabelData{
		// 			Id:          "label3",
		// 			Language:    "en",
		// 			Description: "third test label",
		// 			Score:       0.44,
		// 			Confidence:  0.44,
		// 			Topicality:  0.44,
		// 		},
		// 		id: "label3",
		// 	},
		// 	want:    "third test label",
		// 	wantErr: nil,
		// },
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostCreateVideoCategory(t *testing.T) {
	type args struct {
		categories []core.Entity
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostLabelbyCategory(t *testing.T) {
	type args struct {
		label      core.VideoLabelData
		categories []core.Entity
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostSegbyVideobyLabel(t *testing.T) {
	type args struct {
		segments []core.LabelSegment
		data     core.VideoLabelData
		video    core.Picture
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostVideobyLabel(t *testing.T) {
	type args struct {
		video core.Picture
		data  core.VideoLabelData
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
