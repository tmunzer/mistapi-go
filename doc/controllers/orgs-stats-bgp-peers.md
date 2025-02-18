# Orgs Stats-BGP Peers

```go
orgsStatsBGPPeers := client.OrgsStatsBGPPeers()
```

## Class Name

`OrgsStatsBGPPeers`

## Methods

* [Count Org Bgp Stats](../../doc/controllers/orgs-stats-bgp-peers.md#count-org-bgp-stats)
* [Search Org Bgp Stats](../../doc/controllers/orgs-stats-bgp-peers.md#search-org-bgp-stats)


# Count Org Bgp Stats

Count Org BGP Stats

```go
CountOrgBgpStats(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsStatsBGPPeers.CountOrgBgpStats(ctx, orgId)
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


# Search Org Bgp Stats

Search Org BGP Stats

```go
SearchOrgBgpStats(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseSearchBgps],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseSearchBgps`](../../doc/models/response-search-bgps.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsStatsBGPPeers.SearchOrgBgpStats(ctx, orgId)
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
  "end": 0,
  "limit": 0,
  "results": [
    {
      "evpn_overlay": true,
      "for_overlay": true,
      "local_as": 65000,
      "mac": "020001c04668",
      "neighbor": "15.8.3.5",
      "neighbor_as": 65000,
      "neighbor_mac": "020001c04600",
      "node": "node0",
      "org_id": "0c160b7f-1027-4cd1-923b-744534c4b070",
      "rx_pkts": 63366,
      "rx_routes": 60,
      "site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "state": "established",
      "timestamp": 1666251056.07,
      "tx_pkts": 1735,
      "tx_routes": 60,
      "up": true,
      "uptime": 31355,
      "vrf_name": "default"
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

