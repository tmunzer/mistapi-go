# Sites Clients-Wireless

```go
sitesClientsWireless := client.SitesClientsWireless()
```

## Class Name

`SitesClientsWireless`

## Methods

* [Count Site Wireless Client Events](../../doc/controllers/sites-clients-wireless.md#count-site-wireless-client-events)
* [Count Site Wireless Client Sessions](../../doc/controllers/sites-clients-wireless.md#count-site-wireless-client-sessions)
* [Count Site Wireless Clients](../../doc/controllers/sites-clients-wireless.md#count-site-wireless-clients)
* [Get Site Events for Client](../../doc/controllers/sites-clients-wireless.md#get-site-events-for-client)
* [Search Site Wireless Client Events](../../doc/controllers/sites-clients-wireless.md#search-site-wireless-client-events)
* [Search Site Wireless Client Sessions](../../doc/controllers/sites-clients-wireless.md#search-site-wireless-client-sessions)
* [Search Site Wireless Clients](../../doc/controllers/sites-clients-wireless.md#search-site-wireless-clients)


# Count Site Wireless Client Events

Count by Distinct Attributes of Client-Events

```go
CountSiteWirelessClientEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteClientEventsCountDistinctEnum,
    mType *string,
    reasonCode *int,
    ssid *string,
    ap *string,
    proto *models.Dot11ProtoEnum,
    band *models.Dot11BandEnum,
    wlanId *string,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteClientEventsCountDistinctEnum`](../../doc/models/site-client-events-count-distinct-enum.md) | Query, Optional | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `reasonCode` | `*int` | Query, Optional | For assoc/disassoc events |
| `ssid` | `*string` | Query, Optional | SSID Name |
| `ap` | `*string` | Query, Optional | AP MAC |
| `proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Query, Optional | a / b / g / n / ac / ax |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `wlanId` | `*string` | Query, Optional | WLAN_id |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteClientEventsCountDistinctEnum_ENUMTYPE



















duration := "10m"

limit := 100

apiResponse, err := sitesClientsWireless.CountSiteWirelessClientEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# Count Site Wireless Client Sessions

Count by Distinct Attributes of Client Sessions

```go
CountSiteWirelessClientSessions(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteClientSessionsCountDistinctEnum,
    ap *string,
    band *models.Dot11BandEnum,
    clientFamily *string,
    clientManufacture *string,
    clientModel *string,
    clientOs *string,
    ssid *string,
    wlanId *string,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteClientSessionsCountDistinctEnum`](../../doc/models/site-client-sessions-count-distinct-enum.md) | Query, Optional | **Default**: `"mac"` |
| `ap` | `*string` | Query, Optional | AP MAC |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `clientFamily` | `*string` | Query, Optional | E.g. "Mac", "iPhone", "Apple watch" |
| `clientManufacture` | `*string` | Query, Optional | E.g. "Apple" |
| `clientModel` | `*string` | Query, Optional | E.g. "8+", "XS" |
| `clientOs` | `*string` | Query, Optional | E.g. "Mojave", "Windows 10", "Linux" |
| `ssid` | `*string` | Query, Optional | SSID |
| `wlanId` | `*string` | Query, Optional | WLAN_id |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteClientSessionsCountDistinctEnum_MAC





















duration := "10m"

limit := 100

apiResponse, err := sitesClientsWireless.CountSiteWirelessClientSessions(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# Count Site Wireless Clients

Count by Distinct Attributes of Clients

```go
CountSiteWirelessClients(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteClientsCountDistinctEnum,
    ssid *string,
    ap *string,
    ipAddress *string,
    vlan *string,
    hostname *string,
    os *string,
    model *string,
    device *string,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteClientsCountDistinctEnum`](../../doc/models/site-clients-count-distinct-enum.md) | Query, Optional | **Default**: `"device"` |
| `ssid` | `*string` | Query, Optional | - |
| `ap` | `*string` | Query, Optional | - |
| `ipAddress` | `*string` | Query, Optional | - |
| `vlan` | `*string` | Query, Optional | - |
| `hostname` | `*string` | Query, Optional | - |
| `os` | `*string` | Query, Optional | - |
| `model` | `*string` | Query, Optional | - |
| `device` | `*string` | Query, Optional | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteClientsCountDistinctEnum_HOSTNAME





ipAddress := "192.168.1.1"















duration := "10m"

limit := 100

apiResponse, err := sitesClientsWireless.CountSiteWirelessClients(ctx, siteId, &distinct, nil, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# Get Site Events for Client

Get the list of events for a specific client

```go
GetSiteEventsForClient(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    mType *string,
    proto *models.Dot11ProtoEnum,
    band *models.Dot11BandEnum,
    channel *string,
    wlanId *string,
    ssid *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseClientEventsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `mType` | `*string` | Query, Optional | E.g. MARVIS_EVENT_CLIENT_DHCP_STUCK |
| `proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Query, Optional | a / b / g / n / ac / ax |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `channel` | `*string` | Query, Optional | - |
| `wlanId` | `*string` | Query, Optional | - |
| `ssid` | `*string` | Query, Optional | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClientEventsSearch](../../doc/models/response-client-events-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

















duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesClientsWireless.GetSiteEventsForClient(ctx, siteId, clientMac, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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
  "end": 1513176951,
  "limit": 10,
  "results": [
    {
      "ap": "5c5b350eb31b",
      "band": "5",
      "bssid": "5c5b350918f1",
      "channel": 149,
      "proto": "ac",
      "ssid": "Guest",
      "text": "Status code 0 \"Successful\"",
      "timestamp": 1513358874.667,
      "type": "CLIENT_DNS_OK",
      "type_code": 15,
      "wlan_id": "be22bba7-8e22-e1cf-5185-b880816fe2cf"
    }
  ],
  "start": 1512572151,
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


# Search Site Wireless Client Events

Get Site Clients Events

```go
SearchSiteWirelessClientEvents(
    ctx context.Context,
    siteId uuid.UUID,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `reasonCode` | `*int` | Query, Optional | For assoc/disassoc events |
| `ssid` | `*string` | Query, Optional | SSID Name |
| `ap` | `*string` | Query, Optional | AP MAC |
| `proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Query, Optional | a / b / g / n / ac / ax |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `wlanId` | `*string` | Query, Optional | WLAN_id |
| `nacruleId` | `*string` | Query, Optional | nacrule_id |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsSearch](../../doc/models/response-events-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

















limit := 100





duration := "10m"

apiResponse, err := sitesClientsWireless.SearchSiteWirelessClientEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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


# Search Site Wireless Client Sessions

Search Client Sessions

```go
SearchSiteWirelessClientSessions(
    ctx context.Context,
    siteId uuid.UUID,
    ap *string,
    band *models.Dot11BandEnum,
    clientFamily *string,
    clientManufacture *string,
    clientModel *string,
    clientUsername *string,
    clientOs *string,
    ssid *string,
    wlanId *string,
    pskId *string,
    pskName *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseClientSessionsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `ap` | `*string` | Query, Optional | AP MAC |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 Band |
| `clientFamily` | `*string` | Query, Optional | E.g. "Mac", "iPhone", "Apple watch" |
| `clientManufacture` | `*string` | Query, Optional | E.g. "Apple" |
| `clientModel` | `*string` | Query, Optional | E.g. "8+", "XS" |
| `clientUsername` | `*string` | Query, Optional | Username |
| `clientOs` | `*string` | Query, Optional | E.g. "Mojave", "Windows 10", "Linux" |
| `ssid` | `*string` | Query, Optional | SSID |
| `wlanId` | `*string` | Query, Optional | WLAN_id |
| `pskId` | `*string` | Query, Optional | PSK ID |
| `pskName` | `*string` | Query, Optional | PSK Name |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClientSessionsSearch](../../doc/models/response-client-sessions-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



















pskId := "000000ab-00ab-00ab-00ab-0000000000ab"



limit := 100





duration := "10m"

apiResponse, err := sitesClientsWireless.SearchSiteWirelessClientSessions(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, &pskId, nil, &limit, nil, nil, &duration)
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


# Search Site Wireless Clients

Search Wireless Clients

**NOTE**: fuzzy logic can be used with ‘*’, supported filters: mac, hostname, device, os, model. E.g. /clients/search?device=Mac*&hostname=jerry

```go
SearchSiteWirelessClients(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    ipAddress *string,
    hostname *string,
    device *string,
    os *string,
    model *string,
    ap *string,
    ssid *string,
    text *string,
    nacruleId *string,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Partial / full MAC address |
| `ipAddress` | `*string` | Query, Optional | - |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `device` | `*string` | Query, Optional | Device type, e.g. Mac, Nvidia, iPhone |
| `os` | `*string` | Query, Optional | OS, e.g. Sierra, Yosemite, Windows 10 |
| `model` | `*string` | Query, Optional | model, e.g. "MBP 15 late 2013", 6, 6s, "8+ GSM" |
| `ap` | `*string` | Query, Optional | AP mac where the client has connected to |
| `ssid` | `*string` | Query, Optional | - |
| `text` | `*string` | Query, Optional | Partial / full MAC address, hostname, username, psk_name or ip |
| `nacruleId` | `*string` | Query, Optional | nacrule_id |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClientSearch](../../doc/models/response-client-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



ipAddress := "192.168.1.1"

















limit := 100





duration := "10m"

apiResponse, err := sitesClientsWireless.SearchSiteWirelessClients(ctx, siteId, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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

