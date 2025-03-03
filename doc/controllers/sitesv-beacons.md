# Sitesv Beacons

```go
sitesvBeacons := client.SitesvBeacons()
```

## Class Name

`SitesvBeacons`

## Methods

* [Create Site V Beacon](../../doc/controllers/sitesv-beacons.md#create-site-v-beacon)
* [Delete Site V Beacon](../../doc/controllers/sitesv-beacons.md#delete-site-v-beacon)
* [Get Site V Beacon](../../doc/controllers/sitesv-beacons.md#get-site-v-beacon)
* [List Site V Beacons](../../doc/controllers/sitesv-beacons.md#list-site-v-beacons)
* [Update Site V Beacon](../../doc/controllers/sitesv-beacons.md#update-site-v-beacon)


# Create Site V Beacon

Create Virtual Beacon

```go
CreateSiteVBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Vbeacon) (
    models.ApiResponse[models.Vbeacon],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Vbeacon`](../../doc/models/vbeacon.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Vbeacon](../../doc/models/vbeacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Vbeacon{
    Major:                models.ToPointer(0),
    Message:              models.ToPointer("string"),
    Minor:                models.ToPointer(0),
    Name:                 models.ToPointer("string"),
    Power:                models.ToPointer(4),
    PowerMode:            models.ToPointer(models.BleConfigPowerModeEnum_ENUMDEFAULT),
    Url:                  models.ToPointer("string"),
    Uuid:                 models.ToPointer(uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9")),
    WayfindingNodename:   models.ToPointer("string"),
    X:                    models.ToPointer(float64(0)),
    Y:                    models.ToPointer(float64(0)),
}

apiResponse, err := sitesVBeacons.CreateSiteVBeacon(ctx, siteId, &body)
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
  "major": 0,
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "message": "string",
  "minor": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "power": 4,
  "power_mode": "default",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "url": "string",
  "uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "wayfinding_nodename": "string",
  "x": 0,
  "y": 0
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


# Delete Site V Beacon

Delete Site Virtual Beacon

```go
DeleteSiteVBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    vbeaconId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `vbeaconId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

vbeaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesVBeacons.DeleteSiteVBeacon(ctx, siteId, vbeaconId)
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


# Get Site V Beacon

Get Site Virtual Beacon Details

```go
GetSiteVBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    vbeaconId uuid.UUID) (
    models.ApiResponse[models.Vbeacon],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `vbeaconId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Vbeacon](../../doc/models/vbeacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

vbeaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesVBeacons.GetSiteVBeacon(ctx, siteId, vbeaconId)
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
  "major": 0,
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "message": "string",
  "minor": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "power": 4,
  "power_mode": "default",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "url": "string",
  "uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "wayfinding_nodename": "string",
  "x": 0,
  "y": 0
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


# List Site V Beacons

Get List of Site Virtual Beacons

```go
ListSiteVBeacons(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Vbeacon],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Vbeacon](../../doc/models/vbeacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesVBeacons.ListSiteVBeacons(ctx, siteId, &limit, &page)
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
    "major": 0,
    "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "message": "string",
    "minor": 0,
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "power": 4,
    "power_mode": "default",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "url": "string",
    "uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "wayfinding_nodename": "string",
    "x": 0,
    "y": 0
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


# Update Site V Beacon

Update Site Virtual Beacon

```go
UpdateSiteVBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    vbeaconId uuid.UUID,
    body *models.Vbeacon) (
    models.ApiResponse[models.Vbeacon],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `vbeaconId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Vbeacon`](../../doc/models/vbeacon.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Vbeacon](../../doc/models/vbeacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

vbeaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Vbeacon{
    Major:                models.ToPointer(0),
    Message:              models.ToPointer("string"),
    Minor:                models.ToPointer(0),
    Name:                 models.ToPointer("string"),
    Power:                models.ToPointer(4),
    PowerMode:            models.ToPointer(models.BleConfigPowerModeEnum_ENUMDEFAULT),
    Url:                  models.ToPointer("string"),
    Uuid:                 models.ToPointer(uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9")),
    WayfindingNodename:   models.ToPointer("string"),
    X:                    models.ToPointer(float64(0)),
    Y:                    models.ToPointer(float64(0)),
}

apiResponse, err := sitesVBeacons.UpdateSiteVBeacon(ctx, siteId, vbeaconId, &body)
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
  "major": 0,
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "message": "string",
  "minor": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "power": 4,
  "power_mode": "default",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "url": "string",
  "uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "wayfinding_nodename": "string",
  "x": 0,
  "y": 0
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

