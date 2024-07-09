# Sites Devices

```go
sitesDevices := client.SitesDevices()
```

## Class Name

`SitesDevices`

## Methods

* [Add Site Device Image](../../doc/controllers/sites-devices.md#add-site-device-image)
* [Count Site Device Config History](../../doc/controllers/sites-devices.md#count-site-device-config-history)
* [Count Site Device Events](../../doc/controllers/sites-devices.md#count-site-device-events)
* [Count Site Device Last Config](../../doc/controllers/sites-devices.md#count-site-device-last-config)
* [Count Site Devices](../../doc/controllers/sites-devices.md#count-site-devices)
* [Delete Site Device Image](../../doc/controllers/sites-devices.md#delete-site-device-image)
* [Export Site Devices](../../doc/controllers/sites-devices.md#export-site-devices)
* [Get Site Device](../../doc/controllers/sites-devices.md#get-site-device)
* [Import Site Devices](../../doc/controllers/sites-devices.md#import-site-devices)
* [List Site Devices](../../doc/controllers/sites-devices.md#list-site-devices)
* [Restart Site Multiple Devices](../../doc/controllers/sites-devices.md#restart-site-multiple-devices)
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
| `file` | `models.FileWrapper` | Form, Required | binary file |
| `json` | `*string` | Form, Optional | JSON string describing your upload |

## Response Type

``

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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Count Site Device Config History

Counts the number of entries in device config history for distinct field with given filters

```go
CountSiteDeviceConfigHistory(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *string,
    mac *string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesDevices.CountSiteDeviceConfigHistory(ctx, siteId, nil, nil, &page, &limit, nil, nil, &duration)
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
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteDeviceEventsCountDistinctEnum`](../../doc/models/site-device-events-count-distinct-enum.md) | Query, Optional | - |
| `model` | `*string` | Query, Optional | - |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `typeCode` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteDeviceEventsCountDistinctEnum("model")







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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Count Site Device Last Config

Counts the number of entries in device config history for distinct field with given filters

```go
CountSiteDeviceLastConfig(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDeviceLastConfigCountDistinctEnum,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteDeviceLastConfigCountDistinctEnum`](../../doc/models/site-device-last-config-count-distinct-enum.md) | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteDeviceLastConfigCountDistinctEnum("mac")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesDevices.CountSiteDeviceLastConfig(ctx, siteId, &distinct, &page, &limit, nil, nil, &duration)
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
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteDevicesCountDistinctEnum`](../../doc/models/site-devices-count-distinct-enum.md) | Query, Optional | - |
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
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteDevicesCountDistinctEnum("model")























page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesDevices.CountSiteDevices(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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

``

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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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

`[]byte`

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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device

Get Device Configuration

```go
GetSiteDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.GetSiteDeviceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.GetSiteDeviceResponse`](../../doc/models/containers/get-site-device-response.md)

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
    "power": 10,
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
    "mtu": 0,
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
  "port_config": {
    "property1": {
      "additional_vlan_ids": [
        55,
        66
      ],
      "authentication_protocol": "pap",
      "disabled": true,
      "dynamic_vlan": {
        "default_vlan_id": 999,
        "enabled": true,
        "type": "string",
        "vlans": {
          "1-10": null,
          "user": null
        }
      },
      "enable_mac_auth": false,
      "forwarding": "all",
      "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
      "mxtunnel_name": "string",
      "port_auth": "none",
      "port_vlan_id": 1,
      "radius_config": {
        "acct_interim_interval": 0,
        "acct_servers": [
          {
            "host": "1.2.3.4",
            "keywrap_enabled": true,
            "keywrap_format": "hex",
            "keywrap_kek": "1122334455",
            "keywrap_mack": "1122334455",
            "port": 1813,
            "secret": "testing123"
          }
        ],
        "auth_servers": [
          {
            "host": "1.2.3.4",
            "keywrap_enabled": true,
            "keywrap_format": "hex",
            "keywrap_kek": "1122334455",
            "keywrap_mack": "1122334455",
            "port": 1812,
            "secret": "testing123"
          }
        ],
        "auth_servers_retries": 3,
        "auth_servers_timeout": 5,
        "coa_enabled": false,
        "coa_port": 3799,
        "network": "string",
        "source_ip": "string"
      },
      "radsec": {
        "enabled": true,
        "idle_timeout": 60,
        "mxcluster_ids": [
          "572586b7-f97b-a22b-526c-8b97a3f609c4"
        ],
        "proxy_hosts": [
          "mxedge1.local"
        ],
        "server_name": "radsec.abc.com",
        "servers": [
          {
            "host": "1.1.1.1",
            "port": 1812
          }
        ],
        "use_mxedge": true,
        "use_site_mxedge": false
      },
      "vlan_id": 9,
      "vland_ids": [
        1,
        10,
        50
      ],
      "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
      "wxtunnel_remote_id": "wifiguest"
    },
    "property2": {
      "additional_vlan_ids": [
        55,
        66
      ],
      "authentication_protocol": "pap",
      "disabled": true,
      "dynamic_vlan": {
        "default_vlan_id": 999,
        "enabled": true,
        "type": "string",
        "vlans": {
          "1-10": null,
          "user": null
        }
      },
      "enable_mac_auth": false,
      "forwarding": "all",
      "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
      "mxtunnel_name": "string",
      "port_auth": "none",
      "port_vlan_id": 1,
      "radius_config": {
        "acct_interim_interval": 0,
        "acct_servers": [
          {
            "host": "1.2.3.4",
            "keywrap_enabled": true,
            "keywrap_format": "hex",
            "keywrap_kek": "1122334455",
            "keywrap_mack": "1122334455",
            "port": 1813,
            "secret": "testing123"
          }
        ],
        "auth_servers": [
          {
            "host": "1.2.3.4",
            "keywrap_enabled": true,
            "keywrap_format": "hex",
            "keywrap_kek": "1122334455",
            "keywrap_mack": "1122334455",
            "port": 1812,
            "secret": "testing123"
          }
        ],
        "auth_servers_retries": 3,
        "auth_servers_timeout": 5,
        "coa_enabled": false,
        "coa_port": 3799,
        "network": "string",
        "source_ip": "string"
      },
      "radsec": {
        "enabled": true,
        "idle_timeout": 60,
        "mxcluster_ids": [
          "572586b7-f97b-a22b-526c-8b97a3f609c4"
        ],
        "proxy_hosts": [
          "mxedge1.local"
        ],
        "server_name": "radsec.abc.com",
        "servers": [
          {
            "host": "1.1.1.1",
            "port": 1812
          }
        ],
        "use_mxedge": true,
        "use_site_mxedge": false
      },
      "vlan_id": 9,
      "vland_ids": [
        1,
        10,
        50
      ],
      "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
      "wxtunnel_remote_id": "wifiguest"
    }
  },
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
    body []models.DeviceAp) (
    models.ApiResponse[[]models.ImportSiteDevicesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`[]models.DeviceAp`](../../doc/models/device-ap.md) | Body, Optional | - |

## Response Type

[`[]models.ImportSiteDevicesResponse`](../../doc/models/containers/import-site-devices-response.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := []models.DeviceAp{
    models.DeviceAp{
        DeviceprofileId:  models.NewOptional(models.ToPointer(uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"))),
        DisableEth1:      models.ToPointer(false),
        DisableEth2:      models.ToPointer(false),
        DisableEth3:      models.ToPointer(false),
        DisableModule:    models.ToPointer(false),
        Height:           models.ToPointer(float64(2.75)),
        MapId:            models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
        Name:             models.ToPointer("conference room"),
        Notes:            models.ToPointer("slightly off center"),
        OrgId:            models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Orientation:      models.ToPointer(45),
        PoePassthrough:   models.ToPointer(false),
        SiteId:           models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Vars:             map[string]string{
            "RADIUS_IP1": "172.31.2.5",
            "RADIUS_SECRET": "11s64632d",
        },
        X:                models.ToPointer(float64(53.5)),
        Y:                models.ToPointer(float64(173.1)),
    },
}

apiResponse, err := sitesDevices.ImportSiteDevices(ctx, siteId, body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Devices

Get list of devices on the site.

```go
ListSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeWithAllEnum,
    name *string,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.ListSiteDevicesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeWithAllEnum`](../../doc/models/device-type-with-all-enum.md) | Query, Optional | - |
| `name` | `*string` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.ListSiteDevicesResponse`](../../doc/models/containers/list-site-devices-response.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeWithAllEnum("ap")



page := 1

limit := 100

apiResponse, err := sitesDevices.ListSiteDevices(ctx, siteId, &mType, nil, &page, &limit)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Restart Site Multiple Devices

Note that only the devices that are connected will be restarted.

```go
RestartSiteMultipleDevices(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UtilsDevicesRestartMulti) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsDevicesRestartMulti`](../../doc/models/utils-devices-restart-multi.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsDevicesRestartMulti{
    DeviceIds: []uuid.UUID{
        uuid.MustParse("00000000-0000-0000-1000-5c5b35584a6f"),
        uuid.MustParse("00000000-0000-0000-1000-5c5b350ea3b3"),
    },
}

resp, err := sitesDevices.RestartSiteMultipleDevices(ctx, siteId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Device Config History

Search for entries in device config history

```go
SearchSiteDeviceConfigHistory(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeEnum,
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
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | - |
| `mac` | `*string` | Query, Optional | Device MAC Address |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseConfigHistorySearch`](../../doc/models/response-config-history-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum("ap")



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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
| `mac` | `*string` | Query, Optional | device mac |
| `model` | `*string` | Query, Optional | device model |
| `text` | `*string` | Query, Optional | event message |
| `timestamp` | `*string` | Query, Optional | event time |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `lastBy` | `*string` | Query, Optional | Return last/recent event for passed in field |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseEventsDevices`](../../doc/models/response-events-devices.md)

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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Device Last Configs

Search Device Last Configs

```go
SearchSiteDeviceLastConfigs(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeEnum,
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
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `version` | `*string` | Query, Optional | - |
| `name` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseConfigHistorySearch`](../../doc/models/response-config-history-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum("ap")







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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Devices

Search Device

```go
SearchSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    hostname *string,
    mType *models.DeviceTypeEnum,
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
    band24Bandwith *int,
    band5Bandwith *int,
    band6Bandwith *int,
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
| `hostname` | `*string` | Query, Optional | partial / full hostname |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | - |
| `model` | `*string` | Query, Optional | device model |
| `mac` | `*string` | Query, Optional | device MAC |
| `version` | `*string` | Query, Optional | version |
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
| `band24Bandwith` | `*int` | Query, Optional | Bandwidth of band_24 |
| `band5Bandwith` | `*int` | Query, Optional | Bandwidth of band_5 |
| `band6Bandwith` | `*int` | Query, Optional | Bandwidth of band_6 |
| `eth0PortSpeed` | `*int` | Query, Optional | Port speed of eth0 |
| `sort` | [`*models.SearchSiteDevicesSortEnum`](../../doc/models/search-site-devices-sort-enum.md) | Query, Optional | sort options |
| `descSort` | [`*models.SearchSiteDevicesDescSortEnum`](../../doc/models/search-site-devices-desc-sort-enum.md) | Query, Optional | sort options in reverse order |
| `stats` | `*bool` | Query, Optional | whether to return device stats |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseDeviceSearch`](../../doc/models/response-device-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



mType := models.DeviceTypeEnum("ap")









ipAddress := "192.168.1.1"



























sort := models.SearchSiteDevicesSortEnum("timestamp")



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
      "lldp_system_desc": "Cisco IOS Software, C2960S Software (C2960S-UNIVERSALK9-M), Version 15.2(1)E1, RELEASE SOFTWARE (fc2)\nTechnical Support: http://www.cisco.com/techsupport\nCopyright (c) 1986-2013 by Cisco Systems, Inc.\nCompiled Fri 22-Nov-13 07:10 by prod_rel_team",
      "lldp_system_name": "ME-DC-1-ACC-SW",
      "mac": "5c5b353e5299",
      "model": "AP41",
      "mxedge_id": "00000000-0000-0000-1000-43a81f238391",
      "mxtunnel_status": "down",
      "org_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
      "site_id": "a8178443-ecb5-461c-b854-f16627619ab3",
      "sku": "AP41-US",
      "timestamp": 1596588619.007,
      "uptime": 85280,
      "version": "0.7.20216"
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


# Update Site Device

Update Device Configuration

```go
UpdateSiteDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.MistDevice) (
    models.ApiResponse[models.UpdateSiteDeviceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MistDevice`](../../doc/models/containers/mist-device.md) | Body, Optional | Request Body |

## Response Type

[`models.UpdateSiteDeviceResponse`](../../doc/models/containers/update-site-device-response.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MistDeviceContainer.FromDeviceAp(models.DeviceAp{
    DeviceprofileId:  models.NewOptional(models.ToPointer(uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"))),
    DisableEth1:      models.ToPointer(false),
    DisableEth2:      models.ToPointer(false),
    DisableEth3:      models.ToPointer(false),
    DisableModule:    models.ToPointer(false),
    Height:           models.ToPointer(float64(2.75)),
    MapId:            models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
    Name:             models.ToPointer("conference room"),
    Notes:            models.ToPointer("slightly off center"),
    OrgId:            models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    Orientation:      models.ToPointer(45),
    PoePassthrough:   models.ToPointer(false),
    SiteId:           models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Vars:             map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    X:                models.ToPointer(float64(53.5)),
    Y:                models.ToPointer(float64(173.1)),
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
    "power": 10,
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
    "mtu": 0,
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
  "port_config": {
    "property1": {
      "additional_vlan_ids": [
        55,
        66
      ],
      "authentication_protocol": "pap",
      "disabled": true,
      "dynamic_vlan": {
        "default_vlan_id": 999,
        "enabled": true,
        "type": "string",
        "vlans": {
          "1-10": null,
          "user": null
        }
      },
      "enable_mac_auth": false,
      "forwarding": "all",
      "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
      "mxtunnel_name": "string",
      "port_auth": "none",
      "port_vlan_id": 1,
      "radius_config": {
        "acct_interim_interval": 0,
        "acct_servers": [
          {
            "host": "1.2.3.4",
            "keywrap_enabled": true,
            "keywrap_format": "hex",
            "keywrap_kek": "1122334455",
            "keywrap_mack": "1122334455",
            "port": 1813,
            "secret": "testing123"
          }
        ],
        "auth_servers": [
          {
            "host": "1.2.3.4",
            "keywrap_enabled": true,
            "keywrap_format": "hex",
            "keywrap_kek": "1122334455",
            "keywrap_mack": "1122334455",
            "port": 1812,
            "secret": "testing123"
          }
        ],
        "auth_servers_retries": 3,
        "auth_servers_timeout": 5,
        "coa_enabled": false,
        "coa_port": 3799,
        "network": "string",
        "source_ip": "string"
      },
      "radsec": {
        "enabled": true,
        "idle_timeout": 60,
        "mxcluster_ids": [
          "572586b7-f97b-a22b-526c-8b97a3f609c4"
        ],
        "proxy_hosts": [
          "mxedge1.local"
        ],
        "server_name": "radsec.abc.com",
        "servers": [
          {
            "host": "1.1.1.1",
            "port": 1812
          }
        ],
        "use_mxedge": true,
        "use_site_mxedge": false
      },
      "vlan_id": 9,
      "vland_ids": [
        1,
        10,
        50
      ],
      "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
      "wxtunnel_remote_id": "wifiguest"
    },
    "property2": {
      "additional_vlan_ids": [
        55,
        66
      ],
      "authentication_protocol": "pap",
      "disabled": true,
      "dynamic_vlan": {
        "default_vlan_id": 999,
        "enabled": true,
        "type": "string",
        "vlans": {
          "1-10": null,
          "user": null
        }
      },
      "enable_mac_auth": false,
      "forwarding": "all",
      "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
      "mxtunnel_name": "string",
      "port_auth": "none",
      "port_vlan_id": 1,
      "radius_config": {
        "acct_interim_interval": 0,
        "acct_servers": [
          {
            "host": "1.2.3.4",
            "keywrap_enabled": true,
            "keywrap_format": "hex",
            "keywrap_kek": "1122334455",
            "keywrap_mack": "1122334455",
            "port": 1813,
            "secret": "testing123"
          }
        ],
        "auth_servers": [
          {
            "host": "1.2.3.4",
            "keywrap_enabled": true,
            "keywrap_format": "hex",
            "keywrap_kek": "1122334455",
            "keywrap_mack": "1122334455",
            "port": 1812,
            "secret": "testing123"
          }
        ],
        "auth_servers_retries": 3,
        "auth_servers_timeout": 5,
        "coa_enabled": false,
        "coa_port": 3799,
        "network": "string",
        "source_ip": "string"
      },
      "radsec": {
        "enabled": true,
        "idle_timeout": 60,
        "mxcluster_ids": [
          "572586b7-f97b-a22b-526c-8b97a3f609c4"
        ],
        "proxy_hosts": [
          "mxedge1.local"
        ],
        "server_name": "radsec.abc.com",
        "servers": [
          {
            "host": "1.1.1.1",
            "port": 1812
          }
        ],
        "use_mxedge": true,
        "use_site_mxedge": false
      },
      "vlan_id": 9,
      "vland_ids": [
        1,
        10,
        50
      ],
      "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
      "wxtunnel_remote_id": "wifiguest"
    }
  },
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
