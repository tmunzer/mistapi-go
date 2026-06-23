# Sites Beacons

```go
sitesBeacons := client.SitesBeacons()
```

## Class Name

`SitesBeacons`

## Methods

* [Create Site Beacon](../../doc/controllers/sites-beacons.md#create-site-beacon)
* [Delete Site Beacon](../../doc/controllers/sites-beacons.md#delete-site-beacon)
* [Get Site Beacon](../../doc/controllers/sites-beacons.md#get-site-beacon)
* [List Site Beacons](../../doc/controllers/sites-beacons.md#list-site-beacons)
* [Update Site Beacon](../../doc/controllers/sites-beacons.md#update-site-beacon)


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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Beacon`](../../doc/models/beacon.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Beacon](../../doc/models/beacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Beacon{
    EddystoneInstance:    models.ToPointer("string"),
    EddystoneNamespace:   models.ToPointer("string"),
    EddystoneUrl:         models.ToPointer("string"),
    IbeaconMajor:         models.NewOptional(models.ToPointer(1)),
    IbeaconMinor:         models.NewOptional(models.ToPointer(1)),
    IbeaconUuid:          models.NewOptional(models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"))),
    Mac:                  models.ToPointer("string"),
    Name:                 models.ToPointer("string"),
    Power:                models.ToPointer(0),
    Type:                 models.ToPointer(models.BeaconTypeEnum_EDDYSTONEUID),
    X:                    models.ToPointer(float64(0)),
    Y:                    models.ToPointer(float64(0)),
}

apiResponse, err := sitesBeacons.CreateSiteBeacon(ctx, siteId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
  "ibeacon_major": 1,
  "ibeacon_minor": 1,
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Site Beacon

Delete Site Beacon

```go
DeleteSiteBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    beaconId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `beaconId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

beaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesBeacons.DeleteSiteBeacon(ctx, siteId, beaconId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `beaconId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Beacon](../../doc/models/beacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

beaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesBeacons.GetSiteBeacon(ctx, siteId, beaconId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
  "ibeacon_major": 1,
  "ibeacon_minor": 1,
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Beacon](../../doc/models/beacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesBeacons.ListSiteBeacons(ctx, siteId, &limit, &page)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
    "ibeacon_major": 1,
    "ibeacon_minor": 1,
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Update Site Beacon

Update Site Beacon

```go
UpdateSiteBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    beaconId uuid.UUID,
    body *models.Beacon) (
    models.ApiResponse[models.Beacon],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `beaconId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Beacon`](../../doc/models/beacon.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Beacon](../../doc/models/beacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

beaconId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Beacon{
    EddystoneInstance:    models.ToPointer("string"),
    EddystoneNamespace:   models.ToPointer("string"),
    EddystoneUrl:         models.ToPointer("string"),
    IbeaconMajor:         models.NewOptional(models.ToPointer(1)),
    IbeaconMinor:         models.NewOptional(models.ToPointer(1)),
    IbeaconUuid:          models.NewOptional(models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"))),
    Mac:                  models.ToPointer("string"),
    Name:                 models.ToPointer("string"),
    Power:                models.ToPointer(0),
    Type:                 models.ToPointer(models.BeaconTypeEnum_EDDYSTONEUID),
    X:                    models.ToPointer(float64(0)),
    Y:                    models.ToPointer(float64(0)),
}

apiResponse, err := sitesBeacons.UpdateSiteBeacon(ctx, siteId, beaconId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
  "ibeacon_major": 1,
  "ibeacon_minor": 1,
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

