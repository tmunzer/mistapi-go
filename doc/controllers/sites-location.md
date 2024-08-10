# Sites Location

```go
sitesLocation := client.SitesLocation()
```

## Class Name

`SitesLocation`

## Methods

* [Clear Site Ml Overwrite for Device](../../doc/controllers/sites-location.md#clear-site-ml-overwrite-for-device)
* [Clear Site Ml Overwrite for Map](../../doc/controllers/sites-location.md#clear-site-ml-overwrite-for-map)
* [Get Site Beam Coverage Overview](../../doc/controllers/sites-location.md#get-site-beam-coverage-overview)
* [Get Site Default Plf for Models](../../doc/controllers/sites-location.md#get-site-default-plf-for-models)
* [Get Site Machine Learning Current Stat](../../doc/controllers/sites-location.md#get-site-machine-learning-current-stat)
* [Overwrite Site Ml for Device](../../doc/controllers/sites-location.md#overwrite-site-ml-for-device)
* [Overwrite Site Ml for Map](../../doc/controllers/sites-location.md#overwrite-site-ml-for-map)
* [Reset Site Ml Stats by Map](../../doc/controllers/sites-location.md#reset-site-ml-stats-by-map)


# Clear Site Ml Overwrite for Device

Clear ML Overwrite for Device

```go
ClearSiteMlOverwriteForDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesLocation.ClearSiteMlOverwriteForDevice(ctx, siteId, deviceId)
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


# Clear Site Ml Overwrite for Map

Clear ML Overwrite for Map

```go
ClearSiteMlOverwriteForMap(
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

resp, err := sitesLocation.ClearSiteMlOverwriteForMap(ctx, siteId, mapId)
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


# Get Site Beam Coverage Overview

Get Beam Coverage Overview

```go
GetSiteBeamCoverageOverview(
    ctx context.Context,
    siteId uuid.UUID,
    mapId *string,
    mType *models.RfClientTypeEnum,
    clientType *string,
    duration *string,
    resolution *models.ResolutionEnum,
    start *int,
    end *int) (
    models.ApiResponse[models.ResponseLocationCoverage],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `*string` | Query, Optional | map_id (filter by map_id) |
| `mType` | [`*models.RfClientTypeEnum`](../../doc/models/rf-client-type-enum.md) | Query, Optional | - |
| `clientType` | `*string` | Query, Optional | client_type (as filter. optional) |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `resolution` | [`*models.ResolutionEnum`](../../doc/models/resolution-enum.md) | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |

## Response Type

[`models.ResponseLocationCoverage`](../../doc/models/response-location-coverage.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := "00000000-0000-0000-0000-000000000000"





duration := "10m"

resolution := models.ResolutionEnum("default")





apiResponse, err := sitesLocation.GetSiteBeamCoverageOverview(ctx, siteId, &mapId, nil, nil, &duration, &resolution, nil, nil)
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
  "beams_means": [
    [
      1,
      3,
      3.2
    ],
    [
      6,
      10,
      6.5
    ]
  ],
  "end": 1428954000,
  "gridsize": 1,
  "result_def": [
    "x",
    "y",
    "beams_mean",
    "beacons_mean",
    "max_rssi",
    "avg_rssi"
  ],
  "results": [
    [
      1,
      3,
      3.2,
      18.5,
      -68,
      -70
    ],
    [
      6,
      10,
      6.5,
      30,
      1,
      -72.5,
      -75
    ]
  ],
  "start": 1428939600
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


# Get Site Default Plf for Models

Get Default PLF for Models

```go
GetSiteDefaultPlfForModels(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]interface{}],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

`[]interface{}`

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesLocation.GetSiteDefaultPlfForModels(ctx, siteId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
[
  {
    "current": {
      "Android": {
        "completed": 36,
        "int": -6,
        "level": 3,
        "ple": -3,
        "quality": "4",
        "src": "device",
        "timestamp": 1442854794
      },
      "iOS": {
        "completed": 16,
        "int": -6,
        "level": 6,
        "ple": -3,
        "quality": "2",
        "src": "default",
        "timestamp": 1442854704
      },
      "iPod": {
        "int": -10,
        "overwrite": true,
        "ple": -5,
        "src": "overwrite"
      }
    },
    "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  },
  {
    "beacon_id": "7913f032-aab4-c3ae-e83e-5a2756ef4d40",
    "current": {
      "iOS": {
        "completed": 16,
        "int": -6,
        "level": 6,
        "ple": -3,
        "quality": "last",
        "src": "device",
        "timestamp": 1442854704
      }
    }
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


# Get Site Machine Learning Current Stat

Get Machine Learning Current Stat
For each VBLE AP, it has ML model parameters (e.g. Path-loss-estimate, Intercept) as well as completion indicators (Level and PercentageComplete). For the completeness, ML takes N sample to finish its first level and use N*0.25 samples to complete each successive level. When a device is moved, the completeness will be reset as it has to re-learn.

```go
GetSiteMachineLearningCurrentStat(
    ctx context.Context,
    siteId uuid.UUID,
    mapId *string) (
    models.ApiResponse[[]interface{}],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `*string` | Query, Optional | map_id (as filter, optional) |

## Response Type

`[]interface{}`

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := "00000000-0000-0000-0000-000000000000"

apiResponse, err := sitesLocation.GetSiteMachineLearningCurrentStat(ctx, siteId, &mapId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
[
  {
    "current": {
      "Android": {
        "completed": 36,
        "int": -6,
        "level": 3,
        "ple": -3,
        "quality": "4",
        "src": "device",
        "timestamp": 1442854794
      },
      "iOS": {
        "completed": 16,
        "int": -6,
        "level": 6,
        "ple": -3,
        "quality": "2",
        "src": "default",
        "timestamp": 1442854704
      },
      "iPod": {
        "int": -10,
        "overwrite": true,
        "ple": -5,
        "src": "overwrite"
      }
    },
    "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  },
  {
    "beacon_id": "7913f032-aab4-c3ae-e83e-5a2756ef4d40",
    "current": {
      "iOS": {
        "completed": 16,
        "int": -6,
        "level": 6,
        "ple": -3,
        "quality": "last",
        "src": "device",
        "timestamp": 1442854704
      }
    }
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


# Overwrite Site Ml for Device

Overwrite ML For Device

```go
OverwriteSiteMlForDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body map[string]models.MlOverwriteAdditionalProperties) (
    models.ApiResponse[[]interface{}],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`map[string]models.MlOverwriteAdditionalProperties`](../../doc/models/ml-overwrite-additional-properties.md) | Body, Optional | Request Body |

## Response Type

`[]interface{}`

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := map[string]models.MlOverwriteAdditionalProperties{
    "iOS": models.MlOverwriteAdditionalProperties{
        Int: models.ToPointer(6),
        Ple: models.ToPointer(-3),
    },
    "iPod": models.MlOverwriteAdditionalProperties{
        Int: models.ToPointer(-10),
        Ple: models.ToPointer(-5),
    },
}

apiResponse, err := sitesLocation.OverwriteSiteMlForDevice(ctx, siteId, deviceId, body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
[
  {
    "current": {
      "Android": {
        "completed": 36,
        "int": -6,
        "level": 3,
        "ple": -3,
        "quality": "4",
        "src": "device",
        "timestamp": 1442854794
      },
      "iOS": {
        "completed": 16,
        "int": -6,
        "level": 6,
        "ple": -3,
        "quality": "2",
        "src": "default",
        "timestamp": 1442854704
      },
      "iPod": {
        "int": -10,
        "overwrite": true,
        "ple": -5,
        "src": "overwrite"
      }
    },
    "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  },
  {
    "beacon_id": "7913f032-aab4-c3ae-e83e-5a2756ef4d40",
    "current": {
      "iOS": {
        "completed": 16,
        "int": -6,
        "level": 6,
        "ple": -3,
        "quality": "last",
        "src": "device",
        "timestamp": 1442854704
      }
    }
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


# Overwrite Site Ml for Map

Overwrite ML For Map

```go
OverwriteSiteMlForMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body map[string]models.MlOverwriteAdditionalProperties) (
    models.ApiResponse[[]interface{}],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`map[string]models.MlOverwriteAdditionalProperties`](../../doc/models/ml-overwrite-additional-properties.md) | Body, Optional | Request Body |

## Response Type

`[]interface{}`

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := map[string]models.MlOverwriteAdditionalProperties{
    "iOS": models.MlOverwriteAdditionalProperties{
        Int: models.ToPointer(6),
        Ple: models.ToPointer(-3),
    },
    "iPod": models.MlOverwriteAdditionalProperties{
        Int: models.ToPointer(-10),
        Ple: models.ToPointer(-5),
    },
}

apiResponse, err := sitesLocation.OverwriteSiteMlForMap(ctx, siteId, mapId, body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
[
  {
    "current": {
      "Android": {
        "completed": 36,
        "int": -6,
        "level": 3,
        "ple": -3,
        "quality": "4",
        "src": "device",
        "timestamp": 1442854794
      },
      "iOS": {
        "completed": 16,
        "int": -6,
        "level": 6,
        "ple": -3,
        "quality": "2",
        "src": "default",
        "timestamp": 1442854704
      },
      "iPod": {
        "int": -10,
        "overwrite": true,
        "ple": -5,
        "src": "overwrite"
      }
    },
    "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  },
  {
    "beacon_id": "7913f032-aab4-c3ae-e83e-5a2756ef4d40",
    "current": {
      "iOS": {
        "completed": 16,
        "int": -6,
        "level": 6,
        "ple": -3,
        "quality": "last",
        "src": "device",
        "timestamp": 1442854704
      }
    }
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


# Reset Site Ml Stats by Map

Reset ML Stats by Map

```go
ResetSiteMlStatsByMap(
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

resp, err := sitesLocation.ResetSiteMlStatsByMap(ctx, siteId, mapId)
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

