# Utilities Mx Edge

```go
utilitiesMxEdge := client.UtilitiesMxEdge()
```

## Class Name

`UtilitiesMxEdge`


# Preempt Sites Mx Tunnel

To preempt AP’s which are not connected to preferred peer to the preferred peer

```go
PreemptSitesMxTunnel(
    ctx context.Context,
    siteId uuid.UUID,
    mxtunnelId uuid.UUID) (
    models.ApiResponse[models.ResponseMxtunnelsPreemptAps],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mxtunnelId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseMxtunnelsPreemptAps`](../../doc/models/response-mxtunnels-preempt-aps.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxtunnelId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesMxEdge.PreemptSitesMxTunnel(ctx, siteId, mxtunnelId)
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

