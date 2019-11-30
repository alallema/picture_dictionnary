package service

import (
	"fmt"
	"log"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/alallema/picture_dictionnary.git/core/service"
)

func (conf Config) GetClient() *vision.ImageAnnotatorClient {
	client, err := vision.NewImageAnnotatorClient(conf.Ctx)
	if err != nil {
		log.Fatalf("Failed to create client:  %v", err)
	}
	return client
}

func DetectLabels(conf Config) ([]service.LabelData, error) {
	var labels []service.LabelData

	image, err := vision.NewImageFromReader(conf.File)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		return labels, err
	}
	annotations, err := conf.Client.DetectLabels(conf.Ctx, image, nil, 10)
	if err != nil {
		return labels, err
	}
	if len(annotations) == 0 {
		fmt.Println("No annotations found.")
	} else {
		for _, annotation := range annotations {
			label := service.LabelData{
				Mid:         annotation.Mid,
				Description: annotation.Description,
				Score:       annotation.Score,
				Confidence:  annotation.Confidence,
				Topicality:  annotation.Topicality,
				Locale:      annotation.Locale,
			}
			labels = append(labels, label)
		}
	}
	// os.File(conf.File).Close()
	return labels, err
}

func LocalizeObjects(conf Config) ([]service.LocalizedObjectData, error) {
	var objects []service.LocalizedObjectData

	image, err := vision.NewImageFromReader(conf.File)
	if err != nil {
		return objects, err
	}
	annotations, err := conf.Client.LocalizeObjects(conf.Ctx, image, nil)
	if err != nil {
		return objects, err
	}

	if len(annotations) == 0 {
		fmt.Println("No objects found.")
		return objects, err
	}

	for _, annotation := range annotations {
		object := service.LocalizedObjectData{
			Mid:   annotation.Mid,
			Name:  annotation.Name,
			Score: annotation.Score,
		}

		for _, v := range annotation.BoundingPoly.NormalizedVertices {
			object.BoundingPoly.NormalizedVertices = append(object.BoundingPoly.NormalizedVertices, &service.NormalizedVertex{
				X: v.X,
				Y: v.Y,
			})
		}
		objects = append(objects, object)
	}

	// os.File(conf.File).Close()
	return objects, err
}
