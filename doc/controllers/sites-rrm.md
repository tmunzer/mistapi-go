# Sites RRM

```go
sitesRRM := client.SitesRRM()
```

## Class Name

`SitesRRM`

## Methods

* [Get Site Channel Scores](../../doc/controllers/sites-rrm.md#get-site-channel-scores)
* [Get Site Current Channel Planning](../../doc/controllers/sites-rrm.md#get-site-current-channel-planning)
* [Get Site Current Rrm Considerations](../../doc/controllers/sites-rrm.md#get-site-current-rrm-considerations)
* [List Site Current Rrm Neighbors](../../doc/controllers/sites-rrm.md#list-site-current-rrm-neighbors)
* [List Site Rrm Events](../../doc/controllers/sites-rrm.md#list-site-rrm-events)


# Get Site Channel Scores

Get Site Channel Scores

```go
GetSiteChannelScores(
    ctx context.Context,
    siteId uuid.UUID,
    band models.Dot11BandEnum,
    start *string,
    end *string) (
    models.ApiResponse[models.ResponseRrmChannelScores],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Template, Required | 802.11 Band |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseRrmChannelScores](../../doc/models/response-rrm-channel-scores.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

band := models.Dot11BandEnum_ENUM5

apiResponse, err := sitesRRM.GetSiteChannelScores(ctx, siteId, band, nil, nil)
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
  "results": [
    {
      "channel": 36,
      "util_score": 0.009,
      "util_score_noise_floor": 0.001,
      "util_score_non_wifi": 0.003,
      "util_score_other": 0.002,
      "util_score_radar": 0.0,
      "util_score_undecodable_wifi": 0.004,
      "util_score_unknown_wifi": 0.0
    }
  ]
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


# Get Site Current Channel Planning

Get Current Channel Planning

```go
GetSiteCurrentChannelPlanning(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.Rrm],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Rrm](../../doc/models/rrm.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesRRM.GetSiteCurrentChannelPlanning(ctx, siteId)
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
  "band_24": {},
  "band_24_metric": {
    "avg_aps_per_channel": 0,
    "channel_distribution_uniformity": 0,
    "cochannel_neighbors": 0,
    "density": 0,
    "naps_by_channel": {},
    "naps_by_power": {},
    "neighbors": 0,
    "noise": 0
  },
  "band_5": {},
  "band_5_metric": {
    "avg_aps_per_channel": 0,
    "channel_distribution_uniformity": 0,
    "cochannel_neighbors": 0,
    "density": 0,
    "naps_by_channel": {},
    "naps_by_power": {},
    "neighbors": 0,
    "noise": 0
  },
  "rftemplate": {
    "band_24": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channels": [
        1,
        6,
        11
      ],
      "disabled": true,
      "power": 5,
      "power_max": 3,
      "power_min": 18,
      "preamble": "auto"
    },
    "band_5": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channels": [
        36,
        40,
        44,
        48,
        52,
        56,
        60,
        64,
        100,
        104,
        108,
        112
      ],
      "disabled": true,
      "power": 9,
      "power_max": 6,
      "power_min": 17,
      "preamble": "auto"
    },
    "country_code": "string",
    "created_time": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "model_specific": {
      "property1": {
        "band_24": {
          "allow_rrm_disable": true,
          "antenna_mode": "default",
          "bandwidth": 20,
          "channels": [
            1,
            6,
            11
          ],
          "disabled": true,
          "power": 9,
          "power_max": 6,
          "power_min": 17,
          "preamble": "auto"
        },
        "band_5": {
          "allow_rrm_disable": true,
          "antenna_mode": "default",
          "bandwidth": 20,
          "channels": [
            36,
            40,
            44,
            48,
            52,
            56,
            60,
            64,
            100,
            104,
            108,
            112
          ],
          "disabled": true,
          "power": 10,
          "power_max": 6,
          "power_min": 15,
          "preamble": "auto"
        }
      }
    },
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  },
  "rftemplate_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "rftemplate_name": "string",
  "status": "updating",
  "timestamp": 0
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


# Get Site Current Rrm Considerations

Get Current RRM Considerations for an AP on a Specific Band

```go
GetSiteCurrentRrmConsiderations(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    band models.Dot11BandEnum) (
    models.ApiResponse[models.ResponseRrmConsideration],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Template, Required | 802.11 Band |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseRrmConsideration](../../doc/models/response-rrm-consideration.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

band := models.Dot11BandEnum_ENUM5

apiResponse, err := sitesRRM.GetSiteCurrentRrmConsiderations(ctx, siteId, deviceId, band)
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
  "results": [
    {
      "channel": 36,
      "noise": -78,
      "other_rssi": -66,
      "other_ssid": "Rivendell5G",
      "rssi": -48,
      "util_score": 0.1,
      "util_score_non_wifi": 0.01,
      "util_score_other": 0.05
    }
  ]
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


# List Site Current Rrm Neighbors

List Current RRM observed neighbors

```go
ListSiteCurrentRrmNeighbors(
    ctx context.Context,
    siteId uuid.UUID,
    band models.Dot11BandEnum,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseRrmNeighbors],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Template, Required | 802.11 Band |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseRrmNeighbors](../../doc/models/response-rrm-neighbors.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

band := models.Dot11BandEnum_ENUM5

limit := 100

page := 1

apiResponse, err := sitesRRM.ListSiteCurrentRrmNeighbors(ctx, siteId, band, &limit, &page)
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
  "end": 1428954000,
  "limit": 100,
  "next": "/api/v1/sites/dca0a44b-324c-11e6-a776-0243ad110007/events/rrm?start=1428939600&end=1428949600&limit=200&token=001a0010000000120010000005005880ec18000004776c616e007fffffeb067ab8e29c1d659b6a7c8cf698bf81490003",
  "results": [
    {
      "mac": "a7c7096d7b8f",
      "neighbors": [
        {
          "mac": "5c5b35000311",
          "rssi": -75
        }
      ]
    }
  ],
  "start": 1428939600
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


# List Site Rrm Events

List Site RRM Events

```go
ListSiteRrmEvents(
    ctx context.Context,
    siteId uuid.UUID,
    band *models.Dot11BandEnum,
    start *string,
    end *string,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseEventsRrm],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 band used to filter radio results. enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsRrm](../../doc/models/response-events-rrm.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesRRM.ListSiteRrmEvents(ctx, siteId, nil, nil, nil, &duration, &limit, &page)
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
  "end": 1428954000,
  "limit": 100,
  "next": "/api/v1/sites/dca0a44b-324c-11e6-a776-0243ad110007/events/rrm?start=1428939600&end=1428949600&limit=200&token=001a0010000000120010000005005880ec18000004776c616e007fffffeb067ab8e29c1d659b6a7c8cf698bf81490003",
  "results": [
    {
      "ap": "5c5b359e4fe0",
      "band": "24",
      "bandwidth": 20,
      "channel": 6,
      "event": "scheduled-site_rrm",
      "power": 5,
      "pre_bandwidth": 20,
      "pre_channel": 1,
      "pre_power": 11,
      "pre_usage": "24",
      "timestamp": 1428939600,
      "usage": "24"
    }
  ],
  "start": 1428939600
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

