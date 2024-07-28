# Orgs Stats-Assets

```go
orgsStatsAssets := client.OrgsStatsAssets()
```

## Class Name

`OrgsStatsAssets`

## Methods

* [Count Org Assets by Distance Field](../../doc/controllers/orgs-stats-assets.md#count-org-assets-by-distance-field)
* [List Org Assets Stats](../../doc/controllers/orgs-stats-assets.md#list-org-assets-stats)
* [Search Org Assets](../../doc/controllers/orgs-stats-assets.md#search-org-assets)


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



apiResponse, err := orgsStatsAssets.CountOrgAssetsByDistanceField(ctx, orgId, nil)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
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

apiResponse, err := orgsStatsAssets.ListOrgAssetsStats(ctx, orgId, &page, &limit, nil, nil, &duration)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
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

apiResponse, err := orgsStatsAssets.SearchOrgAssets(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
