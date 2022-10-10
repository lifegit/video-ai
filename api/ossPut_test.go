package api_test

import (
	"github.com/lifegit/video/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOssPut(t *testing.T) {
	err := api.OssPut()
	assert.NoError(t, err)
}
