package androidtest

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetByProjectIdentifier(t *testing.T) {
	input := eas.GetByProjectIdentifierGoogleServiceAccountKeyData{
		ProjectIdentifier: "playstore-release-5nszcyhabt",
		AccountId:         utils.AccountId,
	}

	actualData, actualErr := utils.Client.Android.GoogleServiceAccountKey.GetByProjectIdentifier(input)

	expectedData := &eas.GoogleServiceAccountKeyData{
		Id:                "36b45ce5-1cf3-4e29-a04d-88fb3c4b5683",
		ProjectIdentifier: "playstore-release-5nszcyhabt",
		ClientEmail:       "play-store-release-5nszcyhabt@playstore-release-5nszcyhabt.iam.gserviceaccount.com",
		ClientIdentifier:  "108020064321595488457",
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
