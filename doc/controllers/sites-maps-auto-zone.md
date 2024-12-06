# Sites Maps-Auto-Zone

```go
sitesMapsAutoZone := client.SitesMapsAutoZone()
```

## Class Name

`SitesMapsAutoZone`

## Methods

* [Delete Site Map Auto Zone](../../doc/controllers/sites-maps-auto-zone.md#delete-site-map-auto-zone)
* [Get Site Map Auto Zone Status](../../doc/controllers/sites-maps-auto-zone.md#get-site-map-auto-zone-status)
* [Start Site Map Auto Zone](../../doc/controllers/sites-maps-auto-zone.md#start-site-map-auto-zone)


# Delete Site Map Auto Zone

This API starts the auto zones service for a specified map. This map must have an image to parse for the auto zones service. Repeated POST requests to this endpoint while the auto zones service is proccessing the map or awaiting review will be rejected.

```go
DeleteSiteMapAutoZone(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMapsAutoZone.DeleteSiteMapAutoZone(ctx, mapId, siteId)
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


# Get Site Map Auto Zone Status

This API provides the current status of the auto zones service for a given map

```go
GetSiteMapAutoZoneStatus(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID) (
    models.ApiResponse[models.ResponseAutoZone],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseAutoZone`](../../doc/models/response-auto-zone.md)

## Example Usage

```go
ctx := context.Background()

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesMapsAutoZone.GetSiteMapAutoZoneStatus(ctx, mapId, siteId)
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
  "status": "awaiting_review",
  "zones": [
    {
      "name": "zone1",
      "vertices": [
        {
          "x": 0,
          "y": 0
        },
        {
          "x": 0,
          "y": 10
        },
        {
          "x": 10,
          "y": 10
        },
        {
          "x": 10,
          "y": 0
        }
      ]
    },
    {
      "name": "zone2",
      "vertices": [
        {
          "x": 0,
          "y": 0
        },
        {
          "x": 0,
          "y": 20
        },
        {
          "x": 20,
          "y": 20
        },
        {
          "x": 20,
          "y": 0
        }
      ]
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


# Start Site Map Auto Zone

This API starts the auto zones service for a specified map. This map must have an image to parse for the auto zones service. Reppeated POST requests to this endpoint while the auto zones service is proccessing the map will be rejected.

```go
StartSiteMapAutoZone(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMapsAutoZone.StartSiteMapAutoZone(ctx, mapId, siteId)
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

