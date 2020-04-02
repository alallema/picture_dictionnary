package repo

import (
	"fmt"

	core "github.com/alallema/picture_dictionnary.git/core/service"
	"github.com/go-redis/redis"
)

func GetVideo(id string) core.Picture {
	var video core.Picture

	resultRequest, err := Client.HGetAll("video:" + id).Result()
	_ = err
	for i, pict := range resultRequest {
		if i == "id" {
			video.Id = pict
		}
		if i == "title" {
			video.Title = pict
		}
		if i == "format" {
			video.Format = pict
		}
		if i == "videoPath" {
			video.PicturePath = pict
		}
		if i == "videoURL" {
			video.PictureURL = pict
		}
		if i == "videoPath" {
			video.PicturePath = pict
		}
		if i == "source" {
			video.Source = pict
		}
	}
	return video
}

func GetCategoryDescription(id string) string {
	resultRequest, err := Client.HGet("categoryDescr:"+id, "en").Result()
	_ = err
	return resultRequest
}

func GetVideoByLabel(id string) []string {
	resultRequest, err := Client.SMembers("VidId:" + id).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		fmt.Println(err)
	}
	return resultRequest
}
