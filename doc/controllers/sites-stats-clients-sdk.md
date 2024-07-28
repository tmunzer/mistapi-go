# Sites Stats-Clients SDK

```go
sitesStatsClientsSDK := client.SitesStatsClientsSDK()
```

## Class Name

`SitesStatsClientsSDK`

## Methods

* [Get Site Sdk Stats](../../doc/controllers/sites-stats-clients-sdk.md#get-site-sdk-stats)
* [Get Site Sdk Stats by Map](../../doc/controllers/sites-stats-clients-sdk.md#get-site-sdk-stats-by-map)


# Get Site Sdk Stats

Get Detail Stats of a SdkClient

```go
GetSiteSdkStats(
    ctx context.Context,
    siteId uuid.UUID,
    sdkclientId uuid.UUID) (
    models.ApiResponse[models.SdkstatsWirelessClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `sdkclientId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SdkstatsWirelessClient`](../../doc/models/sdkstats-wireless-client.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sdkclientId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsClientsSDK.GetSiteSdkStats(ctx, siteId, sdkclientId)
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
  "id": "d56bd5fa-0a0a-4861-a9df-5ac83d3a2eeb",
  "last_seen": 1428939600,
  "name": "John's iPhone",
  "network_connection": {
    "mac": "c3-b6-e5-af-41-15",
    "rssi": -75,
    "signal_level": 3,
    "type": "WiFi"
  },
  "uuid": "ada72f8f-1643-e5c6-94db-f2a5636f1a64",
  "vbeacons": [
    {
      "id": "d379d29d-24b4-96c5-5dd4-6f2a2dc5aaeb",
      "since": 1428939300
    }
  ],
  "x": 60,
  "y": 80,
  "zones": [
    {
      "id": "8ac84899-32db-6327-334c-9b6d58544cfe",
      "since": 1428939600
    }
  ]
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


# Get Site Sdk Stats by Map

Get SdkClient Stats By Map

```go
GetSiteSdkStatsByMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    models.ApiResponse[[]models.SdkclientStat],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.SdkclientStat`](../../doc/models/sdkclient-stat.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsClientsSDK.GetSiteSdkStatsByMap(ctx, siteId, mapId)
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

