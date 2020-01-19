package service

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
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
	conf.BucketName = "picture-dictionnary-bucket"
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
		picture.PicturePath = "gs://picture-dictionnary-bucket/" + file
		picture.PictureURL = "https://storage.cloud.google.com/picture-dictionnary-bucket/" + file
		picture.CreatedDate = time.Now()
		picture.Id = guuid.New().String()
	} else {
		log.Printf("Not good formating picture path")
		return picture, fmt.Errorf("Not good formating picture path")
	}
	return picture, nil
}

// func ExtractPictureFromDirectory(dir string) []core.Picture {
// 	var pictureFile []core.Picture

// 	files := GetDirectory(dir)
// 	for _, filename := range files {
// 		picture := CreatePictureData(filename.Name())
// 		pictureFile = append(pictureFile, picture)
// 	}

// 	return pictureFile
// }

func GetDirectory(dirPath string) []os.FileInfo {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Println(err)
	}
	return files
}

func GetFile(file string) io.Reader {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
	}
	// defer f.Close()
	return f
}

func ReadFile(file string) []byte {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
	}
	// defer f.Close()
	return f
}
