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

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgTunnelCountDistinctEnum`](../../doc/models/org-tunnel-count-distinct-enum.md) | Query, Optional | - If `type`==`wxtunnel`: wxtunnel_id / ap / remote_ip / remote_port / state / mxedge_id / mxcluster_id / site_id / peer_mxedge_id; default is wxtunnel_id<br>- If `type`==`wan`: mac / site_id / node / peer_ip / peer_host/ ip / tunnel_name / protocol / auth_algo / encrypt_algo / ike_version / last_event / up<br><br>**Default**: `"wxtunnel_id"` |
| `mType` | [`*models.OrgTunnelTypeCountEnum`](../../doc/models/org-tunnel-type-count-enum.md) | Query, Optional | **Default**: `"wxtunnel"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

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

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `*string` | Query, Optional | If `type`==`wxtunnel` |
| `siteId` | `*string` | Query, Optional | - |
| `wxtunnelId` | `*string` | Query, Optional | If `type`==`wxtunnel` |
| `ap` | `*string` | Query, Optional | If `type`==`wxtunnel` |
| `mac` | `*string` | Query, Optional | If `type`==`wan` |
| `node` | `*string` | Query, Optional | If `type`==`wan` |
| `peerIp` | `*string` | Query, Optional | If `type`==`wan` |
| `peerHost` | `*string` | Query, Optional | If `type`==`wan` |
| `ip` | `*string` | Query, Optional | If `type`==`wan` |
| `tunnelName` | `*string` | Query, Optional | If `type`==`wan` |
| `protocol` | `*string` | Query, Optional | If `type`==`wan` |
| `authAlgo` | `*string` | Query, Optional | If `type`==`wan` |
| `encryptAlgo` | `*string` | Query, Optional | If `type`==`wan` |
| `ikeVersion` | `*string` | Query, Optional | If `type`==`wan` |
| `up` | `*string` | Query, Optional | If `type`==`wan` |
| `mType` | [`*models.TunnelTypeEnum`](../../doc/models/tunnel-type-enum.md) | Query, Optional | **Default**: `"wxtunnel"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"5m"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

