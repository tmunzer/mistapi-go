# Sites Clients-NAC

```go
sitesClientsNAC := client.SitesClientsNAC()
```

## Class Name

`SitesClientsNAC`

## Methods

* [Count Site Nac Client Events](../../doc/controllers/sites-clients-nac.md#count-site-nac-client-events)
* [Count Site Nac Clients](../../doc/controllers/sites-clients-nac.md#count-site-nac-clients)
* [Search Site Nac Client Events](../../doc/controllers/sites-clients-nac.md#search-site-nac-client-events)
* [Search Site Nac Clients](../../doc/controllers/sites-clients-nac.md#search-site-nac-clients)


# Count Site Nac Client Events

Count by Distinct Attributes of NAC Client-Events

```go
CountSiteNacClientEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteNacClientEventsCountDistinctEnum,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteNacClientEventsCountDistinctEnum`](../../doc/models/site-nac-client-events-count-distinct-enum.md) | Query, Optional | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions) |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

apiResponse, err := sitesClientsNAC.CountSiteNacClientEvents(ctx, siteId, nil, nil, nil, nil, &duration, &limit)
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


# Count Site Nac Clients

Count by Distinct Attributes of NAC Clients

```go
CountSiteNacClients(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteNacClientsCountDistinctEnum,
    lastNacruleId *string,
    nacruleMatched *bool,
    authType *string,
    lastVlanId *string,
    lastNasVendor *string,
    idpId *string,
    lastSsid *string,
    lastUsername *string,
    timestamp *float64,
    lastAp *string,
    mac *string,
    lastStatus *string,
    mType *string,
    mdmComplianceStatus *string,
    mdmProvider *string,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteNacClientsCountDistinctEnum`](../../doc/models/site-nac-clients-count-distinct-enum.md) | Query, Optional | NAC Policy Rule ID, if matched<br><br>**Default**: `"type"` |
| `lastNacruleId` | `*string` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `lastVlanId` | `*string` | Query, Optional | Vlan ID |
| `lastNasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `idpId` | `*string` | Query, Optional | SSO ID, if present and used |
| `lastSsid` | `*string` | Query, Optional | SSID |
| `lastUsername` | `*string` | Query, Optional | Username presented by the client |
| `timestamp` | `*float64` | Query, Optional | Start time, in epoch |
| `lastAp` | `*string` | Query, Optional | AP MAC connected to by client |
| `mac` | `*string` | Query, Optional | MAC address |
| `lastStatus` | `*string` | Query, Optional | Connection status of client i.e "permitted", "denied, "session_ended" |
| `mType` | `*string` | Query, Optional | Client type i.e. "wireless", "wired" etc. |
| `mdmComplianceStatus` | `*string` | Query, Optional | MDM compliance of client i.e "compliant", "not compliant" |
| `mdmProvider` | `*string` | Query, Optional | MDM provider of client’s organisation eg "intune", "jamf" |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteNacClientsCountDistinctEnum_ENUMTYPE

duration := "10m"

limit := 100

apiResponse, err := sitesClientsNAC.CountSiteNacClients(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# Search Site Nac Client Events

Search NAC Client Events

```go
SearchSiteNacClientEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    nacruleId *uuid.UUID,
    nacruleMatched *bool,
    dryrunNacruleId *string,
    dryrunNacruleMatched *bool,
    authType *string,
    vlan *int,
    nasVendor *string,
    bssid *string,
    idpId *uuid.UUID,
    idpRole *string,
    idpUsername *string,
    respAttrs []string,
    ssid *string,
    username *string,
    ap *string,
    randomMac *bool,
    mac *string,
    timestamp *float64,
    usermacLabel *string,
    text *string,
    nasIp *string,
    ingressVlan *string,
    start *string,
    end *string,
    duration *string,
    limit *int,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseEventsNacClientSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions) |
| `nacruleId` | `*uuid.UUID` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `dryrunNacruleId` | `*string` | Query, Optional | NAC Policy Dry Run Rule ID, if present and matched |
| `dryrunNacruleMatched` | `*bool` | Query, Optional | True - if dryrun rule present and matched with priority, False - if not matched or not present |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `vlan` | `*int` | Query, Optional | Vlan ID |
| `nasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `bssid` | `*string` | Query, Optional | BSSID |
| `idpId` | `*uuid.UUID` | Query, Optional | SSO ID, if present and used |
| `idpRole` | `*string` | Query, Optional | IDP returned roles/groups for the user |
| `idpUsername` | `*string` | Query, Optional | Username presented to the Identity Provider |
| `respAttrs` | `[]string` | Query, Optional | Radius attributes returned by NAC to NAS Devive<br><br>**Constraints**: *Unique Items Required* |
| `ssid` | `*string` | Query, Optional | SSID |
| `username` | `*string` | Query, Optional | Username presented by the client |
| `ap` | `*string` | Query, Optional | AP MAC |
| `randomMac` | `*bool` | Query, Optional | AP random macMAC |
| `mac` | `*string` | Query, Optional | MAC address |
| `timestamp` | `*float64` | Query, Optional | Time, in epoch |
| `usermacLabel` | `*string` | Query, Optional | Labels derived from usermac entry |
| `text` | `*string` | Query, Optional | Partial / full MAC address, username, device_mac or ap |
| `nasIp` | `*string` | Query, Optional | IP address of NAS device |
| `ingressVlan` | `*string` | Query, Optional | Vendor specific Vlan ID in radius requests |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order.<br><br>**Default**: `"wxid"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsNacClientSearch](../../doc/models/response-events-nac-client-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

respAttrs := []string{
    "Tunnel-Type=VLAN",
    "Tunnel-Medium-Type=IEEE-802",
    "Tunnel-Private-Group-Id=750",
    "User-Name=anonymous",
}

duration := "10m"

limit := 100

sort := "-site_id"

apiResponse, err := sitesClientsNAC.SearchSiteNacClientEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, respAttrs, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &sort, nil)
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


# Search Site Nac Clients

Search Site NAC Clients

```go
SearchSiteNacClients(
    ctx context.Context,
    siteId uuid.UUID,
    ap *string,
    authType *string,
    edrManaged *bool,
    edrProvider *models.EdrProviderEnum,
    edrStatus *models.EdrStatusEnum,
    family *string,
    hostname *string,
    idpId *string,
    mac *string,
    mdmManaged *bool,
    mdmCompliance *string,
    mdmProvider *string,
    mfg *string,
    model *string,
    mxedgeId *string,
    nacruleId *string,
    nacruleMatched *bool,
    nacruleName *string,
    nasVendor *string,
    nasIp *string,
    ingressVlan *string,
    os *string,
    ssid *string,
    status *models.NacClientLastStatusEnum,
    text *string,
    timestamp *float64,
    mType *string,
    usermacLabel []string,
    username *string,
    vlan *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseClientNacSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `ap` | `*string` | Query, Optional | AP MAC connected to by client |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `edrManaged` | `*bool` | Query, Optional | Filters NAC clients that are integrated with EDR providers |
| `edrProvider` | [`*models.EdrProviderEnum`](../../doc/models/edr-provider-enum.md) | Query, Optional | EDR provider of client's organization |
| `edrStatus` | [`*models.EdrStatusEnum`](../../doc/models/edr-status-enum.md) | Query, Optional | EDR Status of the NAC client |
| `family` | `*string` | Query, Optional | Client family, e.g. "Phone/Tablet/Wearable", "Access Point" |
| `hostname` | `*string` | Query, Optional | Client hostname, e.g. "my-laptop", "my-phone" |
| `idpId` | `*string` | Query, Optional | SSO ID, if present and used |
| `mac` | `*string` | Query, Optional | MAC address |
| `mdmManaged` | `*bool` | Query, Optional | Filters NAC clients that are managed by MDM providers |
| `mdmCompliance` | `*string` | Query, Optional | MDM compliance of client i.e "compliant", "not compliant" |
| `mdmProvider` | `*string` | Query, Optional | MDM provider of client’s organisation eg "intune", "jamf" |
| `mfg` | `*string` | Query, Optional | Client manufacturer, e.g. "apple", "cisco", "juniper" |
| `model` | `*string` | Query, Optional | Client model, e.g. "iPhone 12", "MX100" |
| `mxedgeId` | `*string` | Query, Optional | ID of Mist Edge that the client is connected through |
| `nacruleId` | `*string` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `nacruleName` | `*string` | Query, Optional | NAC Policy Rule Name matched |
| `nasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `nasIp` | `*string` | Query, Optional | IP address of NAS device |
| `ingressVlan` | `*string` | Query, Optional | Vendor specific Vlan ID in radius requests |
| `os` | `*string` | Query, Optional | Client OS, e.g. "iOS 18.1", "Android", "Windows", "Linux" |
| `ssid` | `*string` | Query, Optional | SSID |
| `status` | [`*models.NacClientLastStatusEnum`](../../doc/models/nac-client-last-status-enum.md) | Query, Optional | Connection status of client i.e "permitted", "denied, "session_ended" |
| `text` | `*string` | Query, Optional | partial / full MAC address, last_username, device_mac, nas_ip or last_ap |
| `timestamp` | `*float64` | Query, Optional | Start time, in epoch |
| `mType` | `*string` | Query, Optional | Client type i.e. "wireless", "wired" etc. |
| `usermacLabel` | `[]string` | Query, Optional | Labels derived from usermac entry<br><br>**Constraints**: *Unique Items Required* |
| `username` | `*string` | Query, Optional | Username presented by the client |
| `vlan` | `*string` | Query, Optional | Vlan name or ID assigned to the client |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order.<br><br>**Default**: `"wxid"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClientNacSearch](../../doc/models/response-client-nac-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

status := models.NacClientLastStatusEnum_PERMITTED

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesClientsNAC.SearchSiteNacClients(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &status, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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

