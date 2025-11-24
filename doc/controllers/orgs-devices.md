# Orgs Devices

```go
orgsDevices := client.OrgsDevices()
```

## Class Name

`OrgsDevices`

## Methods

* [Count Org Device Events](../../doc/controllers/orgs-devices.md#count-org-device-events)
* [Count Org Device Last Configs](../../doc/controllers/orgs-devices.md#count-org-device-last-configs)
* [Count Org Devices](../../doc/controllers/orgs-devices.md#count-org-devices)
* [Get Org Juniper Devices Command](../../doc/controllers/orgs-devices.md#get-org-juniper-devices-command)
* [List Org Aps Macs](../../doc/controllers/orgs-devices.md#list-org-aps-macs)
* [List Org Devices](../../doc/controllers/orgs-devices.md#list-org-devices)
* [List Org Devices Summary](../../doc/controllers/orgs-devices.md#list-org-devices-summary)
* [Search Org Device Events](../../doc/controllers/orgs-devices.md#search-org-device-events)
* [Search Org Device Last Configs](../../doc/controllers/orgs-devices.md#search-org-device-last-configs)
* [Search Org Devices](../../doc/controllers/orgs-devices.md#search-org-devices)


# Count Org Device Events

Count by Distinct Attributes of Org Devices Events

```go
CountOrgDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgDevicesEventsCountDistinctEnum,
    siteId *uuid.UUID,
    ap *string,
    apfw *string,
    model *string,
    text *string,
    timestamp *string,
    mType *string,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgDevicesEventsCountDistinctEnum`](../../doc/models/org-devices-events-count-distinct-enum.md) | Query, Optional | **Default**: `"model"` |
| `siteId` | `*uuid.UUID` | Query, Optional | Site id |
| `ap` | `*string` | Query, Optional | AP mac |
| `apfw` | `*string` | Query, Optional | AP Firmware |
| `model` | `*string` | Query, Optional | Device model |
| `text` | `*string` | Query, Optional | Event message |
| `timestamp` | `*string` | Query, Optional | Event time |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgDevicesEventsCountDistinctEnum_MODEL

siteId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

ap := "5c5b53010101"

apfw := "10.0.0"

model := "AP43"

text := "Device connected"

timestamp := "1703003296"

duration := "10m"

limit := 100

apiResponse, err := orgsDevices.CountOrgDeviceEvents(ctx, orgId, &distinct, &siteId, &ap, &apfw, &model, &text, &timestamp, nil, nil, nil, &duration, &limit)
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


# Count Org Device Last Configs

Counts the number of entries in device config history for distinct field with given filters

```go
CountOrgDeviceLastConfigs(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeDefaultApEnum,
    distinct *models.OrgDevicesLastConfigsCountDistinctEnum,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | **Default**: `"ap"` |
| `distinct` | [`*models.OrgDevicesLastConfigsCountDistinctEnum`](../../doc/models/org-devices-last-configs-count-distinct-enum.md) | Query, Optional | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeDefaultApEnum_AP

duration := "10m"

limit := 100

apiResponse, err := orgsDevices.CountOrgDeviceLastConfigs(ctx, orgId, &mType, nil, nil, nil, &duration, &limit)
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


# Count Org Devices

Count by Distinct Attributes of Org Devices

```go
CountOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgDevicesCountDistinctEnum,
    hostname *string,
    siteId *uuid.UUID,
    model *string,
    managed *string,
    mac *string,
    version *string,
    ip *string,
    mxtunnelStatus *models.CountOrgDevicesMxtunnelStatusEnum,
    mxedgeId *uuid.UUID,
    lldpSystemName *string,
    lldpSystemDesc *string,
    lldpPortId *string,
    lldpMgmtAddr *string,
    mType *models.DeviceTypeDefaultApEnum,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgDevicesCountDistinctEnum`](../../doc/models/org-devices-count-distinct-enum.md) | Query, Optional | **Default**: `"model"` |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `siteId` | `*uuid.UUID` | Query, Optional | Site id |
| `model` | `*string` | Query, Optional | Device model |
| `managed` | `*string` | Query, Optional | for switches and gateways, to filter on managed/unmanaged devices. enum: `true`, `false` |
| `mac` | `*string` | Query, Optional | AP mac |
| `version` | `*string` | Query, Optional | Version |
| `ip` | `*string` | Query, Optional | - |
| `mxtunnelStatus` | [`*models.CountOrgDevicesMxtunnelStatusEnum`](../../doc/models/count-org-devices-mxtunnel-status-enum.md) | Query, Optional | MxTunnel status, enum: `up`, `down` |
| `mxedgeId` | `*uuid.UUID` | Query, Optional | Mist Edge id, if AP is connecting to a Mist Edge |
| `lldpSystemName` | `*string` | Query, Optional | LLDP system name |
| `lldpSystemDesc` | `*string` | Query, Optional | LLDP system description |
| `lldpPortId` | `*string` | Query, Optional | LLDP port id |
| `lldpMgmtAddr` | `*string` | Query, Optional | LLDP management ip address |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | **Default**: `"ap"` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgDevicesCountDistinctEnum_MODEL

hostname := "my-hostname"

siteId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

model := "MR84"

managed := "true"

mac := "5c5b53010101"

version := "10.0.0"

ip := "192.168.1.1"

mxedgeId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

lldpSystemName := "my-lldp-system"

lldpSystemDesc := "my-lldp-system-description"

lldpPortId := "ge-0/0/1"

lldpMgmtAddr := "10.4.2.3"

mType := models.DeviceTypeDefaultApEnum_AP

duration := "10m"

limit := 100

apiResponse, err := orgsDevices.CountOrgDevices(ctx, orgId, &distinct, &hostname, &siteId, &model, &managed, &mac, &version, &ip, nil, &mxedgeId, &lldpSystemName, &lldpSystemDesc, &lldpPortId, &lldpMgmtAddr, &mType, nil, nil, &duration, &limit)
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


# Get Org Juniper Devices Command

Get Org Juniper Devices command

Juniper devices can be managed/adopted by Mist. Currently outbound-ssh + netconf is used.
A few lines of CLI commands are generated per-Org, allowing the Juniper devices to phone home to Mist.

```go
GetOrgJuniperDevicesCommand(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string) (
    models.ApiResponse[models.ResponseDeviceConfigCmd],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | Site_id would be used for proxy config check of the site and automatic site assignment |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceConfigCmd](../../doc/models/response-device-config-cmd.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevices.GetOrgJuniperDevicesCommand(ctx, orgId, nil)
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
  "cmd": "set system services ssh...\n...\nset system services outbound-ssh client mist ..."
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


# List Org Aps Macs

For some scenarios like E911 or security systems, the BSSIDs are required to identify which AP the client is connecting to. Then the location of the AP can be used as the approximate location of the client.

Each radio MAC can have 16 BSSIDs (enumerate the last octet from 0-F)

```go
ListOrgApsMacs(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.ApRadioMac],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ApRadioMac](../../doc/models/ap-radio-mac.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsDevices.ListOrgApsMacs(ctx, orgId, &limit, &page)
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
    "mac": "5c5b35000001",
    "radio_macs": [
      "5c5b35000040",
      "5c5b35000050",
      "5c5b35000060"
    ]
  },
  {
    "mac": "5c5b45000001",
    "radio_macs": [
      "5c5b45000040",
      "5c5b45000050",
      "5c5b45000060"
    ]
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


# List Org Devices

Get List of Org Devices

```go
ListOrgDevices(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseOrgDevices],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseOrgDevices](../../doc/models/response-org-devices.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevices.ListOrgDevices(ctx, orgId)
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
      "mac": "string",
      "name": "string"
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


# List Org Devices Summary

Get Org Devices Summary

```go
ListOrgDevicesSummary(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseOrgDevicesSummary],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseOrgDevicesSummary](../../doc/models/response-org-devices-summary.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevices.ListOrgDevicesSummary(ctx, orgId)
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
  "num_aps": 630,
  "num_gateways": 6,
  "num_mxedges": 1,
  "num_switches": 30,
  "num_unassigned_aps": 5,
  "num_unassigned_gateways": 0,
  "num_unassigned_switches": 0
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


# Search Org Device Events

Search Org Devices Events

```go
SearchOrgDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    model *string,
    deviceType *models.DeviceTypeWithAllEnum,
    text *string,
    timestamp *string,
    mType *string,
    lastBy *string,
    includes *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseDeviceEventsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Device mac |
| `model` | `*string` | Query, Optional | Device model |
| `deviceType` | [`*models.DeviceTypeWithAllEnum`](../../doc/models/device-type-with-all-enum.md) | Query, Optional | **Default**: `"ap"` |
| `text` | `*string` | Query, Optional | Event message |
| `timestamp` | `*string` | Query, Optional | Event time |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `lastBy` | `*string` | Query, Optional | Return last/recent event for passed in field |
| `includes` | `*string` | Query, Optional | Keyword to include events from additional indices (e.g. ext_tunnel for prisma events) |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceEventsSearch](../../doc/models/response-device-events-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "5c5b53010101"

model := "AP43"

deviceType := models.DeviceTypeWithAllEnum_AP

text := "Device connected"

timestamp := "1703003296"

lastBy := "port_id"

includes := "ext_tunnel"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsDevices.SearchOrgDeviceEvents(ctx, orgId, &mac, &model, &deviceType, &text, &timestamp, nil, &lastBy, &includes, &limit, nil, nil, &duration, &sort, nil)
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
      "ap": "5c5b351e13b5",
      "apfw": "5c5b351e13b5",
      "model": "BT11-WW",
      "org_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862a",
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "text": "Succeeding DNS query from 172.29.101.134 to 172.29.101.7 for \"portal.mistsys.com\" on vlan 1, id 60224",
      "timestamp": 1547235620.89,
      "type": "CLIENT_DNS_OK"
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


# Search Org Device Last Configs

Search Device Last Configs

```go
SearchOrgDeviceLastConfigs(
    ctx context.Context,
    orgId uuid.UUID,
    deviceType *models.LastConfigDeviceTypeEnum,
    mac *string,
    name *string,
    version *string,
    certExpiryDuration *string,
    start *string,
    end *string,
    limit *int,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseConfigHistorySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceType` | [`*models.LastConfigDeviceTypeEnum`](../../doc/models/last-config-device-type-enum.md) | Query, Optional | **Default**: `"ap"` |
| `mac` | `*string` | Query, Optional | Device MAC address |
| `name` | `*string` | Query, Optional | Devices Name |
| `version` | `*string` | Query, Optional | Device Version |
| `certExpiryDuration` | `*string` | Query, Optional | Duration for expiring cert queries (format: 2d/3h/172800 seconds) |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseConfigHistorySearch](../../doc/models/response-config-history-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceType := models.LastConfigDeviceTypeEnum_AP

mac := "5c5b53010101"

name := "My AP"

version := "10.0.0"

certExpiryDuration := "2d"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsDevices.SearchOrgDeviceLastConfigs(ctx, orgId, &deviceType, &mac, &name, &version, &certExpiryDuration, nil, nil, &limit, &duration, &sort, nil)
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


# Search Org Devices

Search Org Devices

```go
SearchOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    band24Bandwidth *int,
    band24Channel *int,
    band24Power *int,
    band5Bandwidth *int,
    band5Channel *int,
    band5Power *int,
    band6Bandwidth *int,
    band6Channel *int,
    band6Power *int,
    cpu *string,
    clustered *string,
    eth0PortSpeed *int,
    evpntopoId *string,
    extIp *string,
    hostname *string,
    ip *string,
    lastConfigStatus *string,
    lastHostname *string,
    lldpMgmtAddr *string,
    lldpPortId *string,
    lldpPowerAllocated *int,
    lldpPowerDraw *int,
    lldpSystemDesc *string,
    lldpSystemName *string,
    mac *string,
    model *string,
    mxedgeId *string,
    mxedgeIds *string,
    mxtunnelStatus *models.SearchOrgDevicesMxtunnelStatusEnum,
    node *string,
    node0Mac *string,
    node1Mac *string,
    powerConstrained *bool,
    siteId *string,
    t128agentVersion *string,
    version *string,
    mType *models.DeviceTypeDefaultApEnum,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseDeviceSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `band24Bandwidth` | `*int` | Query, Optional | If `type`==`ap`, Bandwidth of band_24 |
| `band24Channel` | `*int` | Query, Optional | If `type`==`ap`, Channel of band_24 |
| `band24Power` | `*int` | Query, Optional | If `type`==`ap`, Power of band_24 |
| `band5Bandwidth` | `*int` | Query, Optional | If `type`==`ap`, Bandwidth of band_5 |
| `band5Channel` | `*int` | Query, Optional | If `type`==`ap`, Channel of band_5 |
| `band5Power` | `*int` | Query, Optional | If `type`==`ap`, Power of band_5 |
| `band6Bandwidth` | `*int` | Query, Optional | If `type`==`ap`, Bandwidth of band_6 |
| `band6Channel` | `*int` | Query, Optional | If `type`==`ap`, Channel of band_6 |
| `band6Power` | `*int` | Query, Optional | If `type`==`ap`, Power of band_6 |
| `cpu` | `*string` | Query, Optional | If `type`==`switch` or `type`==`gateway`, max cpu usage |
| `clustered` | `*string` | Query, Optional | If `type`==`gateway`, true / false |
| `eth0PortSpeed` | `*int` | Query, Optional | If `type`==`ap`, Port speed of eth0 |
| `evpntopoId` | `*string` | Query, Optional | If `type`==`switch`, EVPN topology id |
| `extIp` | `*string` | Query, Optional | External IP Address |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `ip` | `*string` | Query, Optional | - |
| `lastConfigStatus` | `*string` | Query, Optional | If `type`==`switch` or `type`==`gateway`, last configuration status |
| `lastHostname` | `*string` | Query, Optional | If `type`==`switch` or `type`==`gateway`, last hostname |
| `lldpMgmtAddr` | `*string` | Query, Optional | If `type`==`ap`, LLDP management ip address |
| `lldpPortId` | `*string` | Query, Optional | If `type`==`ap`, LLDP port id |
| `lldpPowerAllocated` | `*int` | Query, Optional | If `type`==`ap`, LLDP Allocated Power |
| `lldpPowerDraw` | `*int` | Query, Optional | If `type`==`ap`, LLDP Negotiated Power |
| `lldpSystemDesc` | `*string` | Query, Optional | If `type`==`ap`, LLDP system description |
| `lldpSystemName` | `*string` | Query, Optional | If `type`==`ap`, LLDP system name |
| `mac` | `*string` | Query, Optional | Device mac |
| `model` | `*string` | Query, Optional | Device model |
| `mxedgeId` | `*string` | Query, Optional | If `type`==`ap`, Mist Edge id, if AP is connecting to a Mist Edge |
| `mxedgeIds` | `*string` | Query, Optional | If `type`==`ap`, Comma separated list of Mist Edge ids, if AP is connecting to a Mist Edge |
| `mxtunnelStatus` | [`*models.SearchOrgDevicesMxtunnelStatusEnum`](../../doc/models/search-org-devices-mxtunnel-status-enum.md) | Query, Optional | If `type`==`ap`, MxTunnel status, up / down |
| `node` | `*string` | Query, Optional | If `type`==`gateway`, `node0` / `node1` |
| `node0Mac` | `*string` | Query, Optional | If `type`==`gateway`, mac for node0 |
| `node1Mac` | `*string` | Query, Optional | If `type`==`gateway`, mac for node1 |
| `powerConstrained` | `*bool` | Query, Optional | If `type`==`ap`, Power_constrained |
| `siteId` | `*string` | Query, Optional | Site id |
| `t128agentVersion` | `*string` | Query, Optional | If `type`==`gateway`,version of 128T agent |
| `version` | `*string` | Query, Optional | Version |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | Type of device. enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceSearch](../../doc/models/response-device-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

band24Bandwidth := 20

band24Channel := 6

band24Power := 8

band5Bandwidth := 20

band5Channel := 50

band5Power := 8

band6Bandwidth := 20

band6Channel := 100

band6Power := 8

cpu := "50"

clustered := "true"

eth0PortSpeed := 1000

evpntopoId := "7dae216d-7c98-a51b-e068-dd7d477b7216"

extIp := "83.42.53.1"

hostname := "my-hostname"

ip := "192.168.1.1"

lastConfigStatus := "success"

lastHostname := "my-last-hostname"

lldpMgmtAddr := "10.4.2.3"

lldpPortId := "ge-0/0/1"

lldpPowerAllocated := 15

lldpPowerDraw := 12

lldpSystemDesc := "my-lldp-system-description"

lldpSystemName := "my-lldp-system"

mac := "5c5b53010101"

model := "AP43"

mxedgeId := "7dae216d-7c98-a51b-e068-dd7d477b7216"

mxedgeIds := "7dae216d-7c98-a51b-e068-dd7d477b7216,7dae216d-7c98-a51b-e068-dd7d477b7217"

node := "node0"

node0Mac := "5c5b350e0001"

node1Mac := "5c5b350e0002"

powerConstrained := true

siteId := "7dae216d-7c98-a51b-e068-dd7d477b7216"

t128agentVersion := "1.2.3"

version := "10.0.0"

mType := models.DeviceTypeDefaultApEnum_AP

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsDevices.SearchOrgDevices(ctx, orgId, &band24Bandwidth, &band24Channel, &band24Power, &band5Bandwidth, &band5Channel, &band5Power, &band6Bandwidth, &band6Channel, &band6Power, &cpu, &clustered, &eth0PortSpeed, &evpntopoId, &extIp, &hostname, &ip, &lastConfigStatus, &lastHostname, &lldpMgmtAddr, &lldpPortId, &lldpPowerAllocated, &lldpPowerDraw, &lldpSystemDesc, &lldpSystemName, &mac, &model, &mxedgeId, &mxedgeIds, nil, &node, &node0Mac, &node1Mac, &powerConstrained, &siteId, &t128agentVersion, &version, &mType, &limit, nil, nil, &duration, &sort, nil)
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

