package images

import (
	"testing"

	"github.com/henrahmagix/go-dep-registry/aws"
	"github.com/henrahmagix/go-dep-registry/dependencies"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func registerAndDeleteDep(dep interface{}) func() {
	dependencies.RegisterGlobal(dep)
	return func() {
		dependencies.DeleteGlobal(dep)
	}
}

func TestUpload(t *testing.T) {
	mockAPI := aws.NewAPI("test", "test")
	teardown := registerAndDeleteDep(&mockAPI)
	defer teardown()

	u, err := NewUploader()
	require.NoError(t, err)
	assert.Equal(t, "Uploading image to AWS: testing test:test", u.Upload())
}

func TestUploadDepError(t *testing.T) {
	_, err := NewUploader()
	require.Error(t, err)
	assert.IsType(t, &dependencies.ErrNotRegistered{}, err)
}
