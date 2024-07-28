# Sites Asset Filters

```go
sitesAssetFilters := client.SitesAssetFilters()
```

## Class Name

`SitesAssetFilters`

## Methods

* [Create Site Asset Filters](../../doc/controllers/sites-asset-filters.md#create-site-asset-filters)
* [Delete Site Asset Filter](../../doc/controllers/sites-asset-filters.md#delete-site-asset-filter)
* [Get Site Asset Filter](../../doc/controllers/sites-asset-filters.md#get-site-asset-filter)
* [List Site Asset Filters](../../doc/controllers/sites-asset-filters.md#list-site-asset-filters)
* [Update Site Asset Filter](../../doc/controllers/sites-asset-filters.md#update-site-asset-filter)


# Create Site Asset Filters

Create Site Asset Filter

```go
CreateSiteAssetFilters(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AssetFilter) (
    models.ApiResponse[models.AssetFilter],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AssetFilter`](../../doc/models/asset-filter.md) | Body, Optional | Request Body |

## Response Type

[`models.AssetFilter`](../../doc/models/asset-filter.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AssetFilter{
    Disabled:              models.ToPointer(false),
    EddystoneUidNamespace: models.ToPointer("2818e3868dec25629ede"),
    EddystoneUrl:          models.ToPointer("https://www.abc.com"),
    IbeaconMajor:          models.ToPointer(13),
    IbeaconUuid:           models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3")),
    MfgCompanyId:          models.ToPointer(935),
    Name:                  "Visitor Tags",
    OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    ServiceUuid:           models.ToPointer(uuid.MustParse("0000fe6a-0000-1000-8000-0030459b3cfb")),
    SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := sitesAssetFilters.CreateSiteAssetFilters(ctx, siteId, &body)
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
  "created_time": 0,
  "disasbled": true,
  "eddystone_uid_namespace": "string",
  "eddystone_url": "string",
  "for_site": true,
  "ibeacon_major": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mfg_company_id": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# Delete Site Asset Filter

Deletes an existing BLE asset filter for the given site.

```go
DeleteSiteAssetFilter(
    ctx context.Context,
    siteId uuid.UUID,
    assetfilterId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `assetfilterId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetfilterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesAssetFilters.DeleteSiteAssetFilter(ctx, siteId, assetfilterId)
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


# Get Site Asset Filter

Get Site Asset Filter Details

```go
GetSiteAssetFilter(
    ctx context.Context,
    siteId uuid.UUID,
    assetfilterId uuid.UUID) (
    models.ApiResponse[models.AssetFilter],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `assetfilterId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.AssetFilter`](../../doc/models/asset-filter.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetfilterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesAssetFilters.GetSiteAssetFilter(ctx, siteId, assetfilterId)
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
  "created_time": 0,
  "disasbled": true,
  "eddystone_uid_namespace": "string",
  "eddystone_url": "string",
  "for_site": true,
  "ibeacon_major": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mfg_company_id": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# List Site Asset Filters

Get List of Site Asset Filters

```go
ListSiteAssetFilters(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.AssetFilter],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.AssetFilter`](../../doc/models/asset-filter.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := sitesAssetFilters.ListSiteAssetFilters(ctx, siteId, &page, &limit)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Site Asset Filter

Updates an existing BLE asset filter for the given site.

```go
UpdateSiteAssetFilter(
    ctx context.Context,
    siteId uuid.UUID,
    assetfilterId uuid.UUID,
    body *models.AssetFilter) (
    models.ApiResponse[models.AssetFilter],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `assetfilterId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AssetFilter`](../../doc/models/asset-filter.md) | Body, Optional | Request Body |

## Response Type

[`models.AssetFilter`](../../doc/models/asset-filter.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetfilterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AssetFilter{
    Disabled:              models.ToPointer(false),
    EddystoneUidNamespace: models.ToPointer("2818e3868dec25629ede"),
    EddystoneUrl:          models.ToPointer("https://www.abc.com"),
    IbeaconMajor:          models.ToPointer(13),
    IbeaconUuid:           models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3")),
    MfgCompanyId:          models.ToPointer(935),
    Name:                  "Visitor Tags",
    OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    ServiceUuid:           models.ToPointer(uuid.MustParse("0000fe6a-0000-1000-8000-0030459b3cfb")),
    SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := sitesAssetFilters.UpdateSiteAssetFilter(ctx, siteId, assetfilterId, &body)
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
  "created_time": 0,
  "disasbled": true,
  "eddystone_uid_namespace": "string",
  "eddystone_url": "string",
  "for_site": true,
  "ibeacon_major": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mfg_company_id": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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

