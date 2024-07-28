# Orgs Asset Filters

```go
orgsAssetFilters := client.OrgsAssetFilters()
```

## Class Name

`OrgsAssetFilters`

## Methods

* [Create Org Asset Filters](../../doc/controllers/orgs-asset-filters.md#create-org-asset-filters)
* [Delete Org Asset Filter](../../doc/controllers/orgs-asset-filters.md#delete-org-asset-filter)
* [Get Org Asset Filter](../../doc/controllers/orgs-asset-filters.md#get-org-asset-filter)
* [List Org Asset Filters](../../doc/controllers/orgs-asset-filters.md#list-org-asset-filters)
* [Update Org Asset Filters](../../doc/controllers/orgs-asset-filters.md#update-org-asset-filters)


# Create Org Asset Filters

Create Asset Filter

Creates a single BLE asset filter for the given site. Any subset of filter properties can be included in the filter. A matching asset must meet the conditions of all given filter properties (logical ‘AND’).

```go
CreateOrgAssetFilters(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AssetFilter) (
    models.ApiResponse[models.AssetFilter],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AssetFilter`](../../doc/models/asset-filter.md) | Body, Optional | - |

## Response Type

[`models.AssetFilter`](../../doc/models/asset-filter.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AssetFilter{
    Disasbled:             models.ToPointer(true),
    EddystoneUidNamespace: models.ToPointer("string"),
    EddystoneUrl:          models.ToPointer("string"),
    IbeaconMajor:          models.ToPointer(0),
    IbeaconUuid:           models.ToPointer(uuid.MustParse("1f89bc00-d0af-481b-82fe-a6629259a39f")),
    MfgCompanyId:          models.ToPointer(0),
    Name:                  "string",
}

apiResponse, err := orgsAssetFilters.CreateOrgAssetFilters(ctx, orgId, &body)
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
  "disasbled": true,
  "eddystone_uid_namespace": "string",
  "eddystone_url": "string",
  "for_site": true,
  "ibeacon_major": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mfg_company_id": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# Delete Org Asset Filter

Deletes an existing BLE asset filter for the given site.

```go
DeleteOrgAssetFilter(
    ctx context.Context,
    orgId uuid.UUID,
    assetfilterId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `assetfilterId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetfilterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsAssetFilters.DeleteOrgAssetFilter(ctx, orgId, assetfilterId)
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


# Get Org Asset Filter

Get Org Asset Filter Details

```go
GetOrgAssetFilter(
    ctx context.Context,
    orgId uuid.UUID,
    assetfilterId uuid.UUID) (
    models.ApiResponse[models.AssetFilter],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `assetfilterId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.AssetFilter`](../../doc/models/asset-filter.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetfilterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsAssetFilters.GetOrgAssetFilter(ctx, orgId, assetfilterId)
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
  "disasbled": true,
  "eddystone_uid_namespace": "string",
  "eddystone_url": "string",
  "for_site": true,
  "ibeacon_major": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mfg_company_id": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# List Org Asset Filters

Get List of Org BLE asset filters.
Each asset filter in the list operates independently. For a filter object to match an asset, all of the filter properties must match (logical ‘AND’ of each filter property). For example, the “Visitor Tags” filter below will match an asset when both the “ibeacon\_uuid” and “ibeacon_major” properties match the asset. All non-matching assets are ignored.

```go
ListOrgAssetFilters(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.AssetFilter],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.AssetFilter`](../../doc/models/asset-filter.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsAssetFilters.ListOrgAssetFilters(ctx, orgId, &page, &limit)
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


# Update Org Asset Filters

Updates an existing BLE asset filter for the given site.

```go
UpdateOrgAssetFilters(
    ctx context.Context,
    orgId uuid.UUID,
    assetfilterId uuid.UUID,
    body *models.AssetFilter) (
    models.ApiResponse[models.AssetFilter],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `assetfilterId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AssetFilter`](../../doc/models/asset-filter.md) | Body, Optional | Request Body |

## Response Type

[`models.AssetFilter`](../../doc/models/asset-filter.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

assetfilterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AssetFilter{
    Disabled:              models.ToPointer(true),
    EddystoneUidNamespace: models.ToPointer("string"),
    EddystoneUrl:          models.ToPointer("string"),
    IbeaconMajor:          models.ToPointer(0),
    IbeaconUuid:           models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab2")),
    MfgCompanyId:          models.ToPointer(0),
    Name:                  "string",
}

apiResponse, err := orgsAssetFilters.UpdateOrgAssetFilters(ctx, orgId, assetfilterId, &body)
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
  "disasbled": true,
  "eddystone_uid_namespace": "string",
  "eddystone_url": "string",
  "for_site": true,
  "ibeacon_major": 0,
  "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mfg_company_id": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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

