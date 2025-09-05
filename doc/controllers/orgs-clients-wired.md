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
| `distinct` | [`*models.OrgWiredClientsCountDistinctEnum`](../../doc/models/org-wired-clients-count-distinct-enum.md) | Query, Optional | **Default**: `"mac"` |
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

distinct := models.OrgWiredClientsCountDistinctEnum_MAC

duration := "10m"

limit := 100

apiResponse, err := orgsClientsWired.CountOrgWiredClients(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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


# Search Org Wired Clients

Search for Wired Clients in org

Note: For list of available `type` values, please refer to [List Client Events Definitions](../../doc/controllers/constants-events.md#list-client-events-definitions)

```go
SearchOrgWiredClients(
    ctx context.Context,
    orgId uuid.UUID,
    authState *string,
    authMethod *string,
    source *models.ClientInfoSourceEnum,
    siteId *string,
    deviceMac *string,
    mac *string,
    portId *string,
    vlan *int,
    ipAddress *string,
    manufacture *string,
    text *string,
    nacruleId *string,
    dhcpHostname *string,
    dhcpFqdn *string,
    dhcpClientIdentifier *string,
    dhcpVendorClassIdentifier *string,
    dhcpRequestParams *string,
    limit *int,
    start *int,
    end *int,
    duration *string,
    sort *string) (
    models.ApiResponse[models.SearchWiredClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `authState` | `*string` | Query, Optional | Authentication state |
| `authMethod` | `*string` | Query, Optional | Authentication method |
| `source` | [`*models.ClientInfoSourceEnum`](../../doc/models/client-info-source-enum.md) | Query, Optional | source from where the client was learned (lldp, mac) |
| `siteId` | `*string` | Query, Optional | Site ID |
| `deviceMac` | `*string` | Query, Optional | Device mac (Gateway/Switch) where the client has connected to |
| `mac` | `*string` | Query, Optional | Partial / full MAC address |
| `portId` | `*string` | Query, Optional | Port id where the client has connected to |
| `vlan` | `*int` | Query, Optional | VLAN |
| `ipAddress` | `*string` | Query, Optional | - |
| `manufacture` | `*string` | Query, Optional | Client manufacturer |
| `text` | `*string` | Query, Optional | Partial / full MAC address, hostname or username |
| `nacruleId` | `*string` | Query, Optional | nacrule_id |
| `dhcpHostname` | `*string` | Query, Optional | DHCP Hostname |
| `dhcpFqdn` | `*string` | Query, Optional | DHCP FQDN |
| `dhcpClientIdentifier` | `*string` | Query, Optional | DHCP Client Identifier |
| `dhcpVendorClassIdentifier` | `*string` | Query, Optional | DHCP Vendor Class Identifier |
| `dhcpRequestParams` | `*string` | Query, Optional | DHCP Request Parameters |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchWiredClient](../../doc/models/search-wired-client.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ipAddress := "192.168.1.1"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsWired.SearchOrgWiredClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

