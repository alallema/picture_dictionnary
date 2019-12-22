package video_api

import (
	"reflect"
	"testing"

	"github.com/alallema/picture_dictionnary.git/core/service"
	config "github.com/alallema/picture_dictionnary.git/vision-client/service"
)

func TestDetectLabelVideo(t *testing.T) {
	type args struct {
		conf config.ConfigVideo
	}
	tests := []struct {
		name  string
		args  args
		want  []service.VideoLabelData
		want1 error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DetectLabelVideo(tt.args.conf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectLabelVideo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DetectLabelVideo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShotChange(t *testing.T) {
	type args struct {
		conf config.ConfigVideo
	}
	tests := []struct {
		name  string
		args  args
		want  []service.Segment
		want1 error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ShotChange(tt.args.conf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShotChange() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShotChange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
