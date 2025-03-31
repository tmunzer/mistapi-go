# Orgs Antivirus Profiles

```go
orgsAntivirusProfiles := client.OrgsAntivirusProfiles()
```

## Class Name

`OrgsAntivirusProfiles`

## Methods

* [Create Org Antivirus Profile](../../doc/controllers/orgs-antivirus-profiles.md#create-org-antivirus-profile)
* [Delete Org Antivirus Profile](../../doc/controllers/orgs-antivirus-profiles.md#delete-org-antivirus-profile)
* [Get Org Antivirus Profile](../../doc/controllers/orgs-antivirus-profiles.md#get-org-antivirus-profile)
* [List Org Antivirus Profiles](../../doc/controllers/orgs-antivirus-profiles.md#list-org-antivirus-profiles)
* [Update Org Antivirus Profile](../../doc/controllers/orgs-antivirus-profiles.md#update-org-antivirus-profile)


# Create Org Antivirus Profile

Create getOrgServices Antivirus Profile

```go
CreateOrgAntivirusProfile(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Avprofile) (
    models.ApiResponse[models.Avprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Avprofile`](../../doc/models/avprofile.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Avprofile](../../doc/models/avprofile.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Avprofile{
    FallbackAction:       models.ToPointer(models.AvprofileFallbackActionEnum_PERMIT),
    MaxFilesize:          models.ToPointer(10000),
    MimeWhitelist:        []string{
    },
    Name:                 "av-custom",
    Protocols:            []models.AvprofileProtocolEnum{
        models.AvprofileProtocolEnum_HTTP,
    },
    UrlWhitelist:         []string{
    },
}

apiResponse, err := orgsAntivirusProfiles.CreateOrgAntivirusProfile(ctx, orgId, &body)
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
  "fallback_action": "permit",
  "max_filesize": 10000,
  "mime_whitelist": [],
  "name": "av-custom",
  "protocols": [
    "http"
  ],
  "url_whitelist": []
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


# Delete Org Antivirus Profile

DeleteOrgAntivirusProfile

```go
DeleteOrgAntivirusProfile(
    ctx context.Context,
    orgId uuid.UUID,
    avprofileId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `avprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

avprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsAntivirusProfiles.DeleteOrgAntivirusProfile(ctx, orgId, avprofileId)
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


# Get Org Antivirus Profile

Get Org Antivirus Profile

```go
GetOrgAntivirusProfile(
    ctx context.Context,
    orgId uuid.UUID,
    avprofileId uuid.UUID) (
    models.ApiResponse[models.Avprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `avprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Avprofile](../../doc/models/avprofile.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

avprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsAntivirusProfiles.GetOrgAntivirusProfile(ctx, orgId, avprofileId)
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
  "fallback_action": "permit",
  "max_filesize": 10000,
  "mime_whitelist": [],
  "name": "av-custom",
  "protocols": [
    "http"
  ],
  "url_whitelist": []
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


# List Org Antivirus Profiles

Get List of Antivirus Profiles

```go
ListOrgAntivirusProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Avprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Avprofile](../../doc/models/avprofile.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsAntivirusProfiles.ListOrgAntivirusProfiles(ctx, orgId, &limit, &page)
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
    "fallback_action": "permit",
    "max_filesize": 10000,
    "mime_whitelist": [],
    "name": "av-custom",
    "protocols": [
      "http"
    ],
    "url_whitelist": []
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


# Update Org Antivirus Profile

Update Org Antivirus Profile

```go
UpdateOrgAntivirusProfile(
    ctx context.Context,
    orgId uuid.UUID,
    avprofileId uuid.UUID,
    body *models.Avprofile) (
    models.ApiResponse[models.Avprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `avprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Avprofile`](../../doc/models/avprofile.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Avprofile](../../doc/models/avprofile.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

avprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Avprofile{
    FallbackAction:       models.ToPointer(models.AvprofileFallbackActionEnum_PERMIT),
    MaxFilesize:          models.ToPointer(10000),
    MimeWhitelist:        []string{
    },
    Name:                 "av-custom",
    Protocols:            []models.AvprofileProtocolEnum{
        models.AvprofileProtocolEnum_HTTP,
    },
    UrlWhitelist:         []string{
    },
}

apiResponse, err := orgsAntivirusProfiles.UpdateOrgAntivirusProfile(ctx, orgId, avprofileId, &body)
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
  "fallback_action": "permit",
  "max_filesize": 10000,
  "mime_whitelist": [],
  "name": "av-custom",
  "protocols": [
    "http"
  ],
  "url_whitelist": []
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

