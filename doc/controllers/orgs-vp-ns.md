# Orgs VP Ns

```go
orgsVPNs := client.OrgsVPNs()
```

## Class Name

`OrgsVPNs`

## Methods

* [Create Org Vpn](../../doc/controllers/orgs-vp-ns.md#create-org-vpn)
* [Delete Org Vpn](../../doc/controllers/orgs-vp-ns.md#delete-org-vpn)
* [Get Org Vpn](../../doc/controllers/orgs-vp-ns.md#get-org-vpn)
* [List Orgs Vpns](../../doc/controllers/orgs-vp-ns.md#list-orgs-vpns)
* [Update Org Vpn](../../doc/controllers/orgs-vp-ns.md#update-org-vpn)


# Create Org Vpn

Create Org VPN

```go
CreateOrgVpn(
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Vpn](../../doc/models/vpn.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Vpn{
    Name:                 "string",
    Paths:                map[string]models.VpnPath{
        "property1": models.VpnPath{
            BfdProfile:           models.ToPointer(models.VpnPathBfdProfileEnum_BROADBAND),
            Ip:                   models.ToPointer("string"),
        },
        "property2": models.VpnPath{
            BfdProfile:           models.ToPointer(models.VpnPathBfdProfileEnum_LTE),
            Ip:                   models.ToPointer("string"),
        },
    },
}

apiResponse, err := orgsVPNs.CreateOrgVpn(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Vpn

Delete Org Vpn

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Vpn

Get Org Vpn

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Vpn](../../doc/models/vpn.md).

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Orgs Vpns

Get List of Org VPNs

```go
ListOrgsVpns(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Vpn],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Vpn](../../doc/models/vpn.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsVPNs.ListOrgsVpns(ctx, orgId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Vpn

Update Org Vpn

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Vpn](../../doc/models/vpn.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

vpnId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Vpn{
    Name:                 "string",
    Paths:                map[string]models.VpnPath{
        "property1": models.VpnPath{
            BfdProfile:           models.ToPointer(models.VpnPathBfdProfileEnum_BROADBAND),
            Ip:                   models.ToPointer("string"),
        },
        "property2": models.VpnPath{
            BfdProfile:           models.ToPointer(models.VpnPathBfdProfileEnum_BROADBAND),
            Ip:                   models.ToPointer("string"),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

