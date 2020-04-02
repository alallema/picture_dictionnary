package main

import (
	"fmt"
	"os"

	collector "github.com/alallema/picture_dictionnary.git/collector/service"
	core "github.com/alallema/picture_dictionnary.git/core/service"
	redis "github.com/alallema/picture_dictionnary.git/redis/repo"
	"github.com/alallema/picture_dictionnary.git/vision-client/service"
	video "github.com/alallema/picture_dictionnary.git/vision-client/video-api"
	vision "github.com/alallema/picture_dictionnary.git/vision-client/vision-api"
	"google.golang.org/api/iterator"
)

func main() {
	confSto := collector.CreateConfStorage()

	for {
		attrs, err := confSto.It.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Println("NAME:", attrs.Name)
		format := core.CheckFormat(attrs.Name)
		fmt.Println("FORMAT:", format)
		if format == 1 {
			fmt.Println("Pict: ", attrs.Name)
			confVis := service.CreateConfVision()
			confVis.GcsURI = "gs://" + os.Getenv("BUCKETNAMETOPROCESS") + "/" + attrs.Name
			picture, err := collector.CreatePictureData(attrs.Name)
			if err == nil && picture.Id != "" && picture.Id != "0" {
				redis.PostPicture(picture)
				labels, err := vision.DetectLabelsFromUri(confVis)
				if err == nil {
					for _, label := range labels {
						if label.Id != "" && label.Id != "0" {
							redis.PostCreateLabel(label)
							redis.PostPicturebyLabel(picture, label)
							redis.PostLabelByPicture(picture, label)
							redis.PostImageUrlbyLabel(picture, label)
						}
					}
				}
				objects, err := vision.LocalizeObjectsFromUri(confVis)
				if err == nil {
					for _, object := range objects {
						if object.Id != "" && object.Id != "0" {
							redis.PostCreateObject(object)
							redis.PostPicturebyObject(picture, object)
							redis.PostLocalizedObjectByPicture(picture, object)
							redis.PostImageUrlbyObject(picture, object)
							for _, v := range object.BoundingPoly.NormalizedVertices {
								fmt.Printf("(%f,%f)\n", v.X, v.Y)
							}
						}
					}
				}
			}
		} else if format == 2 {
			fmt.Println("Video: ", attrs.Name)
			confVid := service.CreateConfVideo()
			confVid.GcsURI = "gs://" + os.Getenv("BUCKETNAMETOPROCESS") + "/" + attrs.Name
			vid, err := collector.CreatePictureData(attrs.Name)
			fmt.Printf("Video Id: %s- Video Name: %s", vid.Id, vid.Title)
			if err == nil && vid.Id != "" && vid.Id != "0" {
				redis.PostVideo(vid)
				labels, err := video.DetectLabelVideo(confVid)
				if err == nil {
					for _, label := range labels {
						if label.Entity.Id != "" && label.Entity.Id != "0" {
							redis.PostCreateVideoLabel(label)
							redis.PostVideobyLabel(vid, label)
							redis.PostLabelbyCategory(label, label.CategoryEntities)
							redis.PostCreateVideoCategory(label.CategoryEntities)
							redis.PostSegbyVideobyLabel(label.Segments, label, vid)
						}
					}
				}
			}
		}
		collector.CopyToBucket(confSto.Client, os.Getenv("BUCKETNAME"), os.Getenv("BUCKETNAMETOPROCESS"), attrs.Name)
		collector.Delete(confSto.Client, os.Getenv("BUCKETNAMETOPROCESS"), attrs.Name)
	}
}
