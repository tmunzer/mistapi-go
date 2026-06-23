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
* [Send Site Nac Client Co A](../../doc/controllers/sites-clients-nac.md#send-site-nac-client-co-a)


# Count Site Nac Client Events

Count NAC client events for a site, optionally grouped by the `distinct` field and filtered by event type and time range. Use [Count Org NAC Client Events](../../doc/controllers/orgs-clients-nac.md#count-org-nac-client-events) to count NAC client events across the organization.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteNacClientEventsCountDistinctEnum`](../../doc/models/site-nac-client-events-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `ap`, `auth_type`, `dryrun_nacrule_id`, `mac`, `nacrule_id`, `nas_vendor`, `ssid`, `type`, `username`, `vlan` |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions) |
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

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

apiResponse, err := sitesClientsNAC.CountSiteNacClientEvents(ctx, siteId, nil, nil, nil, nil, &duration, &limit)
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


# Count Site Nac Clients

Count NAC clients for a site, optionally grouped by the `distinct` field and filtered by authentication, identity, endpoint, network, and time attributes. Use [Count Org NAC Clients](../../doc/controllers/orgs-clients-nac.md#count-org-nac-clients) to count NAC clients across the organization.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteNacClientsCountDistinctEnum`](../../doc/models/site-nac-clients-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `ap`, `auth_type`, `device_mac`, `edr_managed`, `edr_provider`, `edr_status`, `family`, `hostname`, `idp_id`, `mfg`, `mdm_compliance`, `mdm_managed`, `mdm_provider`, `model`, `mxedge_id`, `nacrule_matched`, `nacrule_name`, `nacrule_id`, `nas_ip`, `nas_vendor`, `os`, `ssid`, `status`, `type`, `usermac_label`, `username`, `vlan`<br><br>**Default**: `"type"` |
| `lastNacruleId` | `*string` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `lastVlanId` | `*string` | Query, Optional | Filter results by last VLAN ID |
| `lastNasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `idpId` | `*string` | Query, Optional | SSO ID, if present and used |
| `lastSsid` | `*string` | Query, Optional | Filter results by last SSID |
| `lastUsername` | `*string` | Query, Optional | Username presented by the client |
| `lastAp` | `*string` | Query, Optional | AP MAC connected to by client |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `lastStatus` | `*string` | Query, Optional | Connection status of client i.e "permitted", "denied, "session_ended" |
| `mType` | `*string` | Query, Optional | Client type i.e. "wireless", "wired" etc. |
| `mdmComplianceStatus` | `*string` | Query, Optional | MDM compliance of client i.e "compliant", "not compliant" |
| `mdmProvider` | `*string` | Query, Optional | MDM provider of client’s organisation eg "intune", "jamf" |
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

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteNacClientsCountDistinctEnum_ENUMTYPE

duration := "10m"

limit := 100

apiResponse, err := sitesClientsNAC.CountSiteNacClients(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# Search Site Nac Client Events

Search NAC client events for a site with filters for authentication, NAC rule, identity provider, RADIUS, network, endpoint, and time attributes. Use [Search Org NAC Client Events](../../doc/controllers/orgs-clients-nac.md#search-org-nac-client-events) to search NAC client events across the organization.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

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
| `vlan` | `*int` | Query, Optional | Filter results by VLAN ID |
| `nasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `bssid` | `*string` | Query, Optional | Filter results by BSSID |
| `idpId` | `*uuid.UUID` | Query, Optional | SSO ID, if present and used |
| `idpRole` | `*string` | Query, Optional | IDP returned roles/groups for the user |
| `idpUsername` | `*string` | Query, Optional | Username presented to the Identity Provider |
| `respAttrs` | `[]string` | Query, Optional | RADIUS attributes returned by NAC to NAS Devive<br><br>**Constraints**: *Unique Items Required* |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `username` | `*string` | Query, Optional | Filter results by username |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address |
| `randomMac` | `*bool` | Query, Optional | Filter results by whether the client is using a randomized MAC address |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `usermacLabel` | `*string` | Query, Optional | Labels derived from usermac entry |
| `text` | `*string` | Query, Optional | Partial / full MAC address, username, device_mac or ap. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*bccddeeff`) are not supported |
| `nasIp` | `*string` | Query, Optional | IP address of NAS device |
| `ingressVlan` | `*string` | Query, Optional | Vendor specific VLAN ID in RADIUS requests |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order.<br><br>**Default**: `"wxid"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: NAC Client Events

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

apiResponse, err := sitesClientsNAC.SearchSiteNacClientEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, respAttrs, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &sort, nil)
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
  "end": 1513176951,
  "limit": 10,
  "results": [
    {
      "ap": "5c5b35513227",
      "auth_type": "eap-ttls",
      "bssid": "5c5b355fafcc",
      "dryrun_nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264012",
      "dryrun_nacrule_matched": true,
      "idp_id": "912ef72e-2239-4996-b81e-469e87a27cd6",
      "idp_role": [
        "itsuperusers",
        "vip"
      ],
      "mac": "ac3eb179e535",
      "nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62",
      "nacrule_matched": true,
      "nas_vendor": "juniper-mist",
      "org_id": "27547ac2-d114-4e04-beb1-f3f1e6e81ec6",
      "random_mac": "false",
      "resp_attrs": [
        "Tunnel-Type=VLAN",
        "Tunnel-Medium-Type=IEEE-802",
        "Tunnel-Private-Group-Id=750",
        "User-Name=anonymous"
      ],
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "ssid": "mist_nac",
      "timestamp": 1691512031.358188,
      "type": "NAC_CLIENT_PERMIT",
      "username": "user@deaflyz.net",
      "vlan": "750"
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Site Nac Clients

Search NAC clients for a site with filters for authentication, endpoint posture, identity, network, NAC rule, and time attributes. Use [Search Org NAC Clients](../../doc/controllers/orgs-clients-nac.md#search-org-nac-clients) to search NAC clients across the organization.

```go
SearchSiteNacClients(
    ctx context.Context,
    siteId uuid.UUID,
    ap *string,
    authType *string,
    certExpiryDuration *string,
    edrManaged *bool,
    edrProvider *models.EdrProviderEnum,
    edrStatus *models.EdrStatusEnum,
    family *string,
    hostname *string,
    idpId *string,
    mac *string,
    mdmCompliance *string,
    mdmProvider *string,
    mdmManaged *bool,
    mfg *string,
    model *string,
    nacruleName *string,
    nacruleId *string,
    nacruleMatched *bool,
    nasVendor *string,
    nasIp *string,
    ingressVlan *string,
    os *string,
    ssid *string,
    status *models.NacClientLastStatusEnum,
    text *string,
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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `ap` | `*string` | Query, Optional | MAC address of the AP the client is/was connected to |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `certExpiryDuration` | `*string` | Query, Optional | Filter by certificate expiry within a specific duration from now (e.g., "7d" for 7 days, "1m" for 1 month) |
| `edrManaged` | `*bool` | Query, Optional | Filters NAC clients that are integrated with EDR providers |
| `edrProvider` | [`*models.EdrProviderEnum`](../../doc/models/edr-provider-enum.md) | Query, Optional | EDR provider used to filter NAC clients. enum: `crowdstrike`, `sentinelone` |
| `edrStatus` | [`*models.EdrStatusEnum`](../../doc/models/edr-status-enum.md) | Query, Optional | EDR status used to filter NAC clients. enum: `sentinelone_healthy`, `sentinelone_infected`, `crowdstrike_low`, `crowdstrike_medium`, `crowdstrike_high`, `crowdstrike_critical`, `crowdstrike_informational` |
| `family` | `*string` | Query, Optional | Partial / full Client family (e.g. "Phone/Tablet/Wearable", "Access Point"). Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `Surveillance*` and `*urveillance*` match `Surveillance Camera`). Suffix-only wildcards (e.g. `*Camera`) are not supported |
| `hostname` | `*string` | Query, Optional | Partial / full Client hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `everest*` and `*rest*` match `my-everest-client`). Suffix-only wildcards (e.g. `*everest`) are not supported |
| `idpId` | `*string` | Query, Optional | SSO ID, if present and used |
| `mac` | `*string` | Query, Optional | Partial / full Client MAC address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*bccddeeff`) are not supported |
| `mdmCompliance` | `*string` | Query, Optional | MDM compliance of client i.e "compliant", "not compliant" |
| `mdmProvider` | `*string` | Query, Optional | MDM provider of client’s organization eg "intune", "jamf" |
| `mdmManaged` | `*bool` | Query, Optional | Filters NAC clients that are managed by MDM providers |
| `mfg` | `*string` | Query, Optional | Partial / full Client manufacturer (e.g. "apple", "cisco", "juniper"). Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `Raspberry Pi*` and `*Pi*` match `Raspberry Pi Trading Ltd`). Suffix-only wildcards (e.g. `*Ltd`) are not supported |
| `model` | `*string` | Query, Optional | Client model, e.g. "iPhone 12", "MX100" |
| `nacruleName` | `*string` | Query, Optional | NAC Policy Rule Name matched |
| `nacruleId` | `*string` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `nasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `nasIp` | `*string` | Query, Optional | IP address of NAS device |
| `ingressVlan` | `*string` | Query, Optional | Vendor specific VLAN ID in RADIUS requests |
| `os` | `*string` | Query, Optional | Client OS, e.g. "iOS 18.1", "Android", "Windows", "Linux" |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `status` | [`*models.NacClientLastStatusEnum`](../../doc/models/nac-client-last-status-enum.md) | Query, Optional | Client connection status used to filter results. enum: `permitted`, `session_started`, `session_stopped`, `denied` |
| `text` | `*string` | Query, Optional | partial / full MAC address, last_username, device_mac, nas_ip. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*bccddeeff`) are not supported. |
| `mType` | `*string` | Query, Optional | Client type i.e. "wireless", "wired" etc. |
| `usermacLabel` | `[]string` | Query, Optional | Labels derived from usermac entry<br><br>**Constraints**: *Unique Items Required* |
| `username` | `*string` | Query, Optional | Filter results by username |
| `vlan` | `*string` | Query, Optional | Filter results by VLAN ID |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order.<br><br>**Default**: `"wxid"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClientNacSearch](../../doc/models/response-client-nac-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

certExpiryDuration := "7d"

family := "Surveillance Camera"

hostname := "my-everest-client"

mac := "aabbccddeeff"

mfg := "Raspberry Pi Trading Ltd"

status := models.NacClientLastStatusEnum_PERMITTED

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesClientsNAC.SearchSiteNacClients(ctx, siteId, nil, nil, &certExpiryDuration, nil, nil, nil, &family, &hostname, nil, &mac, nil, nil, nil, &mfg, nil, nil, nil, nil, nil, nil, nil, nil, nil, &status, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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
  "end": 1513362753,
  "limit": 3,
  "results": [
    {
      "ap": [
        "5c5b35bf16bb",
        "d4dc090041b4"
      ],
      "auth_type": "eap-tls",
      "cert_cn": [
        "john@mycorp.net"
      ],
      "cert_issuer": [
        "/C=US/ST=CA/CN=MyCorp"
      ],
      "client_ip": [
        "10.7.51.74"
      ],
      "edr_managed": true,
      "edr_provider": "sentinelone",
      "edr_status": "sentinelone_healthy",
      "idp_id": "string",
      "idp_role": [
        "string"
      ],
      "last_ap": "string",
      "last_cert_cn": "john@mycorp.net",
      "last_cert_expiry": 1746711240,
      "last_cert_issuer": "/C=US/ST=CA/CN=MyCorp",
      "last_cert_serial": "2c63510123456789",
      "last_cert_subject": "/C=US/O=MyCorp/CN=john@mycorp.net/emailAddress=john@mycorp.net",
      "last_client_ip": "10.7.51.74",
      "last_nacrule_id": "603b62db-d839-4152-9f7f-f2578443de8d",
      "last_nacrule_name": "Wireless Cert Auth",
      "last_nas_vendor": "juniper-mist",
      "last_ssid": "string",
      "last_status": "permitted",
      "mac": "string",
      "nacrule_id": [
        "603b62db-d839-4152-9f7f-f2578443de8d"
      ],
      "nacrule_matched": true,
      "nacrule_name": [
        "Wireless Cert Auth"
      ],
      "nas_vendor": [
        "juniper-mist"
      ],
      "org_id": "31f27122-68a9-47a4-b526-8fb8a62a8acb",
      "random_mac": "true",
      "site_id": "832b1d74-9531-409b-ae37-4d7f3edbde92",
      "ssid": [
        "string"
      ],
      "timestamp": 1694689718.612,
      "type": "wireless",
      "usermac_label": [
        "non-compliant",
        "building26",
        "floor52"
      ]
    }
  ],
  "start": 1513276353,
  "total": 2
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


# Send Site Nac Client Co A

Sends CoA (Change of Authorization) command to a NAC client.

```go
SendSiteNacClientCoA(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    body *models.NacClientCoa) (
    models.ApiResponse[models.NacClientCoaResponse],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `body` | [`*models.NacClientCoa`](../../doc/models/nac-client-coa.md) | Body, Optional | Request Body |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.NacClientCoaResponse](../../doc/models/nac-client-coa-response.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

body := models.NacClientCoa{
    CoaType:              models.ToPointer(models.NacCoaTypeEnum_REAUTH),
}

apiResponse, err := sitesClientsNAC.SendSiteNacClientCoA(ctx, siteId, clientMac, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

