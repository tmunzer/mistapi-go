# Orgs Clients-Wired

```go
orgsClientsWired := client.OrgsClientsWired()
```

## Class Name

`OrgsClientsWired`

## Methods

* [Count Org Wired Clients](../../doc/controllers/orgs-clients-wired.md#count-org-wired-clients)
* [Search Org Wired Clients](../../doc/controllers/orgs-clients-wired.md#search-org-wired-clients)


# Count Org Wired Clients

Count by Distinct Attributes of Clients

Note: For list of available `type` values, please refer to [List Client Events Definitions](../../doc/controllers/constants-events.md#list-client-events-definitions)

```go
CountOrgWiredClients(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgWiredClientsCountDistinctEnum,
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
| `distinct` | [`*models.OrgWiredClientsCountDistinctEnum`](../../doc/models/org-wired-clients-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `device_mac`, `mac`, `port_id`, `site_id`, `type`, `vlan`<br><br>**Default**: `"mac"` |
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

distinct := models.OrgWiredClientsCountDistinctEnum_MAC

duration := "10m"

limit := 100

apiResponse, err := orgsClientsWired.CountOrgWiredClients(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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


# Search Org Wired Clients

Search for Wired Clients in org

Note: For list of available `type` values, please refer to [List Client Events Definitions](../../doc/controllers/constants-events.md#list-client-events-definitions)

```go
SearchOrgWiredClients(
    ctx context.Context,
    orgId uuid.UUID,
    authState *string,
    authMethod *string,
    source *string,
    siteId *string,
    deviceMac *string,
    mac *string,
    portId *string,
    vlan *string,
    ip *string,
    manufacture *string,
    text *string,
    nacruleId *string,
    dhcpHostname *string,
    dhcpFqdn *string,
    dhcpClientIdentifier *string,
    dhcpVendorClassIdentifier *string,
    dhcpRequestParams *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.SearchWiredClient],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `authState` | `*string` | Query, Optional | Filter results by auth state |
| `authMethod` | `*string` | Query, Optional | Filter results by authentication method. Accepts multiple comma-separated values. |
| `source` | `*string` | Query, Optional | Filter results by client learning source. enum: `lldp`, `mac`. Accepts multiple comma-separated values. |
| `siteId` | `*string` | Query, Optional | Filter results by site identifier |
| `deviceMac` | `*string` | Query, Optional | Filter results by one or more gateway or switch MAC addresses where the client has connected. Supports comma-separated values |
| `mac` | `*string` | Query, Optional | Partial / full Client MAC address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*bccddeeff`) are not supported. Accepts multiple comma-separated values. |
| `portId` | `*string` | Query, Optional | Filter results by one or more port identifiers where the client has connected. Supports comma-separated values |
| `vlan` | `*string` | Query, Optional | Filter results by one or more VLAN IDs. Supports comma-separated values |
| `ip` | `*string` | Query, Optional | Filter results by one or more IPv4 addresses. Supports comma-separated values |
| `manufacture` | `*string` | Query, Optional | Filter results by manufacturer. Accepts multiple comma-separated values. |
| `text` | `*string` | Query, Optional | Partial / full Client MAC address, hostname or username. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*ddeeff`) are not supported |
| `nacruleId` | `*string` | Query, Optional | Filter results by NAC rule identifier |
| `dhcpHostname` | `*string` | Query, Optional | Filter results by DHCP hostname. Accepts multiple comma-separated values. |
| `dhcpFqdn` | `*string` | Query, Optional | Filter results by DHCP FQDN |
| `dhcpClientIdentifier` | `*string` | Query, Optional | Filter results by DHCP client identifier. Accepts multiple comma-separated values. |
| `dhcpVendorClassIdentifier` | `*string` | Query, Optional | DHCP Vendor Class Identifier. Accepts multiple comma-separated values. |
| `dhcpRequestParams` | `*string` | Query, Optional | Filter results by DHCP request parameters. Accepts multiple comma-separated values. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchWiredClient](../../doc/models/search-wired-client.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

authMethod := "server_reject,mac_auth"

source := "lldp,mac"

mac := "aabbccddeeff,aabbcc*"

vlan := "1"

ip := "192.168.1.1"

manufacture := "Unknown,GIFA"

dhcpHostname := "client-a,client-b"

dhcpClientIdentifier := "MAC address a8f7d982288f,MAC address 5c5b351e120c"

dhcpVendorClassIdentifier := "Mist AP34-WW,Mist BT11-WW"

dhcpRequestParams := "1 121 3 6 12 15 28 42 43 180,1 3 6 12 15 28 42 43 180"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsWired.SearchOrgWiredClients(ctx, orgId, nil, &authMethod, &source, nil, nil, &mac, nil, &vlan, &ip, &manufacture, nil, nil, &dhcpHostname, nil, &dhcpClientIdentifier, &dhcpVendorClassIdentifier, &dhcpRequestParams, &limit, nil, nil, &duration, &sort, nil)
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
  "end": 1648529800.8221116,
  "limit": 1000,
  "results": [
    {
      "auth_method": "mac_auth",
      "auth_state": "authenticated",
      "device_mac": [
        "001122334455"
      ],
      "dhcp_client_identifier": "MAC address 00155df6d500",
      "dhcp_client_options": [
        {
          "code": "DHO_DHCP_MESSAGE_TYPE(53)",
          "data": "DHCPREQUEST"
        },
        {
          "code": "DHO_DHCP_CLIENT_IDENTIFIER(61)",
          "data": "MAC address 00155df6d500"
        },
        {
          "code": "DHO_DHCP_REQUESTED_ADDRESS(50)",
          "data": "172.17.10.255"
        },
        {
          "code": "DHO_DHCP_SERVER_IDENTIFIER(54)",
          "data": "172.17.8.1"
        },
        {
          "code": "DHO_DHCP_MAX_MESSAGE_SIZE(57)",
          "data": "1280"
        },
        {
          "code": "DHO_DHCP_PARAMETER_REQUEST_LIST(55)",
          "data": " 1 3 6 12 15 28 43 180"
        },
        {
          "code": "DHO_VENDOR_CLASS_IDENTIFIER(60)",
          "data": "MSFT 5.0"
        },
        {
          "code": "DHO_HOST_NAME(12)",
          "data": "ITS-VMMT0-D1N02"
        }
      ],
      "dhcp_fqdn": "ITS-VMMT0-D1N02.mgthub.local",
      "dhcp_hostname": "ITS-VMMT0-D1N02",
      "dhcp_request_params": "1 3 6 15 31 33 43 44 46 47 119 121 249 252",
      "dhcp_vendor_class_identifier": "MSFT 5.0",
      "mac": "112233445566",
      "org_id": "c168ddee-c14c-11e5-8e81-1258369c38a9",
      "port_id": [
        "et-0/0/1"
      ],
      "site_id": "c168ddee-c14c-11e5-8e81-1258369c38a9",
      "timestamp": 1571174567.807,
      "vlan": [
        0,
        1001
      ]
    }
  ],
  "start": 1648443400.8221116,
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

