# Orgs Devices-WAN Cluster

```go
orgsDevicesWANCluster := client.OrgsDevicesWANCluster()
```

## Class Name

`OrgsDevicesWANCluster`

## Methods

* [Create Org Gateway Ha Cluster](../../doc/controllers/orgs-devices-wan-cluster.md#create-org-gateway-ha-cluster)
* [Delete Org Gateway Ha Cluster](../../doc/controllers/orgs-devices-wan-cluster.md#delete-org-gateway-ha-cluster)


# Create Org Gateway Ha Cluster

Create HA Cluster from unassigned Gateways

```go
CreateOrgGatewayHaCluster(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.HaClusterConfig) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.HaClusterConfig`](../../doc/models/ha-cluster-config.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.HaClusterConfig{
    DisableAutoConfig: models.ToPointer(true),
    Managed:           models.ToPointer(true),
    Nodes:             []models.HaClusterConfigNode{
        models.HaClusterConfigNode{
            Mac: models.ToPointer("aff827549235"),
        },
        models.HaClusterConfigNode{
            Mac: models.ToPointer("8396cd006c8c"),
        },
    },
    SiteId:            models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
}

resp, err := orgsDevicesWANCluster.CreateOrgGatewayHaCluster(ctx, orgId, &body)
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


# Delete Org Gateway Ha Cluster

Delete HA Cluster

After HA cluster deleted, both of the nodes will be unassigned.

```go
DeleteOrgGatewayHaCluster(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.HaClusterDelete) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.HaClusterDelete`](../../doc/models/ha-cluster-delete.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.HaClusterDelete{
    Mac: models.ToPointer("aff827549235"),
}

resp, err := orgsDevicesWANCluster.DeleteOrgGatewayHaCluster(ctx, orgId, &body)
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

