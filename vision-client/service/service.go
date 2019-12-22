package service

import (
	video "cloud.google.com/go/videointelligence/apiv1"
	vision "cloud.google.com/go/vision/apiv1"
	"log"
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
