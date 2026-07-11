package eas

import (
	"terraform-provider-eas/internal/eas/internal/api/account"
	"terraform-provider-eas/internal/eas/internal/api/accountvariable"
	androidappbuildcredentials "terraform-provider-eas/internal/eas/internal/api/android/appbuildcredentials"
	androidappcredentials "terraform-provider-eas/internal/eas/internal/api/android/appcredentials"
	"terraform-provider-eas/internal/eas/internal/api/android/fcmkey"
	"terraform-provider-eas/internal/eas/internal/api/android/googleserviceaccountkey"
	"terraform-provider-eas/internal/eas/internal/api/android/keystore"
	"terraform-provider-eas/internal/eas/internal/api/app"
	appleappbuildcredentials "terraform-provider-eas/internal/eas/internal/api/apple/appbuildcredentials"
	appleappcredentials "terraform-provider-eas/internal/eas/internal/api/apple/appcredentials"
	"terraform-provider-eas/internal/eas/internal/api/apple/appidentifier"
	"terraform-provider-eas/internal/eas/internal/api/apple/appstoreapikey"
	"terraform-provider-eas/internal/eas/internal/api/apple/certificate"
	"terraform-provider-eas/internal/eas/internal/api/apple/provisioningprofile"
	"terraform-provider-eas/internal/eas/internal/api/apple/pushkey"
	"terraform-provider-eas/internal/eas/internal/api/apple/team"
	"terraform-provider-eas/internal/eas/internal/api/appvariable"
	"terraform-provider-eas/internal/eas/internal/api/environmentvariable"
	"terraform-provider-eas/internal/eas/internal/api/me"
)

type MeData = me.Data
type AccountData = account.Data

type EnvironmentVariableData = environmentvariable.Data
type UpdateEnvironmentVariableData = environmentvariable.UpdateData

type AppData = app.Data
type CreateAppData = app.CreateData
type UpdateAppData = app.UpdateData

type AppVariableData = appvariable.Data
type CreateAppVariableData = appvariable.CreateData
type UpdateAppVariableData = appvariable.UpdateData
type GetAppVariableData = appvariable.GetData
type GetByNameAppVariableData = appvariable.GetByNameData

type AccountVariableData = accountvariable.Data
type CreateAccountVariableData = accountvariable.CreateData
type UpdateAccountVariableData = accountvariable.UpdateData
type GetAccountVariableData = accountvariable.GetData

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
