# Orgs Assets

```go
orgsAssets := client.OrgsAssets()
```

## Class Name

`OrgsAssets`

## Methods

* [Count Org Assets by Distance Field](../../doc/controllers/orgs-assets.md#count-org-assets-by-distance-field)
* [Create Org Asset](../../doc/controllers/orgs-assets.md#create-org-asset)
* [Delete Org Asset](../../doc/controllers/orgs-assets.md#delete-org-asset)
* [Get Org Asset](../../doc/controllers/orgs-assets.md#get-org-asset)
* [Import Org Assets](../../doc/controllers/orgs-assets.md#import-org-assets)
* [List Org Assets](../../doc/controllers/orgs-assets.md#list-org-assets)
* [List Org Assets Stats](../../doc/controllers/orgs-assets.md#list-org-assets-stats)
* [Search Org Assets](../../doc/controllers/orgs-assets.md#search-org-assets)
* [Update Org Asset](../../doc/controllers/orgs-assets.md#update-org-asset)


# Count Org Assets by Distance Field

Count Org Assets

```go
CountOrgAssetsByDistanceField(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgAssetCountDistinctEnum) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgAssetCountDistinctEnum`](../../doc/models/org-asset-count-distinct-enum.md) | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := orgsAssets.CountOrgAssetsByDistanceField(ctx, orgId, nil)
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


# Create Org Asset

Create Org Asset

```go
CreateOrgAsset(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Asset) (
    models.ApiResponse[models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Asset`](../../doc/models/asset.md) | Body, Optional | Request Body |

## Response Type

[`models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Asset{
    Mac:          "mac0",
    Name:         "name6",
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:       models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := orgsAssets.CreateOrgAsset(ctx, orgId, &body)
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


# Delete Org Asset

Delete Org Asset

```go
DeleteOrgAsset(
    ctx context.Context,
    orgId uuid.UUID,
    assetId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `assetId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsAssets.DeleteOrgAsset(ctx, orgId, assetId)
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


# Get Org Asset

Get Org Asset Details

```go
GetOrgAsset(
    ctx context.Context,
    orgId uuid.UUID,
    assetId uuid.UUID) (
    models.ApiResponse[models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `assetId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsAssets.GetOrgAsset(ctx, orgId, assetId)
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


# Import Org Assets

Impert Org Assets.

It can be done via a CSV file or a JSON payload.

#### CSV File Format

```csv
name,mac
"asset_name",5c5b53010101
```

```go
ImportOrgAssets(
    ctx context.Context,
    orgId uuid.UUID,
    file *models.FileWrapper) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `file` | `*models.FileWrapper` | Form, Optional | CSV file |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



resp, err := orgsAssets.ImportOrgAssets(ctx, orgId, nil)
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


# List Org Assets

Get List of Org Assets

```go
ListOrgAssets(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsAssets.ListOrgAssets(ctx, orgId, &page, &limit)
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


# List Org Assets Stats

Get List of Org Assets Stats

```go
ListOrgAssetsStats(
    ctx context.Context,
    orgId uuid.UUID,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
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

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsAssets.ListOrgAssetsStats(ctx, orgId, &page, &limit, nil, nil, &duration)
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


# Search Org Assets

Search for Org Assets

```go
SearchOrgAssets(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    mac *string,
    deviceName *string,
    name *string,
    mapId *string,
    ibeaconUuid *string,
    ibeaconMajor *string,
    ibeaconMinor *string,
    eddystoneUidNamespace *string,
    eddystoneUidInstance *string,
    eddystoneUrl *string,
    apMac *string,
    beam *int,
    rssi *int,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `deviceName` | `*string` | Query, Optional | - |
| `name` | `*string` | Query, Optional | - |
| `mapId` | `*string` | Query, Optional | - |
| `ibeaconUuid` | `*string` | Query, Optional | - |
| `ibeaconMajor` | `*string` | Query, Optional | - |
| `ibeaconMinor` | `*string` | Query, Optional | - |
| `eddystoneUidNamespace` | `*string` | Query, Optional | - |
| `eddystoneUidInstance` | `*string` | Query, Optional | - |
| `eddystoneUrl` | `*string` | Query, Optional | - |
| `apMac` | `*string` | Query, Optional | - |
| `beam` | `*int` | Query, Optional | - |
| `rssi` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseStatsAssets`](../../doc/models/response-stats-assets.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





























limit := 100





duration := "10m"

apiResponse, err := orgsAssets.SearchOrgAssets(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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


# Update Org Asset

Update Org Asset

```go
UpdateOrgAsset(
    ctx context.Context,
    orgId uuid.UUID,
    assetId uuid.UUID,
    body *models.Asset) (
    models.ApiResponse[models.Asset],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `assetId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Asset`](../../doc/models/asset.md) | Body, Optional | Request Body |

## Response Type

[`models.Asset`](../../doc/models/asset.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Asset{
    Mac:          "mac0",
    Name:         "name6",
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:       models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := orgsAssets.UpdateOrgAsset(ctx, orgId, assetId, &body)
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

