# Orgs Stats-Mx Edges

```go
orgsStatsMxEdges := client.OrgsStatsMxEdges()
```

## Class Name

`OrgsStatsMxEdges`

## Methods

* [Get Org Mx Edge Stats](../../doc/controllers/orgs-stats-mx-edges.md#get-org-mx-edge-stats)
* [List Org Mx Edges Stats](../../doc/controllers/orgs-stats-mx-edges.md#list-org-mx-edges-stats)


# Get Org Mx Edge Stats

Get Org MxEdge Details Stats

```go
GetOrgMxEdgeStats(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    forSite *bool) (
    models.ApiResponse[models.StatsMxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `forSite` | `*bool` | Query, Optional | **Default**: `false` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.StatsMxedge](../../doc/models/stats-mxedge.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

forSite := false

apiResponse, err := orgsStatsMxEdges.GetOrgMxEdgeStats(ctx, orgId, mxedgeId, &forSite)
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


# List Org Mx Edges Stats

Get List of Org MxEdge Stats

```go
ListOrgMxEdgesStats(
    ctx context.Context,
    orgId uuid.UUID,
    forSite *models.ForSiteEnum,
    start *string,
    end *string,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsMxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `forSite` | [`*models.ForSiteEnum`](../../doc/models/for-site-enum.md) | Query, Optional | Filter for site level mist edges |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.StatsMxedge](../../doc/models/stats-mxedge.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsStatsMxEdges.ListOrgMxEdgesStats(ctx, orgId, nil, nil, nil, &duration, &limit, &page)
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
[
  {
    "cpu_stat": {
      "cpus": {
        "property1": {
          "idle": 0,
          "interrupt": 0,
          "system": 0,
          "usage": 0,
          "user": 0
        },
        "property2": {
          "idle": 0,
          "interrupt": 0,
          "system": 0,
          "usage": 0,
          "user": 0
        }
      },
      "idle": 0,
      "interrupt": 0,
      "system": 0,
      "usage": 0,
      "user": 0
    },
    "ext_ip": "string",
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ip_stat": {
      "ip": "string",
      "ips": {
        "property1": "string",
        "property2": "string"
      }
    },
    "lag_stat": {
      "property1": {
        "active_ports": [
          "string"
        ]
      },
      "property2": {
        "active_ports": [
          "string"
        ]
      }
    },
    "last_seen": 0,
    "magic": "string",
    "memory_stats": {
      "active": 0,
      "available": 0,
      "buffers": 0,
      "cached": 0,
      "free": 0,
      "inactive": 0,
      "swap_cached": 0,
      "swap_free": 0,
      "swap_total": 0,
      "total": 0,
      "usage": 0
    },
    "model": "string",
    "mxagent_registered": true,
    "mxcluster_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "name": "string",
    "num_tunnels": 0,
    "port_stat": {
      "property1": {
        "full_duplex": true,
        "lldp_stats": {
          "mgmt_addr": "string",
          "port_desc": "string",
          "port_id": "string",
          "system_desc": "string",
          "system_name": "string"
        },
        "rx_bytes": 0,
        "rx_errors": 0,
        "rx_pkts": 0,
        "speed": 0,
        "tx_bytes": 0,
        "tx_pkts": 0,
        "up": true
      },
      "property2": {
        "full_duplex": true,
        "lldp_stats": {
          "mgmt_addr": "string",
          "port_desc": "string",
          "port_id": "string",
          "system_desc": "string",
          "system_name": "string"
        },
        "rx_bytes": 0,
        "rx_errors": 0,
        "rx_pkts": 0,
        "speed": 0,
        "tx_bytes": 0,
        "tx_pkts": 0,
        "up": true
      }
    },
    "status": "string",
    "tunterm_registered": true,
    "tunterm_stat": {
      "monitoring_failed": true
    },
    "uptime": 0,
    "version": "string",
    "virtualization_type": "string"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

