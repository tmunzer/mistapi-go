# Sites Beacons

```go
sitesBeacons := client.SitesBeacons()
```

## Class Name

`SitesBeacons`

## Methods

* [Create Site Beacon](../../doc/controllers/sites-beacons.md#create-site-beacon)
* [Delete Site Beacons](../../doc/controllers/sites-beacons.md#delete-site-beacons)
* [Get Site Beacon](../../doc/controllers/sites-beacons.md#get-site-beacon)
* [List Site Beacons](../../doc/controllers/sites-beacons.md#list-site-beacons)
* [Update Site Beacons](../../doc/controllers/sites-beacons.md#update-site-beacons)


# Create Site Beacon

Create Site Beacon

```go
CreateSiteBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Beacon) (
    models.ApiResponse[models.Beacon],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Beacon`](../../doc/models/beacon.md) | Body, Optional | Request Body |

## Response Type

[`models.Beacon`](../../doc/models/beacon.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Beacon{
    EddystoneInstance:    models.ToPointer("string"),
    EddystoneNamespace:   models.ToPointer("string"),
    EddystoneUrl:         models.ToPointer("string"),
    IbeaconMajor:         models.ToPointer(0),
    IbeaconMinor:         models.ToPointer(0),
    IbeaconUuid:          models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
    Mac:                  models.ToPointer("string"),
    Name:                 models.ToPointer("string"),
    Power:                models.ToPointer(0),
    Type:                 models.ToPointer(models.BeaconTypeEnum_EDDYSTONEUID),
    X:                    models.ToPointer(float64(0)),
    Y:                    models.ToPointer(float64(0)),
}

apiResponse, err := sitesBeacons.CreateSiteBeacon(ctx, siteId, &body)
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
  "eddystone_instance": "string",
  "eddystone_namespace": "string",
  "eddystone_url": "string",
  "ibeacon_major": 0,
  "ibeacon_minor": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mac": "string",
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "power": 0,
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "type": "eddystone-uid",
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


# Delete Site Beacons

Delete Site Beacon

```go
DeleteSiteBeacons(
    ctx context.Context,
    siteId uuid.UUID,
    beaconId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `beaconId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

beaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesBeacons.DeleteSiteBeacons(ctx, siteId, beaconId)
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


# Get Site Beacon

Get Site Beacon Details

```go
GetSiteBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    beaconId uuid.UUID) (
    models.ApiResponse[models.Beacon],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `beaconId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Beacon`](../../doc/models/beacon.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

beaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesBeacons.GetSiteBeacon(ctx, siteId, beaconId)
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
  "eddystone_instance": "string",
  "eddystone_namespace": "string",
  "eddystone_url": "string",
  "ibeacon_major": 0,
  "ibeacon_minor": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mac": "string",
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "power": 0,
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "type": "eddystone-uid",
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


# List Site Beacons

Get List of Site Beacons

```go
ListSiteBeacons(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Beacon],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Beacon`](../../doc/models/beacon.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesBeacons.ListSiteBeacons(ctx, siteId, &limit, &page)
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
    "eddystone_instance": "string",
    "eddystone_namespace": "string",
    "eddystone_url": "string",
    "ibeacon_major": 0,
    "ibeacon_minor": 0,
    "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "power": 0,
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "type": "eddystone-uid",
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


# Update Site Beacons

Update Site Beacon

```go
UpdateSiteBeacons(
    ctx context.Context,
    siteId uuid.UUID,
    beaconId uuid.UUID,
    body *models.Beacon) (
    models.ApiResponse[models.Beacon],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `beaconId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Beacon`](../../doc/models/beacon.md) | Body, Optional | Request Body |

## Response Type

[`models.Beacon`](../../doc/models/beacon.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

beaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Beacon{
    EddystoneInstance:    models.ToPointer("string"),
    EddystoneNamespace:   models.ToPointer("string"),
    EddystoneUrl:         models.ToPointer("string"),
    IbeaconMajor:         models.ToPointer(0),
    IbeaconMinor:         models.ToPointer(0),
    IbeaconUuid:          models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
    Mac:                  models.ToPointer("string"),
    Name:                 models.ToPointer("string"),
    Power:                models.ToPointer(0),
    Type:                 models.ToPointer(models.BeaconTypeEnum_EDDYSTONEUID),
    X:                    models.ToPointer(float64(0)),
    Y:                    models.ToPointer(float64(0)),
}

apiResponse, err := sitesBeacons.UpdateSiteBeacons(ctx, siteId, beaconId, &body)
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
  "eddystone_instance": "string",
  "eddystone_namespace": "string",
  "eddystone_url": "string",
  "ibeacon_major": 0,
  "ibeacon_minor": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mac": "string",
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "power": 0,
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "type": "eddystone-uid",
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

