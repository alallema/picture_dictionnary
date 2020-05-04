package server

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

func (server *Server) AllLabels() *[]Tag {
	var labelList []Tag

	resultRequest, err := server.Client.SMembers("labelId").Result()
	if err == redis.Nil {
		log.Warn().Msg("key does not exist")
	} else if err != nil {
		log.Error().Err(err)
	}
	for _, labelId := range resultRequest {
		keys := server.Client.HKeys("labelDescr:" + labelId)
		label := server.Client.HGet("labelDescr:"+labelId, keys.Val()[0])
		if label.Err() == redis.Nil {
			msg := fmt.Sprintf("Label Id has no description %s", labelId)
			log.Error().Err(err).Msg(msg)
		} else {
			tag := Tag{
				Id:    labelId,
				Title: label.Val(),
				Type:  "label",
			}
			labelList = append(labelList, tag)
		}
	}
	return &labelList
}

func (server *Server) AllObjects() *[]Tag {
	var objectList []Tag

	resultRequest, err := server.Client.SMembers("objectId").Result()
	if err == redis.Nil {
		log.Warn().Msg("key does not exist")
	} else if err != nil {
		log.Error().Err(err)
	}
	for _, objectId := range resultRequest {
		if (objectId != "") {
			object := server.Client.HGet("objectDescr:"+objectId, "en")
			if object.Err() == redis.Nil {
				log.Error().Err(err).Msg("object Id has no description")
			} else {
				tag := Tag{
					Id:    objectId,
					Title: object.Val(),
					Type:  "category",
				}
				objectList = append(objectList, tag)
			}
		}
	}
	return &objectList
}

func (server *Server) AllCategories() *[]Tag {
	var catList []Tag

	resultRequest, err := server.Client.SMembers("categoryId").Result()
	if err == redis.Nil {
		log.Warn().Msg("key does not exist")
	} else if err != nil {
		log.Error().Err(err)
	}
	for _, catId := range resultRequest {
		cat := server.Client.HGet("categoryDescr:"+catId, "en-US")
		if cat.Err() == redis.Nil {
			log.Error().Err(err).Msg("cat Id has no description")
		} else {
			tag := Tag{
				Id:    catId,
				Title: cat.Val(),
				Type:  "object",
			}
			catList = append(catList, tag)
		}
	}
	return &catList
}

func (server *Server) PictureByTag(id string) *[]Picture {
	var pictureList []Picture

	resultRequest, err := server.Client.SMembers("Id:" + string(id)).Result()
	if err == redis.Nil {
		log.Warn().Msg("key does not exist")
	} else if err != nil {
		log.Error().Err(err)
	} else if len(resultRequest) == 0 {
		log.Info().Str("picture", string(id)).Msg("No picture result for this tag")
	}
	for _, pictureId := range resultRequest {
		var picture Picture

		if pictureId != "" {
			pictures, err := server.Client.HGetAll("picture:" + pictureId).Result()
			if err == redis.Nil {
				log.Error().Err(err).Msg("Picture doesn't exist")
			} else if err != nil {
				log.Error().Err(err)
			} else {
				for i, pict := range pictures {
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
			}
			pictureList = append(pictureList, picture)
		}
	}
	return &pictureList
}

func (server *Server) VideoByTag(id string) *[]Picture {
	var videoList []Picture

	resultRequest, err := server.Client.SMembers("VidId:" + string(id)).Result()
	if err == redis.Nil {
		log.Warn().Msg("key does not exist")
	} else if err != nil {
		log.Error().Err(err)
	} else if len(resultRequest) == 0 {
		log.Info().Str("video", string(id)).Msg("No video result for this tag")
	}
	for _, videoId := range resultRequest {
		var video Picture

		if videoId != "" {
			videos, err := server.Client.HGetAll("video:" + videoId).Result()
			if err == redis.Nil {
				log.Error().Err(err).Msg("Picture doesn't exist")
			} else if err != nil {
				log.Error().Err(err)
			} else {
				for i, pict := range videos {
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
			}
			videoList = append(videoList, video)
		}
	}
	return &videoList
}

func (server *Server) PictureUrlByTag(tag string) *[]string {
	var urlList []string

	resultRequest, err := server.Client.SMembers("URLId" + tag).Result()
	if err == redis.Nil {
		log.Warn().Msg("key does not exist")
	} else if err != nil {
		log.Error().Err(err)
	}
	for _, url := range resultRequest {
		urlList = append(urlList, url)
	}
	return &urlList
}
