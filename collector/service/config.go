package service

import (
	"cloud.google.com/go/storage"
	"context"
)

type ConfigStorage struct {
	Ctx        context.Context
	Client     *storage.Client
	BucketName string
	It         *storage.ObjectIterator
}
