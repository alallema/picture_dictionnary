package vision_api

import (
	"context"
	"fmt"
	"testing"
	config "github.com/alallema/picture_dictionnary.git/vision-client/service"
)

func Test_getClient(t *testing.T) {

	var conf config.Config
	conf.Ctx = context.Background()
	conf.Client = conf.GetClient()

	if conf.Client != nil {
		fmt.Printf("client\n")
	} else {
		fmt.Printf("no client\n")
	}
}
