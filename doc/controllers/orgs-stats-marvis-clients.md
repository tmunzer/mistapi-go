# Orgs Stats-Marvis Clients

```go
orgsStatsMarvisClients := client.OrgsStatsMarvisClients()
```

## Class Name

`OrgsStatsMarvisClients`

## Methods

* [Count Org Marvis Clients Stats](../../doc/controllers/orgs-stats-marvis-clients.md#count-org-marvis-clients-stats)
* [Search Org Marvis Clients Stats](../../doc/controllers/orgs-stats-marvis-clients.md#search-org-marvis-clients-stats)


# Count Org Marvis Clients Stats

Count Marvis Client stats records by a distinct field.

```go
CountOrgMarvisClientsStats(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *string,
    deviceId *string,
    wifiMac *string,
    wifiIp *string,
    hostname *string,
    model *string,
    mfg *string,
    serial *string,
    osType *string,
    osVersion *string,
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
| `distinct` | `*string` | Query, Optional | Field to count by. enum: `device_id`, `wifi_mac`, `wifi_ip`, `hostname`, `model`, `mfg`, `serial`, `os_type`, `os_version`<br><br>**Default**: `"os_type"` |
| `deviceId` | `*string` | Query, Optional | Filter by Marvis Client installation device UUID |
| `wifiMac` | `*string` | Query, Optional | Filter by device Wi-Fi MAC address |
| `wifiIp` | `*string` | Query, Optional | Filter by device Wi-Fi IP address |
| `hostname` | `*string` | Query, Optional | Filter by device hostname |
| `model` | `*string` | Query, Optional | Filter by device model |
| `mfg` | `*string` | Query, Optional | Filter by device manufacturer |
| `serial` | `*string` | Query, Optional | Filter by device serial number |
| `osType` | `*string` | Query, Optional | Filter by device OS type or platform |
| `osVersion` | `*string` | Query, Optional | Filter by device OS version |
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

distinct := "os_type"

limit := 100

duration := "10m"

apiResponse, err := orgsStatsMarvisClients.CountOrgMarvisClientsStats(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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


# Search Org Marvis Clients Stats

Search Marvis Client stats records across the organization.

```go
SearchOrgMarvisClientsStats(
    ctx context.Context,
    orgId uuid.UUID,
    deviceId *string,
    wifiMac *string,
    wifiIp *string,
    hostname *string,
    model *string,
    mfg *string,
    serial *string,
    osType *string,
    osVersion *string,
    limit *int,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.StatsMarvisClientsSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `*string` | Query, Optional | Filter by Marvis Client installation device UUID |
| `wifiMac` | `*string` | Query, Optional | Filter by device Wi-Fi MAC address |
| `wifiIp` | `*string` | Query, Optional | Filter by device Wi-Fi IP address |
| `hostname` | `*string` | Query, Optional | Filter by device hostname |
| `model` | `*string` | Query, Optional | Filter by device model |
| `mfg` | `*string` | Query, Optional | Filter by device manufacturer |
| `serial` | `*string` | Query, Optional | Filter by device serial number |
| `osType` | `*string` | Query, Optional | Filter by device OS type or platform |
| `osVersion` | `*string` | Query, Optional | Filter by device OS version |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |

## Response Type

**200**: Paginated Marvis Client stats search results

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.StatsMarvisClientsSearch](../../doc/models/stats-marvis-clients-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

duration := "10m"

apiResponse, err := orgsStatsMarvisClients.SearchOrgMarvisClientsStats(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "limit": 100,
  "results": [
    {
      "battery_charging": false,
      "battery_level": 82,
      "cpu_background": 2.1,
      "cpu_idle": 78.3,
      "cpu_system": 5.4,
      "cpu_user": 14.2,
      "device_id": "0c2a2c6c-5a95-4956-a02d-e1b39c2e5c6e",
      "hostname": "my-android-phone",
      "location": {
        "map_id": "b4695157-0d1d-4da0-8f9e-5c38149d8b81",
        "site_id": "d14f16d8-d14f-11e5-8e81-1258369c38a9",
        "timestamp": 1717027100,
        "x": 423.5,
        "y": 201.0
      },
      "memory_total": 8589934592,
      "memory_usage": 3758096384,
      "mfg": "Samsung",
      "model": "Galaxy S23",
      "org_id": "c168ddee-c14c-11e5-8e81-1258369c38a9",
      "os_type": "Android",
      "os_version": "14",
      "serial": "R5CX123456A",
      "storage_total": 128849018880,
      "storage_usage": 52428800000,
      "timestamp": 1717027200,
      "wifi_band": "5GHz",
      "wifi_bssid": "00:11:22:33:44:55",
      "wifi_channel": 36,
      "wifi_ip": "192.168.1.55",
      "wifi_mac": "00:aa:bb:cc:dd:ee",
      "wifi_rssi": -58,
      "wifi_ssid": "Corp-WiFi"
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

