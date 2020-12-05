package images

import (
	"github.com/henrahmagix/go-dep-registry/aws"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpload(t *testing.T) {
	api := aws.NewAPI("test", "test")
	mockDeps := map[string]interface{}{
		"*aws.API": &api,
	}

	u, err := NewUploader(mockDeps)
	require.NoError(t, err)
	assert.Equal(t, "Uploading image to AWS: testing test:test", u.Upload())
}
