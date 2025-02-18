# Orgs SDK Templates

```go
orgsSDKTemplates := client.OrgsSDKTemplates()
```

## Class Name

`OrgsSDKTemplates`

## Methods

* [Create Sdk Template](../../doc/controllers/orgs-sdk-templates.md#create-sdk-template)
* [Delete Sdk Template](../../doc/controllers/orgs-sdk-templates.md#delete-sdk-template)
* [Get Sdk Template](../../doc/controllers/orgs-sdk-templates.md#get-sdk-template)
* [List Sdk Templates](../../doc/controllers/orgs-sdk-templates.md#list-sdk-templates)
* [Update Sdk Template](../../doc/controllers/orgs-sdk-templates.md#update-sdk-template)


# Create Sdk Template

Create SDK Template

```go
CreateSdkTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Sdktemplate) (
    models.ApiResponse[models.Sdktemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sdktemplate`](../../doc/models/sdktemplate.md) | Body, Optional | Request Body |

## Response Type

[`models.Sdktemplate`](../../doc/models/sdktemplate.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sdktemplate{
    Name:                 "name6",
}

apiResponse, err := orgsSDKTemplates.CreateSdkTemplate(ctx, orgId, &body)
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
  "bg_image": "http://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg",
  "btn_flr_bgcolor": "#282828",
  "default": true,
  "header_txt": "Mist",
  "name": "default",
  "search_txtcolor": "#282828",
  "welcome_msg": "Welcome to Mist"
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


# Delete Sdk Template

Delete SDK Template

```go
DeleteSdkTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sdktemplateId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdktemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdktemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSDKTemplates.DeleteSdkTemplate(ctx, orgId, sdktemplateId)
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


# Get Sdk Template

Get SDK Template Details

```go
GetSdkTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sdktemplateId uuid.UUID) (
    models.ApiResponse[models.Sdktemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdktemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Sdktemplate`](../../doc/models/sdktemplate.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdktemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSDKTemplates.GetSdkTemplate(ctx, orgId, sdktemplateId)
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
  "bg_image": "http://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg",
  "btn_flr_bgcolor": "#282828",
  "default": true,
  "header_txt": "Mist",
  "name": "default",
  "search_txtcolor": "#282828",
  "welcome_msg": "Welcome to Mist"
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


# List Sdk Templates

Get List of Org SDK Templates

```go
ListSdkTemplates(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.Sdktemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.Sdktemplate`](../../doc/models/sdktemplate.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSDKTemplates.ListSdkTemplates(ctx, orgId)
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
    "bg_image": "http://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg",
    "btn_flr_bgcolor": "#282828",
    "default": true,
    "header_txt": "Mist",
    "name": "default",
    "search_txtcolor": "#282828",
    "welcome_msg": "Welcome to Mist"
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


# Update Sdk Template

Update SDK Template

```go
UpdateSdkTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sdktemplateId uuid.UUID,
    body *models.Sdktemplate) (
    models.ApiResponse[models.Sdktemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sdktemplateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sdktemplate`](../../doc/models/sdktemplate.md) | Body, Optional | Request Body |

## Response Type

[`models.Sdktemplate`](../../doc/models/sdktemplate.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdktemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sdktemplate{
    Name:                 "name6",
}

apiResponse, err := orgsSDKTemplates.UpdateSdkTemplate(ctx, orgId, sdktemplateId, &body)
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
  "bg_image": "http://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg",
  "btn_flr_bgcolor": "#282828",
  "default": true,
  "header_txt": "Mist",
  "name": "default",
  "search_txtcolor": "#282828",
  "welcome_msg": "Welcome to Mist"
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

