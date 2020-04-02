package service

import (
	"context"
	"os"
	"reflect"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/alallema/picture_dictionnary.git/core/service"
)

func TestConfigStorage_GetStorageClient(t *testing.T) {
	type fields struct {
		Ctx        context.Context
		Client     *storage.Client
		BucketName string
		It         *storage.ObjectIterator
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "No client",
			fields: fields{
				Ctx:        context.Background(),
				Client:     nil,
				BucketName: "No Name",
				It:         nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := ConfigStorage{
				Ctx:        tt.fields.Ctx,
				Client:     tt.fields.Client,
				BucketName: tt.fields.BucketName,
				It:         tt.fields.It,
			}
			if got := conf.GetStorageClient(); got == nil {
				t.Errorf("GetStorageClient() = %v ", got)
			}
		})
	}
}

func TestCreateConfStorage(t *testing.T) {
	tests := []struct {
		name string
		want ConfigStorage
	}{
		{
			name: "Test1",
			want: ConfigStorage{
				Ctx:        context.Background(),
				Client:     nil,
				BucketName: os.Getenv("BUCKETNAMETOPROCESS"),
				It:         nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateConfStorage(); got.Ctx == nil || got.BucketName != os.Getenv("BUCKETNAMETOPROCESS") {
				t.Errorf("CreateConfStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreatePictureData(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    service.Picture
		wantErr bool
	}{
		{
			name: "TestVideoNotGoodFormat",
			args: args{
				file: "test1-file.mp4",
			},
			want:    service.Picture{},
			wantErr: true,
		},
		{
			name: "TestVideo",
			args: args{
				file: "Test/Test/test1-file.mp4",
			},
			want: service.Picture{
				Title:       "test1-file.mp4",
				Format:      "mp4",
				Source:      "Test",
				PicturePath: "gs://picture-dictionnary-bucket/Test/Test/test1-file.mp4",
				PictureURL:  "https://storage.cloud.google.com/picture-dictionnary-bucket/Test/Test/test1-file.mp4",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreatePictureData(tt.args.file)
			println(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePictureData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.Id = got.Id
			tt.want.CreatedDate = got.CreatedDate
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePictureData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyToBucket(t *testing.T) {
	clientTest, _ := storage.NewClient(context.Background())
	type args struct {
		client    *storage.Client
		dstBucket string
		srcBucket string
		srcObject string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNoBucket",
			args: args{
				client:    clientTest,
				dstBucket: "",
				srcBucket: "",
				srcObject: "test1-file",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyToBucket(tt.args.client, tt.args.dstBucket, tt.args.srcBucket, tt.args.srcObject); (err != nil) != tt.wantErr {
				t.Errorf("CopyToBucket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	clientTest, _ := storage.NewClient(context.Background())
	type args struct {
		client *storage.Client
		bucket string
		object string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNoBucket",
			args: args{
				client: clientTest,
				bucket: "",
				object: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.client, tt.args.bucket, tt.args.object); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
