package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	core "github.com/alallema/picture_dictionnary.git/core/service"
	guuid "github.com/google/uuid"
)

func (conf ConfigStorage) GetStorageClient() *storage.Client {
	client, err := storage.NewClient(conf.Ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func CreateConfStorage() ConfigStorage {
	var conf ConfigStorage
	conf.Ctx = context.Background()
	conf.Client = conf.GetStorageClient()
	conf.BucketName = os.Getenv("BUCKETNAMETOPROCESS")
	conf.It = conf.Client.Bucket(conf.BucketName).Objects(conf.Ctx, nil)
	return conf
}

// Path prototype :
// collector/directory_source/title.format

func CreatePictureData(file string) (core.Picture, error) {
	var picture core.Picture

	picData := strings.Split(file, "/")
	lenData := len(picData)
	if lenData >= 3 {
		pic := strings.Split(picData[lenData-1], ".")
		picture.Title = picData[lenData-1]
		picture.Source = picData[lenData-2]
		if len(pic) == 2 {
			picture.Format = pic[1]
		}
		picture.PicturePath = "gs://" + os.Getenv("BUCKETNAME") + "/" + file
		picture.PictureURL = "https://storage.cloud.google.com/" + os.Getenv("BUCKETNAME") + "/" + file
		picture.CreatedDate = time.Now()
		picture.Id = guuid.New().String()
	} else {
		log.Printf("Not good formating picture path")
		return picture, fmt.Errorf("Not good formating picture path")
	}
	return picture, nil
}

func CopyToBucket(client *storage.Client, dstBucket, srcBucket, srcObject string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	dstObject := srcObject
	src := client.Bucket(srcBucket).Object(srcObject)
	dst := client.Bucket(dstBucket).Object(dstObject)

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return err
	}
	return nil
}

func Delete(client *storage.Client, bucket, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	o := client.Bucket(bucket).Object(object)
	if err := o.Delete(ctx); err != nil {
		return err
	}
	return nil
}
