package repo

import (
	"fmt"
	"log"

	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/go-redis/redis"
)

func PostPicture(picture core.Picture) error {
	var err error

	if Client.HSet("picture:"+string(picture.Id), "id", string(picture.Id)).Val() != true {
		log.Fatalf("Failed to push PictureId:  %v", err)
	}
	if Client.HSet("picture:"+string(picture.Id), "title", picture.Title).Val() != true {
		log.Fatalf("Failed to push PictureTitle:  %v", err)
	}
	if Client.HSet("picture:"+string(picture.Id), "format", picture.Format).Val() != true {
		log.Fatalf("Failed to push PictureFormat:  %v", err)
	}
	if Client.HSet("picture:"+string(picture.Id), "pictureURL", picture.PictureURL).Val() != true {
		log.Fatalf("Failed to push PictureUrl:  %v", err)
	}
	return err
}

func PostLabelByPicture(picture core.Picture, data core.LabelData) error {

	vals, err := Client.ZAdd("pictureIdLabel:"+string(picture.Id), redis.Z{
		Score:  float64(data.Score),
		Member: string(data.Id),
	}).Result()
	fmt.Println(vals)

	return err
}

func PostLocalizedObjectByPicture(picture core.Picture, data core.LocalizedObjectData) error {

	vals, err := Client.ZAdd("pictureIdObject:"+string(picture.Id), redis.Z{
		Score:  float64(data.Score),
		Member: string(data.Id),
	}).Result()
	fmt.Println(vals)

	return err
}

func PostCreateLabel(data core.LabelData) error {
	var err error

	if Client.HSet("labelDescr:"+data.Id, data.Language, data.Description).Val() != true {
		log.Fatalf("Failed to push labelDescr:  %v", err)
	}
	if Client.Set("labelId", data.Id, 0).Err() != nil {
		log.Fatalf("Failed to push labelId:  %v", err)
	}

	return err
}

func PostCreateObject(data core.LabelData) error {
	var err error

	if Client.HSet("objectDescr:"+data.Id, data.Language, data.Description).Val() != true {
		log.Fatalf("Failed to push objectDescr:  %v", err)
	}
	if Client.Set("objectId", data.Id, 0).Err() != nil {
		log.Fatalf("Failed to push objectId:  %v", err)
	}

	return err
}

func PostPicturebyLabel(picture core.Picture, data core.LabelData) {
	err := Client.Set("Id:"+string(data.Id), picture.Id, 0).Err()
	if err != nil {
		log.Fatalf("Failed to push objectId:  %v", err)
		panic(err)
	}
}
