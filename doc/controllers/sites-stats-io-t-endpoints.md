# Sites Stats-Io T Endpoints

```go
sitesStatsIoTEndpoints := client.SitesStatsIoTEndpoints()
```

## Class Name

`SitesStatsIoTEndpoints`


# Search Site Iot Endpoints

Search IoT Endpoints

```go
SearchSiteIotEndpoints(
    ctx context.Context,
    siteId uuid.UUID,
    apMac *string,
    mac *string,
    mType *string,
    mfg *string,
    limit *int,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.ResponseIotEndpointsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `apMac` | `*string` | Query, Optional | AP MAC address |
| `mac` | `*string` | Query, Optional | Device MAC address |
| `mType` | `*string` | Query, Optional | IoT endpoint type. enum: `zigbee` |
| `mfg` | `*string` | Query, Optional | Manufacturer name |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseIotEndpointsSearch](../../doc/models/response-iot-endpoints-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apMac := "5c5b350e0001"

mac := "63f9e299182b63f9"

mType := "zigbee"

mfg := "Assa Abloy"

limit := 100

duration := "10m"

apiResponse, err := sitesStatsIoTEndpoints.SearchSiteIotEndpoints(ctx, siteId, &apMac, &mac, &mType, &mfg, &limit, nil, nil, &duration)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401Error:
            log.Fatalln("ResponseHttp401ErrorException: ", typedErr)
        case *errors.ResponseHttp403Error:
            log.Fatalln("ResponseHttp403ErrorException: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429Error:
            log.Fatalln("ResponseHttp429ErrorException: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "end": 1531862583.0,
  "results": [
    {
      "ap_mac": "5c5b350e0001",
      "id": "63f9e299182b63f9",
      "lqi": 178,
      "mac": "63f9e299182b63f9",
      "mfg": "Assa Abloy",
      "model": "Assa Abloy",
      "timestamp": 1531782218,
      "type": "zigbee"
    }
  ],
  "start": 1531776183.0,
  "total": 2
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

