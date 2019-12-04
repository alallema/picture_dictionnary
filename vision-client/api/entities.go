package api

import (
	"context"
	"io"

	vision "cloud.google.com/go/vision/apiv1"
)

type Config struct {
	Ctx      context.Context
	Client   *vision.ImageAnnotatorClient
	Filename string
	W        io.Writer
	File     io.Reader
}
