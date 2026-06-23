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

Count device event records across the organization, optionally grouped by `distinct` and filtered by site, AP, firmware, model, event text, event type, and time range.

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
    mType *string,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgDevicesEventsCountDistinctEnum`](../../doc/models/org-devices-events-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `ap`, `apfw`, `model`, `org_id`, `site_id`, `text`, `timestamp`, `type`<br><br>**Default**: `"model"` |
| `siteId` | `*uuid.UUID` | Query, Optional | Filter results by site identifier |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address |
| `apfw` | `*string` | Query, Optional | Filter results by AP firmware version |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `text` | `*string` | Query, Optional | Filter results by event message text |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

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

duration := "10m"

limit := 100

apiResponse, err := orgsDevices.CountOrgDeviceEvents(ctx, orgId, &distinct, &siteId, &ap, &apfw, &model, &text, nil, nil, nil, &duration, &limit)
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


# Count Org Device Last Configs

Count device config history records across the organization, optionally grouped by `distinct` and filtered by device type and time range.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | Filter results by type. enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `distinct` | [`*models.OrgDevicesLastConfigsCountDistinctEnum`](../../doc/models/org-devices-last-configs-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `mac`, `name`, `site_id`, `version` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

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


# Count Org Devices

Count organization device records, optionally grouped by `distinct` and filtered by device identifiers, model, LLDP attributes, Mist Edge, tunnel status, device type, and time range.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgDevicesCountDistinctEnum`](../../doc/models/org-devices-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `hostname`, `ip`, `lldp_mgmt_addr`, `lldp_port_id`, `lldp_system_desc`, `lldp_system_name`, `mac`, `model`, `mxedge_id`, `mxtunnel_status`, `site_id`, `version`<br><br>**Default**: `"model"` |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `siteId` | `*uuid.UUID` | Query, Optional | Filter results by site identifier |
| `model` | `*string` | Query, Optional | Filter results by device model. Accepts multiple comma-separated values. |
| `managed` | `*string` | Query, Optional | for switches and gateways, to filter on managed/unmanaged devices. Deprecated in favour of mist_configured. enum: `true`, `false` |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `version` | `*string` | Query, Optional | Filter results by software version |
| `ip` | `*string` | Query, Optional | Filter results by IPv4 address |
| `mxtunnelStatus` | [`*models.CountOrgDevicesMxtunnelStatusEnum`](../../doc/models/count-org-devices-mxtunnel-status-enum.md) | Query, Optional | MxTunnel status, enum: `up`, `down` |
| `mxedgeId` | `*uuid.UUID` | Query, Optional | Mist Edge id, if AP is connecting to a Mist Edge |
| `lldpSystemName` | `*string` | Query, Optional | Filter results by LLDP system name |
| `lldpSystemDesc` | `*string` | Query, Optional | Filter results by LLDP system description |
| `lldpPortId` | `*string` | Query, Optional | Filter results by LLDP port identifier |
| `lldpMgmtAddr` | `*string` | Query, Optional | LLDP management IP address |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | Filter results by type. enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgDevicesCountDistinctEnum_MODEL

hostname := "my-hostname"

siteId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

model := "AP45,AP47D"

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | Filter results by site identifier |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceConfigCmd](../../doc/models/response-device-config-cmd.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevices.GetOrgJuniperDevicesCommand(ctx, orgId, nil)
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
  "cmd": "set system services ssh...\n...\nset system services outbound-ssh client mist ..."
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


# List Org Aps Macs

For some scenarios like E911 or security systems, the BSSIDs are required to identify which AP the client is connecting to. Then the location of the AP can be used as the approximate location of the client.

Each radio MAC can have up to 16 BSSIDs. These are derived by incrementing the least significant hexadecimal digit (last nibble) of the MAC address from 0 to F, while keeping the remaining bits unchanged.

```go
ListOrgApsMacs(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.ApRadioMac],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ApRadioMac](../../doc/models/ap-radio-mac.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsDevices.ListOrgApsMacs(ctx, orgId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Devices

List devices in the organization, including APs, switches and gateways managed or monitored by Mist.

```go
ListOrgDevices(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseOrgDevices],
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseOrgDevices](../../doc/models/response-org-devices.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevices.ListOrgDevices(ctx, orgId)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Devices Summary

Return aggregate organization device counts by device category and assignment state.

```go
ListOrgDevicesSummary(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseOrgDevicesSummary],
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseOrgDevicesSummary](../../doc/models/response-org-devices-summary.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevices.ListOrgDevicesSummary(ctx, orgId)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Device Events

Search device event records across the organization with filters for MAC address, model, device type, event text, event type, additional event indices, and time range.

```go
SearchOrgDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    model *string,
    deviceType *string,
    text *string,
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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Filter results by MAC address. Accepts multiple comma-separated values. |
| `model` | `*string` | Query, Optional | Filter results by device model. Accepts multiple comma-separated values. |
| `deviceType` | `*string` | Query, Optional | Filter results by device type. Accepts multiple comma-separated values.<br><br>**Default**: `"ap"` |
| `text` | `*string` | Query, Optional | Filter results by event message text |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions). Accepts multiple comma-separated values. |
| `lastBy` | `*string` | Query, Optional | Return last/recent event for passed in field |
| `includes` | `*string` | Query, Optional | Keyword to include events from additional indices (e.g. ext_tunnel for prisma events) |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceEventsSearch](../../doc/models/response-device-events-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "5c5b53010101,5c5b53020202"

model := "EX4100-48MP,AP17"

deviceType := "switch,ap"

text := "Device connected"

mType := "SW_PORT_DOWN,AP_CONFIGURED"

lastBy := "port_id"

includes := "ext_tunnel"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsDevices.SearchOrgDeviceEvents(ctx, orgId, &mac, &model, &deviceType, &text, &mType, &lastBy, &includes, &limit, nil, nil, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Device Last Configs

Search device config history records across the organization with filters for device type, MAC address, name, software version, certificate-expiry duration, and time range.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceType` | [`*models.LastConfigDeviceTypeEnum`](../../doc/models/last-config-device-type-enum.md) | Query, Optional | Filter results by device type. enum: `ap`, `gateway`, `switch`, `mxedge`<br><br>**Default**: `"ap"` |
| `mac` | `*string` | Query, Optional | Filter results by MAC address. Accepts multiple comma-separated values. |
| `name` | `*string` | Query, Optional | Filter results by name. Accepts multiple comma-separated values. |
| `version` | `*string` | Query, Optional | Filter results by software version. Accepts multiple comma-separated values. |
| `certExpiryDuration` | `*string` | Query, Optional | Duration for expiring cert queries (format: 2d/3h/172800 seconds) |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseConfigHistorySearch](../../doc/models/response-config-history-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceType := models.LastConfigDeviceTypeEnum_AP

mac := "5c5b53010101,5c5b53020202"

name := "name-a,name-b"

version := "apfw-0.15.34615-noro-0a3c,apfw-0.15.34436-noro-d810"

certExpiryDuration := "2d"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsDevices.SearchOrgDeviceLastConfigs(ctx, orgId, &deviceType, &mac, &name, &version, &certExpiryDuration, nil, nil, &limit, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Devices

Search organization devices with filters for AP radio attributes, gateway HA attributes, switch EVPN attributes, LLDP data, MAC address, IP address, model, software version, site, Mist Edge, and time range. Set `stats=true` to include device stats in the response.

```go
SearchOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    band24Channel *int,
    band5Channel *int,
    band6Channel *int,
    band24Bandwidth *int,
    band5Bandwidth *int,
    band6Bandwidth *int,
    band24Power *int,
    band5Power *int,
    band6Power *int,
    clustered *bool,
    eth0PortSpeed *int,
    evpntopoId *uuid.UUID,
    extIp *string,
    hostname *string,
    ip *string,
    lastConfigStatus *string,
    lastHostname *string,
    lldpMgmtAddr *string,
    lldpPortId *string,
    lldpSystemDesc *string,
    lldpSystemName *string,
    mac *string,
    model *string,
    mxedgeId *uuid.UUID,
    mxedgeIds *string,
    mxtunnelStatus *models.SearchOrgDevicesMxtunnelStatusEnum,
    node *models.HaClusterNodeEnum,
    node0Mac *string,
    node1Mac *string,
    powerConstrained *bool,
    radiusStats *string,
    siteId *string,
    stats *bool,
    t128agentVersion *string,
    mType *models.DeviceTypeDefaultApEnum,
    version *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseDeviceSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `band24Channel` | `*int` | Query, Optional | When `type`==`ap`, Channel of band_24. Accepts multiple comma-separated integer values. |
| `band5Channel` | `*int` | Query, Optional | When `type`==`ap`, Channel of band_5. Accepts multiple comma-separated integer values. |
| `band6Channel` | `*int` | Query, Optional | When `type`==`ap`, Channel of band_6. Accepts multiple comma-separated integer values. |
| `band24Bandwidth` | `*int` | Query, Optional | When `type`==`ap`, Bandwidth of band_24. Accepts multiple comma-separated integer values. |
| `band5Bandwidth` | `*int` | Query, Optional | When `type`==`ap`, Bandwidth of band_5. Accepts multiple comma-separated integer values. |
| `band6Bandwidth` | `*int` | Query, Optional | When `type`==`ap`, Bandwidth of band_6. Accepts multiple comma-separated integer values. |
| `band24Power` | `*int` | Query, Optional | When `type`==`ap`, Power of band_24. Accepts multiple comma-separated integer values. |
| `band5Power` | `*int` | Query, Optional | When `type`==`ap`, Power of band_5. Accepts multiple comma-separated integer values. |
| `band6Power` | `*int` | Query, Optional | When `type`==`ap`, Power of band_6. Accepts multiple comma-separated integer values. |
| `clustered` | `*bool` | Query, Optional | When `type`==`gateway`, true / false |
| `eth0PortSpeed` | `*int` | Query, Optional | When `type`==`ap`, Port speed of eth0. Accepts multiple comma-separated integer values. |
| `evpntopoId` | `*uuid.UUID` | Query, Optional | When `type`==`switch`, EVPN topology id |
| `extIp` | `*string` | Query, Optional | Partial / full Device external ip. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `1.2.3.*` and `*.2.3.*` match `1.2.3.4`). Suffix-only wildcards (e.g. `*.2.3.4`) are not supported. Accepts multiple comma-separated values. |
| `hostname` | `*string` | Query, Optional | Partial / full Device hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `my-london*` and `*london*` match `my-london-1`). Suffix-only wildcards (e.g. `*london-1`) are not supported. Accepts multiple comma-separated values. |
| `ip` | `*string` | Query, Optional | Partial / full Device IP address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `10.100.10.*` and `*100.10.*` match `10.100.10.54`). Suffix-only wildcards (e.g. `*.54`) are not supported. Accepts multiple comma-separated values. |
| `lastConfigStatus` | `*string` | Query, Optional | When `type`==`switch` or `type`==`gateway`, last configuration status |
| `lastHostname` | `*string` | Query, Optional | Last hostname of the device. Accepts multiple comma-separated values. |
| `lldpMgmtAddr` | `*string` | Query, Optional | When `type`==`ap`, LLDP management IP address. Accepts multiple comma-separated values. |
| `lldpPortId` | `*string` | Query, Optional | When `type`==`ap`, LLDP port id. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `ge-0/0/*` and `*-0/0/*` match `ge-0/0/30`). Suffix-only wildcards (e.g. `*switch-01`) are not supported. Accepts multiple comma-separated values. |
| `lldpSystemDesc` | `*string` | Query, Optional | When `type`==`ap`, LLDP system description. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `Juniper Networks*` and `*Networks*` match `Juniper Networks, Inc.`). Suffix-only wildcards (e.g. `*switch-01`) are not supported |
| `lldpSystemName` | `*string` | Query, Optional | When `type`==`ap`, LLDP system name. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `my-switch*` and `*switch*` match `my-switch-01`). Suffix-only wildcards (e.g. `*switch-01`) are not supported. Accepts multiple comma-separated values. |
| `mac` | `*string` | Query, Optional | Partial / full Device MAC address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `001122*` and `*1122*` match `001122334455`). Suffix-only wildcards (e.g. `*4455`) are not supported. Accepts multiple comma-separated values. |
| `model` | `*string` | Query, Optional | Partial / full Device model. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `AP4*` and `*P4*` match `AP43`). Suffix-only wildcards (e.g. `*43`) are not supported. Accepts multiple comma-separated values. |
| `mxedgeId` | `*uuid.UUID` | Query, Optional | When `type`==`ap`, Mist Edge id, if AP is connecting to a Mist Edge. Accepts multiple comma-separated values. |
| `mxedgeIds` | `*string` | Query, Optional | When `type`==`ap`, Comma separated list of Mist Edge id, if AP is connecting to a Mist Edge |
| `mxtunnelStatus` | [`*models.SearchOrgDevicesMxtunnelStatusEnum`](../../doc/models/search-org-devices-mxtunnel-status-enum.md) | Query, Optional | When `type`==`ap`, Mist Tunnel status used to filter results. enum: `down`, `up` |
| `node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Query, Optional | When `type`==`gateway`. enum: `node0`, `node1` |
| `node0Mac` | `*string` | Query, Optional | When `type`==`gateway`, node0 MAC address |
| `node1Mac` | `*string` | Query, Optional | When `type`==`gateway`, node1 MAC address |
| `powerConstrained` | `*bool` | Query, Optional | When `type`==`ap`, whether the AP is power constrained. |
| `radiusStats` | `*string` | Query, Optional | When `type`==`switch` or `type`==`gateway`, Key-value pairs where the key<br>is the RADIUS server address and the value contains authentication statistics:<br><br>* <server_address> (string): IP address of the RADIUS server as the key<br>* `auth_accepts` (long): Number of accepted authentication requests<br>* `auth_rejects` (long): Number of rejected authentication requests<br>* `auth_timeouts` (long): Number of authentication timeouts<br>* `auth_server_status` (string): Status of the server. Possible values: `up`, `down`, `unreachable` |
| `siteId` | `*string` | Query, Optional | Filter results by site identifier |
| `stats` | `*bool` | Query, Optional | Whether to return device stats<br><br>**Default**: `false` |
| `t128agentVersion` | `*string` | Query, Optional | When `type`==`gateway` (SSR only), version of 128T agent |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | Device type used to filter results. enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `version` | `*string` | Query, Optional | Filter results by software version. Accepts multiple comma-separated values. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceSearch](../../doc/models/response-device-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

eth0PortSpeed := 100

extIp := "1.2.3.4,1.2.3.*"

hostname := "my-london-1,my-london*"

ip := "10.100.10.54,10.100.10.*"

lastConfigStatus := "success"

lastHostname := "ap-01,ap-02"

lldpMgmtAddr := "192.0.2.10,192.0.2.11"

lldpPortId := "00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"

lldpSystemName := "richmond-hill-switch,Phoenix_Switch"

mac := "aabbccddeeff,aabbcc*"

model := "AP43,AP4*"

siteId := "7dae216d-7c98-a51b-e068-dd7d477b7216"

stats := false

mType := models.DeviceTypeDefaultApEnum_AP

version := "0.15.34615,0.14.29967"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsDevices.SearchOrgDevices(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &eth0PortSpeed, nil, &extIp, &hostname, &ip, &lastConfigStatus, &lastHostname, &lldpMgmtAddr, &lldpPortId, nil, &lldpSystemName, &mac, &model, nil, nil, nil, nil, nil, nil, nil, nil, &siteId, &stats, nil, &mType, &version, &limit, nil, nil, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

