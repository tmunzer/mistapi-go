# Orgs Cradlepoint

```go
orgsCradlepoint := client.OrgsCradlepoint()
```

## Class Name

`OrgsCradlepoint`

## Methods

* [Delete Org Cradlepoint Connection](../../doc/controllers/orgs-cradlepoint.md#delete-org-cradlepoint-connection)
* [Setup Org Cradlepoint Connection to Mist](../../doc/controllers/orgs-cradlepoint.md#setup-org-cradlepoint-connection-to-mist)
* [Sync Org Cradlepoint Routers](../../doc/controllers/orgs-cradlepoint.md#sync-org-cradlepoint-routers)
* [Update Org Cradlepoint Connection to Mist](../../doc/controllers/orgs-cradlepoint.md#update-org-cradlepoint-connection-to-mist)


# Delete Org Cradlepoint Connection

This deletes the Cradlepoint integration in Mist

```go
DeleteOrgCradlepointConnection(
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

resp, err := orgsCradlepoint.DeleteOrgCradlepointConnection(ctx, orgId)
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


# Setup Org Cradlepoint Connection to Mist

This sets up cradlepoint webhooks to send events to Mist

```go
SetupOrgCradlepointConnectionToMist(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountCradlepointConfig) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountCradlepointConfig`](../../doc/models/account-cradlepoint-config.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AccountCradlepointConfig{
    CpApiId:              models.ToPointer("84446d61-2206-4ea5-855a-0043f980be54"),
    CpApiKey:             models.ToPointer("79c329da9893e34099c7d8ad5cb9c941"),
    EcmApiId:             models.ToPointer("73446d61-2206-4ea5-855a-0043f980be62"),
    EcmApiKey:            models.ToPointer("68b329da9893e34099c7d8ad5cb9c9405"),
}

resp, err := orgsCradlepoint.SetupOrgCradlepointConnectionToMist(ctx, orgId, &body)
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


# Sync Org Cradlepoint Routers

This syncs cradlepoint devices with Mist. We’ll also attempt to use the LLDP data from cradlepoint to identify the linkage against Mist Site / Device

```go
SyncOrgCradlepointRouters(
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

resp, err := orgsCradlepoint.SyncOrgCradlepointRouters(ctx, orgId)
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


# Update Org Cradlepoint Connection to Mist

This updates the Cradlepoint integration settings in Mist

```go
UpdateOrgCradlepointConnectionToMist(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountCradlepointConfig) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountCradlepointConfig`](../../doc/models/account-cradlepoint-config.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AccountCradlepointConfig{
    CpApiId:              models.ToPointer("84446d61-2206-4ea5-855a-0043f980be54"),
    CpApiKey:             models.ToPointer("79c329da9893e34099c7d8ad5cb9c941"),
    EcmApiId:             models.ToPointer("73446d61-2206-4ea5-855a-0043f980be62"),
    EcmApiKey:            models.ToPointer("68b329da9893e34099c7d8ad5cb9c9405"),
}

resp, err := orgsCradlepoint.UpdateOrgCradlepointConnectionToMist(ctx, orgId, &body)
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

