package repo

import (
	"fmt"
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes"
	"log"
)

func PostVideo(video core.Picture) {
	if Client.HSet("video:"+string(video.Id), "id", string(video.Id)).Val() != true {
		log.Printf("Failed to push PictureId: %s", string(video.Id))
		return
	}
	if Client.HSet("video:"+string(video.Id), "title", video.Title).Val() != true {
		log.Printf("Failed to push PictureTitle: %s", string(video.Title))
	}
	if Client.HSet("video:"+string(video.Id), "format", video.Format).Val() != true {
		log.Printf("Failed to push PictureFormat: %s", string(video.Format))
	}
	if Client.HSet("video:"+string(video.Id), "videoPath", video.PicturePath).Val() != true {
		log.Printf("Failed to push PicturePath: %s", string(video.PicturePath))
	}
	if Client.HSet("video:"+string(video.Id), "videoURL", video.PictureURL).Val() != true {
		log.Printf("Failed to push PictureURL: %s", string(video.PictureURL))
	}
	if Client.HSet("video:"+string(video.Id), "source", video.Source).Val() != true {
		log.Printf("Failed to push Source: %s", string(video.Source))
	}
	return
}

func PostCreateVideoLabel(data core.VideoLabelData) error {
	var err error

	if Client.HExists("labelDescr:"+data.Entity.Id, data.Entity.Language).Val() == false {
		if Client.HSet("labelDescr:"+data.Entity.Id, data.Entity.Language, data.Entity.Description).Val() != true {
			log.Printf("Failed to push")
		}
		err = Client.SAdd("labelId", data.Entity.Id).Err()
		if err == redis.Nil {
			log.Printf("Failed to push labelId: %s", string(data.Entity.Id))
		} else if err != nil {
			log.Printf("Error: %v", err)
		}
	}
	return err
}

func PostCreateVideoCategory(categories []core.Entity) error {
	var err error

	for _, category := range categories {
		if Client.HExists("categoryDescr:"+category.Id, category.Language).Val() == false {
			if Client.HSet("categoryDescr:"+category.Id, category.Language, category.Description).Val() != true {
				log.Printf("Failed to push")
			}
			err = Client.SAdd("categoryId", category.Id).Err()
			if err == redis.Nil {
				log.Printf("Failed to push labelId: %s", string(category.Id))
			} else if err != nil {
				log.Printf("Error: %v", err)
			}
		}
	}
	return err
}

func PostVideobyLabel(video core.Picture, data core.VideoLabelData) error {
	err := Client.SAdd("VidId:"+string(data.Entity.Id), video.Id).Err()
	if err == redis.Nil {
		log.Printf("Failed to push label: %s", string(data.Entity.Id))
	} else if err != nil {
		log.Printf("Error: %v", err)
	}
	return err
}

func PostSegbyVideobyLabel(segments []core.LabelSegment, data core.VideoLabelData, video core.Picture) error {
	var err error

	for _, segment := range segments {
		start, _ := ptypes.Duration(segment.Segment.StartTimeOffset)
		end, _ := ptypes.Duration(segment.Segment.EndTimeOffset)
		duration := fmt.Sprintf("%s-%s", start, end)
		err = Client.ZAdd("seg:"+string(video.Id)+":"+string(data.Entity.Id), redis.Z{
			Score:  float64(segment.Confidence),
			Member: duration,
		}).Err()
		if err == redis.Nil {
			log.Printf("Failed to push Segment: %s of %s", duration, string(video.Id))
		} else if err != nil {
			log.Printf("Error: %v", err)
		}
	}
	return err
}

func PostLabelbyCategory(label core.VideoLabelData, categories []core.Entity) error {
	var err error

	for _, category := range categories {
		err = Client.SAdd("categoryId:"+string(category.Id), label.Entity.Id).Err()
		if err == redis.Nil {
			log.Printf("Failed to push label in category: %s", string(label.Entity.Id))
		} else if err != nil {
			log.Printf("Error: %v", err)
		}
	}
	return err
}
