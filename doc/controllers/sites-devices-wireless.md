# Sites Devices-Wireless

```go
sitesDevicesWireless := client.SitesDevicesWireless()
```

## Class Name

`SitesDevicesWireless`

## Methods

* [Get Site Device Iot Port](../../doc/controllers/sites-devices-wireless.md#get-site-device-iot-port)
* [Get Site Device Radio Channels](../../doc/controllers/sites-devices-wireless.md#get-site-device-radio-channels)
* [Set Site Device Iot Port](../../doc/controllers/sites-devices-wireless.md#set-site-device-iot-port)


# Get Site Device Iot Port

Returns the current state of each enabled IoT pin configured as an output.

```go
GetSiteDeviceIotPort(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[map[string]int],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type map[string]int.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevicesWireless.GetSiteDeviceIotPort(ctx, siteId, deviceId)
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
{
  "A1": 1,
  "DO": 0
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


# Get Site Device Radio Channels

Get a list of allowed channels (per channel width)

```go
GetSiteDeviceRadioChannels(
    ctx context.Context,
    siteId uuid.UUID,
    countryCode *string) (
    models.ApiResponse[models.ResponseDeviceRadioChannels],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `countryCode` | `*string` | Query, Optional | Country code for the site (for AP config generation), in [two-character](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceRadioChannels](../../doc/models/response-device-radio-channels.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

countryCode := "US"

apiResponse, err := sitesDevicesWireless.GetSiteDeviceRadioChannels(ctx, siteId, &countryCode)
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
  "band24_40mhz_allowed": false,
  "band24_channels": {
    "20": [
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9,
      10,
      11
    ],
    "40": [
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9,
      10,
      11
    ]
  },
  "band24_enabled": true,
  "band5_channels": {
    "20": [
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
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161,
      165
    ],
    "40": [
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
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161
    ],
    "80": [
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
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161
    ],
    "dfs": [
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144
    ],
    "outdoor": [
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
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161,
      165
    ]
  },
  "band5_enabled": true,
  "certified": true,
  "code": 840,
  "dfs_ok": true,
  "key": "US",
  "name": "United States",
  "uses": "US_FCC"
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


# Set Site Device Iot Port

**Note**: For each IoT pin referenced:

* The pin must be enabled using the Device `iot_config` API
* The pin must support the output direction

```go
SetSiteDeviceIotPort(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body map[string]int) (
    models.ApiResponse[map[string]int],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | `map[string]int` | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type map[string]int.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := sitesDevicesWireless.SetSiteDeviceIotPort(ctx, siteId, deviceId, nil)
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
{
  "A1": 1,
  "DO": 0
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

