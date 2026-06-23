# Sites Clients-Wired

```go
sitesClientsWired := client.SitesClientsWired()
```

## Class Name

`SitesClientsWired`

## Methods

* [Count Site Wired Clients](../../doc/controllers/sites-clients-wired.md#count-site-wired-clients)
* [Search Site Wired Clients](../../doc/controllers/sites-clients-wired.md#search-site-wired-clients)


# Count Site Wired Clients

Count wired clients for a site, optionally grouped by the `distinct` field and filtered by MAC address, switch port, VLAN, and time range. Use [Count Org Wired Clients](../../doc/controllers/orgs-clients-wired.md#count-org-wired-clients) to count wired clients across the organization.

```go
CountSiteWiredClients(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteWiredClientsCountDistinctEnum,
    mac *string,
    deviceMac *string,
    portId *string,
    vlan *string,
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
| `distinct` | [`*models.SiteWiredClientsCountDistinctEnum`](../../doc/models/site-wired-clients-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `mac`, `port_id`, `vlan`<br><br>**Default**: `"mac"` |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `deviceMac` | `*string` | Query, Optional | Filter results by device MAC address |
| `portId` | `*string` | Query, Optional | Filter results by port identifier |
| `vlan` | `*string` | Query, Optional | Filter results by VLAN ID |
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

distinct := models.SiteWiredClientsCountDistinctEnum_MAC

mac := "0123456789ab"

deviceMac := "0123456789ab"

portId := "ge-1/1/1"

vlan := "10"

duration := "10m"

limit := 100

apiResponse, err := sitesClientsWired.CountSiteWiredClients(ctx, siteId, &distinct, &mac, &deviceMac, &portId, &vlan, nil, nil, &duration, &limit)
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


# Search Site Wired Clients

Search wired clients for a site with filters for device MAC address, client MAC address, IP address, switch port, VLAN, manufacturer, DHCP attributes, NAC rule, and time range. Use [Search Org Wired Clients](../../doc/controllers/orgs-clients-wired.md#search-org-wired-clients) to search wired clients across the organization.

```go
SearchSiteWiredClients(
    ctx context.Context,
    siteId uuid.UUID,
    deviceMac *string,
    mac *string,
    ip *string,
    portId *string,
    source *models.ClientInfoSourceEnum,
    vlan *string,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `*string` | Query, Optional | Filter results by device MAC address |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `ip` | `*string` | Query, Optional | Filter results by IP address |
| `portId` | `*string` | Query, Optional | Filter results by port identifier |
| `source` | [`*models.ClientInfoSourceEnum`](../../doc/models/client-info-source-enum.md) | Query, Optional | Filter results by client learning source. enum: `lldp`, `mac` |
| `vlan` | `*string` | Query, Optional | Filter results by VLAN ID |
| `manufacture` | `*string` | Query, Optional | Filter results by manufacturer |
| `text` | `*string` | Query, Optional | Single entry of hostname/mac |
| `nacruleId` | `*string` | Query, Optional | Filter results by NAC rule identifier |
| `dhcpHostname` | `*string` | Query, Optional | Filter results by DHCP hostname |
| `dhcpFqdn` | `*string` | Query, Optional | Filter results by DHCP FQDN |
| `dhcpClientIdentifier` | `*string` | Query, Optional | Filter results by DHCP client identifier |
| `dhcpVendorClassIdentifier` | `*string` | Query, Optional | DHCP Vendor Class Identifier |
| `dhcpRequestParams` | `*string` | Query, Optional | Filter results by DHCP request parameters |
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

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0123456789ab"

mac := "0123456789ab"

ip := "10.3.5.12"

portId := "ge-1/1/1"

vlan := "10"

manufacture := "Juniper-Mist"

text := "client-hostname"

nacruleId := "8bfc2490-d726-3587-038d-cb2e71bd2330"

dhcpHostname := "client-hostname"

dhcpFqdn := "client.example.com"

dhcpClientIdentifier := "01:23:45:67:89:ab"

dhcpVendorClassIdentifier := "Juniper-Mist-AP,Juniper-Mist-Client"

dhcpRequestParams := "hostname,domain-name,domain-name-servers"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesClientsWired.SearchSiteWiredClients(ctx, siteId, &deviceMac, &mac, &ip, &portId, nil, &vlan, &manufacture, &text, &nacruleId, &dhcpHostname, &dhcpFqdn, &dhcpClientIdentifier, &dhcpVendorClassIdentifier, &dhcpRequestParams, &limit, nil, nil, &duration, &sort, nil)
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

