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

Create an organization Mist Tunnel configuration, including hosting Mist Edge clusters, VLANs, heartbeat settings, and optional IPsec settings.

```go
CreateOrgMxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Mxtunnel) (
    models.ApiResponse[models.Mxtunnel],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxtunnel`](../../doc/models/mxtunnel.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxtunnel](../../doc/models/mxtunnel.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxtunnel{
    HelloInterval:        models.NewOptional(models.ToPointer(60)),
    HelloRetries:         models.NewOptional(models.ToPointer(7)),
    Ipsec:                models.ToPointer(models.MxtunnelIpsec{
        DnsServers:           models.NewOptional(models.ToPointer([]string{
            "string",
        })),
        Enabled:              models.ToPointer(true),
        ExtraRoutes:          []models.MxtunnelIpsecExtraRoute{
            models.MxtunnelIpsecExtraRoute{
                Dest:                 models.ToPointer("string"),
                NextHop:              models.ToPointer("192.168.0.1"),
            },
        },
        SplitTunnel:          models.ToPointer(true),
        UseMxedge:            models.ToPointer(true),
    }),
    VlanIds:              []int{
        0,
    },
    AdditionalProperties: map[string]interface{}{
        "cluster_ids": interface{}("string"),
    },
}

apiResponse, err := orgsMxTunnels.CreateOrgMxTunnel(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Org Mx Tunnel

Delete an organization Mist Tunnel configuration by Mist Tunnel ID.

```go
DeleteOrgMxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    mxtunnelId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxTunnels.DeleteOrgMxTunnel(ctx, orgId, mxtunnelId)
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
    fmt.Println(resp.StatusCode)
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


# Get Org Mx Tunnel

Retrieve configuration details for a specific organization Mist Tunnel, including cluster, VLAN, heartbeat, IPsec, and preemption settings.

```go
GetOrgMxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    mxtunnelId uuid.UUID) (
    models.ApiResponse[models.Mxtunnel],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxtunnel](../../doc/models/mxtunnel.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsMxTunnels.GetOrgMxTunnel(ctx, orgId, mxtunnelId)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Mx Tunnels

List organization Mist Tunnel configurations used to carry AP user VLANs to Mist Edge clusters.

```go
ListOrgMxTunnels(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Mxtunnel],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Mxtunnel](../../doc/models/mxtunnel.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsMxTunnels.ListOrgMxTunnels(ctx, orgId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Update Org Mx Tunnel

Update an organization Mist Tunnel configuration, including cluster membership, VLANs, heartbeat settings, IPsec, and preemption behavior.

```go
UpdateOrgMxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    mxtunnelId uuid.UUID,
    body *models.Mxtunnel) (
    models.ApiResponse[models.Mxtunnel],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxtunnelId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxtunnel`](../../doc/models/mxtunnel.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxtunnel](../../doc/models/mxtunnel.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxtunnel{
    HelloInterval:        models.NewOptional(models.ToPointer(60)),
    HelloRetries:         models.NewOptional(models.ToPointer(7)),
    Mtu:                  models.ToPointer(0),
    Protocol:             models.ToPointer(models.MxtunnelProtocolEnum_UDP),
}

apiResponse, err := orgsMxTunnels.UpdateOrgMxTunnel(ctx, orgId, mxtunnelId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

