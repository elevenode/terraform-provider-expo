package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"terraform-provider-eas/internal/eas/internal/api/me"
	"terraform-provider-eas/internal/eas/test/utils"
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
