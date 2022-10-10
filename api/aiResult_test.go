package api_test

import (
	"github.com/lifegit/video/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAIResult(t *testing.T) {
	err := api.AIResult()
	assert.NoError(t, err)
}
