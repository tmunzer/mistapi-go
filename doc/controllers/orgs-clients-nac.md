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


# Count Org Nac Client Events

Count by Distinct Attributes of NAC Client-Events

```go
CountOrgNacClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgNacClientEventsCountDistinctEnum,
    mType *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgNacClientEventsCountDistinctEnum`](../../doc/models/org-nac-client-events-count-distinct-enum.md) | Query, Optional | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions) |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")









duration := "10m"

limit := 100

apiResponse, err := orgsClientsNAC.CountOrgNacClientEvents(ctx, orgId, nil, nil, nil, nil, &duration, &limit)
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


# Count Org Nac Clients

Count by Distinct Attributes of NAC Clients

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
    timestamp *float64,
    siteId *string,
    lastAp *string,
    mac *string,
    lastStatus *string,
    mType *string,
    mdmComplianceStatus *string,
    mdmProvider *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgNacClientsCountDistinctEnum`](../../doc/models/org-nac-clients-count-distinct-enum.md) | Query, Optional | NAC Policy Rule ID, if matched<br><br>**Default**: `"type"` |
| `lastNacruleId` | `*string` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `lastVlanId` | `*string` | Query, Optional | Vlan ID |
| `lastNasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `idpId` | `*string` | Query, Optional | SSO ID, if present and used |
| `lastSsid` | `*string` | Query, Optional | SSID |
| `lastUsername` | `*string` | Query, Optional | Username presented by the client |
| `timestamp` | `*float64` | Query, Optional | Start time, in epoch |
| `siteId` | `*string` | Query, Optional | Site id if assigned, null if not assigned |
| `lastAp` | `*string` | Query, Optional | AP MAC connected to by client |
| `mac` | `*string` | Query, Optional | MAC address |
| `lastStatus` | `*string` | Query, Optional | Connection status of client i.e "permitted", "denied, "session_ended" |
| `mType` | `*string` | Query, Optional | Client type i.e. "wireless", "wired" etc. |
| `mdmComplianceStatus` | `*string` | Query, Optional | MDM compliance of client i.e "compliant", "not compliant" |
| `mdmProvider` | `*string` | Query, Optional | MDM provider of client’s organization eg "intune", "jamf" |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgNacClientsCountDistinctEnum_ENUMTYPE





































duration := "10m"

limit := 100

apiResponse, err := orgsClientsNAC.CountOrgNacClients(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# Search Org Nac Client Events

Search NAC Client Events

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
    timestamp *float64,
    usermacLabel *string,
    text *string,
    nasIp *string,
    sort *string,
    ingressVlan *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsNacClientSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions) |
| `nacruleId` | `*uuid.UUID` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `dryrunNacruleId` | `*string` | Query, Optional | NAC Policy Dry Run Rule ID, if present and matched |
| `dryrunNacruleMatched` | `*bool` | Query, Optional | True - if dryrun rule present and matched with priority, False - if not matched or not present |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `vlan` | `*int` | Query, Optional | Vlan name or ID assigned to the client |
| `nasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `bssid` | `*string` | Query, Optional | BSSID |
| `idpId` | `*uuid.UUID` | Query, Optional | SSO ID, if present and used |
| `idpRole` | `*string` | Query, Optional | IDP returned roles/groups for the user |
| `idpUsername` | `*string` | Query, Optional | Username presented to the Identity Provider |
| `respAttrs` | `[]string` | Query, Optional | Radius attributes returned by NAC to NAS derive<br><br>**Constraints**: *Unique Items Required* |
| `ssid` | `*string` | Query, Optional | SSID |
| `username` | `*string` | Query, Optional | Username presented by the client |
| `siteId` | `*string` | Query, Optional | Site id |
| `ap` | `*string` | Query, Optional | AP MAC |
| `randomMac` | `*bool` | Query, Optional | AP random macMAC |
| `mac` | `*string` | Query, Optional | MAC address |
| `timestamp` | `*float64` | Query, Optional | Start time, in epoch |
| `usermacLabel` | `*string` | Query, Optional | Labels derived from usermac entry |
| `text` | `*string` | Query, Optional | Partial / full MAC address, username, device_mac or ap |
| `nasIp` | `*string` | Query, Optional | IP address of NAS device |
| `sort` | `*string` | Query, Optional | Sort options, ‘-‘ prefix represents DESC order, default is wcid in ASC order |
| `ingressVlan` | `*string` | Query, Optional | Vendor specific Vlan ID in radius requests |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsNacClientSearch](../../doc/models/response-events-nac-client-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

























respAttrs := []string{
    "Tunnel-Type=VLAN",
    "Tunnel-Medium-Type=IEEE-802",
    "Tunnel-Private-Group-Id=750",
    "User-Name=anonymous",
}





























duration := "10m"

limit := 100

apiResponse, err := orgsClientsNAC.SearchOrgNacClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, respAttrs, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
      "random_mac": false,
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Search Org Nac Clients

Search Org NAC Clients

```go
SearchOrgNacClients(
    ctx context.Context,
    orgId uuid.UUID,
    nacruleId *string,
    nacruleMatched *bool,
    authType *string,
    vlan *string,
    nasVendor *string,
    idpId *string,
    ssid *string,
    username *string,
    timestamp *float64,
    siteId *string,
    ap *string,
    mac *string,
    mdmManaged *bool,
    status *string,
    mType *string,
    mdmCompliance *string,
    mdmProvider *string,
    sort *string,
    usermacLabel []string,
    ingressVlan *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseClientNacSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nacruleId` | `*string` | Query, Optional | NAC Policy Rule ID, if matched |
| `nacruleMatched` | `*bool` | Query, Optional | NAC Policy Rule Matched |
| `authType` | `*string` | Query, Optional | Authentication type, e.g. "eap-tls", "eap-peap", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `vlan` | `*string` | Query, Optional | Vlan name or ID assigned to the client |
| `nasVendor` | `*string` | Query, Optional | Vendor of NAS device |
| `idpId` | `*string` | Query, Optional | SSO ID, if present and used |
| `ssid` | `*string` | Query, Optional | SSID |
| `username` | `*string` | Query, Optional | Username presented by the client |
| `timestamp` | `*float64` | Query, Optional | Start time, in epoch |
| `siteId` | `*string` | Query, Optional | Site id if assigned, null if not assigned |
| `ap` | `*string` | Query, Optional | AP MAC connected to by client |
| `mac` | `*string` | Query, Optional | MAC address |
| `mdmManaged` | `*bool` | Query, Optional | Filters NAC clients that are managed by MDM providers |
| `status` | `*string` | Query, Optional | Connection status of client i.e "permitted", "denied, "session_ended" |
| `mType` | `*string` | Query, Optional | Client type i.e. "wireless", "wired" etc. |
| `mdmCompliance` | `*string` | Query, Optional | MDM compliance of client i.e "compliant", "not compliant" |
| `mdmProvider` | `*string` | Query, Optional | MDM provider of client’s organization eg "intune", "jamf" |
| `sort` | `*string` | Query, Optional | Sort options, ‘-‘ prefix represents DESC order, default is wcid in ASC order |
| `usermacLabel` | `[]string` | Query, Optional | Labels derived from usermac entry<br><br>**Constraints**: *Unique Items Required* |
| `ingressVlan` | `*string` | Query, Optional | Vendor specific Vlan ID in radius requests |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClientNacSearch](../../doc/models/response-client-nac-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













































duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsClientsNAC.SearchOrgNacClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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
      "random_mac": true,
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

