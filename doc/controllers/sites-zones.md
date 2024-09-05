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
* [List Site Zones](../../doc/controllers/sites-zones.md#list-site-zones)
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
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
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
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

zoneType := models.ZoneTypeEnum("rssizones")

distinct := models.SiteZoneCountDistinctEnum("scope_id")







scope := models.ZoneScopeEnum("site")





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesZones.CountSiteZoneSessions(ctx, siteId, zoneType, &distinct, nil, nil, nil, &scope, nil, nil, &duration, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Site Zones

Get List of Site Zones

```go
ListSiteZones(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Zone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Zone`](../../doc/models/zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesZones.ListSiteZones(ctx, siteId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
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
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`models.ResponseZoneSearch`](../../doc/models/response-zone-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

zoneType := models.ZoneTypeEnum("rssizones")







scope := models.VisitsScopeEnum("site")





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesZones.SearchSiteZoneSessions(ctx, siteId, zoneType, nil, nil, nil, &scope, nil, nil, &duration, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

