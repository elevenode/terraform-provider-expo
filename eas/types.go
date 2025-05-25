package eas

import (
	"github.com/fintreal/eas-sdk-go/internal/api/account"
	androidappbuildcredentials "github.com/fintreal/eas-sdk-go/internal/api/android/appbuildcredentials"
	androidappcredentials "github.com/fintreal/eas-sdk-go/internal/api/android/appcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/android/fcmkey"
	"github.com/fintreal/eas-sdk-go/internal/api/android/googleserviceaccountkey"
	"github.com/fintreal/eas-sdk-go/internal/api/android/keystore"
	"github.com/fintreal/eas-sdk-go/internal/api/app"
	appleappbuildcredentials "github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"
	appleappcredentials "github.com/fintreal/eas-sdk-go/internal/api/apple/appcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appidentifier"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appstoreapikey"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/certificate"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/provisioningprofile"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/pushkey"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/team"
	"github.com/fintreal/eas-sdk-go/internal/api/appvariable"
	"github.com/fintreal/eas-sdk-go/internal/api/me"
)

type MeData = me.Data
type AccountData = account.Data

type AppData = app.Data
type CreateAppData = app.CreateData
type UpdateAppData = app.UpdateData

type AppVariableData = appvariable.Data
type CreateAppVariableData = appvariable.CreateData
type UpdateAppVariableData = appvariable.UpdateData
type GetAppVariableData = appvariable.GetData
type GetByNameAppVariableData = appvariable.GetByNameData

type AppleTeamData = team.Data
type CreateAppleTeamData = team.CreateData
type UpdateAppleTeamData = team.UpdateData
type GetByIdentifierAppleTeamData = team.GetByIdentifierData

type AppleAppIdentifierData = appidentifier.Data
type CreateAppleAppIdentifierData = appidentifier.CreateData
type GetByIdentifierAppleAppIdentifierData = appidentifier.GetByIdentifierData

type AppleCertificateData = certificate.Data
type GetBySerialNumberAppleCertificateData = certificate.GetBySerialNumberData

type ProvisioningProfileData = provisioningprofile.Data
type CreateProvisioningProfileData = provisioningprofile.CreateData
type GetProvisioningProfileData = provisioningprofile.GetData

type AppStoreApiKeyData = appstoreapikey.Data
type GetByIdentifierAppStoreApiKeyData = appstoreapikey.GeyByIdentifierData

type AppCredentialsData = appleappcredentials.Data
type CreateAppCredentialsData = appleappcredentials.CreateData
type UpdateAppCredentialsData = appleappcredentials.UpdateData
type GetAppCredentialsData = appleappcredentials.GetData

type AppBuildCredentialsData = appleappbuildcredentials.Data
type CreateAppBuildCredentialsData = appleappbuildcredentials.CreateData
type GetAppBuildCredentialsData = appleappbuildcredentials.GetData

type GetByProjectIdentifierGoogleServiceAccountKeyData = googleserviceaccountkey.GetByProjectIdentifierData
type GoogleServiceAccountKeyData = googleserviceaccountkey.Data

type GetByIdentifierPushKeyData = pushkey.GeyByIdentifierData
type PushKeyData = pushkey.Data

type AndroidAppCredentialsData = androidappcredentials.Data
type CreateAndroidAppCredentialsData = androidappcredentials.CreateData
type GetAndroidAppCredentialsData = androidappcredentials.GetData

type AndroidAppBuildCredentialsData = androidappbuildcredentials.Data
type GetAndroidAppBuildCredentialsData = androidappbuildcredentials.GetData
type CreateAndroidAppBuildCredentialsData = androidappbuildcredentials.CreateData

type CreateFCMKey = fcmkey.CreateData

type AndroidKeystoreData = keystore.Data
type CreateAndroidKeystoreData = keystore.CreateData
