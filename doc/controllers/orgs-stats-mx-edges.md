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
    mxedgeId uuid.UUID) (
    models.ApiResponse[models.MxedgeStats],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.MxedgeStats`](../../doc/models/mxedge-stats.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsStatsMxEdges.GetOrgMxEdgeStats(ctx, orgId, mxedgeId)
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
  "cpu_stat": {
    "cpus": {
      "cpu0": {
        "idle": 89,
        "interrupt": 0,
        "system": 8,
        "usage": 10,
        "user": 1
      },
      "cpu1": {
        "idle": 81,
        "interrupt": 0,
        "system": 4,
        "usage": 18,
        "user": 13
      },
      "cpu2": {
        "idle": 81,
        "interrupt": 0,
        "system": 4,
        "usage": 18,
        "user": 13
      },
      "cpu3": {
        "idle": 2,
        "interrupt": 0,
        "system": 50,
        "usage": 97,
        "user": 46
      }
    },
    "idle": 62,
    "interrupt": 0,
    "system": 17,
    "usage": 37,
    "user": 19
  },
  "created_time": 1632684398,
  "for_site": false,
  "id": "00000000-0000-0000-1000-020000a80cb4",
  "ip_stat": {
    "ip": "192.168.1.244",
    "ips": {
      "ens18": "192.168.1.244/24,fe80::104c:ffff:fee0:caf8/64"
    },
    "macs": {
      "ens18": "e4434b217044"
    }
  },
  "lag_stat": {
    "lacp0": {
      "active_ports": [
        "port0",
        "port1"
      ]
    }
  },
  "last_seen": 1633721215,
  "mac": "020000a80cb4",
  "memory_stat": {
    "active": 394936320,
    "available": 4699291648,
    "buffers": 107646976,
    "cached": 478060544,
    "free": 4330659840,
    "inactive": 211980288,
    "swap_cached": 0,
    "swap_free": 1022357504,
    "swap_total": 1022357504,
    "total": 8365957120,
    "usage": 48
  },
  "model": "ME-VM",
  "modified_time": 1633643629,
  "mxagent_registered": true,
  "mxcluster_id": "678bc339-7635-4556-bbc0-e77ad493ef8b",
  "name": "me-vm-1",
  "num_tunnels": 0,
  "oob_ip_config": {
    "dns": [
      "8.8.8.8",
      "1.1.1.1"
    ],
    "gateway": "10.0.0.1",
    "ip": "10.0.0.10",
    "netmask": "255.255.255.0",
    "type": "static"
  },
  "org_id": "11b08247-b1ee-4152-9b25-312b323ce480",
  "port_stat": {
    "port0": {
      "full_duplex": true,
      "mac": "9e294e49091d",
      "rx_bytes": 646898375700,
      "rx_errors": 0,
      "rx_pkts": 8784449574,
      "speed": 10000,
      "state": "forwarding",
      "tx_bytes": 647200748038,
      "tx_errors": 0,
      "tx_pkts": 8788647466,
      "up": true
    },
    "port1": {
      "full_duplex": true,
      "mac": "a270fe53437e",
      "rx_bytes": 647200437652,
      "rx_errors": 0,
      "rx_pkts": 8788644886,
      "speed": 10000,
      "state": "forwarding",
      "tx_bytes": 646898681650,
      "tx_errors": 0,
      "tx_pkts": 8784452092,
      "up": true
    }
  },
  "sensor_stat": {},
  "serial": "string",
  "service_stat": {
    "mxagent": {
      "ext_ip": "99.0.86.164",
      "last_seen": 1633721215,
      "package_state": "Installed",
      "package_version": "3.1.1037-1",
      "running_state": "Running",
      "uptime": 21240
    },
    "tunterm": {
      "ext_ip": "99.0.86.164",
      "last_seen": 1633721203,
      "package_state": "Installed",
      "package_version": "0.1.2449+deb10",
      "running_state": "Running",
      "uptime": 76261
    }
  },
  "services": [
    "tunterm"
  ],
  "site_id": "00000000-0000-0000-0000-000000000000",
  "status": "connected",
  "tunterm_ip_config": {
    "gateway": "192.168.11.1",
    "ip": "192.168.11.91",
    "netmask": "255.255.255.0"
  },
  "tunterm_port_config": {
    "downstream_ports": [
      "0",
      "1"
    ],
    "separate_upstream_downstream": false,
    "upstream_ports": [
      "0",
      "1"
    ]
  },
  "tunterm_registered": true,
  "tunterm_stat": {
    "monitoring_failed": false
  },
  "uptime": 76281,
  "virtualization_type": "KVM"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Mx Edges Stats

Get List of Org MxEdge Stats

```go
ListOrgMxEdgesStats(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string,
    forSite *bool) (
    models.ApiResponse[[]models.MxedgeStats],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `forSite` | `*bool` | Query, Optional | filter for site level mist edges |

## Response Type

[`[]models.MxedgeStats`](../../doc/models/mxedge-stats.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100





duration := "10m"

forSite := false

apiResponse, err := orgsStatsMxEdges.ListOrgMxEdgesStats(ctx, orgId, &page, &limit, nil, nil, &duration, &forSite)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
