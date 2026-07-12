package test

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/account"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
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
