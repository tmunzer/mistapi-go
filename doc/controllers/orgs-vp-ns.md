# Orgs VP Ns

```go
orgsVPNs := client.OrgsVPNs()
```

## Class Name

`OrgsVPNs`

## Methods

* [Count Org Peer Path Stats](../../doc/controllers/orgs-vp-ns.md#count-org-peer-path-stats)
* [Create Org Vpns](../../doc/controllers/orgs-vp-ns.md#create-org-vpns)
* [Delete Org Vpn](../../doc/controllers/orgs-vp-ns.md#delete-org-vpn)
* [Get Org Vpn](../../doc/controllers/orgs-vp-ns.md#get-org-vpn)
* [List Orgs Vpns](../../doc/controllers/orgs-vp-ns.md#list-orgs-vpns)
* [Search Org Peer Path Stats](../../doc/controllers/orgs-vp-ns.md#search-org-peer-path-stats)
* [Update Org Vpn](../../doc/controllers/orgs-vp-ns.md#update-org-vpn)


# Count Org Peer Path Stats

Count Org Peer Path Statgs

```go
CountOrgPeerPathStats(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | `*string` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsVPNs.CountOrgPeerPathStats(ctx, orgId, nil, &page, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Create Org Vpns

Create Org VPN

```go
CreateOrgVpns(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Vpn) (
    models.ApiResponse[models.Vpn],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Vpn`](../../doc/models/vpn.md) | Body, Optional | - |

## Response Type

[`models.Vpn`](../../doc/models/vpn.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Vpn{
    Name:         "string",
    Paths:        map[string]models.VpnPath{
        "property1": models.VpnPath{
            BfdProfile: models.ToPointer(models.VpnPathBfdProfileEnum("broadband")),
            Ip:         models.ToPointer("string"),
        },
        "property2": models.VpnPath{
            BfdProfile: models.ToPointer(models.VpnPathBfdProfileEnum("lte")),
            Ip:         models.ToPointer("string"),
        },
    },
}

apiResponse, err := orgsVPNs.CreateOrgVpns(ctx, orgId, &body)
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
  "id": "497f6eca-6276-5009-bfeb-53cbbbba6f1b",
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "paths": {
    "property1": {
      "bfd_profile": "broadband",
      "ip": "string"
    },
    "property2": {
      "bfd_profile": "broadband",
      "ip": "string"
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Delete Org Vpn

delete Org Vpn

```go
DeleteOrgVpn(
    ctx context.Context,
    orgId uuid.UUID,
    vpnId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `vpnId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

vpnId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsVPNs.DeleteOrgVpn(ctx, orgId, vpnId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Org Vpn

getOrgVpn

```go
GetOrgVpn(
    ctx context.Context,
    orgId uuid.UUID,
    vpnId uuid.UUID) (
    models.ApiResponse[models.Vpn],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `vpnId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Vpn`](../../doc/models/vpn.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

vpnId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsVPNs.GetOrgVpn(ctx, orgId, vpnId)
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
  "id": "497f6eca-6276-5009-bfeb-53cbbbba6f1b",
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "paths": {
    "property1": {
      "bfd_profile": "broadband",
      "ip": "string"
    },
    "property2": {
      "bfd_profile": "broadband",
      "ip": "string"
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Orgs Vpns

Get List of Org VPNs

```go
ListOrgsVpns(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Vpn],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Vpn`](../../doc/models/vpn.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsVPNs.ListOrgsVpns(ctx, orgId, &page, &limit)
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
    "name": "string",
    "paths": {
      "property1": {
        "bfd_profile": "broadband",
        "ip": "string"
      },
      "property2": {
        "bfd_profile": "lte",
        "ip": "string"
      }
    }
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Org Peer Path Stats

Search Org Peer Path Stats

```go
SearchOrgPeerPathStats(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.VpnPeerStatSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`models.VpnPeerStatSearch`](../../doc/models/vpn-peer-stat-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

limit := 100

apiResponse, err := orgsVPNs.SearchOrgPeerPathStats(ctx, orgId, nil, nil, &duration, &limit)
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
  "end": 1619518989.4989712,
  "limit": 10,
  "results": [
    {
      "auth_algo": "hmac-sha1-96",
      "enc_algo": "aes-cbc-128",
      "ike_version": "1",
      "is_active": true,
      "last_seen": 1619518709.222,
      "mac": "020001c04668",
      "org_id": "0c160b7f-1027-4cd1-923b-744534c4b070",
      "peer_mac": "020001367edd",
      "peer_port_id": "DC_Internet",
      "peer_site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "port_id": "Lte",
      "site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "type": "svr",
      "up": true,
      "uptime": 1527128046
    },
    {
      "is_active": true,
      "last_seen": 1619518709.222,
      "latency": 91,
      "mac": "020001c04668",
      "mos": 436,
      "mtu": 1500,
      "org_id": "0c160b7f-1027-4cd1-923b-744534c4b070",
      "peer_mac": "020001367edd",
      "peer_port_id": "DC_Internet",
      "peer_router_name": "RIDCBBP1",
      "peer_site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "port_id": "Lte",
      "router_name": "RIST01544AP1",
      "site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "type": "svr",
      "up": true,
      "uptime": 1527128046
    }
  ],
  "start": 1619518689.4989705,
  "total": 2
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Org Vpn

update Org Vpn

```go
UpdateOrgVpn(
    ctx context.Context,
    orgId uuid.UUID,
    vpnId uuid.UUID,
    body *models.Vpn) (
    models.ApiResponse[models.Vpn],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `vpnId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Vpn`](../../doc/models/vpn.md) | Body, Optional | - |

## Response Type

[`models.Vpn`](../../doc/models/vpn.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

vpnId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Vpn{
    Name:         "string",
    Paths:        map[string]models.VpnPath{
        "property1": models.VpnPath{
            BfdProfile: models.ToPointer(models.VpnPathBfdProfileEnum("broadband")),
            Ip:         models.ToPointer("string"),
        },
        "property2": models.VpnPath{
            BfdProfile: models.ToPointer(models.VpnPathBfdProfileEnum("broadband")),
            Ip:         models.ToPointer("string"),
        },
    },
}

apiResponse, err := orgsVPNs.UpdateOrgVpn(ctx, orgId, vpnId, &body)
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
  "id": "497f6eca-6276-5009-bfeb-53cbbbba6f1b",
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "paths": {
    "property1": {
      "bfd_profile": "broadband",
      "ip": "string"
    },
    "property2": {
      "bfd_profile": "broadband",
      "ip": "string"
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

