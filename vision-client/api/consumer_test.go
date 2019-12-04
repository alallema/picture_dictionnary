package api

import (
	"context"
	"fmt"
	"testing"
)

func Test_getClient(t *testing.T) {

	var conf Config
	conf.Ctx = context.Background()
	conf.Client = conf.GetClient()

	if conf.Client != nil {
		fmt.Printf("client\n")
	} else {
		fmt.Printf("no client\n")
	}
}
