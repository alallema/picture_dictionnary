package video_api

import (
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/alallema/picture_dictionnary.git/vision-client/service"
	"reflect"
	"testing"
)

func TestDetectLabelVideo(t *testing.T) {
	type args struct {
		conf service.ConfigVideo
	}
	tests := []struct {
		name    string
		args    args
		want    []core.VideoLabelData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectLabelVideo(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectLabelVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectLabelVideo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShotChange(t *testing.T) {
	type args struct {
		conf service.ConfigVideo
	}
	tests := []struct {
		name    string
		args    args
		want    []core.Segment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShotChange(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShotChange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShotChange() got = %v, want %v", got, tt.want)
			}
		})
	}
}
