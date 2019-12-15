package repo

import (
	"fmt"
	// "log"
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/go-redis/redis"
)

func PostPicture(picture core.Picture) error {
	
	if Client.HSet("picture:"+string(picture.Id), "id", string(picture.Id)).Val() != true {
		// log.Printf("Failed to push PictureId:  %v", err)
		return fmt.Errorf("Failed to push PictureId: %s", string(picture.Id))
	}
	if Client.HSet("picture:"+string(picture.Id), "title", picture.Title).Val() != true {
		// log.Printf("Failed to push PictureTitle:  %v", err)
		return fmt.Errorf("Failed to push PictureTitle: %s", string(picture.Title))
	}
	if Client.HSet("picture:"+string(picture.Id), "format", picture.Format).Val() != true {
		// log.Printf("Failed to push PictureFormat:  %v", err)
		return fmt.Errorf("Failed to push PictureFormat: %s", string(picture.Format))
	}
	if Client.HSet("picture:"+string(picture.Id), "picturePath", picture.PicturePath).Val() != true {
		// log.Printf("Failed to push PicturePath:  %v", err)
		return fmt.Errorf("Failed to push PicturePath: %s", string(picture.PicturePath))
	}
	if Client.HSet("picture:"+string(picture.Id), "pictureURL", picture.PictureURL).Val() != true {
		// log.Printf("Failed to push PictureUrl:  %v", err)
		return fmt.Errorf("Failed to push PictureURL: %s", string(picture.PictureURL))
	}
	if Client.HSet("picture:"+string(picture.Id), "source", picture.Source).Val() != true {
		// log.Printf("Failed to push PictureSource:  %v", err)
		return fmt.Errorf("Failed to push Source: %s", string(picture.Source))
	}
	return nil
}

func PostLabelByPicture(picture core.Picture, data core.LabelData) error {

	resulRequest, err := Client.ZAdd("pictureIdLabel:"+string(picture.Id), redis.Z{
		Score:  float64(data.Score),
		Member: string(data.Id),
	}).Result()
	if err == redis.Nil {
		// log.Printf("Failed to push PictureIdLabel:  %v", err)
		return fmt.Errorf("Failed to push PictureIdLabel: %s", string(picture.Id))
	} else if err != nil {
		return err
	}
	fmt.Println(resulRequest)

	return err
}

func PostLocalizedObjectByPicture(picture core.Picture, data core.LocalizedObjectData) error {

	resulRequest, err := Client.ZAdd("pictureIdObject:"+string(picture.Id), redis.Z{
		Score:  float64(data.Score),
		Member: string(data.Id),
	}).Result()
	if err == redis.Nil {
		// log.Printf("Failed to push PictureIdObject:  %v", err)
		return fmt.Errorf("Failed to push PictureIdObject: %s", string(picture.Id))
	} else if err != nil {
		return err
	}
	fmt.Println(resulRequest)

	return err
}

func PostCreateLabel(data core.LabelData) error {
	var err error

	if Client.HSet("labelDescr:"+data.Id, data.Language, data.Description).Val() != true {
		fmt.Println("already in or failed to push")
	}
	resulRequest, err := Client.SAdd("labelId", data.Id, 0).Result()
	if err == redis.Nil {
		// log.Printf("Failed to push labelId:  %v", err)
		return fmt.Errorf("Failed to push labelId: %s", string(data.Id))
	} else if err != nil {
		// log.Print(err)
		return err
	}
	fmt.Println(resulRequest)

	return err
}

func PostCreateObject(data core.LocalizedObjectData) error {
	var err error

	if Client.HSet("objectDescr:"+data.Id, data.Language, data.Name).Val() != true {
		// log.Fatalf("Failed to push objectDescr:  %v", err)
		fmt.Println("already in or failed to push")
	}
	resulRequest, err := Client.SAdd("objectId", data.Id, 0).Result()
	if err == redis.Nil {
		// log.Printf("Failed to push objectId:  %v", err)
		return fmt.Errorf("Failed to push objectId: %s", string(data.Id))
	} else if err != nil {
		// log.Print(err)
		return err 
	}
	fmt.Println(resulRequest)

	return err
}

func PostPicturebyLabel(picture core.Picture, data core.LabelData) error {
	err := Client.SAdd("Id:"+string(data.Id), picture.Id, 0).Err()
	if err == redis.Nil {
		// log.Printf("Failed to push Picture by Label:  %v", err)
		return fmt.Errorf("Failed to push label: %s", string(data.Id))
	} else if err != nil {
		// log.Print(err)
		return err
	}
	return err
}

func PostPicturebyObject(picture core.Picture, data core.LocalizedObjectData) error {
	err := Client.SAdd("Id:"+string(data.Id), picture.Id, 0).Err()
	if err == redis.Nil {
		// log.Printf("Failed to push Picture by Label:  %v", err)
		return fmt.Errorf("Failed to push label: %s", string(data.Id))
	} else if err != nil {
		// log.Print(err)
		return err
	}
	return err
}
