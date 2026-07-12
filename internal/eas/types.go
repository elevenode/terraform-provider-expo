package eas

import (
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/account"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/accountvariable"
	androidappbuildcredentials "github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/android/appbuildcredentials"
	androidappcredentials "github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/android/appcredentials"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/android/fcmkey"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/android/googleserviceaccountkey"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/android/keystore"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/app"
	appleappbuildcredentials "github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appbuildcredentials"
	appleappcredentials "github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appcredentials"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appidentifier"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appstoreapikey"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/certificate"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/provisioningprofile"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/pushkey"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/team"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/appvariable"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/environmentvariable"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/me"
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
