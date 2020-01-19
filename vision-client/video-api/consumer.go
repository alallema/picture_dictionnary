package video_api

import (
	"github.com/alallema/picture_dictionnary.git/core/service"
	config "github.com/alallema/picture_dictionnary.git/vision-client/service"
	videopb "google.golang.org/genproto/googleapis/cloud/videointelligence/v1"
	"log"
)

func DetectLabelVideo(conf config.ConfigVideo) ([]service.VideoLabelData, error) {
	var videoLabels []service.VideoLabelData

	op, err := conf.Client.AnnotateVideo(conf.Ctx, &videopb.AnnotateVideoRequest{
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
		// InputContent: conf.File,
		InputUri: conf.GcsURI,
	})
	if err != nil {
		log.Printf("Failed to start annotation job: %v on: %s", err, conf.GcsURI)
		return videoLabels, err
	}
	resp, err := op.Wait(conf.Ctx)
	if err != nil {
		log.Printf("Failed to annotate: %v on: %s", err, conf.GcsURI)
		return videoLabels, err
	}

	printLabels := func(labels []*videopb.LabelAnnotation) {
		for _, label := range labels {

			videoLabel := service.VideoLabelData{
				Entity: service.Entity{
					Id:          label.Entity.EntityId,
					Description: label.Entity.Description,
					Language:    label.Entity.LanguageCode,
				},
			}
			for _, category := range label.CategoryEntities {
				categoryEntity := service.Entity{
					Id:          category.EntityId,
					Description: category.Description,
					Language:    category.LanguageCode,
				}
				videoLabel.CategoryEntities = append(videoLabel.CategoryEntities, categoryEntity)
			}
			for _, segment := range label.Segments {
				Segment := service.Segment{
					StartTimeOffset: segment.Segment.StartTimeOffset,
					EndTimeOffset:   segment.Segment.EndTimeOffset,
				}
				labelSegment := service.LabelSegment{
					Segment:    Segment,
					Confidence: segment.Confidence,
				}
				videoLabel.Segments = append(videoLabel.Segments, labelSegment)
			}
			videoLabels = append(videoLabels, videoLabel)
		}
	}

	// A single video was processed. Get the first result.
	result := resp.AnnotationResults[0]

	// printLabels(result.SegmentLabelAnnotations)
	printLabels(result.ShotLabelAnnotations)
	// printLabels(result.FrameLabelAnnotations)

	return videoLabels, nil
}

func ShotChange(conf config.ConfigVideo) ([]service.Segment, error) {
	var shots []service.Segment

	op, err := conf.Client.AnnotateVideo(conf.Ctx, &videopb.AnnotateVideoRequest{
		Features: []videopb.Feature{
			videopb.Feature_SHOT_CHANGE_DETECTION,
		},
		// InputContent: conf.File,
		InputUri: conf.GcsURI,
	})
	if err != nil {
		log.Printf("Failed to start annotation job: %v on: %s", err, conf.GcsURI)
		return shots, err
	}
	resp, err := op.Wait(conf.Ctx)
	if err != nil {
		log.Printf("Failed to annotate: %v on: %s", err, conf.GcsURI)
		return shots, err
	}

	// A single video was processed. Get the first result.
	result := resp.AnnotationResults[0].ShotAnnotations

	for _, shot := range result {

		shots = append(shots, service.Segment{
			StartTimeOffset: shot.StartTimeOffset,
			EndTimeOffset:   shot.EndTimeOffset,
		})
	}

	return shots, nil
}
