package api_test

import (
	"github.com/lifegit/video/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTran(t *testing.T) {
	err := api.Tran()
	assert.NoError(t, err)
}
