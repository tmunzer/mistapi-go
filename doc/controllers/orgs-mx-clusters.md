# Orgs Mx Clusters

```go
orgsMxClusters := client.OrgsMxClusters()
```

## Class Name

`OrgsMxClusters`

## Methods

* [Create Org Mx Edge Cluster](../../doc/controllers/orgs-mx-clusters.md#create-org-mx-edge-cluster)
* [Delete Org Mx Edge Cluster](../../doc/controllers/orgs-mx-clusters.md#delete-org-mx-edge-cluster)
* [Get Org Mx Edge Cluster](../../doc/controllers/orgs-mx-clusters.md#get-org-mx-edge-cluster)
* [List Org Mx Edge Clusters](../../doc/controllers/orgs-mx-clusters.md#list-org-mx-edge-clusters)
* [Update Org Mx Edge Cluster](../../doc/controllers/orgs-mx-clusters.md#update-org-mx-edge-cluster)


# Create Org Mx Edge Cluster

Create MxCluster

```go
CreateOrgMxEdgeCluster(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Mxcluster) (
    models.ApiResponse[models.Mxcluster],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxcluster`](../../doc/models/mxcluster.md) | Body, Optional | Request Body |

## Response Type

[`models.Mxcluster`](../../doc/models/mxcluster.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxcluster{
    Name:                      models.ToPointer("string"),
    Radsec:                    models.ToPointer(models.MxclusterRadsec{
        Enabled:         models.ToPointer(true),
    }),
    TuntermApSubnets:          []string{
        "string",
    },
    TuntermHosts:              []string{
        "string",
    },
}

apiResponse, err := orgsMxClusters.CreateOrgMxEdgeCluster(ctx, orgId, &body)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Delete Org Mx Edge Cluster

Delete Org MXEdge Cluster

```go
DeleteOrgMxEdgeCluster(
    ctx context.Context,
    orgId uuid.UUID,
    mxclusterId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxclusterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxClusters.DeleteOrgMxEdgeCluster(ctx, orgId, mxclusterId)
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


# Get Org Mx Edge Cluster

Get Org MxEdge Cluster Details

```go
GetOrgMxEdgeCluster(
    ctx context.Context,
    orgId uuid.UUID,
    mxclusterId uuid.UUID) (
    models.ApiResponse[models.Mxcluster],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Mxcluster`](../../doc/models/mxcluster.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxclusterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsMxClusters.GetOrgMxEdgeCluster(ctx, orgId, mxclusterId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Mx Edge Clusters

Get List of Org MxEdge Clusters

```go
ListOrgMxEdgeClusters(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Mxcluster],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Mxcluster`](../../doc/models/mxcluster.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsMxClusters.ListOrgMxEdgeClusters(ctx, orgId, &page, &limit)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Org Mx Edge Cluster

Update Org MxEdge Cluster

```go
UpdateOrgMxEdgeCluster(
    ctx context.Context,
    orgId uuid.UUID,
    mxclusterId uuid.UUID,
    body *models.Mxcluster) (
    models.ApiResponse[models.Mxcluster],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxcluster`](../../doc/models/mxcluster.md) | Body, Optional | Request Body |

## Response Type

[`models.Mxcluster`](../../doc/models/mxcluster.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxclusterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxcluster{
    Name:                      models.ToPointer("string"),
    Radsec:                    models.ToPointer(models.MxclusterRadsec{
        Enabled:         models.ToPointer(true),
    }),
    TuntermApSubnets:          []string{
        "string",
    },
    TuntermHosts:              []string{
        "string",
    },
}

apiResponse, err := orgsMxClusters.UpdateOrgMxEdgeCluster(ctx, orgId, mxclusterId, &body)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
