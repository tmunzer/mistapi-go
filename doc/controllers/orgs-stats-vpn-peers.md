# Orgs Stats-VPN Peers

```go
orgsStatsVPNPeers := client.OrgsStatsVPNPeers()
```

## Class Name

`OrgsStatsVPNPeers`

## Methods

* [Count Org Peer Path Stats](../../doc/controllers/orgs-stats-vpn-peers.md#count-org-peer-path-stats)
* [Search Org Peer Path Stats](../../doc/controllers/orgs-stats-vpn-peers.md#search-org-peer-path-stats)


# Count Org Peer Path Stats

Count Org Peer Path Statgs

```go
CountOrgPeerPathStats(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | `*string` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")







duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsStatsVPNPeers.CountOrgPeerPathStats(ctx, orgId, nil, nil, nil, &duration, &limit, &page)
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


# Search Org Peer Path Stats

Search Org Peer Path Stats

```go
SearchOrgPeerPathStats(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.VpnPeerStatSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |

## Response Type

[`models.VpnPeerStatSearch`](../../doc/models/vpn-peer-stat-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

limit := 100

apiResponse, err := orgsStatsVPNPeers.SearchOrgPeerPathStats(ctx, orgId, nil, nil, &duration, &limit)
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
  "end": 1619518989.4989712,
  "limit": 10,
  "results": [
    {
      "auth_algo": "hmac-sha1-96",
      "enc_algo": "aes-cbc-128",
      "ike_version": "1",
      "is_active": true,
      "last_seen": 1619518709.222,
      "mac": "020001c04668",
      "org_id": "0c160b7f-1027-4cd1-923b-744534c4b070",
      "peer_mac": "020001367edd",
      "peer_port_id": "DC_Internet",
      "peer_site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "port_id": "Lte",
      "site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "type": "svr",
      "up": true,
      "uptime": 1527128046
    },
    {
      "is_active": true,
      "last_seen": 1619518709.222,
      "latency": 91,
      "mac": "020001c04668",
      "mos": 436,
      "mtu": 1500,
      "org_id": "0c160b7f-1027-4cd1-923b-744534c4b070",
      "peer_mac": "020001367edd",
      "peer_port_id": "DC_Internet",
      "peer_router_name": "RIDCBBP1",
      "peer_site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "port_id": "Lte",
      "router_name": "RIST01544AP1",
      "site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "type": "svr",
      "up": true,
      "uptime": 1527128046
    }
  ],
  "start": 1619518689.4989705,
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

