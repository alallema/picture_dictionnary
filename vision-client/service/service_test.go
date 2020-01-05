package service

import (
	"context"
	"io"
	"reflect"
	"testing"

	video "cloud.google.com/go/videointelligence/apiv1"
	vision "cloud.google.com/go/vision/apiv1"
)

func TestConfigVideo_GetVideoClient(t *testing.T) {
	type fields struct {
		Ctx      context.Context
		Client   *video.Client
		Filename string
		File     []byte
		GcsURI   string
	}
	tests := []struct {
		name   string
		fields fields
		want   *video.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := ConfigVideo{
				Ctx:      tt.fields.Ctx,
				Client:   tt.fields.Client,
				Filename: tt.fields.Filename,
				File:     tt.fields.File,
				GcsURI:   tt.fields.GcsURI,
			}
			if got := conf.GetVideoClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVideoClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigVision_GetVisionClient(t *testing.T) {
	type fields struct {
		Ctx      context.Context
		Client   *vision.ImageAnnotatorClient
		Filename string
		File     io.Reader
		GcsURI   string
	}
	tests := []struct {
		name   string
		fields fields
		want   *vision.ImageAnnotatorClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := ConfigVision{
				Ctx:      tt.fields.Ctx,
				Client:   tt.fields.Client,
				Filename: tt.fields.Filename,
				File:     tt.fields.File,
				GcsURI:   tt.fields.GcsURI,
			}
			if got := conf.GetVisionClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVisionClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateConfVideo(t *testing.T) {
	tests := []struct {
		name string
		want ConfigVideo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateConfVideo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConfVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateConfVision(t *testing.T) {
	tests := []struct {
		name string
		want ConfigVision
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateConfVision(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConfVision() = %v, want %v", got, tt.want)
			}
		})
	}
}
