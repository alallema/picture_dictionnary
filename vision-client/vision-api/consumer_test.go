package vision_api

import (
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/alallema/picture_dictionnary.git/vision-client/service"
	"reflect"
	"testing"
)

func TestDetectLabelsFromFile(t *testing.T) {
	type args struct {
		conf service.ConfigVision
	}
	tests := []struct {
		name    string
		args    args
		want    []core.LabelData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectLabelsFromFile(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectLabelsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectLabelsFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectLabelsFromUri(t *testing.T) {
	type args struct {
		conf service.ConfigVision
	}
	tests := []struct {
		name    string
		args    args
		want    []core.LabelData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectLabelsFromUri(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectLabelsFromUri() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectLabelsFromUri() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizeObjectsFromFile(t *testing.T) {
	type args struct {
		conf service.ConfigVision
	}
	tests := []struct {
		name    string
		args    args
		want    []core.LocalizedObjectData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LocalizeObjectsFromFile(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalizeObjectsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalizeObjectsFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizeObjectsFromUri(t *testing.T) {
	type args struct {
		conf service.ConfigVision
	}
	tests := []struct {
		name    string
		args    args
		want    []core.LocalizedObjectData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LocalizeObjectsFromUri(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalizeObjectsFromUri() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalizeObjectsFromUri() got = %v, want %v", got, tt.want)
			}
		})
	}
}
