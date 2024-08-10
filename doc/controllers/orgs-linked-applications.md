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
* [Link Org to Juniper Juniper Account](../../doc/controllers/orgs-linked-applications.md#link-org-to-juniper-juniper-account)
* [Unlink Org From Juniper Customer Id](../../doc/controllers/orgs-linked-applications.md#unlink-org-from-juniper-customer-id)
* [Update Org Oauth App Accounts](../../doc/controllers/orgs-linked-applications.md#update-org-oauth-app-accounts)


# Add Org Oauth App Accounts

Add Jamf, VMware Authorization With Mist Portal

```go
AddOrgOauthAppAccounts(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    body *models.AccountOauthAdd) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `appName` | [`models.OauthAppNameEnum`](../../doc/models/oauth-app-name-enum.md) | Template, Required | OAuth application name |
| `body` | [`*models.AccountOauthAdd`](../../doc/models/containers/account-oauth-add.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

appName := models.OauthAppNameEnum("mobicontrol")

body := models.AccountOauthAddContainer.FromAccountJamfConfig(models.AccountJamfConfig{
    ClientId:       "client_id0",
    ClientSecret:   "client_secret6",
    InstanceUrl:    "junipertest.jamfcloud.com",
    SmartgroupName: "CompliantGroup1",
})

resp, err := orgsLinkedApplications.AddOrgOauthAppAccounts(ctx, orgId, appName, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Unsuccessful | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Oauth App Authorization

Delete Org Level OAuth Application Authorization With Mist Portal

```go
DeleteOrgOauthAppAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    accountId string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `appName` | [`models.OauthAppNameEnum`](../../doc/models/oauth-app-name-enum.md) | Template, Required | OAuth application name |
| `accountId` | `string` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

appName := models.OauthAppNameEnum("mobicontrol")

accountId := "iojzXIJWEuiD73ZvydOfg"

resp, err := orgsLinkedApplications.DeleteOrgOauthAppAuthorization(ctx, orgId, appName, accountId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Unsuccessful | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Oauth App Linked Status

Get Org Level OAuth Application Linked Status

```go
GetOrgOauthAppLinkedStatus(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    forward string) (
    models.ApiResponse[models.ResponseOauthAppLink],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `appName` | [`models.OauthAppNameEnum`](../../doc/models/oauth-app-name-enum.md) | Template, Required | OAuth application name |
| `forward` | `string` | Query, Required | Mist portal url to which backend needs to redirect after succesful OAuth authorization. **Required** to get the `authorization_url` |

## Response Type

[`models.ResponseOauthAppLink`](../../doc/models/response-oauth-app-link.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

appName := models.OauthAppNameEnum("mobicontrol")

forward := "forward6"

apiResponse, err := orgsLinkedApplications.GetOrgOauthAppLinkedStatus(ctx, orgId, appName, forward)
if err != nil {
    log.Fatalln(err)
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
      "error": "OAuth token refresh failed, please re-link your account",
      "instance_url": "junipertest.jamfcloud.com",
      "last_status": "failed",
      "last_sync": 1665465339000,
      "linked_by": "Testname1",
      "name": "Test Compay1 Ltd",
      "smartgroup_name": "CompliantGroup1"
    }
  ],
  "linked": true
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Link Org to Juniper Juniper Account

Link Juniper Accounts

```go
LinkOrgToJuniperJuniperAccount(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJuniperConfig) (
    models.ApiResponse[models.AccountJuniperInfo],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountJuniperConfig`](../../doc/models/account-juniper-config.md) | Body, Optional | - |

## Response Type

[`models.AccountJuniperInfo`](../../doc/models/account-juniper-info.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AccountJuniperConfig{
    Password: "password",
    Username: "john@nmo.com",
}

apiResponse, err := orgsLinkedApplications.LinkOrgToJuniperJuniperAccount(ctx, orgId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "account": {
    "linked_by": "John Smith (john@abccorp.com)",
    "name": "ABC Corp"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | account already linked | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Unlink Org From Juniper Customer Id

Unlink Juniper Customer ID
`linked_by` field is only required if there are duplicate account_names.

```go
UnlinkOrgFromJuniperCustomerId(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJuniperInfo) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountJuniperInfo`](../../doc/models/account-juniper-info.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



resp, err := orgsLinkedApplications.UnlinkOrgFromJuniperCustomerId(ctx, orgId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Oauth App Accounts

Update Zoom, Teams, Intune Authorization.

Request Payload, These Field And Values Will Be Specific To Each Of The Third Party Apps Accounts.

```go
UpdateOrgOauthAppAccounts(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    body *models.AccountOauthConfig) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `appName` | [`models.OauthAppNameEnum`](../../doc/models/oauth-app-name-enum.md) | Template, Required | OAuth application name |
| `body` | [`*models.AccountOauthConfig`](../../doc/models/account-oauth-config.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

appName := models.OauthAppNameEnum("mobicontrol")

body := models.AccountOauthConfig{
    AccountId:           "iojzXIJWEuiD73ZvydOfg",
    MaxDailyApiRequests: models.ToPointer(5000),
}

resp, err := orgsLinkedApplications.UpdateOrgOauthAppAccounts(ctx, orgId, appName, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

