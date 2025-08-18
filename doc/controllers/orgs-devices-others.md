# Orgs Devices-Others

```go
orgsDevicesOthers := client.OrgsDevicesOthers()
```

## Class Name

`OrgsDevicesOthers`

## Methods

* [Count Org Other Device Events](../../doc/controllers/orgs-devices-others.md#count-org-other-device-events)
* [Delete Org Other Device](../../doc/controllers/orgs-devices-others.md#delete-org-other-device)
* [Get Org Other Device](../../doc/controllers/orgs-devices-others.md#get-org-other-device)
* [List Org Other Devices](../../doc/controllers/orgs-devices-others.md#list-org-other-devices)
* [Reboot Org Other Device](../../doc/controllers/orgs-devices-others.md#reboot-org-other-device)
* [Search Org Other Device Events](../../doc/controllers/orgs-devices-others.md#search-org-other-device-events)
* [Update Org Other Device](../../doc/controllers/orgs-devices-others.md#update-org-other-device)
* [Update Org Other Devices](../../doc/controllers/orgs-devices-others.md#update-org-other-devices)


# Count Org Other Device Events

Count by Distinct Attributes of Org OtherDevices Events

```go
CountOrgOtherDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgOtherdevicesEventsCountDistinctEnum,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgOtherdevicesEventsCountDistinctEnum`](../../doc/models/org-otherdevices-events-count-distinct-enum.md) | Query, Optional | **Default**: `"mac"` |
| `mType` | `*string` | Query, Optional | See  [List Device Events Definitions](../../doc/controllers/constants-events.md#list-other-device-events-definitions) |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgOtherdevicesEventsCountDistinctEnum_MAC

duration := "10m"

limit := 100

apiResponse, err := orgsDevicesOthers.CountOrgOtherDeviceEvents(ctx, orgId, &distinct, nil, nil, nil, &duration, &limit)
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


# Delete Org Other Device

Delete OtherDevice

```go
DeleteOrgOtherDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

resp, err := orgsDevicesOthers.DeleteOrgOtherDevice(ctx, orgId, deviceMac)
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


# Get Org Other Device

Get Org other device (3rd party device)

```go
GetOrgOtherDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.DeviceOther],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.DeviceOther](../../doc/models/device-other.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

apiResponse, err := orgsDevicesOthers.GetOrgOtherDevice(ctx, orgId, deviceMac)
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
  "device_mac": "00112233abcd",
  "id": "ae9dee49-69e7-4710-a114-5b827a777738",
  "mac": "5c5b35000018",
  "model": "W2005",
  "name": "W2005-268",
  "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
  "serial": "WB23015E025468",
  "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
  "state": "online",
  "vendor": "cradlepoint",
  "vendor_api_id": "4658714"
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


# List Org Other Devices

Get List of Org other devices (3rd party devices)

```go
ListOrgOtherDevices(
    ctx context.Context,
    orgId uuid.UUID,
    vendor *string,
    mac *string,
    serial *string,
    model *string,
    name *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.DeviceOther],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `vendor` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `serial` | `*string` | Query, Optional | - |
| `model` | `*string` | Query, Optional | - |
| `name` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.DeviceOther](../../doc/models/device-other.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsDevicesOthers.ListOrgOtherDevices(ctx, orgId, nil, nil, nil, nil, nil, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Reboot Org Other Device

Reboot OtherDevice

```go
RebootOrgOtherDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

resp, err := orgsDevicesOthers.RebootOrgOtherDevice(ctx, orgId, deviceMac)
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


# Search Org Other Device Events

Search Org OtherDevices Events

```go
SearchOrgOtherDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    mac *string,
    deviceMac *string,
    model *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | Site id |
| `mac` | `*string` | Query, Optional | MAC |
| `deviceMac` | `*string` | Query, Optional | MAC of attached device |
| `model` | `*string` | Query, Optional | Device model |
| `vendor` | `*string` | Query, Optional | Vendor name |
| `mType` | `*string` | Query, Optional | See  [List Device Events Definitions](../../doc/controllers/constants-events.md#list-other-device-events-definitions) |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsOtherDevicesSearch](../../doc/models/response-events-other-devices-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

apiResponse, err := orgsDevicesOthers.SearchOrgOtherDeviceEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Other Device

If the Site / Device cannot be identified, a manual association can be made

```go
UpdateOrgOtherDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string,
    body *models.OtherDeviceUpdate) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `body` | [`*models.OtherDeviceUpdate`](../../doc/models/other-device-update.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

body := models.OtherDeviceUpdate{
    DeviceMac:            models.ToPointer("0adfea67e65b"),
    SiteId:               models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
}

resp, err := orgsDevicesOthers.UpdateOrgOtherDevice(ctx, orgId, deviceMac, &body)
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


# Update Org Other Devices

Assign or unassign OtherDevices to and from a site.

```go
UpdateOrgOtherDevices(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OtherDeviceUpdateMulti) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OtherDeviceUpdateMulti`](../../doc/models/other-device-update-multi.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OtherDeviceUpdateMulti{
    Macs:                 []string{
        "5c5b350e0001",
        "5c5b350e0003",
    },
    Op:                   models.OtherDeviceUpdateOperationEnum_ASSIGN,
    SiteId:               models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
    AdditionalProperties: map[string]interface{}{
        "device_mac": interface{}("0adfea67e65b"),
    },
}

resp, err := orgsDevicesOthers.UpdateOrgOtherDevices(ctx, orgId, &body)
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

