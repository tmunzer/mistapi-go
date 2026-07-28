# Orgs Clients-Marvis

```go
orgsClientsMarvis := client.OrgsClientsMarvis()
```

## Class Name

`OrgsClientsMarvis`

## Methods

* [Count Org Marvis Client Events](../../doc/controllers/orgs-clients-marvis.md#count-org-marvis-client-events)
* [Delete Org Marvis Client](../../doc/controllers/orgs-clients-marvis.md#delete-org-marvis-client)
* [Get Org Marvis Client Insights](../../doc/controllers/orgs-clients-marvis.md#get-org-marvis-client-insights)
* [Search Org Marvis Client Events](../../doc/controllers/orgs-clients-marvis.md#search-org-marvis-client-events)


# Count Org Marvis Client Events

Count Marvis Client events by a distinct field.

```go
CountOrgMarvisClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *string,
    mType *string,
    deviceId *string,
    wifiMac *string,
    wifiIp *string,
    hostname *string,
    ssid *string,
    bssid *string,
    channel *string,
    preBssid *string,
    preChannel *string,
    limit *int,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | `*string` | Query, Optional | Field to count by. enum: `type`, `device_id`, `wifi_mac`, `wifi_ip`, `hostname`, `ssid`, `bssid`, `channel`, `pre_bssid`, `pre_channel`<br><br>**Default**: `"type"` |
| `mType` | `*string` | Query, Optional | Filter by event type |
| `deviceId` | `*string` | Query, Optional | Filter by Marvis Client installation device UUID |
| `wifiMac` | `*string` | Query, Optional | Filter by device Wi-Fi MAC address |
| `wifiIp` | `*string` | Query, Optional | Filter by device Wi-Fi IP address |
| `hostname` | `*string` | Query, Optional | Filter by device hostname |
| `ssid` | `*string` | Query, Optional | Filter by SSID involved in roam events |
| `bssid` | `*string` | Query, Optional | Filter by BSSID the client roamed to |
| `channel` | `*string` | Query, Optional | Filter by channel the client roamed to |
| `preBssid` | `*string` | Query, Optional | Filter by BSSID the client roamed from |
| `preChannel` | `*string` | Query, Optional | Filter by channel the client roamed from |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |

## Response Type

**200**: Count result

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := "type"

limit := 100

duration := "10m"

apiResponse, err := orgsClientsMarvis.CountOrgMarvisClientEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
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


# Delete Org Marvis Client

Delete Marvis Client

```go
DeleteOrgMarvisClient(
    ctx context.Context,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsClientsMarvis.DeleteOrgMarvisClient(ctx, orgId)
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


# Get Org Marvis Client Insights

Return time-series metrics for a specific Marvis Client device. For the full list of supported metric field names and example values, refer to [List Insight Metrics](../../doc/controllers/constants-definitions.md#list-insight-metrics) under `/api/v1/const/insight_metrics`.

```go
GetOrgMarvisClientInsights(
    ctx context.Context,
    orgId uuid.UUID,
    marvisclientId uuid.UUID,
    duration *string,
    interval *string,
    start *string,
    end *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.MarvisClientInsights],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `marvisclientId` | `uuid.UUID` | Template, Required | Marvis Client device UUID |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: Marvis Client time-series metrics

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.MarvisClientInsights](../../doc/models/marvis-client-insights.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

marvisclientId := uuid.MustParse("00000e1c-0000-0000-0000-000000000000")

duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := orgsClientsMarvis.GetOrgMarvisClientInsights(ctx, orgId, marvisclientId, &duration, &interval, nil, nil, &limit, &page)
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
  "avg_battery": [
    82,
    79
  ],
  "avg_cellular_rssi": [
    -95,
    -92
  ],
  "avg_cpu": [
    12.5,
    15.2
  ],
  "avg_memory": [
    43.1,
    44.8
  ],
  "avg_wifi_rssi": [
    -62,
    -60
  ],
  "end": 1717113600,
  "interval": 3600,
  "limit": 100,
  "page": 1,
  "rt": [
    "2026-05-30T16:00:00Z",
    "2026-05-30T17:00:00Z"
  ],
  "start": 1717027200
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


# Search Org Marvis Client Events

Search Marvis Client events across the organization.

```go
SearchOrgMarvisClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    deviceId *string,
    wifiMac *string,
    wifiIp *string,
    hostname *string,
    ssid *string,
    bssid *string,
    channel *string,
    preBssid *string,
    preChannel *string,
    limit *int,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.MarvisClientEventsSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | Filter by event type |
| `deviceId` | `*string` | Query, Optional | Filter by Marvis Client installation device UUID |
| `wifiMac` | `*string` | Query, Optional | Filter by device Wi-Fi MAC address |
| `wifiIp` | `*string` | Query, Optional | Filter by device Wi-Fi IP address |
| `hostname` | `*string` | Query, Optional | Filter by device hostname |
| `ssid` | `*string` | Query, Optional | Filter by SSID involved in roam events |
| `bssid` | `*string` | Query, Optional | Filter by BSSID the client roamed to |
| `channel` | `*string` | Query, Optional | Filter by channel the client roamed to |
| `preBssid` | `*string` | Query, Optional | Filter by BSSID the client roamed from |
| `preChannel` | `*string` | Query, Optional | Filter by channel the client roamed from |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |

## Response Type

**200**: Paginated Marvis Client events search results

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.MarvisClientEventsSearch](../../doc/models/marvis-client-events-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

duration := "10m"

apiResponse, err := orgsClientsMarvis.SearchOrgMarvisClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "limit": 1000,
  "results": [
    {
      "bssid": "5c5b35000002",
      "channel": 11,
      "device_id": "da088609-8e8d-6d8b-0e40-fe1dc94b9218",
      "hostname": "jdoe123-dell",
      "location": {
        "map_id": "7735ef91-83b4-6116-1c1a-57819af48867",
        "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
        "timestamp": 1678377926,
        "x": 12.5,
        "y": 45.0
      },
      "neighbor_ap_report": [
        {
          "band": "5",
          "bssid": "5c5b35000003",
          "channel": 44,
          "rssi": -55
        },
        {
          "band": "5",
          "bssid": "5c5b35000004",
          "channel": 36,
          "rssi": -70
        }
      ],
      "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
      "pre_bssid": "5c5b35000001",
      "pre_channel": 7,
      "pre_rssi": -76,
      "rssi": -53,
      "ssid": "Corp",
      "timestamp": 1678377926,
      "type": "MARVISCLIENT_ROAMED",
      "wifi_ip": "10.10.20.1",
      "wifi_mac": "f01c2df166e0"
    },
    {
      "device_id": "da088609-8e8d-6d8b-0e40-fe1dc94b9218",
      "hostname": "jdoe123-dell",
      "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
      "percent": 9,
      "timestamp": 1678378040,
      "type": "MARVISCLIENT_LOW_BATTERY",
      "wifi_ip": "10.10.20.1",
      "wifi_mac": "f01c2df166e0"
    }
  ],
  "total": 1
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

