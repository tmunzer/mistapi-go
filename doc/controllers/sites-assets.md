# Sites Assets

```go
sitesAssets := client.SitesAssets()
```

## Class Name

`SitesAssets`

## Methods

* [Count Site Assets](../../doc/controllers/sites-assets.md#count-site-assets)
* [Create Site Asset](../../doc/controllers/sites-assets.md#create-site-asset)
* [Delete Site Asset](../../doc/controllers/sites-assets.md#delete-site-asset)
* [Get Site Asset](../../doc/controllers/sites-assets.md#get-site-asset)
* [Get Site Asset Stats](../../doc/controllers/sites-assets.md#get-site-asset-stats)
* [Get Site Assets of Interest](../../doc/controllers/sites-assets.md#get-site-assets-of-interest)
* [Get Site Discovered Asset by Map](../../doc/controllers/sites-assets.md#get-site-discovered-asset-by-map)
* [Import Site Assets](../../doc/controllers/sites-assets.md#import-site-assets)
* [List Site Assets](../../doc/controllers/sites-assets.md#list-site-assets)
* [List Site Assets Stats](../../doc/controllers/sites-assets.md#list-site-assets-stats)
* [List Site Discovered Assets](../../doc/controllers/sites-assets.md#list-site-discovered-assets)
* [Search Site Assets](../../doc/controllers/sites-assets.md#search-site-assets)
* [Update Site Asset](../../doc/controllers/sites-assets.md#update-site-asset)


# Count Site Assets

Count Asset by distinct field

```go
CountSiteAssets(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteAssetsCountDistinctEnum) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteAssetsCountDistinctEnum`](../../doc/models/site-assets-count-distinct-enum.md) | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteAssetsCountDistinctEnum("map_id")

apiResponse, err := sitesAssets.CountSiteAssets(ctx, siteId, &distinct)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Create Site Asset

Create Site Asset

```go
CreateSiteAsset(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Asset) (
    models.ApiResponse[models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Asset`](../../doc/models/asset.md) | Body, Optional | Request Body |

## Response Type

[`models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Asset{
    Mac:          "mac0",
    Name:         "name6",
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:       models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := sitesAssets.CreateSiteAsset(ctx, siteId, &body)
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
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mac": "string",
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "tag_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Delete Site Asset

Delete Site Asset

```go
DeleteSiteAsset(
    ctx context.Context,
    siteId uuid.UUID,
    assetId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `assetId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesAssets.DeleteSiteAsset(ctx, siteId, assetId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Asset

Get Site Asset Details

```go
GetSiteAsset(
    ctx context.Context,
    siteId uuid.UUID,
    assetId uuid.UUID) (
    models.ApiResponse[models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `assetId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesAssets.GetSiteAsset(ctx, siteId, assetId)
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
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mac": "string",
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "tag_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Asset Stats

Get Site Asset Details

```go
GetSiteAssetStats(
    ctx context.Context,
    siteId uuid.UUID,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.StatsAsset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.StatsAsset`](../../doc/models/stats-asset.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

apiResponse, err := sitesAssets.GetSiteAssetStats(ctx, siteId, nil, nil, &duration)
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
  "battery_voltage": 0,
  "eddystone_uid_instance": "string",
  "eddystone_uid_namespace": "string",
  "eddystone_url_url": "string",
  "ibeacon_major": 0,
  "ibeacon_minor": 0,
  "ibeacon_uuid": "1f89bc00-d0af-481b-82fe-a6629259a39f",
  "last_seen": 0,
  "mac": "string",
  "map_id": "09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1",
  "name": "string",
  "rssizones": [
    {
      "id": "480f6eca-6276-4993-bfeb-59cbbbbaaf08",
      "since": 0
    }
  ],
  "x": 0,
  "y": 0,
  "zones": [
    {
      "id": "479f6eca-6276-4993-bfeb-5acbbbbabf08",
      "since": 0
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Assets of Interest

Get a list of BLE beacons that matches Asset or AssetFilter

```go
GetSiteAssetsOfInterest(
    ctx context.Context,
    siteId uuid.UUID,
    duration *string,
    start *int,
    end *int,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.AssetOfInterest],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.AssetOfInterest`](../../doc/models/asset-of-interest.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"





page := 1

limit := 100

apiResponse, err := sitesAssets.GetSiteAssetsOfInterest(ctx, siteId, &duration, nil, nil, &page, &limit)
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
    "_checkpoint_prep": 0,
    "_checkpoint_preparer": 0,
    "_checkpoint_scan": 0,
    "_id": "string",
    "_timestamp": 0,
    "_ttl": 0,
    "ap_mac": "string",
    "beam": 0,
    "by": "string",
    "curr_site": "string",
    "device_name": "string",
    "id": "string",
    "last_seen": 0,
    "mac": "string",
    "manufacture": "string",
    "map_id": "string",
    "name": "string",
    "rssi": 0
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Discovered Asset by Map

Get a list of BLE beacons that we discovered (whether they’re defined as assets or not)

```go
GetSiteDiscoveredAssetByMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    models.ApiResponse[[]models.StatsAsset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.StatsAsset`](../../doc/models/stats-asset.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesAssets.GetSiteDiscoveredAssetByMap(ctx, siteId, mapId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Import Site Assets

Impert Site Assets.

It can be done via a CSV file or a JSON payload.

## CSV File Format

```csv
name,mac
"asset_name",5c5b53010101
```

```go
ImportSiteAssets(
    ctx context.Context,
    siteId uuid.UUID,
    upsert *models.ImportSiteAssetsUpsertEnum,
    file *models.FileWrapper) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `upsert` | [`*models.ImportSiteAssetsUpsertEnum`](../../doc/models/import-site-assets-upsert-enum.md) | Query, Optional | API will replace the assets with same mac if provided `upsert`==`True`, otherwise will report in errors in response. |
| `file` | `*models.FileWrapper` | Form, Optional | CSV file |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upsert := models.ImportSiteAssetsUpsertEnum("False")



resp, err := sitesAssets.ImportSiteAssets(ctx, siteId, &upsert, nil)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Assets

Get List of Site Assets

```go
ListSiteAssets(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := sitesAssets.ListSiteAssets(ctx, siteId, &page, &limit)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Assets Stats

Get List of Site Assets Stats

```go
ListSiteAssetsStats(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.StatsAsset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`[]models.StatsAsset`](../../doc/models/stats-asset.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesAssets.ListSiteAssetsStats(ctx, siteId, &page, &limit, nil, nil, &duration)
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
    "battery_voltage": 0,
    "eddystone_uid_instance": "string",
    "eddystone_uid_namespace": "string",
    "eddystone_url_url": "string",
    "ibeacon_major": 0,
    "ibeacon_minor": 0,
    "ibeacon_uuid": "1f89bc00-d0af-481b-82fe-a6629259a39f",
    "last_seen": 0,
    "mac": "string",
    "map_id": "09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1",
    "name": "string",
    "rssizones": [
      {
        "id": "478f6eca-6276-4993-bfeb-5bcbbbbacf08",
        "since": 0
      }
    ],
    "x": 0,
    "y": 0,
    "zones": [
      {
        "id": "477f6eca-6276-4993-bfeb-5ccbbbbadf08",
        "since": 0
      }
    ]
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Discovered Assets

Get List of Site Discovered BLE Assets that doesn’t match any of the Asset / Assetfilters

```go
ListSiteDiscoveredAssets(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`[]models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesAssets.ListSiteDiscoveredAssets(ctx, siteId, &page, &limit, nil, nil, &duration)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Assets

Assets Search

```go
SearchSiteAssets(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    mapId *uuid.UUID,
    ibeaconUuid *uuid.UUID,
    ibeaconMajor *int,
    ibeaconMinor *int,
    eddystoneUidNamespace *string,
    eddystoneUidInstance *string,
    eddystoneUrl *string,
    deviceName *string,
    by *string,
    name *string,
    apMac *string,
    beam *string,
    rssi *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseStatsAssets],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | - |
| `mapId` | `*uuid.UUID` | Query, Optional | - |
| `ibeaconUuid` | `*uuid.UUID` | Query, Optional | - |
| `ibeaconMajor` | `*int` | Query, Optional | - |
| `ibeaconMinor` | `*int` | Query, Optional | - |
| `eddystoneUidNamespace` | `*string` | Query, Optional | - |
| `eddystoneUidInstance` | `*string` | Query, Optional | - |
| `eddystoneUrl` | `*string` | Query, Optional | - |
| `deviceName` | `*string` | Query, Optional | - |
| `by` | `*string` | Query, Optional | - |
| `name` | `*string` | Query, Optional | - |
| `apMac` | `*string` | Query, Optional | - |
| `beam` | `*string` | Query, Optional | - |
| `rssi` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseStatsAssets`](../../doc/models/response-stats-assets.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





























limit := 100





duration := "10m"

apiResponse, err := sitesAssets.SearchSiteAssets(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 0,
  "limit": 0,
  "next": "string",
  "results": [
    {
      "battery_voltage": 0,
      "eddystone_uid_instance": "string",
      "eddystone_uid_namespace": "string",
      "eddystone_url_url": "string",
      "ibeacon_major": 0,
      "ibeacon_minor": 0,
      "ibeacon_uuid": "1f89bc00-d0af-481b-82fe-a6629259a39f",
      "last_seen": 0,
      "mac": "string",
      "map_id": "09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1",
      "name": "string",
      "rssizones": [
        {
          "id": "476f6eca-6276-4993-bfeb-5dcbbbbaef08",
          "since": 0
        }
      ],
      "x": 0,
      "y": 0,
      "zones": [
        {
          "id": "475f6eca-6276-4993-bfeb-5ecbbbbf6f08",
          "since": 0
        }
      ]
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Site Asset

Update Site Asset

```go
UpdateSiteAsset(
    ctx context.Context,
    siteId uuid.UUID,
    assetId uuid.UUID,
    body *models.Asset) (
    models.ApiResponse[models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `assetId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Asset`](../../doc/models/asset.md) | Body, Optional | Request Body |

## Response Type

[`models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Asset{
    Mac:          "mac0",
    Name:         "name6",
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:       models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := sitesAssets.UpdateSiteAsset(ctx, siteId, assetId, &body)
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
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mac": "string",
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "tag_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

