package vision_api

import (
	"log"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/alallema/picture_dictionnary.git/core/service"
	config "github.com/alallema/picture_dictionnary.git/vision-client/service"
)

func DetectLabelsFromFile(conf config.ConfigVision) ([]service.LabelData, error) {
	var labels []service.LabelData

	image, err := vision.NewImageFromReader(conf.File)
	if err != nil {
		log.Printf("Failed to read file: %v on file: %s", err, conf.Filename)
		return labels, err
	}
	annotations, err := conf.Client.DetectLabels(conf.Ctx, image, nil, 10)
	if err != nil {
		log.Printf("Failed to detect label: %v on file: %s", err, conf.Filename)
		return labels, err
	}
	if len(annotations) == 0 {
		log.Printf("No annotations found.")
	} else {
		for _, annotation := range annotations {
			label := service.LabelData{
				Id:          annotation.Mid,
				Description: annotation.Description,
				Score:       annotation.Score,
				Confidence:  annotation.Confidence,
				Topicality:  annotation.Topicality,
				Language:    "en",
			}
			labels = append(labels, label)
		}
	}
	// os.File(conf.File).Close()
	// TODO close file
	return labels, err
}

func DetectLabelsFromUri(conf config.ConfigVision) ([]service.LabelData, error) {
	var labels []service.LabelData

	image := vision.NewImageFromURI(conf.GcsURI)
	annotations, err := conf.Client.DetectLabels(conf.Ctx, image, nil, 10)
	if err != nil {
		log.Printf("Failed to detect label: %v on file: %s", err, conf.GcsURI)
		return labels, err
	}
	if len(annotations) == 0 {
		log.Printf("No annotations found.")
	} else {
		for _, annotation := range annotations {
			label := service.LabelData{
				Id:          annotation.Mid,
				Description: annotation.Description,
				Score:       annotation.Score,
				Confidence:  annotation.Confidence,
				Topicality:  annotation.Topicality,
				Language:    "en",
			}
			labels = append(labels, label)
		}
	}
	return labels, err
}

func LocalizeObjectsFromFile(conf config.ConfigVision) ([]service.LocalizedObjectData, error) {
	var objects []service.LocalizedObjectData

	image, err := vision.NewImageFromReader(conf.File)
	if err != nil {
		log.Printf("Failed to read file: %v on file: %s", err, conf.Filename)
		return objects, err
	}
	annotations, err := conf.Client.LocalizeObjects(conf.Ctx, image, nil)
	if err != nil {
		log.Printf("Failed to detect label: %v on file: %s", err, conf.Filename)
		return objects, err
	}

	if len(annotations) == 0 {
		log.Printf("No objects found.")
		return objects, err
	}

	for _, annotation := range annotations {
		object := service.LocalizedObjectData{
			Id:       annotation.Mid,
			Name:     annotation.Name,
			Score:    annotation.Score,
			Language: "en",
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
	// TODO close file
	return objects, err
}

func LocalizeObjectsFromUri(conf config.ConfigVision) ([]service.LocalizedObjectData, error) {
	var objects []service.LocalizedObjectData

	image := vision.NewImageFromURI(conf.GcsURI)
	annotations, err := conf.Client.LocalizeObjects(conf.Ctx, image, nil)
	if err != nil {
		log.Printf("Failed to detect label: %v on file: %s", err, conf.GcsURI)
		return objects, err
	}

	if len(annotations) == 0 {
		log.Printf("No objects found.")
		return objects, err
	}

	for _, annotation := range annotations {
		object := service.LocalizedObjectData{
			Id:       annotation.Mid,
			Name:     annotation.Name,
			Score:    annotation.Score,
			Language: "en",
		}

		for _, v := range annotation.BoundingPoly.NormalizedVertices {
			object.BoundingPoly.NormalizedVertices = append(object.BoundingPoly.NormalizedVertices, &service.NormalizedVertex{
				X: v.X,
				Y: v.Y,
			})
		}
		objects = append(objects, object)
	}
	return objects, err
}
