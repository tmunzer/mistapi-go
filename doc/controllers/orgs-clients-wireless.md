# Orgs Clients-Wireless

```go
orgsClientsWireless := client.OrgsClientsWireless()
```

## Class Name

`OrgsClientsWireless`

## Methods

* [Count Org Wireless Client Events](../../doc/controllers/orgs-clients-wireless.md#count-org-wireless-client-events)
* [Count Org Wireless Clients](../../doc/controllers/orgs-clients-wireless.md#count-org-wireless-clients)
* [Count Org Wireless Clients Sessions](../../doc/controllers/orgs-clients-wireless.md#count-org-wireless-clients-sessions)
* [Search Org Wireless Client Events](../../doc/controllers/orgs-clients-wireless.md#search-org-wireless-client-events)
* [Search Org Wireless Client Sessions](../../doc/controllers/orgs-clients-wireless.md#search-org-wireless-client-sessions)
* [Search Org Wireless Clients](../../doc/controllers/orgs-clients-wireless.md#search-org-wireless-clients)


# Count Org Wireless Client Events

Count by Distinct Attributes of Client-Events

```go
CountOrgWirelessClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.SiteClientEventsCountDistinctEnum,
    mType *string,
    reasonCode *int,
    ssid *string,
    ap *string,
    proto *models.Dot11ProtoEnum,
    band *models.Dot11BandEnum,
    wlanId *string,
    siteId *uuid.UUID,
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
| `distinct` | [`*models.SiteClientEventsCountDistinctEnum`](../../doc/models/site-client-events-count-distinct-enum.md) | Query, Optional | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `reasonCode` | `*int` | Query, Optional | For assoc/disassoc events |
| `ssid` | `*string` | Query, Optional | SSID Name |
| `ap` | `*string` | Query, Optional | AP MAC |
| `proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Query, Optional | a / b / g / n / ac / ax |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `wlanId` | `*string` | Query, Optional | WLAN ID |
| `siteId` | `*uuid.UUID` | Query, Optional | Site ID |
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

distinct := models.SiteClientEventsCountDistinctEnum_ENUMTYPE

siteId := uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")

duration := "10m"

limit := 100

apiResponse, err := orgsClientsWireless.CountOrgWirelessClientEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, &siteId, nil, nil, &duration, &limit)
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


# Count Org Wireless Clients

Count by Distinct Attributes of Org Wireless Clients

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
    ip *string,
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
| `distinct` | [`*models.OrgClientsCountDistinctEnum`](../../doc/models/org-clients-count-distinct-enum.md) | Query, Optional | **Default**: `"device"` |
| `mac` | `*string` | Query, Optional | Partial / full MAC address |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `device` | `*string` | Query, Optional | Device type, e.g. Mac, Nvidia, iPhone |
| `os` | `*string` | Query, Optional | OS, e.g. Sierra, Yosemite, Windows 10 |
| `model` | `*string` | Query, Optional | Model, e.g. "MBP 15 late 2013", 6, 6s, "8+ GSM" |
| `ap` | `*string` | Query, Optional | AP mac where the client has connected to |
| `vlan` | `*string` | Query, Optional | VLAN |
| `ssid` | `*string` | Query, Optional | SSID |
| `ip` | `*string` | Query, Optional | - |
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

distinct := models.OrgClientsCountDistinctEnum_DEVICE

mac := "5c5b53010101"

hostname := "my-hostname"

device := "iPhone"

os := "Windows 10"

model := "iPhone 8"

ap := "5c5b53010101"

vlan := "10"

ssid := "MySSID"

ip := "192.168.1.1"

duration := "10m"

limit := 100

apiResponse, err := orgsClientsWireless.CountOrgWirelessClients(ctx, orgId, &distinct, &mac, &hostname, &device, &os, &model, &ap, &vlan, &ssid, &ip, nil, nil, &duration, &limit)
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


# Count Org Wireless Clients Sessions

Count by Distinct Attributes of Org Wireless Clients Sessions

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
    wlanId *uuid.UUID,
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
| `distinct` | [`*models.OrgClientSessionsCountDistinctEnum`](../../doc/models/org-client-sessions-count-distinct-enum.md) | Query, Optional | **Default**: `"device"` |
| `ap` | `*string` | Query, Optional | AP MAC |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `clientFamily` | `*string` | Query, Optional | E.g. "Mac", "iPhone", "Apple watch" |
| `clientManufacture` | `*string` | Query, Optional | E.g. "Apple" |
| `clientModel` | `*string` | Query, Optional | E.g. "8+", "XS" |
| `clientOs` | `*string` | Query, Optional | E.g. "Mojave", "Windows 10", "Linux" |
| `ssid` | `*string` | Query, Optional | SSID |
| `wlanId` | `*uuid.UUID` | Query, Optional | WLAN_id |
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

distinct := models.OrgClientSessionsCountDistinctEnum_DEVICE

ap := "5c5b53010101"

clientFamily := "iPhone"

clientManufacture := "Apple"

clientModel := "iPhone 8"

clientOs := "Windows 10"

ssid := "MySSID"

wlanId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

duration := "10m"

limit := 100

apiResponse, err := orgsClientsWireless.CountOrgWirelessClientsSessions(ctx, orgId, &distinct, &ap, nil, &clientFamily, &clientManufacture, &clientModel, &clientOs, &ssid, &wlanId, nil, nil, &duration, &limit)
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
    keyMgmt *models.ClientKeyMgmtEnum,
    proto *models.Dot11ProtoEnum,
    band *models.Dot11BandEnum,
    wlanId *uuid.UUID,
    nacruleId *uuid.UUID,
    start *string,
    end *string,
    duration *string,
    sort *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `reasonCode` | `*int` | Query, Optional | For assoc/disassoc events |
| `ssid` | `*string` | Query, Optional | SSID Name |
| `ap` | `*string` | Query, Optional | AP MAC |
| `keyMgmt` | [`*models.ClientKeyMgmtEnum`](../../doc/models/client-key-mgmt-enum.md) | Query, Optional | Key Management Protocol, e.g. WPA2-PSK, WPA3-SAE, WPA2-Enterprise |
| `proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Query, Optional | a / b / g / n / ac / ax |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `wlanId` | `*uuid.UUID` | Query, Optional | WLAN_id |
| `nacruleId` | `*uuid.UUID` | Query, Optional | Nacrule_id |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsSearch](../../doc/models/response-events-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

reasonCode := 7

ssid := "MySSID"

ap := "5c5b53010101"

keyMgmt := models.ClientKeyMgmtEnum_WPA2PSK

wlanId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

nacruleId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

duration := "10m"

sort := "-site_id"

limit := 100

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientEvents(ctx, orgId, nil, &reasonCode, &ssid, &ap, &keyMgmt, nil, nil, &wlanId, &nacruleId, nil, nil, &duration, &sort, &limit)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
    wlanId *uuid.UUID,
    pskId *string,
    pskName *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string) (
    models.ApiResponse[models.SearchWirelessClientSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ap` | `*string` | Query, Optional | AP MAC |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `clientFamily` | `*string` | Query, Optional | E.g. "Mac", "iPhone", "Apple watch" |
| `clientManufacture` | `*string` | Query, Optional | E.g. "Apple" |
| `clientModel` | `*string` | Query, Optional | E.g. "8+", "XS" |
| `clientUsername` | `*string` | Query, Optional | Username |
| `clientOs` | `*string` | Query, Optional | E.g. "Mojave", "Windows 10", "Linux" |
| `ssid` | `*string` | Query, Optional | SSID |
| `wlanId` | `*uuid.UUID` | Query, Optional | WLAN_id |
| `pskId` | `*string` | Query, Optional | PSK ID |
| `pskName` | `*string` | Query, Optional | PSK Name |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchWirelessClientSession](../../doc/models/search-wireless-client-session.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ap := "5c5b53010101"

clientFamily := "iPhone"

clientManufacture := "Apple"

clientModel := "iPhone 8"

clientUsername := "john.doe"

clientOs := "Windows 10"

ssid := "MySSID"

wlanId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

pskId := "000000ab-00ab-00ab-00ab-0000000000ab"

pskName := "MyPPSK"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientSessions(ctx, orgId, &ap, nil, &clientFamily, &clientManufacture, &clientModel, &clientUsername, &clientOs, &ssid, &wlanId, &pskId, &pskName, &limit, nil, nil, &duration, &sort)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Search Org Wireless Clients

Search Org Wireless Clients

```go
SearchOrgWirelessClients(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *uuid.UUID,
    mac *string,
    ip *string,
    hostname *string,
    band *string,
    device *string,
    os *string,
    model *string,
    ap *string,
    pskId *string,
    pskName *string,
    username *string,
    vlan *string,
    ssid *string,
    text *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string) (
    models.ApiResponse[models.ResponseClientSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*uuid.UUID` | Query, Optional | Site ID |
| `mac` | `*string` | Query, Optional | Partial / full MAC address |
| `ip` | `*string` | Query, Optional | - |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `band` | `*string` | Query, Optional | Radio band. enum: `24`, `5`, `6` |
| `device` | `*string` | Query, Optional | Device type, e.g. Mac, Nvidia, iPhone |
| `os` | `*string` | Query, Optional | Only available for clients running the Marvis Client app, os, e.g. Sierra, Yosemite, Windows 10 |
| `model` | `*string` | Query, Optional | Only available for clients running the Marvis Client app, model, e.g. "MBP 15 late 2013", 6, 6s, "8+ GSM" |
| `ap` | `*string` | Query, Optional | AP mac where the client has connected to |
| `pskId` | `*string` | Query, Optional | PSK ID |
| `pskName` | `*string` | Query, Optional | Only available for clients using PPSK authentication, the Name of the PSK |
| `username` | `*string` | Query, Optional | Only available for clients using 802.1X authentication, partial / full username |
| `vlan` | `*string` | Query, Optional | VLAN |
| `ssid` | `*string` | Query, Optional | SSID |
| `text` | `*string` | Query, Optional | Partial / full MAC address, hostname, username, psk_name or ip |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClientSearch](../../doc/models/response-client-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

mac := "5c5b53010101"

ip := "192.168.1.1"

hostname := "my-hostname"

band := "5"

device := "iPhone"

os := "Windows 10"

model := "iPhone 8"

ap := "5c5b53010101"

pskId := "000000ab-00ab-00ab-00ab-0000000000ab"

pskName := "MyPPSK"

username := "john.doe"

vlan := "10"

ssid := "MySSID"

text := "5c5b530"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClients(ctx, orgId, &siteId, &mac, &ip, &hostname, &band, &device, &os, &model, &ap, &pskId, &pskName, &username, &vlan, &ssid, &text, &limit, nil, nil, &duration, &sort)
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
      "last_device": "Mac",
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

