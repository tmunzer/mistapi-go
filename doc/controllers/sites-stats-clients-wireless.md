# Sites Stats-Clients Wireless

```go
sitesStatsClientsWireless := client.SitesStatsClientsWireless()
```

## Class Name

`SitesStatsClientsWireless`

## Methods

* [Get Site Wireless Client Stats](../../doc/controllers/sites-stats-clients-wireless.md#get-site-wireless-client-stats)
* [Get Site Wireless Clients Stats by Map](../../doc/controllers/sites-stats-clients-wireless.md#get-site-wireless-clients-stats-by-map)
* [List Site Unconnected Client Stats](../../doc/controllers/sites-stats-clients-wireless.md#list-site-unconnected-client-stats)
* [List Site Wireless Clients Stats](../../doc/controllers/sites-stats-clients-wireless.md#list-site-wireless-clients-stats)


# Get Site Wireless Client Stats

Get Site Client Stats Details

```go
GetSiteWirelessClientStats(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    wired *bool) (
    models.ApiResponse[models.StatsClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `wired` | `*bool` | Query, Optional | **Default**: `false` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type models.StatsClient.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

wired := false

apiResponse, err := sitesStatsClientsWireless.GetSiteWirelessClientStats(ctx, siteId, clientMac, &wired)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsStatsWirelessClient(); ok {
        fmt.Println("Value narrowed down to models.StatsWirelessClient: ", *r)
    } else if r, ok := responseBody.AsStatsWiredClient(); ok {
        fmt.Println("Value narrowed down to models.StatsWiredClient: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "annotation": "unknown",
  "ap_id": "00000000-0000-0000-1000-5c5b35963d70",
  "ap_mac": "5c5b358e6fea",
  "assoc_time": 1741152905,
  "band": "5",
  "bssid": "5c5b358298f2",
  "channel": 157,
  "dual_band": true,
  "family": "",
  "group": "",
  "hostname": "android-9b228dc33690",
  "idle_time": 5,
  "ip": "10.100.0.47",
  "is_guest": false,
  "key_mgmt": "WPA3-SAE-FT/CCMP",
  "last_seen": 1741257505,
  "mac": "dadbfc123456",
  "manufacture": "Unknown",
  "map_id": "ed7a0a4e-8835-4c94-ba78-6c1169c9f135",
  "model": "",
  "num_locating_aps": 2,
  "os": "Android 10",
  "proto": "ac",
  "rssi": -39,
  "rx_bps": 0,
  "rx_bytes": 14451780,
  "rx_pkts": 44175,
  "rx_rate": 6,
  "rx_retries": 2010,
  "site_id": "96c348a9-d6d7-4732-b4f5-23350a2843cd",
  "snr": 47,
  "ssid": "Live_demo_only",
  "tx_bps": 0,
  "tx_bytes": 56364072,
  "tx_pkts": 43685,
  "tx_rate": 173.3,
  "tx_retries": 5413,
  "uptime": 104600,
  "vlan_id": "1",
  "wlan_id": "497fc18a-79b5-405a-bf5a-192eed31ea60",
  "x": 695.3357339330526,
  "x_m": 24.086588,
  "y": 760.6746524247893,
  "y_m": 26.349943
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


# Get Site Wireless Clients Stats by Map

Get Site Clients Stats By Map

```go
GetSiteWirelessClientsStatsByMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsWirelessClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.StatsWirelessClient](../../doc/models/stats-wireless-client.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesStatsClientsWireless.GetSiteWirelessClientsStatsByMap(ctx, siteId, mapId, nil, nil, &duration, &limit, &page)
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
    "annotation": "unknown",
    "ap_id": "00000000-0000-0000-1000-5c5b35963d70",
    "ap_mac": "5c5b358e6fea",
    "assoc_time": 1741152905,
    "band": "5",
    "bssid": "5c5b358298f2",
    "channel": 157,
    "dual_band": true,
    "family": "",
    "group": "",
    "hostname": "android-9b228dc33690",
    "idle_time": 5,
    "ip": "10.100.0.47",
    "is_guest": false,
    "key_mgmt": "WPA3-SAE-FT/CCMP",
    "last_seen": 1741257505,
    "mac": "dadbfc123456",
    "manufacture": "Unknown",
    "map_id": "ed7a0a4e-8835-4c94-ba78-6c1169c9f135",
    "model": "",
    "num_locating_aps": 2,
    "os": "Android 10",
    "proto": "ac",
    "rssi": -39,
    "rx_bps": 0,
    "rx_bytes": 14451780,
    "rx_pkts": 44175,
    "rx_rate": 6,
    "rx_retries": 2010,
    "site_id": "96c348a9-d6d7-4732-b4f5-23350a2843cd",
    "snr": 47,
    "ssid": "Live_demo_only",
    "tx_bps": 0,
    "tx_bytes": 56364072,
    "tx_pkts": 43685,
    "tx_rate": 173.3,
    "tx_retries": 5413,
    "uptime": 104600,
    "vlan_id": "1",
    "wlan_id": "497fc18a-79b5-405a-bf5a-192eed31ea60",
    "x": 695.3357339330526,
    "x_m": 24.086588,
    "y": 760.6746524247893,
    "y_m": 26.349943
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


# List Site Unconnected Client Stats

Get List of Site Unconnected Client Location

```go
ListSiteUnconnectedClientStats(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    models.ApiResponse[[]models.StatsUnconnectedClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.StatsUnconnectedClient](../../doc/models/stats-unconnected-client.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsClientsWireless.ListSiteUnconnectedClientStats(ctx, siteId, mapId)
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
    "ap_mac": "5c5b350e0410",
    "last_seen": 1428939600,
    "mac": "5684dae9ac8b",
    "manufacture": "Apple",
    "rssi": -75,
    "x": 60,
    "y": 80
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


# List Site Wireless Clients Stats

Get List of Site All Clients Stats Details

```go
ListSiteWirelessClientsStats(
    ctx context.Context,
    siteId uuid.UUID,
    wired *bool,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.StatsClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wired` | `*bool` | Query, Optional | **Default**: `false` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []models.StatsClient.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wired := false

limit := 100





duration := "10m"

apiResponse, err := sitesStatsClientsWireless.ListSiteWirelessClientsStats(ctx, siteId, &wired, &limit, nil, nil, &duration)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsStatsWirelessClient(); ok {
            fmt.Println("Value narrowed down to models.StatsWirelessClient: ", *i)
        } else if i, ok := item.AsStatsWiredClient(); ok {
            fmt.Println("Value narrowed down to models.StatsWiredClient: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
[
  {
    "annotation": "unknown",
    "ap_id": "00000000-0000-0000-1000-5c5b35963d70",
    "ap_mac": "5c5b358e6fea",
    "assoc_time": 1741152905,
    "band": "5",
    "bssid": "5c5b358298f2",
    "channel": 157,
    "dual_band": true,
    "family": "",
    "group": "",
    "hostname": "android-9b228dc33690",
    "idle_time": 5,
    "ip": "10.100.0.47",
    "is_guest": false,
    "key_mgmt": "WPA3-SAE-FT/CCMP",
    "last_seen": 1741257505,
    "mac": "dadbfc123456",
    "manufacture": "Unknown",
    "map_id": "ed7a0a4e-8835-4c94-ba78-6c1169c9f135",
    "model": "",
    "num_locating_aps": 2,
    "os": "Android 10",
    "proto": "ac",
    "rssi": -39,
    "rx_bps": 0,
    "rx_bytes": 14451780,
    "rx_pkts": 44175,
    "rx_rate": 6,
    "rx_retries": 2010,
    "site_id": "96c348a9-d6d7-4732-b4f5-23350a2843cd",
    "snr": 47,
    "ssid": "Live_demo_only",
    "tx_bps": 0,
    "tx_bytes": 56364072,
    "tx_pkts": 43685,
    "tx_rate": 173.3,
    "tx_retries": 5413,
    "uptime": 104600,
    "vlan_id": "1",
    "wlan_id": "497fc18a-79b5-405a-bf5a-192eed31ea60",
    "x": 695.3357339330526,
    "x_m": 24.086588,
    "y": 760.6746524247893,
    "y_m": 26.349943
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

