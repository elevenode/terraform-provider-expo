# eas-sdk-go

Expo Application Services SDK for Go.
- Manage Expo apps and environment variables from code
- Uses Expo GraphQL API
- Builds on [machinebox/graphql](https://github.com/machinebox/graphql). Big Thanks!

## Installation
```
go get github.com/fintreal/eas-sdk-go
```

## Usage

### Initialize
```go
client := eas.NewEASClient(token)

data, _ := client.Me.Get()

fmt.Printf("%+v\n", data)
```

### Account

##### Get By Name
```go
client := eas.NewEASClient(token)

data, _ := client.Account.GetByName(accountName)

fmt.Printf("%+v\n", data)
```

### App

##### Get
```go
client := NewEASClient(token)

data, err := client.App.Get(appId)

fmt.Printf("%+v", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAppData{
	AccountId: accountId,
	Name:      "Test App Name",
	Slug:      "Test Slug",
}

data, _ := client.App.Create(inputData)

fmt.Printf("%+v\n", data)
```

##### Update
```go
client := eas.NewEASClient(token)

inputData := eas.UpdateAppData{
	AccountId: accountId,
	Name:      "Test App Name",
}

data, _ := client.App.Update(inputData)

fmt.Printf("%+v\n", data)
```

### App Environment Variable

##### Get
```go
client := eas.NewEASClient(token)
	
data, _ := client.AppVariable.Get(appVariableId, appId)

fmt.Printf("%+v\n", data)
```

##### Get By Name
```go
client := eas.NewEASClient(token)

data, _ := client.AppVariable.GetByName(appVariableName, appId)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)
inputData := eas.CreateAppVariableData{
	AppId:        appId,
	Name:         "API_URL",
	Value:        "http://localhost:3000/api",
	Visibility:   "PUBLIC",
	Environments: []string{"PREVIEW", "PRODUCTION"},
}

data, _ := client.AppVariable.Create(inputData)

fmt.Printf("%+v\n", data)
```

##### Update
```go
client := eas.NewEASClient(token)
inputData := eas.UpdateAppVariableData{
	Id:           id,
	Name:         "API_URL",
	Value:        "http://localhost:3000/api",
	Visibility:   "PUBLIC",
	Environments: []string{"DEVELOPMENT", "PREVIEW"},
}

data, _ := client.AppVariable.Create(inputData)

fmt.Printf("%+v\n", data)
```

##### Delete
```go
client := eas.NewEASClient(token)

data, _ := client.AppVariable.Delete(appVariableId)

fmt.Printf("%+v\n", data)
```

### Apple Service

##### Team

##### Get By Identifier
```go
client := eas.NewEASClient(token)

data, _ := client.Apple.Team.GetByIdentifier(teamIdentifier)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAppleTeamData{
    AccountId: accountId,
    Name:      "Team Name",
}

data, _ := client.Apple.Team.Create(inputData)

fmt.Printf("%+v\n", data)
```

##### Update
```go
client := eas.NewEASClient(token)

inputData := eas.UpdateAppleTeamData{
    Id:   teamId,
    Name: "Updated Team Name",
}

data, _ := client.Apple.Team.Update(inputData)

fmt.Printf("%+v\n", data)
```

### Apple App Identifier

##### Get By Identifier
```go
client := eas.NewEASClient(token)

data, _ := client.Apple.AppIdentifier.GetByIdentifier(identifier)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAppleAppIdentifierData{
    AccountId: accountId,
    BundleId:  "com.example.app",
}

data, _ := client.Apple.AppIdentifier.Create(inputData)

fmt.Printf("%+v\n", data)
```

### Apple Certificate

##### Get By Serial Number
```go
client := eas.NewEASClient(token)

data, _ := client.Apple.Certificate.GetBySerialNumber(serialNumber)

fmt.Printf("%+v\n", data)
```

### Apple Provisioning Profile

##### Get
```go
client := eas.NewEASClient(token)

data, _ := client.Apple.ProvisioningProfile.Get(profileId)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateProvisioningProfileData{
    AccountId: accountId,
    AppId:     appId,
    // Add other required fields
}

data, _ := client.Apple.ProvisioningProfile.Create(inputData)

fmt.Printf("%+v\n", data)
```

### Apple App Store API Key

##### Get By Identifier
```go
client := eas.NewEASClient(token)

data, _ := client.Apple.AppStoreApiKey.GetByIdentifier(identifier)

fmt.Printf("%+v\n", data)
```

### Apple App Credentials

##### Get
```go
client := eas.NewEASClient(token)

data, _ := client.Apple.AppCredentials.Get(credentialsId)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAppCredentialsData{
    AccountId: accountId,
    AppId:     appId,
    // Add other required fields
}

data, _ := client.Apple.AppCredentials.Create(inputData)

fmt.Printf("%+v\n", data)
```

##### Update
```go
client := eas.NewEASClient(token)

inputData := eas.UpdateAppCredentialsData{
    Id:   credentialsId,
    // Add other fields to update
}

data, _ := client.Apple.AppCredentials.Update(inputData)

fmt.Printf("%+v\n", data)
```

### Apple App Build Credentials

##### Get
```go
client := eas.NewEASClient(token)

data, _ := client.Apple.AppBuildCredentials.Get(buildCredentialsId)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAppBuildCredentialsData{
    AccountId: accountId,
    AppId:     appId,
    // Add other required fields
}

data, _ := client.Apple.AppBuildCredentials.Create(inputData)

fmt.Printf("%+v\n", data)
```

### Android Service

### Android App Credentials

##### Get
```go
client := eas.NewEASClient(token)

data, _ := client.Android.AppCredentials.Get(credentialsId)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAndroidAppCredentialsData{
    AccountId: accountId,
    AppId:     appId,
    // Add other required fields
}

data, _ := client.Android.AppCredentials.Create(inputData)

fmt.Printf("%+v\n", data)
```

### Android App Build Credentials

##### Get
```go
client := eas.NewEASClient(token)

data, _ := client.Android.AppBuildCredentials.Get(buildCredentialsId)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAndroidAppBuildCredentialsData{
    AccountId: accountId,
    AppId:     appId,
    // Add other required fields
}

data, _ := client.Android.AppBuildCredentials.Create(inputData)

fmt.Printf("%+v\n", data)
```

### Android FCM Key

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateFCMKey{
    AccountId: accountId,
    AppId:     appId,
    // Add other required fields
}

data, _ := client.Android.FCMKey.Create(inputData)

fmt.Printf("%+v\n", data)
```

### Android Keystore

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAndroidKeystoreData{
    AccountId: accountId,
    AppId:     appId,
    // Add other required fields
}

data, _ := client.Android.Keystore.Create(inputData)

fmt.Printf("%+v\n", data)
```

### Google Service Account Key

##### Get By Project Identifier
```go
client := eas.NewEASClient(token)

data, _ := client.Android.GoogleServiceAccountKey.GetByProjectIdentifier(projectIdentifier)

fmt.Printf("%+v\n", data)
```

### Apple Push Key

##### Get By Identifier
```go
client := eas.NewEASClient(token)

data, _ := client.Apple.PushKey.GetByIdentifier(identifier)

fmt.Printf("%+v\n", data)
```
