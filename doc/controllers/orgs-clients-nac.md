# Orgs Clients-NAC

```go
orgsClientsNAC := client.OrgsClientsNAC()
```

## Class Name

`OrgsClientsNAC`

## Methods

* [Count Org Nac Client Events](../../doc/controllers/orgs-clients-nac.md#count-org-nac-client-events)
* [Count Org Nac Clients](../../doc/controllers/orgs-clients-nac.md#count-org-nac-clients)
* [Search Org Nac Client Events](../../doc/controllers/orgs-clients-nac.md#search-org-nac-client-events)
* [Search Org Nac Clients](../../doc/controllers/orgs-clients-nac.md#search-org-nac-clients)
* [Send Org Nac Client Co A](../../doc/controllers/orgs-clients-nac.md#send-org-nac-client-co-a)


# Count Org Nac Client Events

Count NAC client events across the organization, optionally grouped by `distinct` and filtered by event type and time range.

```go
CountOrgNacClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgNacClientEventsCountDistinctEnum,
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
| `distinct` | [`*models.OrgNacClientEventsCountDistinctEnum`](../../doc/models/org-nac-client-events-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `ap`, `auth_type`, `dryrun_nacrule_id`, `mac`, `nacrule_id`, `nas_vendor`, `ssid`, `type`, `username`, `vlan` |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions). Accepts multiple comma-separated values. |
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

mType := "NAC_CLIENT_PERMIT,NAC_SESSION_STARTED"

duration := "10m"

limit := 100

apiResponse, err := orgsClientsNAC.CountOrgNacClientEvents(ctx, orgId, nil, &mType, nil, nil, &duration, &limit)
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


# Count Org Nac Clients

Count NAC clients across the organization, optionally grouped by `distinct` and filtered by authentication, identity, endpoint, network, site, and time attributes.

```go
CountOrgNacClients(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgNacClientsCountDistinctEnum,
    lastNacruleId *string,
    nacruleMatched *bool,
    authType *string,
    lastVlanId *string,
    lastNasVendor *string,
    idpId *string,
    lastSsid *string,
    lastUsername *string,
    siteId *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgNacClientsCountDistinctEnum`](../../doc/models/org-nac-clients-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `ap`, `auth_type`, `device_mac`, `edr_managed`, `edr_provider`, `edr_status`, `family`, `hostname`, `idp_id`, `mfg`, `mdm_compliance`, `mdm_managed`, `mdm_provider`, `model`, `mxedge_id`, `nacrule_matched`, `nacrule_name`, `nacrule_id`, `nas_ip`, `nas_vendor`, `os`, `site_id`, `ssid`, `status`, `type`, `usermac_label`, `username`, `vlan`<br><br>**Default**: `"type"` |
| `lastNacruleId` | `*string` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `lastVlanId` | `*string` | Query, Optional | Filter results by last VLAN ID |
| `lastNasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `idpId` | `*string` | Query, Optional | SSO ID, if present and used |
| `lastSsid` | `*string` | Query, Optional | Filter results by last SSID |
| `lastUsername` | `*string` | Query, Optional | Username presented by the client |
| `siteId` | `*string` | Query, Optional | Filter results by site identifier |
| `lastAp` | `*string` | Query, Optional | AP MAC connected to by client |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `lastStatus` | `*string` | Query, Optional | Connection status of client i.e "permitted", "denied, "session_ended" |
| `mType` | `*string` | Query, Optional | Client type i.e. "wireless", "wired" etc. Accepts multiple comma-separated values. |
| `mdmComplianceStatus` | `*string` | Query, Optional | MDM compliance of client i.e "compliant", "not compliant" |
| `mdmProvider` | `*string` | Query, Optional | MDM provider of client’s organization eg "intune", "jamf" |
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

distinct := models.OrgNacClientsCountDistinctEnum_ENUMTYPE

mType := "wired,wireless"

duration := "10m"

limit := 100

apiResponse, err := orgsClientsNAC.CountOrgNacClients(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &mType, nil, nil, nil, nil, &duration, &limit)
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


# Search Org Nac Client Events

Search NAC client authentication event records across the organization with filters for authentication, NAC rule, identity provider, RADIUS, network, endpoint, site, and time attributes.

```go
SearchOrgNacClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
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
    siteId *string,
    ap *string,
    randomMac *bool,
    mac *string,
    usermacLabel *string,
    text *string,
    nasIp *string,
    ingressVlan *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions). Accepts multiple comma-separated values. |
| `nacruleId` | `*uuid.UUID` | Query, Optional | NAC Policy Rule ID, if matched. Accepts multiple comma-separated values. |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `dryrunNacruleId` | `*string` | Query, Optional | NAC Policy Dry Run Rule ID, if present and matched |
| `dryrunNacruleMatched` | `*bool` | Query, Optional | True - if dryrun rule present and matched with priority, False - if not matched or not present |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth". Accepts multiple comma-separated values. |
| `vlan` | `*int` | Query, Optional | Filter results by VLAN ID. Accepts multiple comma-separated integer values. |
| `nasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `bssid` | `*string` | Query, Optional | Filter results by BSSID |
| `idpId` | `*uuid.UUID` | Query, Optional | SSO ID, if present and used |
| `idpRole` | `*string` | Query, Optional | IDP returned roles/groups for the user |
| `idpUsername` | `*string` | Query, Optional | Username presented to the Identity Provider |
| `respAttrs` | `[]string` | Query, Optional | RADIUS attributes returned by NAC to NAS derive<br><br>**Constraints**: *Unique Items Required* |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `username` | `*string` | Query, Optional | Filter results by username. Accepts multiple comma-separated values. |
| `siteId` | `*string` | Query, Optional | Filter results by one site identifier. Use a single value; comma-separated values are not supported |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address |
| `randomMac` | `*bool` | Query, Optional | Filter results by whether the client is using a randomized MAC address. Accepts multiple comma-separated boolean values. |
| `mac` | `*string` | Query, Optional | Filter results by one MAC address. Use a single value; comma-separated values are not supported |
| `usermacLabel` | `*string` | Query, Optional | Labels derived from usermac entry |
| `text` | `*string` | Query, Optional | Partial / full MAC address, username, device_mac or ap |
| `nasIp` | `*string` | Query, Optional | IP address of NAS device. Accepts multiple comma-separated values. |
| `ingressVlan` | `*string` | Query, Optional | Vendor specific VLAN ID in RADIUS requests |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order.<br><br>**Default**: `"wxid"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: NAC Client Events

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsNacClientSearch](../../doc/models/response-events-nac-client-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := "NAC_CLIENT_PERMIT,NAC_SESSION_STARTED"

authType := "mab,eap-tls"

respAttrs := []string{
    "Tunnel-Type=VLAN",
    "Tunnel-Medium-Type=IEEE-802",
    "Tunnel-Private-Group-Id=750",
    "User-Name=anonymous",
}

username := "john.doe,jane.doe"

nasIp := "192.0.2.10,192.0.2.11"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsNAC.SearchOrgNacClientEvents(ctx, orgId, &mType, nil, nil, nil, nil, &authType, nil, nil, nil, nil, nil, nil, respAttrs, nil, &username, nil, nil, nil, nil, nil, nil, &nasIp, nil, &limit, nil, nil, &duration, &sort, nil)
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


# Search Org Nac Clients

Search NAC client records across the organization with filters for authentication, endpoint posture, identity, network, NAC rule, site, and time attributes.

```go
SearchOrgNacClients(
    ctx context.Context,
    orgId uuid.UUID,
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
    siteId *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ap` | `*string` | Query, Optional | MAC address of the AP the client is/was connected to |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth". Accepts multiple comma-separated values. |
| `certExpiryDuration` | `*string` | Query, Optional | Filter by certificate expiry within a specific duration from now (e.g., "7d" for 7 days, "1m" for 1 month). Accepts multiple comma-separated values. |
| `edrManaged` | `*bool` | Query, Optional | Filters NAC clients that are integrated with EDR providers |
| `edrProvider` | [`*models.EdrProviderEnum`](../../doc/models/edr-provider-enum.md) | Query, Optional | EDR provider used to filter NAC clients. enum: `crowdstrike`, `sentinelone` |
| `edrStatus` | [`*models.EdrStatusEnum`](../../doc/models/edr-status-enum.md) | Query, Optional | EDR status used to filter NAC clients. enum: `sentinelone_healthy`, `sentinelone_infected`, `crowdstrike_low`, `crowdstrike_medium`, `crowdstrike_high`, `crowdstrike_critical`, `crowdstrike_informational` |
| `family` | `*string` | Query, Optional | Partial / full Client family (e.g. "Phone/Tablet/Wearable", "Access Point"). Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `Surveillance*` and `*urveillance*` match `Surveillance Camera`). Suffix-only wildcards (e.g. `*Camera`) are not supported. Accepts multiple comma-separated values. |
| `hostname` | `*string` | Query, Optional | Partial / full Client hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `everest*` and `*rest*` match `my-everest-client`). Suffix-only wildcards (e.g. `*everest`) are not supported. Accepts multiple comma-separated values. |
| `idpId` | `*string` | Query, Optional | SSO ID, if present and used |
| `mac` | `*string` | Query, Optional | Partial / full Client MAC address. Use a single value; comma-separated values are not supported. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*bccddeeff`) are not supported |
| `mdmCompliance` | `*string` | Query, Optional | MDM compliance of client i.e "compliant", "not compliant" |
| `mdmProvider` | `*string` | Query, Optional | MDM provider of client’s organization eg "intune", "jamf" |
| `mdmManaged` | `*bool` | Query, Optional | Filters NAC clients that are managed by MDM providers |
| `mfg` | `*string` | Query, Optional | Partial / full Client manufacturer (e.g. "apple", "cisco", "juniper"). Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `Raspberry Pi*` and `*Pi*` match `Raspberry Pi Trading Ltd`). Suffix-only wildcards (e.g. `*Ltd`) are not supported. Accepts multiple comma-separated values. |
| `model` | `*string` | Query, Optional | Client model, e.g. "iPhone 12", "MX100" |
| `nacruleName` | `*string` | Query, Optional | NAC Policy Rule Name matched |
| `nacruleId` | `*string` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `nasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `nasIp` | `*string` | Query, Optional | IP address of NAS device. Accepts multiple comma-separated values. |
| `ingressVlan` | `*string` | Query, Optional | Vendor specific VLAN ID in RADIUS requests |
| `os` | `*string` | Query, Optional | Client OS, e.g. "iOS 18.1", "Android", "Windows", "Linux" |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `status` | [`*models.NacClientLastStatusEnum`](../../doc/models/nac-client-last-status-enum.md) | Query, Optional | Client connection status used to filter results. enum: `permitted`, `session_started`, `session_stopped`, `denied` |
| `text` | `*string` | Query, Optional | partial / full MAC address, last_username, device_mac, nas_ip or last_ap |
| `mType` | `*string` | Query, Optional | Client type i.e. "wireless", "wired" etc. Accepts multiple comma-separated values. |
| `usermacLabel` | `[]string` | Query, Optional | Labels derived from usermac entry<br><br>**Constraints**: *Unique Items Required* |
| `username` | `*string` | Query, Optional | Filter results by username |
| `vlan` | `*string` | Query, Optional | Filter results by VLAN ID |
| `siteId` | `*string` | Query, Optional | Filter results by one site identifier. Use a single value; comma-separated values are not supported |
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

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

authType := "mab,eap-tls"

certExpiryDuration := "7d,1m"

family := "Surveillance Camera,Surveillance*"

hostname := "my-everest-client,my-everest*"

mac := "aabbccddeeff"

mfg := "Raspberry Pi Trading Ltd,Raspberry Pi*"

nasIp := "192.0.2.10,192.0.2.11"

status := models.NacClientLastStatusEnum_PERMITTED

mType := "wired,wireless"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsNAC.SearchOrgNacClients(ctx, orgId, nil, &authType, &certExpiryDuration, nil, nil, nil, &family, &hostname, nil, &mac, nil, nil, nil, &mfg, nil, nil, nil, nil, nil, &nasIp, nil, nil, nil, &status, nil, &mType, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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


# Send Org Nac Client Co A

Sends CoA (Change of Authorization) command to a NAC client.

```go
SendOrgNacClientCoA(
    ctx context.Context,
    orgId uuid.UUID,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `body` | [`*models.NacClientCoa`](../../doc/models/nac-client-coa.md) | Body, Optional | Request Body |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.NacClientCoaResponse](../../doc/models/nac-client-coa-response.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

body := models.NacClientCoa{
    CoaType:              models.ToPointer(models.NacCoaTypeEnum_REAUTH),
}

apiResponse, err := orgsClientsNAC.SendOrgNacClientCoA(ctx, orgId, clientMac, &body)
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

