package aws

import (
	"fmt"
)

// API is a third-party thing that doesn't know about our dependency registry.
type API struct {
	key    string
	secret string
}

// NewAPI returns an API.
func NewAPI(key, secret string) API {
	return API{key: key, secret: secret}
}

// UploadImage returns the data along with the credentials (😬 naughty!)
func (api API) UploadImage(data []byte) string {
	return fmt.Sprintf("Uploading image to AWS: %s %s:%s", data, api.key, api.secret)
}
