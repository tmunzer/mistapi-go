# Sites Wx Tunnels

```go
sitesWxTunnels := client.SitesWxTunnels()
```

## Class Name

`SitesWxTunnels`

## Methods

* [Create Site Wx Tunnel](../../doc/controllers/sites-wx-tunnels.md#create-site-wx-tunnel)
* [Delete Site Wx Tunnel](../../doc/controllers/sites-wx-tunnels.md#delete-site-wx-tunnel)
* [Get Site Wx Tunnel](../../doc/controllers/sites-wx-tunnels.md#get-site-wx-tunnel)
* [List Site Wx Tunnels](../../doc/controllers/sites-wx-tunnels.md#list-site-wx-tunnels)
* [Update Site Wx Tunnel](../../doc/controllers/sites-wx-tunnels.md#update-site-wx-tunnel)


# Create Site Wx Tunnel

Create Site WxLan Tunnel

```go
CreateSiteWxTunnel(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.WxlanTunnel) (
    models.ApiResponse[models.WxlanTunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanTunnel`](../../doc/models/wxlan-tunnel.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WxlanTunnel](../../doc/models/wxlan-tunnel.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanTunnel{
    Dmvpn:                models.ToPointer(models.WxlanTunnelDmvpn{
        Enabled:              models.ToPointer(true),
        HoldingTime:          models.ToPointer(0),
        HostRoutes:           []string{
            "string",
        },
    }),
    ForMgmt:              models.ToPointer(true),
    HelloInterval:        models.ToPointer(1),
    HelloRetries:         models.ToPointer(3),
    Hostname:             models.ToPointer("string"),
    Ipsec:                models.ToPointer(models.WxlanTunnelIpsec{
        Enabled:              models.ToPointer(true),
        Psk:                  "string123",
    }),
    IsStatic:             models.ToPointer(true),
    Mtu:                  models.ToPointer(0),
    Name:                 "string",
    Peers:                []string{
        "string",
    },
    RouterId:             models.ToPointer("string"),
    Secret:               models.ToPointer("string"),
    Sessions:             []models.WxlanTunnelSession{
        models.WxlanTunnelSession{
            ApAsSessionId:        models.ToPointer("string"),
            Comment:              models.ToPointer("string"),
            EnableCookie:         models.ToPointer(true),
            Ethertype:            models.ToPointer(models.WxlanTunnelSessionEthertypeEnum_ETHERNET),
            LocalSessionId:       models.ToPointer(1),
            Pseudo8021adEnabled:  models.ToPointer(true),
            RemoteId:             models.ToPointer("string"),
            RemoteSessionId:      models.ToPointer(1),
            UseApAsSessionIds:    models.ToPointer(true),
        },
    },
    UdpPort:              models.ToPointer(0),
    UseUdp:               models.ToPointer(true),
}

apiResponse, err := sitesWxTunnels.CreateSiteWxTunnel(ctx, siteId, &body)
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
  "mtu": 0,
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


# Delete Site Wx Tunnel

Delete Site WxLan Tunnel

```go
DeleteSiteWxTunnel(
    ctx context.Context,
    siteId uuid.UUID,
    wxtunnelId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesWxTunnels.DeleteSiteWxTunnel(ctx, siteId, wxtunnelId)
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


# Get Site Wx Tunnel

Get Site WxLan tunnel Details

```go
GetSiteWxTunnel(
    ctx context.Context,
    siteId uuid.UUID,
    wxtunnelId uuid.UUID) (
    models.ApiResponse[models.WxlanTunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WxlanTunnel](../../doc/models/wxlan-tunnel.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesWxTunnels.GetSiteWxTunnel(ctx, siteId, wxtunnelId)
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
  "mtu": 0,
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


# List Site Wx Tunnels

Get List of Site WxLan Tunnels

```go
ListSiteWxTunnels(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.WxlanTunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.WxlanTunnel](../../doc/models/wxlan-tunnel.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesWxTunnels.ListSiteWxTunnels(ctx, siteId, &limit, &page)
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


# Update Site Wx Tunnel

Update Site WxLan Tunnel

```go
UpdateSiteWxTunnel(
    ctx context.Context,
    siteId uuid.UUID,
    wxtunnelId uuid.UUID,
    body *models.WxlanTunnel) (
    models.ApiResponse[models.WxlanTunnel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxtunnelId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanTunnel`](../../doc/models/wxlan-tunnel.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WxlanTunnel](../../doc/models/wxlan-tunnel.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanTunnel{
    Dmvpn:                models.ToPointer(models.WxlanTunnelDmvpn{
        Enabled:              models.ToPointer(true),
        HoldingTime:          models.ToPointer(0),
        HostRoutes:           []string{
            "string",
        },
    }),
    ForMgmt:              models.ToPointer(true),
    HelloInterval:        models.ToPointer(1),
    HelloRetries:         models.ToPointer(3),
    Hostname:             models.ToPointer("string"),
    Ipsec:                models.ToPointer(models.WxlanTunnelIpsec{
        Enabled:              models.ToPointer(true),
        Psk:                  "string123",
    }),
    IsStatic:             models.ToPointer(true),
    Mtu:                  models.ToPointer(0),
    Name:                 "string",
    Peers:                []string{
        "string",
    },
    RouterId:             models.ToPointer("string"),
    Secret:               models.ToPointer("string"),
    Sessions:             []models.WxlanTunnelSession{
        models.WxlanTunnelSession{
            ApAsSessionId:        models.ToPointer("string"),
            Comment:              models.ToPointer("string"),
            EnableCookie:         models.ToPointer(true),
            Ethertype:            models.ToPointer(models.WxlanTunnelSessionEthertypeEnum_ETHERNET),
            LocalSessionId:       models.ToPointer(1),
            Pseudo8021adEnabled:  models.ToPointer(true),
            RemoteId:             models.ToPointer("string"),
            RemoteSessionId:      models.ToPointer(1),
            UseApAsSessionIds:    models.ToPointer(true),
        },
    },
    UdpPort:              models.ToPointer(0),
    UseUdp:               models.ToPointer(true),
}

apiResponse, err := sitesWxTunnels.UpdateSiteWxTunnel(ctx, siteId, wxtunnelId, &body)
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
  "mtu": 0,
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

