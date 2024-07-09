# Sites Devices-Others

```go
sitesDevicesOthers := client.SitesDevicesOthers()
```

## Class Name

`SitesDevicesOthers`

## Methods

* [Count Site Other Device Events](../../doc/controllers/sites-devices-others.md#count-site-other-device-events)
* [List Site Other Devices](../../doc/controllers/sites-devices-others.md#list-site-other-devices)
* [Search Site Other Device Events](../../doc/controllers/sites-devices-others.md#search-site-other-device-events)


# Count Site Other Device Events

Count Site OtherDevices Events

```go
CountSiteOtherDeviceEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteOtherDeviceEventsCountDistinctEnum,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteOtherDeviceEventsCountDistinctEnum`](../../doc/models/site-other-device-events-count-distinct-enum.md) | Query, Optional | - |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteOtherDeviceEventsCountDistinctEnum("mac")







duration := "10m"

limit := 100

apiResponse, err := sitesDevicesOthers.CountSiteOtherDeviceEvents(ctx, siteId, &distinct, nil, nil, nil, &duration, &limit)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Other Devices

Get List of Site other devices (3rd party devices)

```go
ListSiteOtherDevices(
    ctx context.Context,
    siteId uuid.UUID,
    vendor *string,
    mac *string,
    serial *string,
    model *string,
    name *string,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.DeviceOther],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `vendor` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `serial` | `*string` | Query, Optional | - |
| `model` | `*string` | Query, Optional | - |
| `name` | `*string` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.DeviceOther`](../../doc/models/device-other.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")











page := 1

limit := 100

apiResponse, err := sitesDevicesOthers.ListSiteOtherDevices(ctx, siteId, nil, nil, nil, nil, nil, &page, &limit)
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
    "created_time": 1676983730,
    "device_mac": "001122334455",
    "id": "ae9dee49-69e7-4710-a114-5b827a777738",
    "mac": "5c5b35000018",
    "model": "AP41",
    "modified_time": 1676983730,
    "name": "hallway",
    "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
    "serial": "FXLH2015150025",
    "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
    "vendor": "cradlepoint"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Other Device Events

Search Site OtherDevices Events

```go
SearchSiteOtherDeviceEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    deviceMac *string,
    vendor *string,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsOtherDevicesSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | mac |
| `deviceMac` | `*string` | Query, Optional | mac of attached device |
| `vendor` | `*string` | Query, Optional | vendor name |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`models.ResponseEventsOtherDevicesSearch`](../../doc/models/response-events-other-devices-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













duration := "10m"

limit := 100

apiResponse, err := sitesDevicesOthers.SearchSiteOtherDeviceEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
  "end": 0,
  "limit": 0,
  "results": {
    "device_mac": "string",
    "mac": "5c5b351e13b5",
    "org_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862a",
    "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
    "text": "Plugged: The Internal 5GB (SIM1) has been inserted into Internal 1.",
    "timestamp": 547235620.89,
    "type": "CELLULAR_EDGE_MODEM_WAN_PLUGGED",
    "vendor": "cradlepoint"
  },
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

