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
* [Search Org Device Events](../../doc/controllers/orgs-devices.md#search-org-device-events)
* [Search Org Device Last Configs](../../doc/controllers/orgs-devices.md#search-org-device-last-configs)
* [Search Org Devices](../../doc/controllers/orgs-devices.md#search-org-devices)


# Count Org Device Events

Count Org Devices Events

```go
CountOrgDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgDevicesEventsCountDistinctEnum,
    siteId *string,
    ap *string,
    apfw *string,
    model *string,
    text *string,
    timestamp *string,
    mType *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgDevicesEventsCountDistinctEnum`](../../doc/models/org-devices-events-count-distinct-enum.md) | Query, Optional | - |
| `siteId` | `*string` | Query, Optional | site id |
| `ap` | `*string` | Query, Optional | AP mac |
| `apfw` | `*string` | Query, Optional | AP Firmware |
| `model` | `*string` | Query, Optional | device model |
| `text` | `*string` | Query, Optional | event message |
| `timestamp` | `*string` | Query, Optional | event time |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgDevicesEventsCountDistinctEnum("model")















limit := 100





duration := "10m"

apiResponse, err := orgsDevices.CountOrgDeviceEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Count Org Device Last Configs

Counts the number of entries in device config history for distinct field with given filters

```go
CountOrgDeviceLastConfigs(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    distinct *models.OrgDevicesLastConfigsCountDistinctEnum,
    start *int,
    end *int,
    limit *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | - |
| `distinct` | [`*models.OrgDevicesLastConfigsCountDistinctEnum`](../../doc/models/org-devices-last-configs-count-distinct-enum.md) | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum("ap")







limit := 100

apiResponse, err := orgsDevices.CountOrgDeviceLastConfigs(ctx, orgId, &mType, nil, nil, nil, &limit)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Count Org Devices

Count Org Devices

```go
CountOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgDevicesCountDistinctEnum,
    hostname *string,
    siteId *string,
    model *string,
    mac *string,
    version *string,
    ipAddress *string,
    mxtunnelStatus *models.CountOrgDevicesMxtunnelStatusEnum,
    mxedgeId *string,
    lldpSystemName *string,
    lldpSystemDesc *string,
    lldpPortId *string,
    lldpMgmtAddr *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgDevicesCountDistinctEnum`](../../doc/models/org-devices-count-distinct-enum.md) | Query, Optional | - |
| `hostname` | `*string` | Query, Optional | partial / full hostname |
| `siteId` | `*string` | Query, Optional | site id |
| `model` | `*string` | Query, Optional | device model |
| `mac` | `*string` | Query, Optional | AP mac |
| `version` | `*string` | Query, Optional | version |
| `ipAddress` | `*string` | Query, Optional | - |
| `mxtunnelStatus` | [`*models.CountOrgDevicesMxtunnelStatusEnum`](../../doc/models/count-org-devices-mxtunnel-status-enum.md) | Query, Optional | MxTunnel status, up / down |
| `mxedgeId` | `*string` | Query, Optional | Mist Edge id, if AP is connecting to a Mist Edge |
| `lldpSystemName` | `*string` | Query, Optional | LLDP system name |
| `lldpSystemDesc` | `*string` | Query, Optional | LLDP system description |
| `lldpPortId` | `*string` | Query, Optional | LLDP port id |
| `lldpMgmtAddr` | `*string` | Query, Optional | LLDP management ip address |
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

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgDevicesCountDistinctEnum("model")











ipAddress := "192.168.1.1"













page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsDevices.CountOrgDevices(ctx, orgId, &distinct, nil, nil, nil, nil, nil, &ipAddress, nil, nil, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
| `siteId` | `*string` | Query, Optional | site_id would be used for proxy config check of the site and automatic site assignment |

## Response Type

[`models.ResponseDeviceConfigCmd`](../../doc/models/response-device-config-cmd.md)

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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Aps Macs

For some scenarios like E911 or security systems, the BSSIDs are required to identify which AP the client is connecting to. Then the location of the AP can be used as the approximate location of the client.

Each radio MAC can have 16 BSSIDs (enumerate the last octet from 0-F)

```go
ListOrgApsMacs(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.ApRadioMac],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.ApRadioMac`](../../doc/models/ap-radio-mac.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsDevices.ListOrgApsMacs(ctx, orgId, &page, &limit)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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

[`models.ResponseOrgDevices`](../../doc/models/response-org-devices.md)

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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Org Device Events

Search Org Devices Events

```go
SearchOrgDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    model *string,
    deviceType *models.DeviceTypeEnum,
    text *string,
    timestamp *string,
    mType *string,
    lastBy *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseDeviceEventsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | device mac |
| `model` | `*string` | Query, Optional | device model |
| `deviceType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | - |
| `text` | `*string` | Query, Optional | event message |
| `timestamp` | `*string` | Query, Optional | event time |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `lastBy` | `*string` | Query, Optional | Return last/recent event for passed in field |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseDeviceEventsSearch`](../../doc/models/response-device-events-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





deviceType := models.DeviceTypeEnum("ap")







lastBy := "port_id"

limit := 100





duration := "10m"

apiResponse, err := orgsDevices.SearchOrgDeviceEvents(ctx, orgId, nil, nil, &deviceType, nil, nil, nil, &lastBy, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Org Device Last Configs

Search Device Last Configs

```go
SearchOrgDeviceLastConfigs(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    mac *string,
    name *string,
    version *string,
    start *int,
    end *int,
    limit *int,
    duration *string) (
    models.ApiResponse[models.ResponseConfigHistorySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | - |
| `mac` | `*string` | Query, Optional | Device MAC address |
| `name` | `*string` | Query, Optional | Devices Name |
| `version` | `*string` | Query, Optional | Device Version |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `limit` | `*int` | Query, Optional | - |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseConfigHistorySearch`](../../doc/models/response-config-history-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum("ap")











limit := 100

duration := "10m"

apiResponse, err := orgsDevices.SearchOrgDeviceLastConfigs(ctx, orgId, &mType, nil, nil, nil, nil, nil, &limit, &duration)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Org Devices

Search Org Devices

```go
SearchOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    hostname *string,
    siteId *string,
    model *string,
    mac *string,
    version *string,
    powerConstrained *bool,
    ipAddress *string,
    mxtunnelStatus *models.SearchOrgDevicesMxtunnelStatusEnum,
    mxedgeId *string,
    lldpSystemName *string,
    lldpSystemDesc *string,
    lldpPortId *string,
    lldpMgmtAddr *string,
    band24Bandwith *int,
    band5Bandwith *int,
    band6Bandwith *int,
    band24Channel *int,
    band5Channel *int,
    band6Channel *int,
    eth0PortSpeed *int,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `hostname` | `*string` | Query, Optional | partial / full hostname |
| `siteId` | `*string` | Query, Optional | site id |
| `model` | `*string` | Query, Optional | device model |
| `mac` | `*string` | Query, Optional | AP mac |
| `version` | `*string` | Query, Optional | version |
| `powerConstrained` | `*bool` | Query, Optional | power_constrained |
| `ipAddress` | `*string` | Query, Optional | - |
| `mxtunnelStatus` | [`*models.SearchOrgDevicesMxtunnelStatusEnum`](../../doc/models/search-org-devices-mxtunnel-status-enum.md) | Query, Optional | MxTunnel status, up / down |
| `mxedgeId` | `*string` | Query, Optional | Mist Edge id, if AP is connecting to a Mist Edge |
| `lldpSystemName` | `*string` | Query, Optional | LLDP system name |
| `lldpSystemDesc` | `*string` | Query, Optional | LLDP system description |
| `lldpPortId` | `*string` | Query, Optional | LLDP port id |
| `lldpMgmtAddr` | `*string` | Query, Optional | LLDP management ip address |
| `band24Bandwith` | `*int` | Query, Optional | Bandwith of band_24 |
| `band5Bandwith` | `*int` | Query, Optional | Bandwith of band_5 |
| `band6Bandwith` | `*int` | Query, Optional | Bandwith of band_6 |
| `band24Channel` | `*int` | Query, Optional | Channel of band_24 |
| `band5Channel` | `*int` | Query, Optional | Channel of band_5 |
| `band6Channel` | `*int` | Query, Optional | Channel of band_6 |
| `eth0PortSpeed` | `*int` | Query, Optional | Port speed of eth0 |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseDeviceSearch`](../../doc/models/response-device-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













ipAddress := "192.168.1.1"



























limit := 100





duration := "10m"

apiResponse, err := orgsDevices.SearchOrgDevices(ctx, orgId, nil, nil, nil, nil, nil, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

