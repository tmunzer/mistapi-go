# Sites Stats-Mx Edges

```go
sitesStatsMxEdges := client.SitesStatsMxEdges()
```

## Class Name

`SitesStatsMxEdges`

## Methods

* [Get Site Mx Edge Stats](../../doc/controllers/sites-stats-mx-edges.md#get-site-mx-edge-stats)
* [List Site Mx Edges Stats](../../doc/controllers/sites-stats-mx-edges.md#list-site-mx-edges-stats)


# Get Site Mx Edge Stats

Get One Site MxEdge Stats

```go
GetSiteMxEdgeStats(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.StatsMxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.StatsMxedge](../../doc/models/stats-mxedge.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

apiResponse, err := sitesStatsMxEdges.GetSiteMxEdgeStats(ctx, siteId, mxedgeId, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Site Mx Edges Stats

Get List of Site MxEdges Stats

```go
ListSiteMxEdgesStats(
    ctx context.Context,
    siteId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsMxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.StatsMxedge](../../doc/models/stats-mxedge.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesStatsMxEdges.ListSiteMxEdgesStats(ctx, siteId, nil, nil, &duration, &limit, &page)
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

