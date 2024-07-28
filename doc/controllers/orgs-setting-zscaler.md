# Orgs Setting Zscaler

```go
orgsSettingZscaler := client.OrgsSettingZscaler()
```

## Class Name

`OrgsSettingZscaler`

## Methods

* [Delete Org Zscaler Credential](../../doc/controllers/orgs-setting-zscaler.md#delete-org-zscaler-credential)
* [Get Org Zscaler Credential](../../doc/controllers/orgs-setting-zscaler.md#get-org-zscaler-credential)
* [Setup Org Zscaler Credential](../../doc/controllers/orgs-setting-zscaler.md#setup-org-zscaler-credential)


# Delete Org Zscaler Credential

To delete Zscaler credential

```go
DeleteOrgZscalerCredential(
    ctx context.Context,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSettingZscaler.DeleteOrgZscalerCredential(ctx, orgId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Org Zscaler Credential

To get Zscaler credential

```go
GetOrgZscalerCredential(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.AccountZscalerInfo],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.AccountZscalerInfo`](../../doc/models/account-zscaler-info.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSettingZscaler.GetOrgZscalerCredential(ctx, orgId)
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
  "cloud_name": "zscalerbeta.net",
  "partner_key": "K35vrZcK3JvrZc",
  "username": "john@nmo.com"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Setup Org Zscaler Credential

To setup Zscaler credential

```go
SetupOrgZscalerCredential(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountZscalerConfig) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountZscalerConfig`](../../doc/models/account-zscaler-config.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AccountZscalerConfig{
    CloudName:  "zscalerbeta.net",
    PartnerKey: "K35vrZcK3JvrZc",
    Password:   "password",
    Username:   "john@nmo.com",
}

resp, err := orgsSettingZscaler.SetupOrgZscalerCredential(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

