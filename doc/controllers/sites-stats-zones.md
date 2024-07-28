# Sites Stats-Zones

```go
sitesStatsZones := client.SitesStatsZones()
```

## Class Name

`SitesStatsZones`

## Methods

* [Get Site Zone Stats](../../doc/controllers/sites-stats-zones.md#get-site-zone-stats)
* [List Site Zones Stats](../../doc/controllers/sites-stats-zones.md#list-site-zones-stats)


# Get Site Zone Stats

Get Detail Zone Stats

```go
GetSiteZoneStats(
    ctx context.Context,
    siteId uuid.UUID,
    zoneType models.ZoneTypeEnum,
    zoneId uuid.UUID) (
    models.ApiResponse[models.ZoneStatsDetails],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `zoneType` | [`models.ZoneTypeEnum`](../../doc/models/zone-type-enum.md) | Template, Required | - |
| `zoneId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ZoneStatsDetails`](../../doc/models/zone-stats-details.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

zoneType := models.ZoneTypeEnum("zones")

zoneId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsZones.GetSiteZoneStats(ctx, siteId, zoneType, zoneId)
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
  "assets": [
    "df8dff06ae90"
  ],
  "client_waits": {
    "avg": 1200,
    "max": 3610,
    "min": 600,
    "p95": 2800
  },
  "clients": [
    "5684dae9ac8b"
  ],
  "id": "8ac84899-32db-6327-334c-9b6d58544cfe",
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "name": "Board Room",
  "num_clients": 80,
  "num_sdkclients": 0,
  "sdkclients": [
    "7e2b463d-c91c-ff7d-f3c0-6eccc6949ff8"
  ]
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


# List Site Zones Stats

Get List of Site Zones Stats

```go
ListSiteZonesStats(
    ctx context.Context,
    siteId uuid.UUID,
    mapId *string) (
    models.ApiResponse[[]models.ZoneStats],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `*string` | Query, Optional | - |

## Response Type

[`[]models.ZoneStats`](../../doc/models/zone-stats.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := "00000000-0000-0000-0000-000000000000"

apiResponse, err := sitesStatsZones.ListSiteZonesStats(ctx, siteId, &mapId)
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
    "assets_waits": {
      "avg": 0,
      "max": 0,
      "min": 0,
      "p95": 0
    },
    "clients_waits": {
      "avg": 1200,
      "max": 3610,
      "min": 600,
      "p95": 2800
    },
    "created_time": 1616625211,
    "id": "123470c7-5d9d-424a-8475-8b344c621234",
    "map_id": "123449d4-d12f-4feb-b40f-5be0e2ae1234",
    "modified_time": 1616625211,
    "name": "Zone A",
    "num_assets": 0,
    "num_clients": 80,
    "num_sdkclients": 10,
    "occupancy_limit": 4,
    "org_id": "1234c1a0-6ef6-11e6-8bbf-02e208b21234",
    "sdkclients_waits": {
      "avg": 1200,
      "max": 3610,
      "min": 600,
      "p95": 2800
    },
    "site_id": "123448e6-6ef6-11e6-8bbf-02e208b21234",
    "vertices": [
      {
        "x": 732,
        "y": 1821
      },
      {
        "x": 732.5,
        "y": 1731
      },
      {
        "x": 837.5,
        "y": 1731.5
      },
      {
        "x": 839,
        "y": 1821
      }
    ],
    "vertices_m": [
      {
        "x": 24.1983341951072,
        "y": 60.198314985369144
      },
      {
        "x": 24.21486311190714,
        "y": 57.22310996138056
      },
      {
        "x": 27.685935639893827,
        "y": 57.23963887818049
      },
      {
        "x": 27.73552239029364,
        "y": 60.198314985369144
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

