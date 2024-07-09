# MS Ps Org Groups

```go
mSPsOrgGroups := client.MSPsOrgGroups()
```

## Class Name

`MSPsOrgGroups`

## Methods

* [Create Msp Org Group](../../doc/controllers/ms-ps-org-groups.md#create-msp-org-group)
* [Delete Msp Org Group](../../doc/controllers/ms-ps-org-groups.md#delete-msp-org-group)
* [Get Msp Org Group](../../doc/controllers/ms-ps-org-groups.md#get-msp-org-group)
* [List Msp Org Groups](../../doc/controllers/ms-ps-org-groups.md#list-msp-org-groups)
* [Update Msp Org Group](../../doc/controllers/ms-ps-org-groups.md#update-msp-org-group)


# Create Msp Org Group

Create MSP Org Group

```go
CreateMspOrgGroup(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.Orggroup) (
    models.ApiResponse[models.Orggroup],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Orggroup`](../../doc/models/orggroup.md) | Body, Optional | Request Body |

## Response Type

[`models.Orggroup`](../../doc/models/orggroup.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Orggroup{
    Name:         "string",
    OrgIds:       []uuid.UUID{
        uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
    },
}

apiResponse, err := mSPsOrgGroups.CreateMspOrgGroup(ctx, mspId, &body)
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
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "org_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ]
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


# Delete Msp Org Group

Delete MSP Org Group

```go
DeleteMspOrgGroup(
    ctx context.Context,
    mspId uuid.UUID,
    orggroupId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `orggroupId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

orggroupId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := mSPsOrgGroups.DeleteMspOrgGroup(ctx, mspId, orggroupId)
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


# Get Msp Org Group

Get MSP Org Group Details

```go
GetMspOrgGroup(
    ctx context.Context,
    mspId uuid.UUID,
    orggroupId uuid.UUID) (
    models.ApiResponse[models.Orggroup],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `orggroupId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Orggroup`](../../doc/models/orggroup.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

orggroupId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsOrgGroups.GetMspOrgGroup(ctx, mspId, orggroupId)
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
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "org_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ]
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


# List Msp Org Groups

Get List of MSP Org Groups

```go
ListMspOrgGroups(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.Orggroup],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.Orggroup`](../../doc/models/orggroup.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsOrgGroups.ListMspOrgGroups(ctx, mspId)
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
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "modified_time": 0,
    "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "name": "string",
    "org_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
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


# Update Msp Org Group

Update MSP Org Group

```go
UpdateMspOrgGroup(
    ctx context.Context,
    mspId uuid.UUID,
    orggroupId uuid.UUID,
    body *models.Orggroup) (
    models.ApiResponse[models.Orggroup],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `orggroupId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Orggroup`](../../doc/models/orggroup.md) | Body, Optional | Request Body |

## Response Type

[`models.Orggroup`](../../doc/models/orggroup.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

orggroupId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := mSPsOrgGroups.UpdateMspOrgGroup(ctx, mspId, orggroupId, nil)
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
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "org_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ]
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

