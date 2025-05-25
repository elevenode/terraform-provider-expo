package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleCertificate(t *testing.T) {
	input := eas.GetBySerialNumberAppleCertificateData{
		SerialNumber: "3D986E25FF1B48C2417853A07AA15C55",
		AccountId:    utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.Certificate.GetBySerialNumber(input)

	expectedData := &eas.AppleCertificateData{
		Id:           utils.ImmutableCertificateId,
		SerialNumber: input.SerialNumber,
	}
	assert.Equal(t, expectedData, actualData)
	assert.NoError(t, actualErr)
}
