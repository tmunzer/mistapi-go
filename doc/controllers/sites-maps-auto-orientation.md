# Sites Maps-Auto-Orientation

```go
sitesMapsAutoOrientation := client.SitesMapsAutoOrientation()
```

## Class Name

`SitesMapsAutoOrientation`

## Methods

* [Clear Site Ap Auto Orient](../../doc/controllers/sites-maps-auto-orientation.md#clear-site-ap-auto-orient)
* [Delete Site Ap Auto Orientation](../../doc/controllers/sites-maps-auto-orientation.md#delete-site-ap-auto-orientation)
* [Start Site Ap Auto Orientation](../../doc/controllers/sites-maps-auto-orientation.md#start-site-ap-auto-orientation)


# Clear Site Ap Auto Orient

This API is used to destroy the autoorientations of a map or subset of APs on a map.

```go
ClearSiteApAutoOrient(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.MacAddresses) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "683b679ac024",
    },
}

resp, err := sitesMapsAutoOrientation.ClearSiteApAutoOrient(ctx, siteId, mapId, &body)
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


# Delete Site Ap Auto Orientation

This API is called to force stop auto placement for a given map

```go
DeleteSiteApAutoOrientation(
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

resp, err := sitesMapsAutoOrientation.DeleteSiteApAutoOrientation(ctx, mapId, siteId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Autoplacement was not triggered | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Start Site Ap Auto Orientation

This API is called to trigger a map for auto orientation. For auto orient feature to work, BLE data needs to be collected from the APs on the map. This precess is not disruptive unlike FTM collection. Repeated POST to this endpoint while a map is still running will be rejected.

List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide auto orient suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default

```go
StartSiteApAutoOrientation(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID,
    body *models.AutoOrient) (
    models.ApiResponse[models.ResponseAutoOrientation],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AutoOrient`](../../doc/models/auto-orient.md) | Body, Optional | - |

## Response Type

[`models.ResponseAutoOrientation`](../../doc/models/response-auto-orientation.md)

## Example Usage

```go
ctx := context.Background()

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AutoOrient{
    ForceCollection: models.ToPointer(false),
}

apiResponse, err := sitesMapsAutoOrientation.StartSiteApAutoOrientation(ctx, mapId, siteId, &body)
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
  "state": "Not Started",
  "time_queued": 1675414259
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Request | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

