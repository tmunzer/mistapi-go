# Sites Stats

```go
sitesStats := client.SitesStats()
```

## Class Name

`SitesStats`


# Get Site Stats

Get Sites Stats

```go
GetSiteStats(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.StatsSite],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.StatsSite`](../../doc/models/stats-site.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStats.GetSiteStats(ctx, siteId)
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
  "address": "string",
  "alarmtemplate_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "country_code": "string",
  "created_time": 0,
  "id": "55c29ce5-7c0f-45b5-b99b-599f805fa3a1",
  "lat": 0,
  "latlng": {
    "lat": 0,
    "lng": 0
  },
  "lng": 0,
  "modified_time": 0,
  "msp_id": "dca3cad3-0c9b-439b-814f-8d5f23797972",
  "name": "string",
  "networktemplate_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "num_ap": 0,
  "num_ap_connected": 0,
  "num_clients": 0,
  "num_devices": 0,
  "num_devices_connected": 0,
  "num_gateway": 0,
  "num_gateway_connected": 0,
  "num_switch": 0,
  "num_switch_connected": 0,
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "rftemplate_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "sitegroup_ids": [
    "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
  ],
  "timezone": "string",
  "tzoffset": 0
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

