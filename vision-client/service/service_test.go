package service

import (
	"context"
	"io"
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
		{
			name: "Test no client",
			fields: fields{
				Ctx:      context.Background(),
				Client:   nil,
				Filename: "",
				File:     []byte{},
				GcsURI:   "",
			},
			want: nil,
		},
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
			if got := conf.GetVideoClient(); got == tt.want {
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
		{
			name: "Test no client",
			fields: fields{
				Ctx:      context.Background(),
				Client:   nil,
				Filename: "",
				File:     nil,
				GcsURI:   "",
			},
			want: nil,
		},
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
			if got := conf.GetVisionClient(); got == tt.want {
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
		{
			name: "Test",
			want: ConfigVideo{
				Ctx:    context.Background(),
				Client: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateConfVideo(); got.Client == tt.want.Client {
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
		{
			name: "Test",
			want: ConfigVision{
				Ctx:    context.Background(),
				Client: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateConfVision(); got.Client == tt.want.Client {
				t.Errorf("CreateConfVision() = %v, want %v", got, tt.want)
			}
		})
	}
}
