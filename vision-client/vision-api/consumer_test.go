package vision_api

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/alallema/picture_dictionnary.git/core/service"
	config "github.com/alallema/picture_dictionnary.git/vision-client/service"
)

func TestGetClient(t *testing.T) {

	var conf config.ConfigVision
	conf.Ctx = context.Background()
	conf.Client = conf.GetVisionClient()

	if conf.Client != nil {
		fmt.Printf("client\n")
	} else {
		fmt.Printf("no client\n")
	}
}

func TestDetectLabels(t *testing.T) {
	type args struct {
		conf config.ConfigVision
	}
	tests := []struct {
		name    string
		args    args
		want    []service.LabelData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectLabelsFromFile(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectLabels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectLabels() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizeObjects(t *testing.T) {
	type args struct {
		conf config.ConfigVision
	}
	tests := []struct {
		name    string
		args    args
		want    []service.LocalizedObjectData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LocalizeObjects(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalizeObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalizeObjects() got = %v, want %v", got, tt.want)
			}
		})
	}
}
