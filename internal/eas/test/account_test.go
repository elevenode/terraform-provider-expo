package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"terraform-provider-eas/internal/eas/internal/api/account"
	"terraform-provider-eas/internal/eas/test/utils"
)

func TestGetAccountByName(t *testing.T) {
	expectedData := &account.Data{
		Id:   utils.AccountId,
		Name: "expo-eas-sdk-go",
	}

	actualData, actualErr := utils.Client.Account.GetByName(expectedData.Name)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
