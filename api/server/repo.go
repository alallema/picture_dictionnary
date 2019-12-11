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
		label := server.Client.HGet("labelDescr:"+labelId, "en")
		if label.Err() == redis.Nil {
			log.Error().Err(err).Msg("Label Id has no description")
		} else {
			tag := Tag{
				Id:    labelId,
				Title: label.Val(),
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
		object := server.Client.HGet("objectDescr:"+objectId, "en")
		if object.Err() == redis.Nil {
			log.Error().Err(err).Msg("object Id has no description")
		} else {
			tag := Tag{
				Id:    objectId,
				Title: object.Val(),
			}
			objectList = append(objectList, tag)
		}
	}
	return &objectList
}

func (server *Server) PictureByTag(id string) *[]Picture {
	var pictureList []Picture

	resultRequest, err := server.Client.ZRangeByScore("Id:"+string(id), redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
	}).Result()
	fmt.Println(resultRequest)
	if err == redis.Nil {
		log.Warn().Msg("key does not exist")
	} else if err != nil {
		log.Error().Err(err)
	} else if len(resultRequest) == 0 {
		log.Info().Str("picture", string(id)).Msg("No result this picture")
	}
	return &pictureList
}
