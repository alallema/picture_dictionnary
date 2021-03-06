package repo

import (
	"fmt"
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/go-redis/redis"
)

func GetPicture(id string) core.Picture {
	var picture core.Picture

	resultRequest, err := Client.HGetAll("picture:" + id).Result()
	_ = err
	for i, pict := range resultRequest {
		if i == "id" {
			picture.Id = pict
		}
		if i == "title" {
			picture.Title = pict
		}
		if i == "format" {
			picture.Format = pict
		}
		if i == "picturePath" {
			picture.PicturePath = pict
		}
		if i == "pictureURL" {
			picture.PictureURL = pict
		}
		if i == "picturePath" {
			picture.PicturePath = pict
		}
		if i == "source" {
			picture.Source = pict
		}
	}
	return picture
}

func GetLabelByPicture(id string) []string {
	resultRequest, err := Client.ZRangeByScore("pictureIdLabel:"+id, redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	_ = err
	return resultRequest
}

func GettLocalizedObjectByPicture(id string) []string {
	resultRequest, err := Client.ZRangeByScore("pictureIdObject:"+id, redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	_ = err
	return resultRequest
}

func GetLabelDescription(id string) string {
	keys := Client.HKeys("labelDescr:" + id)
	if len(keys.Val()) > 0 {
		resultRequest, err := Client.HGet("labelDescr:"+id, keys.Val()[0]).Result()
		_ = err
		return resultRequest
	}
	return "Not found"
}

func GetObjectDescription(id string) string {
	resultRequest, err := Client.HGet("objectDescr:"+id, "en").Result()
	_ = err
	return resultRequest
}

func GetPictureByLabel(id string) []string {
	resultRequest, err := Client.SMembers("Id:" + id).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		fmt.Println(err)
	}
	return resultRequest
}

func GetPictureUrlByLabel(id string) []string {
	resultRequest, err := Client.SMembers("URLId:" + id).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		fmt.Println(err)
	}
	return resultRequest
}
