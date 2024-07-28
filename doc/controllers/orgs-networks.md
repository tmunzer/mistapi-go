# Orgs Networks

```go
orgsNetworks := client.OrgsNetworks()
```

## Class Name

`OrgsNetworks`

## Methods

* [Create Org Network](../../doc/controllers/orgs-networks.md#create-org-network)
* [Delete Org Network](../../doc/controllers/orgs-networks.md#delete-org-network)
* [Get Org Network](../../doc/controllers/orgs-networks.md#get-org-network)
* [List Org Networks](../../doc/controllers/orgs-networks.md#list-org-networks)
* [Update Org Network](../../doc/controllers/orgs-networks.md#update-org-network)


# Create Org Network

Create Organization Network

```go
CreateOrgNetwork(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Network) (
    models.ApiResponse[models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Network`](../../doc/models/network.md) | Body, Optional | - |

## Response Type

[`models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Network{
    DisallowMistServices: models.ToPointer(false),
    Gateway:              models.ToPointer("192.168.70.1"),
    Gateway6:             models.ToPointer("fdad:b0bc:f29e::1"),
    Name:                 "name6",
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    Subnet:               models.ToPointer("192.168.70.0/24"),
    Subnet6:              models.ToPointer("fdad:b0bc:f29e::/32"),
}

apiResponse, err := orgsNetworks.CreateOrgNetwork(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Delete Org Network

Delete Organization Network

```go
DeleteOrgNetwork(
    ctx context.Context,
    orgId uuid.UUID,
    networkId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networkId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networkId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsNetworks.DeleteOrgNetwork(ctx, orgId, networkId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Org Network

Get Organization Network Details

```go
GetOrgNetwork(
    ctx context.Context,
    orgId uuid.UUID,
    networkId uuid.UUID) (
    models.ApiResponse[models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networkId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networkId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsNetworks.GetOrgNetwork(ctx, orgId, networkId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Networks

Get List of Org Networks

```go
ListOrgNetworks(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsNetworks.ListOrgNetworks(ctx, orgId, &page, &limit)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Org Network

Update Organization Network

```go
UpdateOrgNetwork(
    ctx context.Context,
    orgId uuid.UUID,
    networkId uuid.UUID,
    body *models.Network) (
    models.ApiResponse[models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networkId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Network`](../../doc/models/network.md) | Body, Optional | - |

## Response Type

[`models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networkId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Network{
    DisallowMistServices: models.ToPointer(false),
    Gateway:              models.ToPointer("192.168.70.1"),
    Gateway6:             models.ToPointer("fdad:b0bc:f29e::1"),
    Name:                 "name6",
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    Subnet:               models.ToPointer("192.168.70.0/24"),
    Subnet6:              models.ToPointer("fdad:b0bc:f29e::/32"),
}

apiResponse, err := orgsNetworks.UpdateOrgNetwork(ctx, orgId, networkId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

