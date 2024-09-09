# Sites Stats-Devices

```go
sitesStatsDevices := client.SitesStatsDevices()
```

## Class Name

`SitesStatsDevices`

## Methods

* [Get Site All Clients Stats by Device](../../doc/controllers/sites-stats-devices.md#get-site-all-clients-stats-by-device)
* [Get Site Device Stats](../../doc/controllers/sites-stats-devices.md#get-site-device-stats)
* [Get Site Gateway Metrics](../../doc/controllers/sites-stats-devices.md#get-site-gateway-metrics)
* [Get Site Switches Metrics](../../doc/controllers/sites-stats-devices.md#get-site-switches-metrics)
* [List Site Devices Stats](../../doc/controllers/sites-stats-devices.md#list-site-devices-stats)


# Get Site All Clients Stats by Device

Get wireless client stat by Device

```go
GetSiteAllClientsStatsByDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[[]models.StatsWirelessClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.StatsWirelessClient`](../../doc/models/stats-wireless-client.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsDevices.GetSiteAllClientsStatsByDevice(ctx, siteId, deviceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# Get Site Device Stats

Get Site Device Stats Details

```go
GetSiteDeviceStats(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.StatsDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.StatsDevice`](../../doc/models/containers/stats-device.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsDevices.GetSiteDeviceStats(ctx, siteId, deviceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsStatsAp(); ok {
        fmt.Println("Value narrowed down to models.StatsAp: ", *r)
    } else if r, ok := responseBody.AsStatsSwitch(); ok {
        fmt.Println("Value narrowed down to models.StatsSwitch: ", *r)
    } else if r, ok := responseBody.AsStatsGateway(); ok {
        fmt.Println("Value narrowed down to models.StatsGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
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


# Get Site Gateway Metrics

Get Site Gateway Metrics

```go
GetSiteGatewayMetrics(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.GatewayMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.GatewayMetrics`](../../doc/models/gateway-metrics.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsDevices.GetSiteGatewayMetrics(ctx, siteId)
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
  "config_success": 99.9,
  "version_compliance": {
    "major_version": {
      "SRX320": {
        "major_count": 0,
        "major_version": "19.4R2-S1.2"
      }
    },
    "score": 99.9,
    "type": "gateway"
  }
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


# Get Site Switches Metrics

Get version compliance metrics for managed or monitored switches

```go
GetSiteSwitchesMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.SwitchMetricTypeEnum,
    scope *models.SwitchMetricScopeEnum,
    switchMac *string) (
    models.ApiResponse[models.ResponseSwitchMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.SwitchMetricTypeEnum`](../../doc/models/switch-metric-type-enum.md) | Query, Optional | - |
| `scope` | [`*models.SwitchMetricScopeEnum`](../../doc/models/switch-metric-scope-enum.md) | Query, Optional | - |
| `switchMac` | `*string` | Query, Optional | switch mac, used only with metric `type`==`active_ports_summary` |

## Response Type

[`models.ResponseSwitchMetrics`](../../doc/models/response-switch-metrics.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")







apiResponse, err := sitesStatsDevices.GetSiteSwitchesMetrics(ctx, siteId, nil, nil, nil)
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
  "active_ports_summary": {
    "details": {
      "active_port_count": 4,
      "total_port_count": 4
    },
    "score": 100,
    "total_switch_count": 2
  },
  "config_success": {
    "details": {
      "config_success_count": 2
    },
    "score": 100,
    "total_switch_count": 2
  },
  "version_compliance": {
    "details": {
      "major_versions": [
        {
          "major_count": 1,
          "major_version": "21.4R3.5",
          "model": "EX2300-C-12P",
          "system_names": []
        },
        {
          "major_count": 1,
          "major_version": "6.0.4-11",
          "model": "SSR120",
          "system_names": []
        }
      ]
    },
    "score": 100,
    "total_switch_count": 2
  }
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


# List Site Devices Stats

Get List of Site Devices Stats

```go
ListSiteDevicesStats(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeWithAllEnum,
    status *models.StatDeviceStatusFilterEnum,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeWithAllEnum`](../../doc/models/device-type-with-all-enum.md) | Query, Optional | **Default**: `"ap"` |
| `status` | [`*models.StatDeviceStatusFilterEnum`](../../doc/models/stat-device-status-filter-enum.md) | Query, Optional | **Default**: `"all"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.StatsDevice`](../../doc/models/containers/stats-device.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeWithAllEnum("ap")

status := models.StatDeviceStatusFilterEnum("all")

limit := 100

page := 1

apiResponse, err := sitesStatsDevices.ListSiteDevicesStats(ctx, siteId, &mType, &status, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsStatsAp(); ok {
            fmt.Println("Value narrowed down to models.StatsAp: ", *i)
        } else if i, ok := item.AsStatsSwitch(); ok {
            fmt.Println("Value narrowed down to models.StatsSwitch: ", *i)
        } else if i, ok := item.AsStatsGateway(); ok {
            fmt.Println("Value narrowed down to models.StatsGateway: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
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

