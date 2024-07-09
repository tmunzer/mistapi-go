# Orgs Clients-Wireless

```go
orgsClientsWireless := client.OrgsClientsWireless()
```

## Class Name

`OrgsClientsWireless`

## Methods

* [Count Org Wireless Clients](../../doc/controllers/orgs-clients-wireless.md#count-org-wireless-clients)
* [Count Org Wireless Clients Sessions](../../doc/controllers/orgs-clients-wireless.md#count-org-wireless-clients-sessions)
* [Search Org Wireless Client Events](../../doc/controllers/orgs-clients-wireless.md#search-org-wireless-client-events)
* [Search Org Wireless Client Sessions](../../doc/controllers/orgs-clients-wireless.md#search-org-wireless-client-sessions)
* [Search Org Wireless Clients](../../doc/controllers/orgs-clients-wireless.md#search-org-wireless-clients)


# Count Org Wireless Clients

Count Org Wireless Clients

```go
CountOrgWirelessClients(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgClientsCountDistinctEnum,
    mac *string,
    hostname *string,
    device *string,
    os *string,
    model *string,
    ap *string,
    vlan *string,
    ssid *string,
    ipAddress *string,
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
| `distinct` | [`*models.OrgClientsCountDistinctEnum`](../../doc/models/org-clients-count-distinct-enum.md) | Query, Optional | - |
| `mac` | `*string` | Query, Optional | partial / full MAC address |
| `hostname` | `*string` | Query, Optional | partial / full hostname |
| `device` | `*string` | Query, Optional | device type, e.g. Mac, Nvidia, iPhone |
| `os` | `*string` | Query, Optional | os, e.g. Sierra, Yosemite, Windows 10 |
| `model` | `*string` | Query, Optional | model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” |
| `ap` | `*string` | Query, Optional | AP mac where the client has connected to |
| `vlan` | `*string` | Query, Optional | vlan |
| `ssid` | `*string` | Query, Optional | SSID |
| `ipAddress` | `*string` | Query, Optional | - |
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

distinct := models.OrgClientsCountDistinctEnum("device")

















ipAddress := "192.168.1.1"

page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsClientsWireless.CountOrgWirelessClients(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, &ipAddress, &page, &limit, nil, nil, &duration)
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


# Count Org Wireless Clients Sessions

Count Org Wireless Clients Sessions

```go
CountOrgWirelessClientsSessions(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgClientSessionsCountDistinctEnum,
    ap *string,
    band *models.Dot11BandEnum,
    clientFamily *string,
    clientManufacture *string,
    clientModel *string,
    clientOs *string,
    ssid *string,
    wlanId *string,
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
| `distinct` | [`*models.OrgClientSessionsCountDistinctEnum`](../../doc/models/org-client-sessions-count-distinct-enum.md) | Query, Optional | - |
| `ap` | `*string` | Query, Optional | AP MAC |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `clientFamily` | `*string` | Query, Optional | E.g. “Mac”, “iPhone”, “Apple watch” |
| `clientManufacture` | `*string` | Query, Optional | E.g. “Apple” |
| `clientModel` | `*string` | Query, Optional | E.g. “8+”, “XS” |
| `clientOs` | `*string` | Query, Optional | E.g. “Mojave”, “Windows 10”, “Linux” |
| `ssid` | `*string` | Query, Optional | SSID |
| `wlanId` | `*string` | Query, Optional | wlan_id |
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

distinct := models.OrgClientSessionsCountDistinctEnum("device")

















page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsClientsWireless.CountOrgWirelessClientsSessions(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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


# Search Org Wireless Client Events

Get Org Clients Events

```go
SearchOrgWirelessClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    reasonCode *int,
    ssid *string,
    ap *string,
    proto *models.Dot11ProtoEnum,
    band *models.Dot11BandEnum,
    wlanId *string,
    nacruleId *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseEventsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `reasonCode` | `*int` | Query, Optional | for assoc/disassoc events |
| `ssid` | `*string` | Query, Optional | SSID Name |
| `ap` | `*string` | Query, Optional | AP MAC |
| `proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Query, Optional | a / b / g / n / ac / ax |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `wlanId` | `*string` | Query, Optional | wlan_id |
| `nacruleId` | `*string` | Query, Optional | nacrule_id |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseEventsSearch`](../../doc/models/response-events-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

















limit := 100





duration := "10m"

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "results": [
    {
      "ap": "string",
      "band": "24",
      "bssid": "string",
      "channel": 0,
      "proto": "a",
      "ssid": "string",
      "text": "string",
      "timestamp": 0,
      "type": "string",
      "type_code": 0,
      "wlan_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# Search Org Wireless Client Sessions

Search Org Wireless Clients Sessions

```go
SearchOrgWirelessClientSessions(
    ctx context.Context,
    orgId uuid.UUID,
    ap *string,
    band *models.Dot11BandEnum,
    clientFamily *string,
    clientManufacture *string,
    clientModel *string,
    clientUsername *string,
    clientOs *string,
    ssid *string,
    wlanId *string,
    pskId *uuid.UUID,
    pskName *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SearchWirelssClientSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ap` | `*string` | Query, Optional | AP MAC |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `clientFamily` | `*string` | Query, Optional | E.g. “Mac”, “iPhone”, “Apple watch” |
| `clientManufacture` | `*string` | Query, Optional | E.g. “Apple” |
| `clientModel` | `*string` | Query, Optional | E.g. “8+”, “XS” |
| `clientUsername` | `*string` | Query, Optional | Username |
| `clientOs` | `*string` | Query, Optional | E.g. “Mojave”, “Windows 10”, “Linux” |
| `ssid` | `*string` | Query, Optional | SSID |
| `wlanId` | `*string` | Query, Optional | wlan_id |
| `pskId` | `*uuid.UUID` | Query, Optional | PSK ID |
| `pskName` | `*string` | Query, Optional | PSK Name |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.SearchWirelssClientSession`](../../doc/models/search-wirelss-client-session.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



















pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



limit := 100





duration := "10m"

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientSessions(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, &pskId, nil, &limit, nil, nil, &duration)
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
  "end": 1513177200,
  "limit": 10,
  "results": [
    {
      "ap": "5c5b350e0262",
      "band": "5",
      "client_manufacture": "Apple",
      "connect": 1565208388,
      "disconnect": 1565208448,
      "duration": 60.09423865,
      "mac": "b019c66c8348",
      "org_id": "3139f2c2-fac6-11e5-8156-0242ac110006",
      "site_id": "70e0f468-fc13-11e5-85ad-0242ac110008",
      "ssid": "Dummy WLAN 2",
      "tags": [
        "disassociate"
      ],
      "timestamp": 1565208448.662,
      "wlan_id": "99bb4c74-f954-4f36-b844-6b030faffabc"
    }
  ],
  "start": 1511967600,
  "total": 100
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


# Search Org Wireless Clients

Search Org Wireless Clients

```go
SearchOrgWirelessClients(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    mac *string,
    ipAddress *string,
    hostname *string,
    device *string,
    os *string,
    model *string,
    ap *string,
    pskId *uuid.UUID,
    pskName *string,
    vlan *string,
    ssid *string,
    text *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseClientSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | Site ID |
| `mac` | `*string` | Query, Optional | partial / full MAC address |
| `ipAddress` | `*string` | Query, Optional | - |
| `hostname` | `*string` | Query, Optional | partial / full hostname |
| `device` | `*string` | Query, Optional | device type, e.g. Mac, Nvidia, iPhone |
| `os` | `*string` | Query, Optional | os, e.g. Sierra, Yosemite, Windows 10 |
| `model` | `*string` | Query, Optional | model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” |
| `ap` | `*string` | Query, Optional | AP mac where the client has connected to |
| `pskId` | `*uuid.UUID` | Query, Optional | PSK ID |
| `pskName` | `*string` | Query, Optional | PSK Name |
| `vlan` | `*string` | Query, Optional | vlan |
| `ssid` | `*string` | Query, Optional | SSID |
| `text` | `*string` | Query, Optional | partial / full MAC address, hostname, username, psk_name or ip |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseClientSearch`](../../doc/models/response-client-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





ipAddress := "192.168.1.1"











pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")









limit := 100





duration := "10m"

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClients(ctx, orgId, nil, nil, &ipAddress, nil, nil, nil, nil, nil, &pskId, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 17141231418.812,
  "limit": 118,
  "next": "next8",
  "results": [
    {
      "ap": [
        "a83a79a947ee",
        "003e73170b4c"
      ],
      "app_version": [
        "0.100.3"
      ],
      "band": "5",
      "device": [
        "Mac"
      ],
      "ftc": false,
      "hardware": "Apple Wi-Fi adapter",
      "hostname": [
        "hostname-a",
        "hostname-b"
      ],
      "ip": [
        "10.5.23.43",
        "192.168.0.2"
      ],
      "last_ap": "a83a79a947ee",
      "last_devuce": "Mac",
      "last_firmware": "wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675",
      "last_hostname": "hostname-a",
      "last_ip": "10.5.23.43",
      "last_model": "MBP 16\\\" M1 2021",
      "last_os": "Sonoma",
      "last_os_version": "14.4.1 (Build 23E224)",
      "last_psk_id": "abf7dc5c-bb51-4bb7-93b6-5547400ffe11",
      "last_psk_name": "iot",
      "last_ssid": "IoT SSID",
      "last_username": "user@corp.com",
      "last_vlan": 10,
      "last_wlan_id": "e5d67b07-aae8-494b-8584-cbc20c8110aa",
      "mac": "bcd074000000",
      "mfg": "Apple",
      "model": "MBP 16\\\" M1 2021",
      "org_id": "1abff1aa-4571-4c1f-a409-153a1e7a7a24",
      "os": [
        "Sonoma"
      ],
      "os_version": [
        "14.4.1 (Build 23E224)"
      ],
      "protocol": "ax",
      "psk_id": [
        "abf7dc5c-bb51-4bb7-93b6-5547400ffe11"
      ],
      "psk_name": [
        "iot"
      ],
      "sdk_version": [
        "0.100.3"
      ],
      "site_id": "25ff5219-9be7-4db9-907d-0c9b60445147",
      "site_ids": [
        "25ff5219-9be7-4db9-907d-0c9b60445147"
      ],
      "ssid": [
        "IoT SSID"
      ],
      "timestamp": 1714124722.113,
      "username": [
        "user@corp.com"
      ],
      "vlan": [
        10
      ]
    }
  ],
  "start": 10,
  "total": 44
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

