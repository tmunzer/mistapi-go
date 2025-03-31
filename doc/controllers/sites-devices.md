# Sites Devices

```go
sitesDevices := client.SitesDevices()
```

## Class Name

`SitesDevices`

## Methods

* [Add Site Device Image](../../doc/controllers/sites-devices.md#add-site-device-image)
* [Change Site Switch Vc Port Mode](../../doc/controllers/sites-devices.md#change-site-switch-vc-port-mode)
* [Count Site Device Config History](../../doc/controllers/sites-devices.md#count-site-device-config-history)
* [Count Site Device Events](../../doc/controllers/sites-devices.md#count-site-device-events)
* [Count Site Device Last Config](../../doc/controllers/sites-devices.md#count-site-device-last-config)
* [Count Site Devices](../../doc/controllers/sites-devices.md#count-site-devices)
* [Delete Site Device Image](../../doc/controllers/sites-devices.md#delete-site-device-image)
* [Export Site Devices](../../doc/controllers/sites-devices.md#export-site-devices)
* [Get Site Device](../../doc/controllers/sites-devices.md#get-site-device)
* [Import Site Devices](../../doc/controllers/sites-devices.md#import-site-devices)
* [List Site Devices](../../doc/controllers/sites-devices.md#list-site-devices)
* [Search Site Device Config History](../../doc/controllers/sites-devices.md#search-site-device-config-history)
* [Search Site Device Events](../../doc/controllers/sites-devices.md#search-site-device-events)
* [Search Site Device Last Configs](../../doc/controllers/sites-devices.md#search-site-device-last-configs)
* [Search Site Devices](../../doc/controllers/sites-devices.md#search-site-devices)
* [Update Site Device](../../doc/controllers/sites-devices.md#update-site-device)


# Add Site Device Image

Attach up to 3 images to a device

```go
AddSiteDeviceImage(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    imageNumber int,
    file models.FileWrapper,
    json *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `imageNumber` | `int` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | Binary file |
| `json` | `*string` | Form, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageNumber := 110

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })



resp, err := sitesDevices.AddSiteDeviceImage(ctx, siteId, deviceId, imageNumber, file, nil)
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


# Change Site Switch Vc Port Mode

Change VCP port mode

Some switch model allows changing VCP port behaviors, e.g. - use them as regular network ports - change vcp protocol Note, this command will reboot the switch

```go
ChangeSiteSwitchVcPortMode(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.VcPort) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.VcPort`](../../doc/models/vc-port.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.VcPort{
    Mode:                 models.ToPointer(models.VcPortModeEnum_NETWORK),
}

resp, err := sitesDevices.ChangeSiteSwitchVcPortMode(ctx, siteId, deviceId, &body)
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


# Count Site Device Config History

Counts the number of entries in device config history for distinct field with given filters

```go
CountSiteDeviceConfigHistory(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *string,
    mac *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")









duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesDevices.CountSiteDeviceConfigHistory(ctx, siteId, nil, nil, nil, nil, &duration, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Count Site Device Events

Counts the number of entries in ap events history for distinct field with given filters

```go
CountSiteDeviceEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDeviceEventsCountDistinctEnum,
    model *string,
    mType *string,
    typeCode *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteDeviceEventsCountDistinctEnum`](../../doc/models/site-device-events-count-distinct-enum.md) | Query, Optional | **Default**: `"model"` |
| `model` | `*string` | Query, Optional | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `typeCode` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteDeviceEventsCountDistinctEnum_MODEL







limit := 100





duration := "10m"

apiResponse, err := sitesDevices.CountSiteDeviceEvents(ctx, siteId, &distinct, nil, nil, nil, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Count Site Device Last Config

Counts the number of entries in device config history for distinct field with given filters

```go
CountSiteDeviceLastConfig(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDeviceLastConfigCountDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteDeviceLastConfigCountDistinctEnum`](../../doc/models/site-device-last-config-count-distinct-enum.md) | Query, Optional | **Default**: `"mac"` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteDeviceLastConfigCountDistinctEnum_MAC





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesDevices.CountSiteDeviceLastConfig(ctx, siteId, &distinct, nil, nil, &duration, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Count Site Devices

Counts the number of entries in ap events history for distinct field with given filters

```go
CountSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDevicesCountDistinctEnum,
    hostname *string,
    model *string,
    mac *string,
    version *string,
    mxtunnelStatus *string,
    mxedgeId *string,
    lldpSystemName *string,
    lldpSystemDesc *string,
    lldpPortId *string,
    lldpMgmtAddr *string,
    mapId *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteDevicesCountDistinctEnum`](../../doc/models/site-devices-count-distinct-enum.md) | Query, Optional | **Default**: `"model"` |
| `hostname` | `*string` | Query, Optional | - |
| `model` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `version` | `*string` | Query, Optional | - |
| `mxtunnelStatus` | `*string` | Query, Optional | - |
| `mxedgeId` | `*string` | Query, Optional | - |
| `lldpSystemName` | `*string` | Query, Optional | - |
| `lldpSystemDesc` | `*string` | Query, Optional | - |
| `lldpPortId` | `*string` | Query, Optional | - |
| `lldpMgmtAddr` | `*string` | Query, Optional | - |
| `mapId` | `*string` | Query, Optional | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteDevicesCountDistinctEnum_MODEL



























duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesDevices.CountSiteDevices(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Site Device Image

Delete image from a device

```go
DeleteSiteDeviceImage(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    imageNumber int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `imageNumber` | `int` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageNumber := 110

resp, err := sitesDevices.DeleteSiteDeviceImage(ctx, siteId, deviceId, imageNumber)
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


# Export Site Devices

To download the exported device information

```go
ExportSiteDevices(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]byte],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []byte.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevices.ExportSiteDevices(ctx, siteId)
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


# Get Site Device

Get Device Configuration

```go
GetSiteDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.MistDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type models.MistDevice.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevices.GetSiteDevice(ctx, siteId, deviceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceAp: ", *r)
    } else if r, ok := responseBody.AsDeviceSwitch(); ok {
        fmt.Println("Value narrowed down to models.DeviceSwitch: ", *r)
    } else if r, ok := responseBody.AsDeviceGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "aeroscout": {
    "enabled": false,
    "host": "aero.pvt.net",
    "locate_connected": true
  },
  "ble_config": {
    "beacon_enabled": false,
    "beacon_rate": 3,
    "beacon_rate_mode": "custom",
    "beam_disabled": [
      1,
      3,
      6
    ],
    "custom_ble_packet_enabled": false,
    "custom_ble_packet_frame": "0x........",
    "custom_ble_packet_freq_msec": 300,
    "eddystone_uid_adv_power": -65,
    "eddystone_uid_beams": "2-4,7",
    "eddystone_uid_enabled": false,
    "eddystone_uid_freq_msec": 200,
    "eddystone_uid_instance": "5c5b35000001",
    "eddystone_uid_namespace": "2818e3868dec25629ede",
    "eddystone_url_adv_power": -65,
    "eddystone_url_beams": "2-4,7",
    "eddystone_url_enabled": true,
    "eddystone_url_freq_msec": 1000,
    "eddystone_url_url": "https://www.abc.com",
    "ibeacon_adv_power": -65,
    "ibeacon_beams": "2-4,7",
    "ibeacon_enabled": false,
    "ibeacon_freq_msec": 0,
    "ibeacon_major": 13,
    "ibeacon_minor": 138,
    "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
    "power": 6,
    "power_mode": "custom"
  },
  "centrak": {
    "enabled": false
  },
  "client_bridge": {
    "auth": {
      "psk": "foryoureyesonly",
      "type": "psk"
    },
    "enabled": false,
    "ssid": "Uplink-SSID"
  },
  "created_time": 0,
  "deviceprofile_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "esl_config": {
    "cacert": "string",
    "channel": 3,
    "enabled": false,
    "host": "1.1.1.1",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true,
    "vlan_id": 1
  },
  "for_site": true,
  "height": 2.75,
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6008",
  "image1_url": "string",
  "image2_url": "string",
  "image3_url": "string",
  "iot_config": {
    "A1": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A2": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A3": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A4": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "DI1": {
      "enabled": false,
      "name": "string",
      "pullup": "internal"
    },
    "DI2": {
      "enabled": false,
      "name": "string",
      "pullup": "internal"
    },
    "DO": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    }
  },
  "ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "dns_suffix": [
      ".mist.local",
      ".mist.com"
    ],
    "gateway": "10.2.1.254",
    "gateway6": "2607:f8b0:4005:808::1",
    "ip": "10.2.1.1",
    "ip6": "2607:f8b0:4005:808::2004",
    "mtu": 1500,
    "netmask": "255.255.255.0",
    "netmask6": "/32",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "led": {
    "brightness": 255,
    "enabled": true
  },
  "locked": true,
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "mesh": {
    "enabled": false,
    "group": 1,
    "role": "base"
  },
  "modified_time": 0,
  "name": "conference room",
  "notes": "slightly off center",
  "ntp_servers": [
    "string"
  ],
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "orientation": 45,
  "poe_passthrough": false,
  "pwr_config": {
    "base": 2000,
    "prefer_usb_over_wifi": false
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "type": "ap",
  "uplink_port_config": {
    "dot1x": false,
    "keep_wlans_up_if_down": false
  },
  "usb_config": {
    "cacert": "string",
    "channel": 3,
    "enabled": true,
    "host": "1.1.1.1",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true,
    "vlan_id": 1
  },
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "x": 53.5,
  "y": 173.1
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


# Import Site Devices

Import Information for Multiple Devices

CSV format:

```csv
mac,name,map_id,x,y,height,orientation,labels,band_24.power,band_24.bandwidth,band_24.channel,band_24.disabled,band_5.power,band_5.bandwidth,band_5.channel,band_5.disabled,band_6.power,band_6.bandwidth,band_6.channel,band_6.disabled
5c5b53010101,"AP 1",845a23bf-bed9-e43c-4c86-6fa474be7ae5,30,10,2.3,45,"guest, campus, vip",1,20,0,false,0,40,0,false,17,80,0,false
```

```go
ImportSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    file models.FileWrapper) (
    models.ApiResponse[[]models.ConfigDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | File to upload |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []models.ConfigDevice.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })

apiResponse, err := sitesDevices.ImportSiteDevices(ctx, siteId, file)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsDeviceAp(); ok {
            fmt.Println("Value narrowed down to models.DeviceAp: ", *i)
        } else if i, ok := item.AsDeviceSwitch(); ok {
            fmt.Println("Value narrowed down to models.DeviceSwitch: ", *i)
        } else if i, ok := item.AsDeviceGateway(); ok {
            fmt.Println("Value narrowed down to models.DeviceGateway: ", *i)
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


# List Site Devices

Get list of devices on the site.

```go
ListSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeWithAllEnum,
    name *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.ConfigDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeWithAllEnum`](../../doc/models/device-type-with-all-enum.md) | Query, Optional | **Default**: `"ap"` |
| `name` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []models.ConfigDevice.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeWithAllEnum_AP



limit := 100

page := 1

apiResponse, err := sitesDevices.ListSiteDevices(ctx, siteId, &mType, nil, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsDeviceAp(); ok {
            fmt.Println("Value narrowed down to models.DeviceAp: ", *i)
        } else if i, ok := item.AsDeviceSwitch(); ok {
            fmt.Println("Value narrowed down to models.DeviceSwitch: ", *i)
        } else if i, ok := item.AsDeviceGateway(); ok {
            fmt.Println("Value narrowed down to models.DeviceGateway: ", *i)
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


# Search Site Device Config History

Search for entries in device config history

```go
SearchSiteDeviceConfigHistory(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeDefaultApEnum,
    mac *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseConfigHistorySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | **Default**: `"ap"` |
| `mac` | `*string` | Query, Optional | Device MAC Address |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseConfigHistorySearch](../../doc/models/response-config-history-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeDefaultApEnum_AP



limit := 100





duration := "10m"

apiResponse, err := sitesDevices.SearchSiteDeviceConfigHistory(ctx, siteId, &mType, nil, &limit, nil, nil, &duration)
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
  "end": 1531862583,
  "limit": 10,
  "results": [
    {
      "channel_24": 11,
      "channel_5": 100,
      "radio_macs": [
        "5c5b352e000a",
        "5c5b352e000b",
        "5c5b352e000c"
      ],
      "radios": [
        {
          "band": "24",
          "channel": 11
        },
        {
          "band": "5",
          "channel": 100
        }
      ],
      "secpolicy_violated": false,
      "ssids": [
        "test24",
        "test5"
      ],
      "ssids_24": [
        "test24"
      ],
      "ssids_5": [
        "test5"
      ],
      "timestamp": 1531855856.643369,
      "version": "apfw-0.2.14754-cersei-75c8",
      "wlans": [
        {
          "auth": "psk",
          "bands": [
            "24"
          ],
          "id": "be22bba7-8e22-e1cf-5185-b880816fe2cf",
          "ssid": "test24",
          "vlan_ids": [
            "1"
          ]
        },
        {
          "auth": "psk",
          "bands": [
            "5"
          ],
          "id": "f8c18724-4118-3487-811a-f98964988604",
          "ssid": "test5",
          "vlan_ids": [
            "1"
          ]
        }
      ]
    }
  ],
  "start": 1531776183,
  "total": 1
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


# Search Site Device Events

Search Devices Events

```go
SearchSiteDeviceEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    model *string,
    text *string,
    timestamp *string,
    mType *string,
    lastBy *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseEventsDevices],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Device mac |
| `model` | `*string` | Query, Optional | Device model |
| `text` | `*string` | Query, Optional | Event message |
| `timestamp` | `*string` | Query, Optional | Event time |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `lastBy` | `*string` | Query, Optional | Return last/recent event for passed in field |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsDevices](../../doc/models/response-events-devices.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")











lastBy := "port_id"

limit := 100





duration := "10m"

apiResponse, err := sitesDevices.SearchSiteDeviceEvents(ctx, siteId, nil, nil, nil, nil, nil, &lastBy, &limit, nil, nil, &duration)
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
  "end": 1531862583,
  "limit": 2,
  "next": "/api/v1/sites/8aaba0aa-09cc-44bd-9709-33b98040550c/devices/events/search?ap=5c5b350e0001&end=1531855849.000&limit=2&start=1531776183.0",
  "results": [
    {
      "last_reboot_time": 1531854327,
      "text": "Success",
      "timestamp": 1531855849.226722,
      "type": "AP_CONNECT_STATUS",
      "type_code": 2002
    },
    {
      "timestamp": 1531854326,
      "type": "AP_CONFIGURED"
    }
  ],
  "start": 1531776183,
  "total": 14
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


# Search Site Device Last Configs

Search Device Last Configs

```go
SearchSiteDeviceLastConfigs(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeDefaultApEnum,
    mac *string,
    version *string,
    name *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseConfigHistorySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | **Default**: `"ap"` |
| `mac` | `*string` | Query, Optional | - |
| `version` | `*string` | Query, Optional | - |
| `name` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseConfigHistorySearch](../../doc/models/response-config-history-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeDefaultApEnum_AP







limit := 100





duration := "10m"

apiResponse, err := sitesDevices.SearchSiteDeviceLastConfigs(ctx, siteId, &mType, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 1531862583,
  "limit": 10,
  "results": [
    {
      "channel_24": 11,
      "channel_5": 100,
      "radio_macs": [
        "5c5b352e000a",
        "5c5b352e000b",
        "5c5b352e000c"
      ],
      "radios": [
        {
          "band": "24",
          "channel": 11
        },
        {
          "band": "5",
          "channel": 100
        }
      ],
      "secpolicy_violated": false,
      "ssids": [
        "test24",
        "test5"
      ],
      "ssids_24": [
        "test24"
      ],
      "ssids_5": [
        "test5"
      ],
      "timestamp": 1531855856.643369,
      "version": "apfw-0.2.14754-cersei-75c8",
      "wlans": [
        {
          "auth": "psk",
          "bands": [
            "24"
          ],
          "id": "be22bba7-8e22-e1cf-5185-b880816fe2cf",
          "ssid": "test24",
          "vlan_ids": [
            "1"
          ]
        },
        {
          "auth": "psk",
          "bands": [
            "5"
          ],
          "id": "f8c18724-4118-3487-811a-f98964988604",
          "ssid": "test5",
          "vlan_ids": [
            "1"
          ]
        }
      ]
    }
  ],
  "start": 1531776183,
  "total": 1
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


# Search Site Devices

Search Device

```go
SearchSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    hostname *string,
    mType *models.DeviceTypeDefaultApEnum,
    model *string,
    mac *string,
    version *string,
    powerConstrained *bool,
    ipAddress *string,
    mxtunnelStatus *models.SearchSiteDevicesMxtunnelStatusEnum,
    mxedgeId *uuid.UUID,
    lldpSystemName *string,
    lldpSystemDesc *string,
    lldpPortId *string,
    lldpMgmtAddr *string,
    band24Channel *int,
    band5Channel *int,
    band6Channel *int,
    band24Bandwidth *int,
    band5Bandwidth *int,
    band6Bandwidth *int,
    eth0PortSpeed *int,
    sort *models.SearchSiteDevicesSortEnum,
    descSort *models.SearchSiteDevicesDescSortEnum,
    stats *bool,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseDeviceSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | **Default**: `"ap"` |
| `model` | `*string` | Query, Optional | Device model |
| `mac` | `*string` | Query, Optional | Device MAC |
| `version` | `*string` | Query, Optional | Version |
| `powerConstrained` | `*bool` | Query, Optional | power_constrained |
| `ipAddress` | `*string` | Query, Optional | - |
| `mxtunnelStatus` | [`*models.SearchSiteDevicesMxtunnelStatusEnum`](../../doc/models/search-site-devices-mxtunnel-status-enum.md) | Query, Optional | MxTunnel status, up / down |
| `mxedgeId` | `*uuid.UUID` | Query, Optional | Mist Edge id, if AP is connecting to a Mist Edge |
| `lldpSystemName` | `*string` | Query, Optional | LLDP system name |
| `lldpSystemDesc` | `*string` | Query, Optional | LLDP system description |
| `lldpPortId` | `*string` | Query, Optional | LLDP port id |
| `lldpMgmtAddr` | `*string` | Query, Optional | LLDP management ip address |
| `band24Channel` | `*int` | Query, Optional | Channel of band_24 |
| `band5Channel` | `*int` | Query, Optional | Channel of band_5 |
| `band6Channel` | `*int` | Query, Optional | Channel of band_6 |
| `band24Bandwidth` | `*int` | Query, Optional | Bandwidth of band_24 |
| `band5Bandwidth` | `*int` | Query, Optional | Bandwidth of band_5 |
| `band6Bandwidth` | `*int` | Query, Optional | Bandwidth of band_6 |
| `eth0PortSpeed` | `*int` | Query, Optional | Port speed of eth0 |
| `sort` | [`*models.SearchSiteDevicesSortEnum`](../../doc/models/search-site-devices-sort-enum.md) | Query, Optional | Sort options<br>**Default**: `"timestamp"` |
| `descSort` | [`*models.SearchSiteDevicesDescSortEnum`](../../doc/models/search-site-devices-desc-sort-enum.md) | Query, Optional | Sort options in reverse order |
| `stats` | `*bool` | Query, Optional | Whether to return device stats<br>**Default**: `false` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceSearch](../../doc/models/response-device-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



mType := models.DeviceTypeDefaultApEnum_AP









ipAddress := "192.168.1.1"



























sort := models.SearchSiteDevicesSortEnum_TIMESTAMP



stats := false

limit := 100





duration := "10m"

apiResponse, err := sitesDevices.SearchSiteDevices(ctx, siteId, nil, &mType, nil, nil, nil, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &sort, nil, &stats, &limit, nil, nil, &duration)
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
  "next": "string",
  "results": [
    {
      "hostname": [
        "AP41-STB-3E5299-WH-2001",
        "AP41-STB-3E5299-WH-50",
        "AP41-STB-3E5299",
        "5c5b353e5299"
      ],
      "ip": "10.2.16.205",
      "lldp_mgmt_addr": "10.2.10.139",
      "lldp_port_desc": "GigabitEthernet1/0/1",
      "lldp_port_id": "Gi1/0/1",
      "lldp_system_desc": "Cisco IOS Software, C2960S Software (C2960S-UNIVERSALK9-M), Version 15.2(1)E1, RELEASE SOFTWARE (fc2)\nTechnical Support: https://www.cisco.com/techsupport\nCopyright (c) 1986-2013 by Cisco Systems, Inc.\nCompiled Fri 22-Nov-13 07:10 by prod_rel_team",
      "lldp_system_name": "ME-DC-1-ACC-SW",
      "mac": "5c5b353e5299",
      "model": "AP41",
      "mxedge_id": "00000000-0000-0000-1000-43a81f238391",
      "mxtunnel_status": "down",
      "org_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
      "power_constrained": false,
      "power_opmode": "",
      "site_id": "a8178443-ecb5-461c-b854-f16627619ab3",
      "sku": "AP41-US",
      "timestamp": 1596588619.007,
      "type": "ap",
      "uptime": 85280,
      "version": "0.7.20216",
      "wlans": [
        {
          "id": "28c36fc7-dc22-4960-9d81-34087511c2e5",
          "ssid": "Live-Demo-NAC"
        },
        {
          "id": "51b82e2b-f9e8-470b-a32a-cecde5501b0f",
          "ssid": "Live-Demo"
        }
      ]
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Site Device

Update Device Configuration

```go
UpdateSiteDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.MistDevice) (
    models.ApiResponse[models.MistDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MistDevice`](../../doc/models/containers/mist-device.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type models.MistDevice.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MistDeviceContainer.FromDeviceAp(models.DeviceAp{
    DeviceprofileId:      models.NewOptional(models.ToPointer(uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"))),
    DisableEth1:          models.ToPointer(false),
    DisableEth2:          models.ToPointer(false),
    DisableEth3:          models.ToPointer(false),
    DisableModule:        models.ToPointer(false),
    FlowControl:          models.ToPointer(false),
    Height:               models.ToPointer(float64(2.75)),
    MapId:                models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
    Name:                 models.ToPointer("conference room"),
    Notes:                models.ToPointer("slightly off center"),
    Orientation:          models.ToPointer(45),
    PoePassthrough:       models.ToPointer(false),
    Vars:                 map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    X:                    models.ToPointer(float64(53.5)),
    Y:                    models.ToPointer(float64(173.1)),
})

apiResponse, err := sitesDevices.UpdateSiteDevice(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceAp: ", *r)
    } else if r, ok := responseBody.AsDeviceSwitch(); ok {
        fmt.Println("Value narrowed down to models.DeviceSwitch: ", *r)
    } else if r, ok := responseBody.AsDeviceGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "aeroscout": {
    "enabled": false,
    "host": "aero.pvt.net",
    "locate_connected": true
  },
  "ble_config": {
    "beacon_enabled": false,
    "beacon_rate": 3,
    "beacon_rate_mode": "custom",
    "beam_disabled": [
      1,
      3,
      6
    ],
    "custom_ble_packet_enabled": false,
    "custom_ble_packet_frame": "0x........",
    "custom_ble_packet_freq_msec": 300,
    "eddystone_uid_adv_power": -65,
    "eddystone_uid_beams": "2-4,7",
    "eddystone_uid_enabled": false,
    "eddystone_uid_freq_msec": 200,
    "eddystone_uid_instance": "5c5b35000001",
    "eddystone_uid_namespace": "2818e3868dec25629ede",
    "eddystone_url_adv_power": -65,
    "eddystone_url_beams": "2-4,7",
    "eddystone_url_enabled": true,
    "eddystone_url_freq_msec": 1000,
    "eddystone_url_url": "https://www.abc.com",
    "ibeacon_adv_power": -65,
    "ibeacon_beams": "2-4,7",
    "ibeacon_enabled": false,
    "ibeacon_freq_msec": 0,
    "ibeacon_major": 13,
    "ibeacon_minor": 138,
    "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
    "power": 6,
    "power_mode": "custom"
  },
  "centrak": {
    "enabled": false
  },
  "client_bridge": {
    "auth": {
      "psk": "foryoureyesonly",
      "type": "psk"
    },
    "enabled": false,
    "ssid": "Uplink-SSID"
  },
  "created_time": 0,
  "deviceprofile_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "esl_config": {
    "cacert": "string",
    "channel": 3,
    "enabled": false,
    "host": "1.1.1.1",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true,
    "vlan_id": 1
  },
  "for_site": true,
  "height": 2.75,
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6008",
  "image1_url": "string",
  "image2_url": "string",
  "image3_url": "string",
  "iot_config": {
    "A1": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A2": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A3": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A4": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "DI1": {
      "enabled": false,
      "name": "string",
      "pullup": "internal"
    },
    "DI2": {
      "enabled": false,
      "name": "string",
      "pullup": "internal"
    },
    "DO": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    }
  },
  "ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "dns_suffix": [
      ".mist.local",
      ".mist.com"
    ],
    "gateway": "10.2.1.254",
    "gateway6": "2607:f8b0:4005:808::1",
    "ip": "10.2.1.1",
    "ip6": "2607:f8b0:4005:808::2004",
    "mtu": 1500,
    "netmask": "255.255.255.0",
    "netmask6": "/32",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "led": {
    "brightness": 255,
    "enabled": true
  },
  "locked": true,
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "mesh": {
    "enabled": false,
    "group": 1,
    "role": "base"
  },
  "modified_time": 0,
  "name": "conference room",
  "notes": "slightly off center",
  "ntp_servers": [
    "string"
  ],
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "orientation": 45,
  "poe_passthrough": false,
  "pwr_config": {
    "base": 2000,
    "prefer_usb_over_wifi": false
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "type": "ap",
  "uplink_port_config": {
    "dot1x": false,
    "keep_wlans_up_if_down": false
  },
  "usb_config": {
    "cacert": "string",
    "channel": 3,
    "enabled": true,
    "host": "1.1.1.1",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true,
    "vlan_id": 1
  },
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "x": 53.5,
  "y": 173.1
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

