package service

import (
	"context"
	"log"

	video "cloud.google.com/go/videointelligence/apiv1"
	vision "cloud.google.com/go/vision/apiv1"
)

func (conf ConfigVision) GetVisionClient() *vision.ImageAnnotatorClient {
	client, err := vision.NewImageAnnotatorClient(conf.Ctx)
	if err != nil {
		log.Fatalf("Failed to create client:  %v", err)
	}
	return client
}

func (conf ConfigVideo) GetVideoClient() *video.Client {
	client, err := video.NewClient(conf.Ctx)
	if err != nil {
		log.Fatalf("Failed to create client:  %v", err)
	}
	return client
}

func CreateConfVideo() ConfigVideo {
	var conf ConfigVideo
	conf.Ctx = context.Background()
	conf.Client = conf.GetVideoClient()
	return conf
}

func CreateConfVision() ConfigVision {
	var conf ConfigVision
	conf.Ctx = context.Background()
	conf.Client = conf.GetVisionClient()
	return conf
}
