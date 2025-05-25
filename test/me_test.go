package test

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/me"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	expectedData := &me.Data{
		Id:          utils.MeId,
		DisplayName: "integration-test",
	}

	actualData, actualErr := utils.Client.Me.Get()
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
