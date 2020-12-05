package images

import (
	"github.com/henrahmagix/go-dep-registry/aws"
	"github.com/henrahmagix/go-dep-registry/dependencies"
)

// Uploader is our internal image uploader that takes an AWS API.
type Uploader struct {
	awsAPI aws.API
}

// NewUploader returns an Uploader, or an error if it can't get an AWS API.
func NewUploader() (Uploader, error) {
	uploader := Uploader{}

	awsAPI := aws.API{}
	err := dependencies.GetGlobal(&awsAPI)
	if err != nil {
		return uploader, err
	}

	uploader.awsAPI = awsAPI
	return uploader, nil
}

// Upload uses the AWS API to upload a test byte array.
func (u Uploader) Upload() string {
	return u.awsAPI.UploadImage([]byte("testing"))
}
