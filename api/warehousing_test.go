package api_test

import (
	"github.com/lifegit/video/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarehousing(t *testing.T) {
	err := api.Warehousing()
	assert.NoError(t, err)
}
