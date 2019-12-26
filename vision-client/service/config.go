package service

import (
	"context"
	"io"

	video "cloud.google.com/go/videointelligence/apiv1"
	vision "cloud.google.com/go/vision/apiv1"
)

type ConfigVision struct {
	Ctx      context.Context
	Client   *vision.ImageAnnotatorClient
	Filename string
	File     io.Reader
	GcsURI   string
}

type ConfigVideo struct {
	Ctx      context.Context
	Client   *video.Client
	Filename string
	File     []byte
	GcsURI   string
}
