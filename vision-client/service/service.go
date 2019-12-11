package service

import (
	"log"
	vision "cloud.google.com/go/vision/apiv1"
)

func (conf Config) GetClient() *vision.ImageAnnotatorClient {
	client, err := vision.NewImageAnnotatorClient(conf.Ctx)
	if err != nil {
		log.Fatalf("Failed to create client:  %v", err)
	}
	return client
}