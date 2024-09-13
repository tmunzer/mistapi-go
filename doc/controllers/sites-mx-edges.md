# Sites Mx Edges

```go
sitesMxEdges := client.SitesMxEdges()
```

## Class Name

`SitesMxEdges`

## Methods

* [Count Site Mx Edge Events](../../doc/controllers/sites-mx-edges.md#count-site-mx-edge-events)
* [Create Site Mx Edge](../../doc/controllers/sites-mx-edges.md#create-site-mx-edge)
* [Delete Site Mx Edge](../../doc/controllers/sites-mx-edges.md#delete-site-mx-edge)
* [Get Site Mx Edge](../../doc/controllers/sites-mx-edges.md#get-site-mx-edge)
* [List Site Mx Edges](../../doc/controllers/sites-mx-edges.md#list-site-mx-edges)
* [Search Site Mist Edge Events](../../doc/controllers/sites-mx-edges.md#search-site-mist-edge-events)
* [Update Site Mx Edge](../../doc/controllers/sites-mx-edges.md#update-site-mx-edge)
* [Upload Site Mx Edge Support Files](../../doc/controllers/sites-mx-edges.md#upload-site-mx-edge-support-files)


# Count Site Mx Edge Events

Count Mist Edge Events

```go
CountSiteMxEdgeEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteMxedgeEventsCountDistinctEnum,
    mxedgeId *string,
    mxclusterId *string,
    mType *string,
    service *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteMxedgeEventsCountDistinctEnum`](../../doc/models/site-mxedge-events-count-distinct-enum.md) | Query, Optional | **Default**: `"mxedge_id"` |
| `mxedgeId` | `*string` | Query, Optional | mist edge id |
| `mxclusterId` | `*string` | Query, Optional | mist edge cluster id |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `service` | `*string` | Query, Optional | service running on mist edge(mxagent, tunterm etc) |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteMxedgeEventsCountDistinctEnum("mxedge_id")













duration := "10m"

limit := 100

apiResponse, err := sitesMxEdges.CountSiteMxEdgeEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# Create Site Mx Edge

Create Site Mist Edge

```go
CreateSiteMxEdge(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Mxedge) (
    models.ApiResponse[models.Mxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxedge`](../../doc/models/mxedge.md) | Body, Optional | - |

## Response Type

[`models.Mxedge`](../../doc/models/mxedge.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxedge{
    Id:                        models.ToPointer(uuid.MustParse("95ddd29a-6a3c-929e-a431-51a5b09f36a6")),
    Magic:                     models.ToPointer("L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD"),
    Model:                     "ME-100",
    MxclusterId:               models.ToPointer(uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4")),
    Name:                      "Guest",
    Note:                      models.ToPointer("note for mxedge"),
    OrgId:                     models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:                    models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := sitesMxEdges.CreateSiteMxEdge(ctx, siteId, &body)
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
  "id": "95ddd29a-6a3c-929e-a431-51a5b09f36a6",
  "magic": "L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD",
  "model": "ME-100",
  "mxagent_registered": true,
  "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
  "mxedge_mgmt": {
    "mist_password": "MIST_PASSWORD",
    "root_password": "ROOT_PASSWORD"
  },
  "name": "Guest",
  "ntp_servers": [],
  "oob_ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.10",
    "netmask": "255.255.255.0",
    "type": "static"
  },
  "tunterm_dhcpd_config": {
    "2": {
      "enabled": true,
      "servers": [
        "11.2.3.44"
      ]
    },
    "enabled": false,
    "servers": [
      "11.2.3.4"
    ]
  },
  "tunterm_extra_routes": {
    "11.0.0.0/8": {
      "via": "10.3.3.1"
    }
  },
  "tunterm_ip_config": {
    "dns": [
      "8.8.8.8"
    ],
    "dns_suffix": [
      ".mist.local"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.1",
    "netmask": "255.255.255.0"
  }
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


# Delete Site Mx Edge

Delete Site Mist Edge

```go
DeleteSiteMxEdge(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMxEdges.DeleteSiteMxEdge(ctx, siteId, mxedgeId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
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


# Get Site Mx Edge

get Site Mist Edge

```go
GetSiteMxEdge(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMxEdges.GetSiteMxEdge(ctx, siteId, mxedgeId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
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


# List Site Mx Edges

Get List of Site Mist Edges

```go
ListSiteMxEdges(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Mxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Mxedge`](../../doc/models/mxedge.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesMxEdges.ListSiteMxEdges(ctx, siteId, &limit, &page)
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
          "idle": 79,
          "interrupt": 0,
          "system": 4,
          "usage": 20,
          "user": 16
        },
        "cpu1": {
          "idle": 93,
          "interrupt": 0,
          "system": 4,
          "usage": 6,
          "user": 1
        }
      },
      "idle": 87,
      "interrupt": 0,
      "system": 5,
      "usage": 12,
      "user": 7
    },
    "ext_ip": "116.187.144.16",
    "id": "387804a7-3474-85ce-15a2-f9a9684c9c90",
    "ip_stat": {
      "ip": "172.16.5.3",
      "ips": {
        "ens192": "172.16.5.3/24,fe81::20c:29ff:fef8:d18e/64"
      }
    },
    "lag_stat": {
      "lag0": {
        "active_ports": [
          "0",
          "1"
        ]
      }
    },
    "last_seen": 1547437078,
    "magic": "ExNpT5gi-ADR8WTFd4EiQPY3cP5WdSoD",
    "memory_stats": {
      "active": 1061085184,
      "available": 4124860416,
      "buffers": 789495808,
      "cached": 718016512,
      "free": 2818838528,
      "inactive": 458158080,
      "swap_cached": 0,
      "swap_free": 8161062912,
      "swap_total": 8161062912,
      "total": 7947616256,
      "usage": 65
    },
    "model": "ME-S2019",
    "mxagent_registered": false,
    "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
    "name": "Guest",
    "num_tunnels": 31,
    "port_stat": {
      "eth0": {
        "full_duplex": true,
        "lldp_stats": {
          "mgmt_addr": "122.16.3.11",
          "port_desc": "GigabitEthernet4/0/16",
          "port_id": "\u0005Gi4/0/16",
          "system_desc": "Cisco IOS Software",
          "system_name": "ME-DC2-DIS-SW"
        },
        "rx_bytes": 2056,
        "rx_errors": 0,
        "rx_pkts": 670,
        "speed": 1000,
        "tx_bytes": 2056,
        "tx_pkts": 670,
        "up": true
      },
      "eth1": {
        "up": false
      },
      "module": {
        "up": false
      }
    },
    "status": "connected",
    "tunterm_registered": false,
    "tunterm_stat": {
      "monitoring_failed": false
    },
    "uptime": 884221,
    "version": "0.1.2",
    "virtualization_type": "VirtualizationVMware"
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


# Search Site Mist Edge Events

Search Site Mist Edge Events

```go
SearchSiteMistEdgeEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId *string,
    mxclusterId *string,
    mType *string,
    service *string,
    component *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseMxedgeEventsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `*string` | Query, Optional | mist edge id |
| `mxclusterId` | `*string` | Query, Optional | mist edge cluster id |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `service` | `*string` | Query, Optional | service running on mist edge(mxagent, tunterm etc) |
| `component` | `*string` | Query, Optional | component like PS1, PS2 |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |

## Response Type

[`models.ResponseMxedgeEventsSearch`](../../doc/models/response-mxedge-events-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")















duration := "10m"

limit := 100

apiResponse, err := sitesMxEdges.SearchSiteMistEdgeEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
  "end": 1694708579,
  "limit": 10,
  "page": 3,
  "results": [
    {
      "mxcluster_id": "2815c917-58e7-472f-a190-bfd44fb58d05",
      "mxedge_id": "00000000-0000-0000-1000-020000dc585c",
      "org_id": "f2695c32-0e83-4936-b1b2-96fc88051213",
      "service": "tunterm",
      "timestamp": 1694678225.927,
      "type": "ME_SERVICE_STOPPED"
    }
  ],
  "start": 1694622179
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


# Update Site Mx Edge

Update Site Mist Edge settings

```go
UpdateSiteMxEdge(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID,
    body *models.Mxedge) (
    models.ApiResponse[models.Mxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxedge`](../../doc/models/mxedge.md) | Body, Optional | - |

## Response Type

[`models.Mxedge`](../../doc/models/mxedge.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxedge{
    Id:                        models.ToPointer(uuid.MustParse("95ddd29a-6a3c-929e-a431-51a5b09f36a6")),
    Magic:                     models.ToPointer("L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD"),
    Model:                     "ME-100",
    MxclusterId:               models.ToPointer(uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4")),
    Name:                      "Guest",
    Note:                      models.ToPointer("note for mxedge"),
    OrgId:                     models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:                    models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := sitesMxEdges.UpdateSiteMxEdge(ctx, siteId, mxedgeId, &body)
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
  "id": "95ddd29a-6a3c-929e-a431-51a5b09f36a6",
  "magic": "L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD",
  "model": "ME-100",
  "mxagent_registered": true,
  "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
  "mxedge_mgmt": {
    "mist_password": "MIST_PASSWORD",
    "root_password": "ROOT_PASSWORD"
  },
  "name": "Guest",
  "ntp_servers": [],
  "oob_ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.10",
    "netmask": "255.255.255.0",
    "type": "static"
  },
  "tunterm_dhcpd_config": {
    "2": {
      "enabled": true,
      "servers": [
        "11.2.3.44"
      ]
    },
    "enabled": false,
    "servers": [
      "11.2.3.4"
    ]
  },
  "tunterm_extra_routes": {
    "11.0.0.0/8": {
      "via": "10.3.3.1"
    }
  },
  "tunterm_ip_config": {
    "dns": [
      "8.8.8.8"
    ],
    "dns_suffix": [
      ".mist.local"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.1",
    "netmask": "255.255.255.0"
  }
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


# Upload Site Mx Edge Support Files

Support / Upload Mist Edge support files

```go
UploadSiteMxEdgeSupportFiles(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMxEdges.UploadSiteMxEdgeSupportFiles(ctx, siteId, mxedgeId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
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

