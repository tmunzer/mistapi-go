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

Count wireless client event records across the organization, optionally grouped by event attributes and filtered by event type, WLAN, radio, site, and time range.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteClientEventsCountDistinctEnum`](../../doc/models/site-client-events-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `band`, `channel`, `proto`, `ssid`, `type`, `wlan_id` |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions). Accepts multiple comma-separated values. |
| `reasonCode` | `*int` | Query, Optional | Reason code filter for association and disassociation events |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address |
| `proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Query, Optional | 802.11 protocol used to filter results. enum: `a`, `ac`, `ax`, `b`, `be`, `g`, `n` |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 band used to filter radio results. enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable` |
| `wlanId` | `*string` | Query, Optional | Filter results by WLAN identifier |
| `siteId` | `*uuid.UUID` | Query, Optional | Filter results by site identifier |
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

distinct := models.SiteClientEventsCountDistinctEnum_ENUMTYPE

mType := "MARVIS_EVENT_CLIENT_AUTH_FAILURE,CLIENT_DEAUTHENTICATION"

siteId := uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")

duration := "10m"

limit := 100

apiResponse, err := orgsClientsWireless.CountOrgWirelessClientEvents(ctx, orgId, &distinct, &mType, nil, nil, nil, nil, nil, nil, &siteId, nil, nil, &duration, &limit)
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


# Count Org Wireless Clients

Count wireless client records across the organization, optionally grouped by `distinct` and filtered by client identity, AP, SSID, VLAN, IP, and time range.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgClientsCountDistinctEnum`](../../doc/models/org-clients-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `ap`, `device`, `hostname`, `ip`, `mac`, `model`, `os`, `ssid`, `vlan`<br><br>**Default**: `"device"` |
| `mac` | `*string` | Query, Optional | Partial / full Client MAC address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*bccddeeff`) are not supported. Accepts multiple comma-separated values. |
| `hostname` | `*string` | Query, Optional | Partial / full Client hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `everest*` and `*rest*` match `my-everest-client`). Suffix-only wildcards (e.g. `*everest`) are not supported. Accepts multiple comma-separated values. |
| `device` | `*string` | Query, Optional | Filter results by device type |
| `os` | `*string` | Query, Optional | Filter results by operating system |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address |
| `vlan` | `*string` | Query, Optional | Filter results by VLAN ID |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `ip` | `*string` | Query, Optional | Filter results by IPv4 address |
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

distinct := models.OrgClientsCountDistinctEnum_DEVICE

mac := "aabbccddeeff,aabbcc*"

hostname := "my-everest-client,my-everest*"

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


# Count Org Wireless Clients Sessions

Count historical wireless client sessions across the organization, optionally grouped by `distinct` and filtered by AP, band, client attributes, SSID, WLAN, and time range.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgClientSessionsCountDistinctEnum`](../../doc/models/org-client-sessions-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `ap`, `device`, `hostname`, `ip`, `model`, `os`, `ssid`, `vlan`<br><br>**Default**: `"device"` |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address |
| `band` | [`*models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Query, Optional | 802.11 band used to filter radio results. enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable` |
| `clientFamily` | `*string` | Query, Optional | E.g. "Mac", "iPhone", "Apple watch" |
| `clientManufacture` | `*string` | Query, Optional | Filter results by client manufacturer, e.g. "Apple" |
| `clientModel` | `*string` | Query, Optional | Filter results by client model, e.g. "8+", "XS" |
| `clientOs` | `*string` | Query, Optional | E.g. "Mojave", "Windows 10", "Linux" |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `wlanId` | `*uuid.UUID` | Query, Optional | Filter results by WLAN identifier |
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


# Search Org Wireless Client Events

Search wireless client event records across the organization with filters for event type, reason code, SSID, AP, key management, protocol, band, WLAN, NAC rule, and time range.

```go
SearchOrgWirelessClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    reasonCode *string,
    ssid *string,
    ap *string,
    keyMgmt *string,
    proto *string,
    band *string,
    wlanId *uuid.UUID,
    nacruleId *uuid.UUID,
    start *string,
    end *string,
    duration *string,
    sort *string,
    limit *int,
    searchAfter *string) (
    models.ApiResponse[models.ResponseEventsSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions). Accepts multiple comma-separated values. |
| `reasonCode` | `*string` | Query, Optional | Reason code filter for association and disassociation events. Accepts multiple comma-separated integer values. |
| `ssid` | `*string` | Query, Optional | Filter results by SSID. Accepts multiple comma-separated values. |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address. Accepts multiple comma-separated values. |
| `keyMgmt` | `*string` | Query, Optional | Key management protocol used to filter client events. enum: `WPA2-PSK`, `WPA2-PSK/CCMP`, `WPA2-PSK-FT`, `WPA2-PSK-SHA256`, `WPA3-EAP-SHA256`, `WPA3-EAP-SHA256/CCMP`, `WPA3-EAP-FT/GCMP256`, `WPA3-SAE-FT`, `WPA3-SAE-PSK`. Accepts multiple comma-separated values. |
| `proto` | `*string` | Query, Optional | 802.11 protocol used to filter results. enum: `a`, `ac`, `ax`, `b`, `be`, `g`, `n`. Accepts multiple comma-separated values. |
| `band` | `*string` | Query, Optional | 802.11 band used to filter radio results. enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable`. Accepts multiple comma-separated values. |
| `wlanId` | `*uuid.UUID` | Query, Optional | Filter results by WLAN identifier |
| `nacruleId` | `*uuid.UUID` | Query, Optional | Filter results by NAC rule identifier. Accepts multiple comma-separated values. |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsSearch](../../doc/models/response-events-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := "CLIENT_IP_ASSIGNED,CLIENT_DEAUTHENTICATION"

reasonCode := "0,14"

ssid := "Corp,Guest"

ap := "5c5b53010101,5c5b53020202"

keyMgmt := "WPA2-PSK,WPA2-PSK/CCMP"

proto := "ax,ac"

band := "6,5"

wlanId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

nacruleId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

duration := "10m"

sort := "-site_id"

limit := 100

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientEvents(ctx, orgId, &mType, &reasonCode, &ssid, &ap, &keyMgmt, &proto, &band, &wlanId, &nacruleId, nil, nil, &duration, &sort, &limit, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Wireless Client Sessions

Search historical wireless client sessions across the organization with filters for AP, band, client family, manufacturer, model, OS, username, SSID, WLAN, PPSK, and time range.

```go
SearchOrgWirelessClientSessions(
    ctx context.Context,
    orgId uuid.UUID,
    ap *string,
    band *string,
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
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.SearchWirelessClientSession],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address |
| `band` | `*string` | Query, Optional | 802.11 band used to filter radio results. enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable`. Accepts multiple comma-separated values. |
| `clientFamily` | `*string` | Query, Optional | E.g. "Mac", "iPhone", "Apple watch". Accepts multiple comma-separated values. |
| `clientManufacture` | `*string` | Query, Optional | Filter results by client manufacturer, e.g. "Apple". Accepts multiple comma-separated values. |
| `clientModel` | `*string` | Query, Optional | Filter results by client model, e.g. "8+", "XS". Accepts multiple comma-separated values. |
| `clientUsername` | `*string` | Query, Optional | Filter results by client username. Accepts multiple comma-separated values. |
| `clientOs` | `*string` | Query, Optional | E.g. "Mojave", "Windows 10", "Linux". Accepts multiple comma-separated values. |
| `ssid` | `*string` | Query, Optional | Filter results by SSID. Accepts multiple comma-separated values. |
| `wlanId` | `*uuid.UUID` | Query, Optional | Filter results by WLAN identifier |
| `pskId` | `*string` | Query, Optional | PSK identifier used to filter the results |
| `pskName` | `*string` | Query, Optional | Filter results by PSK name. Accepts multiple comma-separated values. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchWirelessClientSession](../../doc/models/search-wireless-client-session.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ap := "5c5b53010101"

band := "5,6"

clientFamily := "Phone/Tablet/Wearable,Monitoring Device"

clientManufacture := "Hewlett Packard Enterprise,Unknown"

clientModel := "iPhone 8,Aruba S0U52A"

clientUsername := "john.doe,jane.doe"

clientOs := "iOS 16.7.16,Unknown"

ssid := "Corp,Guest"

wlanId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

pskId := "000000ab-00ab-00ab-00ab-0000000000ab"

pskName := "psk-a,psk-b"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientSessions(ctx, orgId, &ap, &band, &clientFamily, &clientManufacture, &clientModel, &clientUsername, &clientOs, &ssid, &wlanId, &pskId, &pskName, &limit, nil, nil, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Wireless Clients

Search wireless client records across the organization with filters for site, AP, band, device identity, hostname, IP, MAC address, username, SSID, PPSK, VLAN, and time range.

```go
SearchOrgWirelessClients(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *uuid.UUID,
    ap *string,
    band *string,
    device *string,
    hostname *string,
    ip *string,
    mac *string,
    model *string,
    os *string,
    pskId *string,
    pskName *string,
    ssid *string,
    text *string,
    username *string,
    vlan *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseClientSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*uuid.UUID` | Query, Optional | Filter results by site identifier |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address |
| `band` | `*string` | Query, Optional | Comma separated list of Radio band (e.g. `24,5`). enum: `24`, `5`, `6` |
| `device` | `*string` | Query, Optional | Comma separated list of Device type (e.g. `Mac,iPhone`). Case sensitive |
| `hostname` | `*string` | Query, Optional | Partial / full Client hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `everest*` and `*rest*` match `my-everest-client`). Suffix-only wildcards (e.g. `*everest`) are not supported. Accepts multiple comma-separated values. |
| `ip` | `*string` | Query, Optional | Partial / full Client IP address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `10.100.10.*` and `*100.10.*` match `10.100.10.54`). Suffix-only wildcards (e.g. `*.54`) are not supported. Accepts multiple comma-separated values. |
| `mac` | `*string` | Query, Optional | Partial / full Client MAC address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*bccddeeff`) are not supported. Accepts multiple comma-separated values. |
| `model` | `*string` | Query, Optional | Only available for clients running the Marvis Client app, model, e.g. "MBP 15 late 2013", 6, 6s, "8+ GSM" |
| `os` | `*string` | Query, Optional | Only available for clients running the Marvis Client app, os, e.g. Sierra, Yosemite, Windows 10 |
| `pskId` | `*string` | Query, Optional | PSK identifier used to filter the results |
| `pskName` | `*string` | Query, Optional | Only available for clients using PPSK authentication, the Name of the PSK |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `text` | `*string` | Query, Optional | Partial / full MAC address, hostname, username, psk_name or ip. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `everest*` and `*rest*` match `my-everest-client`). Suffix-only wildcards (e.g. `*everest`) are not supported |
| `username` | `*string` | Query, Optional | Partial / full username. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `johndoe*` and `*mycorp*` match `johndoe@mycorp.com`). Suffix-only wildcards (e.g. `*mycorp.com`) are not supported. Accepts multiple comma-separated values. |
| `vlan` | `*string` | Query, Optional | Filter results by VLAN ID |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClientSearch](../../doc/models/response-client-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")

ap := "5c5b53010101"

band := "5,6"

device := "iPhone"

hostname := "my-everest-client,my-everest*"

ip := "10.100.10.54,10.100.10.*"

mac := "aabbccddeeff,aabbcc*"

model := "iPhone 8"

os := "Windows 10"

pskId := "000000ab-00ab-00ab-00ab-0000000000ab"

pskName := "MyPPSK"

ssid := "MySSID"

text := "5c5b530"

username := "johndoe,johnd*"

vlan := "10"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsWireless.SearchOrgWirelessClients(ctx, orgId, &siteId, &ap, &band, &device, &hostname, &ip, &mac, &model, &os, &pskId, &pskName, &ssid, &text, &username, &vlan, &limit, nil, nil, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

