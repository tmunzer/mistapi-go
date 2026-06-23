# Orgs Linked Applications

```go
orgsLinkedApplications := client.OrgsLinkedApplications()
```

## Class Name

`OrgsLinkedApplications`

## Methods

* [Add Org Oauth App Accounts](../../doc/controllers/orgs-linked-applications.md#add-org-oauth-app-accounts)
* [Delete Org Oauth App Authorization](../../doc/controllers/orgs-linked-applications.md#delete-org-oauth-app-authorization)
* [Get Org Oauth App Linked Status](../../doc/controllers/orgs-linked-applications.md#get-org-oauth-app-linked-status)
* [Update Org Oauth App Account](../../doc/controllers/orgs-linked-applications.md#update-org-oauth-app-account)


# Add Org Oauth App Accounts

Add a linked account for the specified OAuth application using the app-specific account configuration payload.

```go
AddOrgOauthAppAccounts(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    body *models.AccountOauthAdd) (
    models.ApiResponse[models.AccountOauthInfo],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `appName` | [`models.OauthAppNameEnum`](../../doc/models/oauth-app-name-enum.md) | Template, Required | OAuth application name |
| `body` | [`*models.AccountOauthAdd`](../../doc/models/containers/account-oauth-add.md) | Body, Optional | Account configuration payload for adding a supported OAuth-linked application |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.AccountOauthInfo](../../doc/models/account-oauth-info.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

appName := models.OauthAppNameEnum_JAMF

body := models.AccountOauthAddContainer.FromAccountJamfConfig(models.AccountJamfConfig{
    ClientId:             "client_id0",
    ClientSecret:         "client_secret6",
    InstanceUrl:          "junipertest.jamfcloud.com",
    SmartgroupName:       "CompliantGroup1",
})

apiResponse, err := orgsLinkedApplications.AddOrgOauthAppAccounts(ctx, orgId, appName, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "accounts": [
    {
      "account_id": "ae9dee49-69e7-4710-a114-5b827a777738",
      "last_status": "success",
      "last_sync": 1665465339000,
      "linked_by": "Testname1",
      "linked_timestamp": 1728275689.284758,
      "name": "Test Company1 Ltd"
    },
    {
      "account_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "error": "Get token failed, please re-link crowdstrike",
      "last_status": "failed",
      "last_sync": 1234567890123,
      "linked_by": "Testname2",
      "linked_timestamp": 1728275689.284758,
      "name": "Test Company2 Ltd"
    }
  ],
  "linked": true
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Unsuccessful | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Org Oauth App Authorization

Remove a linked account authorization for the specified OAuth application and account ID.

```go
DeleteOrgOauthAppAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    accountId string) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `appName` | [`models.OauthAppNameEnum`](../../doc/models/oauth-app-name-enum.md) | Template, Required | OAuth application name |
| `accountId` | `string` | Template, Required | - |

## Response Type

**200**: Successful

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

appName := models.OauthAppNameEnum_JAMF

accountId := "iojzXIJWEuiD73ZvydOfg"

resp, err := orgsLinkedApplications.DeleteOrgOauthAppAuthorization(ctx, orgId, appName, accountId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Unsuccessful | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Org Oauth App Linked Status

Return linked-account status for the specified organization OAuth application and the authorization URL used to start account linking.

```go
GetOrgOauthAppLinkedStatus(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    forward string) (
    models.ApiResponse[models.AccountOauthInfo],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `appName` | [`models.OauthAppNameEnum`](../../doc/models/oauth-app-name-enum.md) | Template, Required | OAuth application name |
| `forward` | `string` | Query, Required | Mist portal url to which backend needs to redirect after successful OAuth authorization. **Required** to get the `authorization_url` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.AccountOauthInfo](../../doc/models/account-oauth-info.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

appName := models.OauthAppNameEnum_JAMF

forward := "forward6"

apiResponse, err := orgsLinkedApplications.GetOrgOauthAppLinkedStatus(ctx, orgId, appName, forward)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "accounts": [
    {
      "account_id": "ae9dee49-69e7-4710-a114-5b827a777738",
      "last_status": "success",
      "last_sync": 1665465339000,
      "linked_by": "Testname1",
      "linked_timestamp": 1728275689.284758,
      "name": "Test Company1 Ltd"
    },
    {
      "account_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "error": "Get token failed, please re-link crowdstrike",
      "last_status": "failed",
      "last_sync": 1234567890123,
      "linked_by": "Testname2",
      "linked_timestamp": 1728275689.284758,
      "name": "Test Company2 Ltd"
    }
  ],
  "linked": true
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Update Org Oauth App Account

Update app-specific settings for a linked OAuth application account, such as Zoom or Teams guest redaction settings or a Zoom daily API request quota.

```go
UpdateOrgOauthAppAccount(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    accountId string,
    body *models.AccountOauthConfig) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `appName` | [`models.OauthAppNameEnum`](../../doc/models/oauth-app-name-enum.md) | Template, Required | OAuth application name |
| `accountId` | `string` | Template, Required | - |
| `body` | [`*models.AccountOauthConfig`](../../doc/models/account-oauth-config.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

appName := models.OauthAppNameEnum_JAMF

accountId := "iojzXIJWEuiD73ZvydOfg"

body := models.AccountOauthConfig{
    AccountId:            "iojzXIJWEuiD73ZvydOfg",
    MaxDailyApiRequests:  models.ToPointer(5000),
}

resp, err := orgsLinkedApplications.UpdateOrgOauthAppAccount(ctx, orgId, appName, accountId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

