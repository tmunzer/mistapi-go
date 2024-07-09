# Sites Zones

```go
sitesZones := client.SitesZones()
```

## Class Name

`SitesZones`

## Methods

* [Count Site Zone Sessions](../../doc/controllers/sites-zones.md#count-site-zone-sessions)
* [Create Site Zone](../../doc/controllers/sites-zones.md#create-site-zone)
* [Delete Site Zone](../../doc/controllers/sites-zones.md#delete-site-zone)
* [Get Site Zone](../../doc/controllers/sites-zones.md#get-site-zone)
* [Get Site Zone Stats](../../doc/controllers/sites-zones.md#get-site-zone-stats)
* [List Site Zones](../../doc/controllers/sites-zones.md#list-site-zones)
* [List Site Zones Stats](../../doc/controllers/sites-zones.md#list-site-zones-stats)
* [Search Site Zone Sessions](../../doc/controllers/sites-zones.md#search-site-zone-sessions)
* [Update Site Zone](../../doc/controllers/sites-zones.md#update-site-zone)


# Count Site Zone Sessions

Count Site Zone Sessions

```go
CountSiteZoneSessions(
    ctx context.Context,
    siteId uuid.UUID,
    zoneType models.ZoneTypeEnum,
    distinct *models.SiteZoneCountDistinctEnum,
    userType *models.RfClientTypeEnum,
    user *string,
    scopeId *string,
    scope *models.ZoneScopeEnum,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `zoneType` | [`models.ZoneTypeEnum`](../../doc/models/zone-type-enum.md) | Template, Required | - |
| `distinct` | [`*models.SiteZoneCountDistinctEnum`](../../doc/models/site-zone-count-distinct-enum.md) | Query, Optional | - |
| `userType` | [`*models.RfClientTypeEnum`](../../doc/models/rf-client-type-enum.md) | Query, Optional | user type |
| `user` | `*string` | Query, Optional | client MAC / Asset MAC / SDK UUID |
| `scopeId` | `*string` | Query, Optional | if `scope`==`map`/`zone`/`rssizone`, the scope id |
| `scope` | [`*models.ZoneScopeEnum`](../../doc/models/zone-scope-enum.md) | Query, Optional | scope |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

zoneType := models.ZoneTypeEnum("zones")

distinct := models.SiteZoneCountDistinctEnum("scope_id")







scope := models.ZoneScopeEnum("site")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesZones.CountSiteZoneSessions(ctx, siteId, zoneType, &distinct, nil, nil, nil, &scope, &page, &limit, nil, nil, &duration)
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


# Create Site Zone

Create Site Zone

```go
CreateSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Zone) (
    models.ApiResponse[models.Zone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Zone`](../../doc/models/zone.md) | Body, Optional | Request Body |

## Response Type

[`models.Zone`](../../doc/models/zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Zone{
    Name:         models.ToPointer("string"),
    Vertices:     []models.ZoneVertex{
        models.ZoneVertex{
            X: float64(0),
            Y: float64(0),
        },
    },
}

apiResponse, err := sitesZones.CreateSiteZone(ctx, siteId, &body)
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
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "vertices": [
    {
      "x": 0,
      "y": 0
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


# Delete Site Zone

Delete Site Zone

```go
DeleteSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    zoneId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `zoneId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

zoneId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesZones.DeleteSiteZone(ctx, siteId, zoneId)
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


# Get Site Zone

Get Site Zone Details

```go
GetSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    zoneId uuid.UUID) (
    models.ApiResponse[models.Zone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `zoneId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Zone`](../../doc/models/zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

zoneId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesZones.GetSiteZone(ctx, siteId, zoneId)
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
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "vertices": [
    {
      "x": 0,
      "y": 0
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

apiResponse, err := sitesZones.GetSiteZoneStats(ctx, siteId, zoneType, zoneId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Zones

Get List of Site Zones

```go
ListSiteZones(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Zone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Zone`](../../doc/models/zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := sitesZones.ListSiteZones(ctx, siteId, &page, &limit)
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
    "created_time": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "vertices": [
      {
        "x": 0,
        "y": 0
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

apiResponse, err := sitesZones.ListSiteZonesStats(ctx, siteId, &mapId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Zone Sessions

Search Zone Sessions

```go
SearchSiteZoneSessions(
    ctx context.Context,
    siteId uuid.UUID,
    zoneType models.ZoneTypeEnum,
    userType *models.RfClientTypeEnum,
    user *string,
    scopeId *string,
    scope *models.VisitsScopeEnum,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseZoneSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `zoneType` | [`models.ZoneTypeEnum`](../../doc/models/zone-type-enum.md) | Template, Required | - |
| `userType` | [`*models.RfClientTypeEnum`](../../doc/models/rf-client-type-enum.md) | Query, Optional | user type, client (default) / sdkclient / asset |
| `user` | `*string` | Query, Optional | client MAC / Asset MAC / SDK UUID |
| `scopeId` | `*string` | Query, Optional | if `scope`==`map`/`zone`/`rssizone`, the scope id |
| `scope` | [`*models.VisitsScopeEnum`](../../doc/models/visits-scope-enum.md) | Query, Optional | scope |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseZoneSearch`](../../doc/models/response-zone-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

zoneType := models.ZoneTypeEnum("zones")







scope := models.VisitsScopeEnum("site")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesZones.SearchSiteZoneSessions(ctx, siteId, zoneType, nil, nil, nil, &scope, &page, &limit, nil, nil, &duration)
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
  "end": 1541705289.769911,
  "limit": 1,
  "next": "/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/zones/visits/search?limit=2&end=1541705247.000&scope_id=85fbba9e-4e12-11e6-9188-0242ac110007&user_type=asset&start=1541618889.77",
  "results": [
    {
      "enter": 1541705254,
      "scope": "map",
      "timestamp": 1541705254,
      "user": "c4b301c81166"
    }
  ],
  "start": 1541618889.769886,
  "total": 5892
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


# Update Site Zone

Update Site Zone

```go
UpdateSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    zoneId uuid.UUID,
    body *models.Zone) (
    models.ApiResponse[models.Zone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `zoneId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Zone`](../../doc/models/zone.md) | Body, Optional | Request Body |

## Response Type

[`models.Zone`](../../doc/models/zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

zoneId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Zone{
    Name:         models.ToPointer("string"),
    Vertices:     []models.ZoneVertex{
        models.ZoneVertex{
            X: float64(0),
            Y: float64(0),
        },
    },
}

apiResponse, err := sitesZones.UpdateSiteZone(ctx, siteId, zoneId, &body)
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
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "vertices": [
    {
      "x": 0,
      "y": 0
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

