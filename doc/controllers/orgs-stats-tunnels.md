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

Count Mist Tunnels Stats

```go
CountOrgTunnelsStats(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgTunnelCountDistinctEnum,
    mType *models.OrgTunnelTypeCountEnum) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgTunnelCountDistinctEnum`](../../doc/models/org-tunnel-count-distinct-enum.md) | Query, Optional | - If `type`==`wxtunnel`: wxtunnel_id / ap / remote_ip / remote_port / state / mxedge_id / mxcluster_id / site_id / peer_mxedge_id; default is wxtunnel_id<br>- If `type`==`wan`: mac / site_id / node / peer_ip / peer_host/ ip / tunnel_name / protocol / auth_algo / encrypt_algo / ike_version / last_event / up |
| `mType` | [`*models.OrgTunnelTypeCountEnum`](../../doc/models/org-tunnel-type-count-enum.md) | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgTunnelCountDistinctEnum("wxtunnel_id")

mType := models.OrgTunnelTypeCountEnum("wxtunnel")

apiResponse, err := orgsStatsTunnels.CountOrgTunnelsStats(ctx, orgId, &distinct, &mType)
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

Search Org Tunnels Stats

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
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseTunnelSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `*string` | Query, Optional | if `type`==`wxtunnel` |
| `siteId` | `*string` | Query, Optional | - |
| `wxtunnelId` | `*string` | Query, Optional | if `type`==`wxtunnel` |
| `ap` | `*string` | Query, Optional | if `type`==`wxtunnel` |
| `mac` | `*string` | Query, Optional | if `type`==`wan` |
| `node` | `*string` | Query, Optional | if `type`==`wan` |
| `peerIp` | `*string` | Query, Optional | if `type`==`wan` |
| `peerHost` | `*string` | Query, Optional | if `type`==`wan` |
| `ip` | `*string` | Query, Optional | if `type`==`wan` |
| `tunnelName` | `*string` | Query, Optional | if `type`==`wan` |
| `protocol` | `*string` | Query, Optional | if `type`==`wan` |
| `authAlgo` | `*string` | Query, Optional | if `type`==`wan` |
| `encryptAlgo` | `*string` | Query, Optional | if `type`==`wan` |
| `ikeVersion` | `*string` | Query, Optional | if `type`==`wan` |
| `up` | `*string` | Query, Optional | if `type`==`wan` |
| `mType` | [`*models.TunnelTypeEnum`](../../doc/models/tunnel-type-enum.md) | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseTunnelSearch`](../../doc/models/response-tunnel-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")































mType := models.TunnelTypeEnum("wxtunnel")

limit := 100





duration := "10m"

apiResponse, err := orgsStatsTunnels.SearchOrgTunnelsStats(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &mType, &limit, nil, nil, &duration)
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

