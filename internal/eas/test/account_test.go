package test

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/fintreal/eas-sdk-go/test/utils"
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
