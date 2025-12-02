# Orgs SDK Invites

```go
orgsSDKInvites := client.OrgsSDKInvites()
```

## Class Name

`OrgsSDKInvites`

## Methods

* [Activate Sdk Invite](../../doc/controllers/orgs-sdk-invites.md#activate-sdk-invite)
* [Create Sdk Invite](../../doc/controllers/orgs-sdk-invites.md#create-sdk-invite)
* [Get Sdk Invite](../../doc/controllers/orgs-sdk-invites.md#get-sdk-invite)
* [Get Sdk Invite Qr Code](../../doc/controllers/orgs-sdk-invites.md#get-sdk-invite-qr-code)
* [List Sdk Invites](../../doc/controllers/orgs-sdk-invites.md#list-sdk-invites)
* [Revoke Sdk Invite](../../doc/controllers/orgs-sdk-invites.md#revoke-sdk-invite)
* [Send Sdk Invite Email](../../doc/controllers/orgs-sdk-invites.md#send-sdk-invite-email)
* [Send Sdk Invite Sms](../../doc/controllers/orgs-sdk-invites.md#send-sdk-invite-sms)
* [Update Sdk Invite](../../doc/controllers/orgs-sdk-invites.md#update-sdk-invite)


# Activate Sdk Invite

Verify secret

```go
ActivateSdkInvite(
    ctx context.Context,
    secret string,
    body *models.DeviceIdString) (
    models.ApiResponse[models.ResponseMobileVerifySecret],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `secret` | `string` | Template, Required | - |
| `body` | [`*models.DeviceIdString`](../../doc/models/device-id-string.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseMobileVerifySecret](../../doc/models/response-mobile-verify-secret.md).

## Example Usage

```go
ctx := context.Background()

secret := "secret4"

body := models.DeviceIdString{
    DeviceId:             uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
}

apiResponse, err := orgsSDKInvites.ActivateSdkInvite(ctx, secret, &body)
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
  "name": "Macy's",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "secret": "device-specific-secret"
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


# Create Sdk Invite

Create SDK Invite

```go
CreateSdkInvite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Sdkinvite) (
    models.ApiResponse[models.Sdkinvite],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sdkinvite`](../../doc/models/sdkinvite.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sdkinvite](../../doc/models/sdkinvite.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sdkinvite{
    Enabled:              models.ToPointer(true),
    Name:                 "string",
    Quota:                models.ToPointer(0),
    QuotaLimited:         models.ToPointer(true),
}

apiResponse, err := orgsSDKInvites.CreateSdkInvite(ctx, orgId, &body)
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
  "created_time": 1428954000,
  "enabled": true,
  "expire_time": 1428954000,
  "id": "5034b980-b49e-501c-66e0-9de4c38f18a2",
  "name": "Macy's",
  "quota": -1
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


# Get Sdk Invite

Get SDK Invite Details

```go
GetSdkInvite(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID) (
    models.ApiResponse[models.Sdkinvite],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdkinviteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sdkinvite](../../doc/models/sdkinvite.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdkinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSDKInvites.GetSdkInvite(ctx, orgId, sdkinviteId)
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
  "created_time": 1428954000,
  "enabled": true,
  "expire_time": 1428954000,
  "id": "5034b980-b49e-501c-66e0-9de4c38f18a2",
  "name": "Macy's",
  "quota": -1
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


# Get Sdk Invite Qr Code

Revoke SDK Invite

```go
GetSdkInviteQrCode(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID) (
    models.ApiResponse[string],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdkinviteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type string.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdkinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSDKInvites.GetSdkInviteQrCode(ctx, orgId, sdkinviteId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# List Sdk Invites

Get List of Org SDK Invites

```go
ListSdkInvites(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.Sdkinvite],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Sdkinvite](../../doc/models/sdkinvite.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSDKInvites.ListSdkInvites(ctx, orgId)
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
[
  {
    "created_time": 1428954000,
    "enabled": true,
    "expire_time": 1428954000,
    "id": "5034b980-b49e-501c-66e0-9de4c38f18a2",
    "name": "Macy's",
    "quota": -1
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Revoke Sdk Invite

Revoke SDK Invite

```go
RevokeSdkInvite(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdkinviteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdkinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSDKInvites.RevokeSdkInvite(ctx, orgId, sdkinviteId)
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


# Send Sdk Invite Email

Send SDK Invite by Email

```go
SendSdkInviteEmail(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID,
    body *models.EmailString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdkinviteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.EmailString`](../../doc/models/email-string.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdkinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.EmailString{
    Email:                "test@abc.com",
}

resp, err := orgsSDKInvites.SendSdkInviteEmail(ctx, orgId, sdkinviteId, &body)
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


# Send Sdk Invite Sms

Send SDK Invite by SMS

```go
SendSdkInviteSms(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID,
    body *models.SdkInviteSms) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdkinviteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SdkInviteSms`](../../doc/models/sdk-invite-sms.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdkinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SdkInviteSms{
    Number:               "14081234567",
}

resp, err := orgsSDKInvites.SendSdkInviteSms(ctx, orgId, sdkinviteId, &body)
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


# Update Sdk Invite

Update SDK Invite

```go
UpdateSdkInvite(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID,
    body *models.Sdkinvite) (
    models.ApiResponse[models.Sdkinvite],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdkinviteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sdkinvite`](../../doc/models/sdkinvite.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sdkinvite](../../doc/models/sdkinvite.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdkinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sdkinvite{
    Enabled:              models.ToPointer(true),
    Name:                 "name6",
    QuotaLimited:         models.ToPointer(false),
}

apiResponse, err := orgsSDKInvites.UpdateSdkInvite(ctx, orgId, sdkinviteId, &body)
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
  "created_time": 1428954000,
  "enabled": true,
  "expire_time": 1428954000,
  "id": "5034b980-b49e-501c-66e0-9de4c38f18a2",
  "name": "Macy's",
  "quota": -1
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

