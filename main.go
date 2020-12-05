package main

import (
	"fmt"

	"github.com/henrahmagix/go-dep-registry/aws"
	"github.com/henrahmagix/go-dep-registry/dependencies"
	"github.com/henrahmagix/go-dep-registry/images"
)

func main() {
	awsAPI := aws.NewAPI("my key", "my secret")

	err := dependencies.RegisterGlobal(&awsAPI)
	if err != nil {
		panic(err)
	}

	uploader, err := images.NewUploader()
	if err != nil {
		panic(err)
	}

	fmt.Println(uploader.Upload())
}
