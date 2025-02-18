# Sites Maps-Auto-Placement

```go
sitesMapsAutoPlacement := client.SitesMapsAutoPlacement()
```

## Class Name

`SitesMapsAutoPlacement`

## Methods

* [Clear Site Ap Auto Orient](../../doc/controllers/sites-maps-auto-placement.md#clear-site-ap-auto-orient)
* [Clear Site Ap Autoplacement](../../doc/controllers/sites-maps-auto-placement.md#clear-site-ap-autoplacement)
* [Confirm Site Ap Localization Data](../../doc/controllers/sites-maps-auto-placement.md#confirm-site-ap-localization-data)
* [Delete Site Ap Auto Orientation](../../doc/controllers/sites-maps-auto-placement.md#delete-site-ap-auto-orientation)
* [Delete Site Ap Autoplacement](../../doc/controllers/sites-maps-auto-placement.md#delete-site-ap-autoplacement)
* [Get Site Ap Auto Placement](../../doc/controllers/sites-maps-auto-placement.md#get-site-ap-auto-placement)
* [Run Site Ap Autoplacement](../../doc/controllers/sites-maps-auto-placement.md#run-site-ap-autoplacement)
* [Start Site Ap Auto Orientation](../../doc/controllers/sites-maps-auto-placement.md#start-site-ap-auto-orientation)


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
    Macs:                 []string{
        "683b679ac024",
    },
}

resp, err := sitesMapsAutoPlacement.ClearSiteApAutoOrient(ctx, siteId, mapId, &body)
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


# Clear Site Ap Autoplacement

This API is used to destroy the cached autoplacement locations of a map or subset of APs on a map.

```go
ClearSiteApAutoplacement(
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
    Macs:                 []string{
        "683b679ac024",
    },
}

resp, err := sitesMapsAutoPlacement.ClearSiteApAutoplacement(ctx, siteId, mapId, &body)
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


# Confirm Site Ap Localization Data

This API is used to accept or reject the cached autoplacement and auto orientation values of a map or subset of APs on a map. A rejected AP will retain its current X,Y and orientation until accpeted.

```go
ConfirmSiteApLocalizationData(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.UseAutoApValues) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UseAutoApValues`](../../doc/models/use-auto-ap-values.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UseAutoApValues{
    Accept:               models.ToPointer(false),
    For:                  models.ToPointer(models.UseAutoApValuesForEnum_PLACEMENT),
    AdditionalProperties: map[string]interface{}{
        "device_macs": interface{}("string"),
    },
}

resp, err := sitesMapsAutoPlacement.ConfirmSiteApLocalizationData(ctx, siteId, mapId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Invalid localization service expected: placement or orientation | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

resp, err := sitesMapsAutoPlacement.DeleteSiteApAutoOrientation(ctx, mapId, siteId)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Site Ap Autoplacement

This API is called to force stop auto placement for a given map

```go
DeleteSiteApAutoplacement(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMapsAutoPlacement.DeleteSiteApAutoplacement(ctx, siteId, mapId)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Ap Auto Placement

This API is called to view the current status of auto placement for a given map.

#### Status Descriptions

| Status | Description |
| --- | --- |
| `pending` | Autoplacement has not been requested for this map |
| `inprogress` | Autoplacement is currently processing |
| `done` | The autoplacement process has completed |
| `data_needed` | Additional position data is required for autoplacement. Users should verify the requested anchor APs have a position on the map |
| `invalid_model` | Autoplacement is not supported on the model of the APs on the map |
| `invalid_version` | Autoplacement is not supported with the APs current firmware version |
| `error` | There was an error in the autoplacement process |

```go
GetSiteApAutoPlacement(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    models.ApiResponse[models.ResponseAutoPlacementInfo],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseAutoPlacementInfo`](../../doc/models/response-auto-placement-info.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesMapsAutoPlacement.GetSiteApAutoPlacement(ctx, siteId, mapId)
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
  "end_time": 1678900362,
  "start_time": 1678900062,
  "status": "done"
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


# Run Site Ap Autoplacement

This API is called to trigger a map for auto placement. For auto placement feature to work, RTT-FTM data need to be collected from the APs on the map. This scan is disruptive and therefore the user must be notified of service disrution during the functioning of auto placement Repeated POST to this endpoint while a map is still running will be rejected.

List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide autoplacement suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default

```go
RunSiteApAutoplacement(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.AutoPlacement) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AutoPlacement`](../../doc/models/auto-placement.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AutoPlacement{
    ForceCollection:      models.ToPointer(false),
}

resp, err := sitesMapsAutoPlacement.RunSiteApAutoplacement(ctx, siteId, mapId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | * Map has less than 3 APs associated with it to perform auto placement<br>* Auto AP Placement is already in progress for this Map<br>* Autoplacement data does not exist or has gone stale | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
    ForceCollection:      models.ToPointer(false),
}

apiResponse, err := sitesMapsAutoPlacement.StartSiteApAutoOrientation(ctx, mapId, siteId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

