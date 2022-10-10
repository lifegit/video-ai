package api_test

import (
	"github.com/lifegit/video/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAIPut(t *testing.T) {
	err := api.AIPut()
	assert.NoError(t, err)
}
