# Orgs Mx Tunnels

```go
orgsMxTunnels := client.OrgsMxTunnels()
```

## Class Name

`OrgsMxTunnels`

## Methods

* [Create Org Mx Tunnel](../../doc/controllers/orgs-mx-tunnels.md#create-org-mx-tunnel)
* [Delete Org Mx Tunnel](../../doc/controllers/orgs-mx-tunnels.md#delete-org-mx-tunnel)
* [Get Org Mx Tunnel](../../doc/controllers/orgs-mx-tunnels.md#get-org-mx-tunnel)
* [List Org Mx Tunnels](../../doc/controllers/orgs-mx-tunnels.md#list-org-mx-tunnels)
* [Update Org Mx Tunnel](../../doc/controllers/orgs-mx-tunnels.md#update-org-mx-tunnel)


# Create Org Mx Tunnel

Create MxTunnel

```go
CreateOrgMxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Mxtunnel) (
    models.ApiResponse[models.Mxtunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxtunnel`](../../doc/models/mxtunnel.md) | Body, Optional | Request Body |

## Response Type

[`models.Mxtunnel`](../../doc/models/mxtunnel.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxtunnel{
    HelloInterval:     models.NewOptional(models.ToPointer(60)),
    HelloRetries:      models.NewOptional(models.ToPointer(7)),
    Ipsec:             models.ToPointer(models.MxtunnelIpsec{
        DnsServers:  models.NewOptional(models.ToPointer([]string{
            "string",
        })),
        Enabled:     models.ToPointer(true),
        ExtraRoutes: []models.MxtunnelIpsecExtraRoute{
            models.MxtunnelIpsecExtraRoute{
                Dest:    models.ToPointer("string"),
                NextHop: models.ToPointer("192.168.0.1"),
            },
        },
        SplitTunnel: models.ToPointer(true),
        UseMxedge:   models.ToPointer(true),
    }),
    VlanIds:           []int{
        0,
    },
}

apiResponse, err := orgsMxTunnels.CreateOrgMxTunnel(ctx, orgId, &body)
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
  "cluster_ids": [
    "string"
  ],
  "created_time": 0,
  "for_site": true,
  "hello_interval": 60,
  "hello_retries": 7,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ipsec": {
    "dns_servers": [
      "string"
    ],
    "enabled": true,
    "extra_routes": [
      {
        "dest": "string",
        "next_hop": "192.168.0.1"
      }
    ],
    "split_tunnel": true,
    "use_mxedge": true
  },
  "modified_time": 0,
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "vlan_ids": [
    0
  ]
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


# Delete Org Mx Tunnel

Delete Org MxTunnel

```go
DeleteOrgMxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    mxtunnelId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxTunnels.DeleteOrgMxTunnel(ctx, orgId, mxtunnelId)
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


# Get Org Mx Tunnel

Get Org MxTunnel Details

```go
GetOrgMxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    mxtunnelId uuid.UUID) (
    models.ApiResponse[models.Mxtunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Mxtunnel`](../../doc/models/mxtunnel.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsMxTunnels.GetOrgMxTunnel(ctx, orgId, mxtunnelId)
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
  "cluster_ids": [
    "string"
  ],
  "created_time": 0,
  "for_site": true,
  "hello_interval": 60,
  "hello_retries": 7,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ipsec": {
    "dns_servers": [
      "string"
    ],
    "enabled": true,
    "extra_routes": [
      {
        "dest": "string",
        "next_hop": "192.168.0.1"
      }
    ],
    "split_tunnel": true,
    "use_mxedge": true
  },
  "modified_time": 0,
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "vlan_ids": [
    0
  ]
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


# List Org Mx Tunnels

Get List of Org MxTunnels

```go
ListOrgMxTunnels(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Mxtunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Mxtunnel`](../../doc/models/mxtunnel.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsMxTunnels.ListOrgMxTunnels(ctx, orgId, &limit, &page)
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
    "hello_interval": 60,
    "hello_retries": 3,
    "ipsec": {
      "dns_servers": [
        "172.16.0.8"
      ],
      "enabled": true,
      "extra_routes": [
        {
          "dest": "172.16.0.0/12",
          "next_hop": "172.16.0.1"
        }
      ],
      "split_tunnel": true
    },
    "mxcluster_ids": [
      "572586b7-f97b-a22b-526c-8b97a3f609c4"
    ],
    "name": "HQ",
    "protocol": "udp",
    "vlan_ids": [
      3,
      4,
      5
    ]
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


# Update Org Mx Tunnel

Update Org MxTunnel

```go
UpdateOrgMxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    mxtunnelId uuid.UUID,
    body *models.Mxtunnel) (
    models.ApiResponse[models.Mxtunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxtunnelId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxtunnel`](../../doc/models/mxtunnel.md) | Body, Optional | Request Body |

## Response Type

[`models.Mxtunnel`](../../doc/models/mxtunnel.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxtunnel{
    HelloInterval:     models.NewOptional(models.ToPointer(60)),
    HelloRetries:      models.NewOptional(models.ToPointer(7)),
    Mtu:               models.ToPointer(0),
    OrgId:             models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    Protocol:          models.ToPointer(models.MxtunnelProtocolEnum("udp")),
    SiteId:            models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := orgsMxTunnels.UpdateOrgMxTunnel(ctx, orgId, mxtunnelId, &body)
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
  "cluster_ids": [
    "string"
  ],
  "created_time": 0,
  "for_site": true,
  "hello_interval": 60,
  "hello_retries": 7,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ipsec": {
    "dns_servers": [
      "string"
    ],
    "enabled": true,
    "extra_routes": [
      {
        "dest": "string",
        "next_hop": "192.168.0.1"
      }
    ],
    "split_tunnel": true,
    "use_mxedge": true
  },
  "modified_time": 0,
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "vlan_ids": [
    0
  ]
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

