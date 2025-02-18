# Sites RRM

```go
sitesRRM := client.SitesRRM()
```

## Class Name

`SitesRRM`

## Methods

* [Get Site Current Channel Planning](../../doc/controllers/sites-rrm.md#get-site-current-channel-planning)
* [Get Site Current Rrm Considerations](../../doc/controllers/sites-rrm.md#get-site-current-rrm-considerations)
* [Get Site Current Rrm Neighbors](../../doc/controllers/sites-rrm.md#get-site-current-rrm-neighbors)
* [Get Site Rrm Events](../../doc/controllers/sites-rrm.md#get-site-rrm-events)


# Get Site Current Channel Planning

Get Current Channel Planning

```go
GetSiteCurrentChannelPlanning(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.Rrm],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Rrm`](../../doc/models/rrm.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesRRM.GetSiteCurrentChannelPlanning(ctx, siteId)
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
      "channel": 6,
      "disabled": true,
      "power": 5,
      "power_max": 3,
      "power_min": 18,
      "preamble": "auto",
      "usage": "24"
    },
    "band_5": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 100,
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
          "channel": 6,
          "disabled": true,
          "power": 9,
          "power_max": 6,
          "power_min": 17,
          "preamble": "auto",
          "usage": "rrm"
        },
        "band_5": {
          "allow_rrm_disable": true,
          "antenna_mode": "default",
          "bandwidth": 20,
          "channel": 112,
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Template, Required | 802.11 Band |

## Response Type

[`models.ResponseRrmConsideration`](../../doc/models/response-rrm-consideration.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

band := models.Dot11BandEnum_ENUM6

apiResponse, err := sitesRRM.GetSiteCurrentRrmConsiderations(ctx, siteId, deviceId, band)
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
  "results": [
    {
      "channel": 36,
      "noise": -78,
      "non_wifi": 0.08,
      "other_rssi": -66,
      "other_ssid": "Rivendell5G",
      "rssi": -48,
      "util_score": 0.1,
      "util_score_non_wifi": 0.01,
      "util_score_other": 0.05,
      "wifi": 0.13
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


# Get Site Current Rrm Neighbors

Get Current RRM observed neighbors

```go
GetSiteCurrentRrmNeighbors(
    ctx context.Context,
    siteId uuid.UUID,
    band models.Dot11BandEnum,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseRrmNeighbors],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Template, Required | 802.11 Band |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.ResponseRrmNeighbors`](../../doc/models/response-rrm-neighbors.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

band := models.Dot11BandEnum_ENUM6

limit := 100

page := 1

apiResponse, err := sitesRRM.GetSiteCurrentRrmNeighbors(ctx, siteId, band, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Rrm Events

Get Site RRM Events

```go
GetSiteRrmEvents(
    ctx context.Context,
    siteId uuid.UUID,
    band *models.Dot11BandEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseEventsRrm],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.ResponseEventsRrm`](../../doc/models/response-events-rrm.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")







duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesRRM.GetSiteRrmEvents(ctx, siteId, nil, nil, nil, &duration, &limit, &page)
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
  "end": 1428954000,
  "limit": 100,
  "next": "/api/v1/sites/dca0a44b-324c-11e6-a776-0243ad110007/events/rrm?start=1428939600&end=1428949600&limit=200&token=001a0010000000120010000005005880ec18000004776c616e007fffffeb067ab8e29c1d659b6a7c8cf698bf81490003",
  "results": [
    {
      "ap_id": "00000000-0000-0000-1000-5c5b359e4fe0",
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

