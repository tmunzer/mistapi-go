# Orgs Clients-Marvis

```go
orgsClientsMarvis := client.OrgsClientsMarvis()
```

## Class Name

`OrgsClientsMarvis`

## Methods

* [Create Org Marvis Client Invites](../../doc/controllers/orgs-clients-marvis.md#create-org-marvis-client-invites)
* [Delete Org Marvis Client Invite](../../doc/controllers/orgs-clients-marvis.md#delete-org-marvis-client-invite)
* [Get Org Marvis Client Invites](../../doc/controllers/orgs-clients-marvis.md#get-org-marvis-client-invites)
* [List Org Marvis Client Invites](../../doc/controllers/orgs-clients-marvis.md#list-org-marvis-client-invites)
* [Update Org Marvis Client Invite](../../doc/controllers/orgs-clients-marvis.md#update-org-marvis-client-invite)


# Create Org Marvis Client Invites

Create Org Marvis Client Invites

```go
CreateOrgMarvisClientInvites(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MarvisClient) (
    models.ApiResponse[models.MarvisClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MarvisClient`](../../doc/models/marvis-client.md) | Body, Optional | - |

## Response Type

[`models.MarvisClient`](../../doc/models/marvis-client.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MarvisClient{
    Name:         models.ToPointer("string"),
}

apiResponse, err := orgsClientsMarvis.CreateOrgMarvisClientInvites(ctx, orgId, &body)
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
  "id": "3a14098f-b995-7552-b0a4-b8ee39b337a6”",
  "name": "Handhelds"
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


# Delete Org Marvis Client Invite

Delete Org Marvis Client Invite

```go
DeleteOrgMarvisClientInvite(
    ctx context.Context,
    orgId uuid.UUID,
    marvisinviteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `marvisinviteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

marvisinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsClientsMarvis.DeleteOrgMarvisClientInvite(ctx, orgId, marvisinviteId)
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


# Get Org Marvis Client Invites

Get Org Marvis Client Invite

```go
GetOrgMarvisClientInvites(
    ctx context.Context,
    orgId uuid.UUID,
    marvisinviteId uuid.UUID) (
    models.ApiResponse[models.MarvisClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `marvisinviteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.MarvisClient`](../../doc/models/marvis-client.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

marvisinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsClientsMarvis.GetOrgMarvisClientInvites(ctx, orgId, marvisinviteId)
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
  "id": "3a14098f-b995-7552-b0a4-b8ee39b337a6”",
  "name": "Handhelds"
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


# List Org Marvis Client Invites

List Org Marvis Client Invites

```go
ListOrgMarvisClientInvites(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.MarvisClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.MarvisClient`](../../doc/models/marvis-client.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsClientsMarvis.ListOrgMarvisClientInvites(ctx, orgId)
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
    "id": "3a14098f-b995-7552-b0a4-b8ee39b337a6”",
    "name": "Handhelds"
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


# Update Org Marvis Client Invite

Update Org Marvis Client Invite

```go
UpdateOrgMarvisClientInvite(
    ctx context.Context,
    orgId uuid.UUID,
    marvisinviteId uuid.UUID,
    body *models.MarvisClient) (
    models.ApiResponse[models.MarvisClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `marvisinviteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MarvisClient`](../../doc/models/marvis-client.md) | Body, Optional | - |

## Response Type

[`models.MarvisClient`](../../doc/models/marvis-client.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

marvisinviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MarvisClient{
    Name:         models.ToPointer("string"),
}

apiResponse, err := orgsClientsMarvis.UpdateOrgMarvisClientInvite(ctx, orgId, marvisinviteId, &body)
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
  "id": "3a14098f-b995-7552-b0a4-b8ee39b337a6”",
  "name": "Handhelds"
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

