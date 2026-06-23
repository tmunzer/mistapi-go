# Orgs Stats-Tunnels

```go
orgsStatsTunnels := client.OrgsStatsTunnels()
```

## Class Name

`OrgsStatsTunnels`

## Methods

* [Count Org Tunnels Stats](../../doc/controllers/orgs-stats-tunnels.md#count-org-tunnels-stats)
* [Search Org Tunnels Stats](../../doc/controllers/orgs-stats-tunnels.md#search-org-tunnels-stats)


# Count Org Tunnels Stats

Count by Distinct Attributes of Mist Tunnels Stats

```go
CountOrgTunnelsStats(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgTunnelCountDistinctEnum,
    mType *models.OrgTunnelTypeCountEnum,
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
| `distinct` | [`*models.OrgTunnelCountDistinctEnum`](../../doc/models/org-tunnel-count-distinct-enum.md) | Query, Optional | Field used to group tunnel statistics count results. enum: `ap`, `auth_algo`, `encrypt_algo`, `ike_version`, `ip`, `last_event`, `mac`, `mxcluster_id`, `mxedge_id`, `node`, `peer_host`, `peer_ip`, `peer_mxedge_id`, `protocol`, `remote_ip`, `remote_port`, `site_id`, `state`, `tunnel_name`, `up`, `wxtunnel_id`<br><br>**Default**: `"wxtunnel_id"` |
| `mType` | [`*models.OrgTunnelTypeCountEnum`](../../doc/models/org-tunnel-type-count-enum.md) | Query, Optional | Filter results by type. enum: `wan`, `wxtunnel`<br><br>**Default**: `"wxtunnel"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgTunnelCountDistinctEnum_WXTUNNELID

mType := models.OrgTunnelTypeCountEnum_WXTUNNEL

limit := 100

apiResponse, err := orgsStatsTunnels.CountOrgTunnelsStats(ctx, orgId, &distinct, &mType, &limit)
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


# Search Org Tunnels Stats

By default the endpoint returns only `wxtunnel` type stats, to get `wan` type stats
you need to specify `type=wan` in the query parameters.

Tunnel types:

- `wxtunnel` (default) - A WxLan Tunnel (WxTunnel) are used to create a secure connection between Juniper Mist Access Points and third-party VPN concentrators using protocols such as L2TPv3 or dmvpn.
- `wan` - A WAN Tunnel is a secure connection between two Gateways, typically used for site-to-site or mesh connectivity. It can be configured with various protocols and encryption methods.

If `type` is not specified or `type`==`wxtunnel`, the following parameters are supported:

- `mxcluster_id` - the MX cluster ID
- `site_id` - the site ID
- `wxtunnel_id` - the WX tunnel ID
- `ap` - the AP MAC address

If `type`==`wan`, the following parameters are supported:

- `mac` - the MAC address of the WAN device
- `node` - the node ID
- `peer_ip` - the peer IP address
- `peer_host` - the peer host name
- `ip` - the IP address of the WAN device
- `tunnel_name` - the name of the tunnel
- `protocol` - the protocol used for the tunnel
- `auth_algo` - the authentication algorithm used for the tunnel
- `encrypt_algo` - the encryption algorithm used for the tunnel
- `ike_version` - the IKE version used for the tunnel
- `up` - the status of the tunnel (up or down)

```go
SearchOrgTunnelsStats(
    ctx context.Context,
    orgId uuid.UUID,
    mxclusterId *string,
    siteId *string,
    wxtunnelId *string,
    ap *string,
    mac *string,
    node *string,
    peerIp *string,
    peerHost *string,
    ip *string,
    tunnelName *string,
    protocol *string,
    authAlgo *string,
    encryptAlgo *string,
    ikeVersion *string,
    up *string,
    mType *models.TunnelTypeEnum,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseTunnelSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `*string` | Query, Optional | Filter results by mxcluster id when `type`==`wxtunnel` |
| `siteId` | `*string` | Query, Optional | Filter results by site identifier |
| `wxtunnelId` | `*string` | Query, Optional | Filter results by wxtunnel id when `type`==`wxtunnel` |
| `ap` | `*string` | Query, Optional | Filter results by AP MAC address when `type`==`wxtunnel` |
| `mac` | `*string` | Query, Optional | Filter results by MAC address when `type`==`wan` |
| `node` | `*string` | Query, Optional | Filter results by node when `type`==`wan` |
| `peerIp` | `*string` | Query, Optional | Filter results by peer ip when `type`==`wan` |
| `peerHost` | `*string` | Query, Optional | Filter results by peer host when `type`==`wan` |
| `ip` | `*string` | Query, Optional | Filter results by IP address when `type`==`wan` |
| `tunnelName` | `*string` | Query, Optional | Filter results by tunnel name when `type`==`wan` |
| `protocol` | `*string` | Query, Optional | Filter results by protocol when `type`==`wan` |
| `authAlgo` | `*string` | Query, Optional | Filter results by auth algo when `type`==`wan` |
| `encryptAlgo` | `*string` | Query, Optional | Filter results by encrypt algo when `type`==`wan` |
| `ikeVersion` | `*string` | Query, Optional | Filter results by ike version when `type`==`wan` |
| `up` | `*string` | Query, Optional | Filter results by up when `type`==`wan` |
| `mType` | [`*models.TunnelTypeEnum`](../../doc/models/tunnel-type-enum.md) | Query, Optional | Filter results by type. enum: `wan`, `wxtunnel`<br><br>**Default**: `"wxtunnel"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"5m"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseTunnelSearch](../../doc/models/response-tunnel-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.TunnelTypeEnum_WXTUNNEL

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsStatsTunnels.SearchOrgTunnelsStats(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &mType, &limit, nil, nil, &duration, &sort, nil)
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
  "end": 1531862583,
  "limit": 10,
  "results": [
    {
      "auth_algo": "hmac-md5-96",
      "encrypt_algo": "aes-256-cbc",
      "ike_version": "2",
      "ip": "192.168.233.0",
      "last_event": "down reason",
      "mac": "020001ae9dd5",
      "node": "node0",
      "org_id": "78c11da8-f984-4425-bedb-a7ddd7d0f6da",
      "peer_host": "sunnyvale1-vpn.zscalerbeta.net",
      "peer_ip": "10.224.8.16",
      "protocol": "ipsec",
      "rx_bytes": 150,
      "rx_pkts": 75,
      "site_id": "e83e7928-eda1-4e93-82db-df3dd42ab726",
      "tunnel_name": "Device-ipsec-1",
      "tx_bytes": 100,
      "tx_pkts": 50,
      "up": true,
      "uptime": 10,
      "wan_name": "wan"
    }
  ],
  "start": 1531776183,
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

