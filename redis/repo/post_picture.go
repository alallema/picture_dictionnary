package repo

import (
	"fmt"
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/go-redis/redis"
)

func PostPicture(picture core.Picture) {
	if Client.HSet("picture:"+string(picture.Id), "id", string(picture.Id)).Val() != true {
		fmt.Printf("Failed to push PictureId: %s", string(picture.Id))
		return
	}
	if Client.HSet("picture:"+string(picture.Id), "title", picture.Title).Val() != true {
		fmt.Printf("Failed to push PictureTitle: %s", string(picture.Title))
	}
	if Client.HSet("picture:"+string(picture.Id), "format", picture.Format).Val() != true {
		fmt.Printf("Failed to push PictureFormat: %s", string(picture.Format))
	}
	if Client.HSet("picture:"+string(picture.Id), "picturePath", picture.PicturePath).Val() != true {
		fmt.Printf("Failed to push PicturePath: %s", string(picture.PicturePath))
	}
	if Client.HSet("picture:"+string(picture.Id), "pictureURL", picture.PictureURL).Val() != true {
		fmt.Printf("Failed to push PictureURL: %s", string(picture.PictureURL))
	}
	if Client.HSet("picture:"+string(picture.Id), "source", picture.Source).Val() != true {
		fmt.Printf("Failed to push Source: %s", string(picture.Source))
	}
	return
}

func PostLabelByPicture(picture core.Picture, data core.LabelData) {

	err := Client.ZAdd("pictureIdLabel:"+string(picture.Id), redis.Z{
		Score:  float64(data.Score),
		Member: string(data.Id),
	}).Err()
	if err == redis.Nil {
		fmt.Printf("Failed to push PictureIdLabel: %s", string(picture.Id))
	} else if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return
}

func PostLocalizedObjectByPicture(picture core.Picture, data core.LocalizedObjectData) {

	err := Client.ZAdd("pictureIdObject:"+string(picture.Id), redis.Z{
		Score:  float64(data.Score),
		Member: string(data.Id),
	}).Err()
	if err == redis.Nil {
		fmt.Printf("Failed to push PictureIdObject: %s", string(picture.Id))
	} else if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return
}

func PostCreateLabel(data core.LabelData) {

	if Client.HExists("labelDescr:"+data.Id, data.Language).Val() == false {
		if Client.HSet("labelDescr:"+data.Id, data.Language, data.Description).Val() != true {
			fmt.Printf("Failed to push labelDescr: %s", data.Id)
		}
		err := Client.SAdd("labelId", data.Id).Err()
		if err == redis.Nil {
			fmt.Printf("Failed to push labelId: %s", string(data.Id))
		} else if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
	return
}

func PostCreateObject(data core.LocalizedObjectData) {
	if Client.HExists("objectDescr:"+data.Id, data.Language).Val() == false {
		if Client.HSet("objectDescr:"+data.Id, data.Language, data.Name).Val() != true {
			fmt.Printf("Failed to push objectDescr: %s", string(data.Id))
		}
		err := Client.SAdd("objectId", data.Id).Err()
		if err == redis.Nil {
			fmt.Printf("Failed to push objectId: %s", string(data.Id))
		} else if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
	return
}

func PostPicturebyLabel(picture core.Picture, data core.LabelData) {
	err := Client.SAdd("Id:"+string(data.Id), picture.Id).Err()
	if err == redis.Nil {
		fmt.Printf("Failed to push label: %s", string(data.Id))
	} else if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return
}

func PostPicturebyObject(picture core.Picture, data core.LocalizedObjectData) {
	err := Client.SAdd("Id:"+string(data.Id), picture.Id).Err()
	if err == redis.Nil {
		fmt.Printf("Failed to push label: %s", string(data.Id))
	} else if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return
}

func PostImageUrlbyLabel(picture core.Picture, data core.LabelData) {
	err := Client.SAdd("URLId:"+string(data.Id), picture.PictureURL).Err()
	if err == redis.Nil {
		fmt.Printf("Failed to push label: %s", string(data.Id))
	} else if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return
}

func PostImageUrlbyObject(picture core.Picture, data core.LocalizedObjectData) {
	err := Client.SAdd("URLId:"+string(data.Id), picture.PictureURL).Err()
	if err == redis.Nil {
		fmt.Printf("Failed to push label: %s", string(data.Id))
	} else if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return
}
