# Sites RSSI Zones

```go
sitesRSSIZones := client.SitesRSSIZones()
```

## Class Name

`SitesRSSIZones`

## Methods

* [Create Site Rssi Zone](../../doc/controllers/sites-rssi-zones.md#create-site-rssi-zone)
* [Delete Site Rssi Zone](../../doc/controllers/sites-rssi-zones.md#delete-site-rssi-zone)
* [Get Site Rssi Zone](../../doc/controllers/sites-rssi-zones.md#get-site-rssi-zone)
* [List Site Rssi Zones](../../doc/controllers/sites-rssi-zones.md#list-site-rssi-zones)
* [Update Site Rssi Zone](../../doc/controllers/sites-rssi-zones.md#update-site-rssi-zone)


# Create Site Rssi Zone

Create RSSI Zone

```go
CreateSiteRssiZone(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.RssiZone) (
    models.ApiResponse[models.RssiZone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.RssiZone`](../../doc/models/rssi-zone.md) | Body, Optional | Request Body |

## Response Type

[`models.RssiZone`](../../doc/models/rssi-zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.RssiZone{
    Devices:      []models.RssiZoneDevice{
        models.RssiZoneDevice{
            DeviceId: uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            Rssi:     0,
        },
    },
    Name:         models.ToPointer("string"),
}

apiResponse, err := sitesRSSIZones.CreateSiteRssiZone(ctx, siteId, &body)
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
  "devices": [
    {
      "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "rssi": 0
    }
  ],
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# Delete Site Rssi Zone

Delete Site RSSI Zone

```go
DeleteSiteRssiZone(
    ctx context.Context,
    siteId uuid.UUID,
    rssizoneId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rssizoneId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rssizoneId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesRSSIZones.DeleteSiteRssiZone(ctx, siteId, rssizoneId)
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


# Get Site Rssi Zone

Get Site RSSI Zone details

```go
GetSiteRssiZone(
    ctx context.Context,
    siteId uuid.UUID,
    rssizoneId uuid.UUID) (
    models.ApiResponse[[]models.RssiZone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rssizoneId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.RssiZone`](../../doc/models/rssi-zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rssizoneId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesRSSIZones.GetSiteRssiZone(ctx, siteId, rssizoneId)
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
    "devices": [
      {
        "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
        "rssi": 0
      }
    ],
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# List Site Rssi Zones

Get List of Site RSSI Zone (RSSI-based)

```go
ListSiteRssiZones(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.RssiZone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.RssiZone`](../../doc/models/rssi-zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesRSSIZones.ListSiteRssiZones(ctx, siteId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Site Rssi Zone

Update Site RSSI Zone

```go
UpdateSiteRssiZone(
    ctx context.Context,
    siteId uuid.UUID,
    rssizoneId uuid.UUID,
    body *models.RssiZone) (
    models.ApiResponse[models.RssiZone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rssizoneId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.RssiZone`](../../doc/models/rssi-zone.md) | Body, Optional | Request Body |

## Response Type

[`models.RssiZone`](../../doc/models/rssi-zone.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rssizoneId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.RssiZone{
    Devices:      []models.RssiZoneDevice{
        models.RssiZoneDevice{
            DeviceId: uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            Rssi:     0,
        },
    },
    Name:         models.ToPointer("string"),
}

apiResponse, err := sitesRSSIZones.UpdateSiteRssiZone(ctx, siteId, rssizoneId, &body)
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
  "devices": [
    {
      "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "rssi": 0
    }
  ],
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

