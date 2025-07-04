# Orgs Wx Tunnels

```go
orgsWxTunnels := client.OrgsWxTunnels()
```

## Class Name

`OrgsWxTunnels`

## Methods

* [Create Org Wx Tunnel](../../doc/controllers/orgs-wx-tunnels.md#create-org-wx-tunnel)
* [Delete Org Wx Tunnel](../../doc/controllers/orgs-wx-tunnels.md#delete-org-wx-tunnel)
* [Get Org Wx Tunnel](../../doc/controllers/orgs-wx-tunnels.md#get-org-wx-tunnel)
* [List Org Wx Tunnels](../../doc/controllers/orgs-wx-tunnels.md#list-org-wx-tunnels)
* [Update Org Wx Tunnel](../../doc/controllers/orgs-wx-tunnels.md#update-org-wx-tunnel)


# Create Org Wx Tunnel

Create Org WxAN Tunnel

```go
CreateOrgWxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.WxlanTunnel) (
    models.ApiResponse[models.WxlanTunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanTunnel`](../../doc/models/wxlan-tunnel.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WxlanTunnel](../../doc/models/wxlan-tunnel.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanTunnel{
    ForMgmt:              models.ToPointer(false),
    HelloInterval:        models.ToPointer(60),
    HelloRetries:         models.ToPointer(7),
    IsStatic:             models.ToPointer(false),
    Mtu:                  models.ToPointer(0),
    Name:                 "name6",
    UseUdp:               models.ToPointer(false),
}

apiResponse, err := orgsWxTunnels.CreateOrgWxTunnel(ctx, orgId, &body)
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
  "created_time": 0,
  "dmvpn": {
    "enabled": true,
    "holding_time": 0,
    "host_routes": [
      "string"
    ]
  },
  "for_mgmt": true,
  "hello_interval": 1,
  "hello_retries": 3,
  "hostname": "string",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ipsec": {
    "enabled": true,
    "psk": "string123"
  },
  "is_static": true,
  "modified_time": 0,
  "mtu": 1500,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "peers": [
    "string"
  ],
  "router_id": "string",
  "secret": "string",
  "sessions": [
    {
      "ap_as_session_id": "string",
      "comment": "string",
      "enable_cookie": true,
      "ethertype": "ethernet",
      "local_session_id": 1,
      "pseudo_802.1ad_enabled": true,
      "remote_id": "string",
      "remote_session_id": 1,
      "use_ap_as_session_ids": true
    }
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "udp_port": 0,
  "use_udp": true
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


# Delete Org Wx Tunnel

Delete Org WxLAN Tunnel

```go
DeleteOrgWxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    wxtunnelId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsWxTunnels.DeleteOrgWxTunnel(ctx, orgId, wxtunnelId)
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


# Get Org Wx Tunnel

Get Org WxLAN Tunnel Details

```go
GetOrgWxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    wxtunnelId uuid.UUID) (
    models.ApiResponse[models.WxlanTunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WxlanTunnel](../../doc/models/wxlan-tunnel.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsWxTunnels.GetOrgWxTunnel(ctx, orgId, wxtunnelId)
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
  "created_time": 0,
  "dmvpn": {
    "enabled": true,
    "holding_time": 0,
    "host_routes": [
      "string"
    ]
  },
  "for_mgmt": true,
  "hello_interval": 1,
  "hello_retries": 3,
  "hostname": "string",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ipsec": {
    "enabled": true,
    "psk": "string123"
  },
  "is_static": true,
  "modified_time": 0,
  "mtu": 1500,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "peers": [
    "string"
  ],
  "router_id": "string",
  "secret": "string",
  "sessions": [
    {
      "ap_as_session_id": "string",
      "comment": "string",
      "enable_cookie": true,
      "ethertype": "ethernet",
      "local_session_id": 1,
      "pseudo_802.1ad_enabled": true,
      "remote_id": "string",
      "remote_session_id": 1,
      "use_ap_as_session_ids": true
    }
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "udp_port": 0,
  "use_udp": true
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


# List Org Wx Tunnels

Get List of Org WxLAN Tunnels

```go
ListOrgWxTunnels(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.WxlanTunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.WxlanTunnel](../../doc/models/wxlan-tunnel.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsWxTunnels.ListOrgWxTunnels(ctx, orgId, &limit, &page)
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
    "created_time": 0,
    "dmvpn": {
      "enabled": true,
      "holding_time": 0,
      "host_routes": [
        "string"
      ]
    },
    "for_mgmt": true,
    "hello_interval": 1,
    "hello_retries": 3,
    "hostname": "string",
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ipsec": {
      "enabled": true,
      "psk": "string123"
    },
    "is_static": true,
    "modified_time": 0,
    "mtu": 1500,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "peers": [
      "string"
    ],
    "router_id": "string",
    "secret": "string",
    "sessions": [
      {
        "ap_as_session_id": "string",
        "comment": "string",
        "enable_cookie": true,
        "ethertype": "ethernet",
        "local_session_id": 1,
        "pseudo_802.1ad_enabled": true,
        "remote_id": "string",
        "remote_session_id": 1,
        "use_ap_as_session_ids": true
      }
    ],
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "udp_port": 0,
    "use_udp": true
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


# Update Org Wx Tunnel

Update Org WxLAN Tunnel

```go
UpdateOrgWxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    wxtunnelId uuid.UUID,
    body *models.WxlanTunnel) (
    models.ApiResponse[models.WxlanTunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wxtunnelId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanTunnel`](../../doc/models/wxlan-tunnel.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WxlanTunnel](../../doc/models/wxlan-tunnel.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanTunnel{
    ForMgmt:              models.ToPointer(false),
    HelloInterval:        models.ToPointer(60),
    HelloRetries:         models.ToPointer(7),
    IsStatic:             models.ToPointer(false),
    Mtu:                  models.ToPointer(0),
    Name:                 "name6",
    UseUdp:               models.ToPointer(false),
}

apiResponse, err := orgsWxTunnels.UpdateOrgWxTunnel(ctx, orgId, wxtunnelId, &body)
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
  "created_time": 0,
  "dmvpn": {
    "enabled": true,
    "holding_time": 0,
    "host_routes": [
      "string"
    ]
  },
  "for_mgmt": true,
  "hello_interval": 1,
  "hello_retries": 3,
  "hostname": "string",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ipsec": {
    "enabled": true,
    "psk": "string123"
  },
  "is_static": true,
  "modified_time": 0,
  "mtu": 1500,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "peers": [
    "string"
  ],
  "router_id": "string",
  "secret": "string",
  "sessions": [
    {
      "ap_as_session_id": "string",
      "comment": "string",
      "enable_cookie": true,
      "ethertype": "ethernet",
      "local_session_id": 1,
      "pseudo_802.1ad_enabled": true,
      "remote_id": "string",
      "remote_session_id": 1,
      "use_ap_as_session_ids": true
    }
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "udp_port": 0,
  "use_udp": true
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

