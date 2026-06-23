# Utilities Location

```go
utilitiesLocation := client.UtilitiesLocation()
```

## Class Name

`UtilitiesLocation`


# Send Site Devices Arbitrary Ble Beacon

Send arbitrary BLE Beacon for a period of time

Note that only the devices that are connected will be restarted.

```go
SendSiteDevicesArbitraryBleBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UtilsSendBleBeacon) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsSendBleBeacon`](../../doc/models/utils-send-ble-beacon.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsSendBleBeacon{
    BeaconFrame:          models.ToPointer("68b329da9893e34099c7d8ad5cb9c940"),
    BeaconFreq:           models.ToPointer(100),
    Duration:             models.ToPointer(10),
    Macs:                 []string{
        "5c5b35584a6f",
        "5c5b350ea3b3",
    },
    MapIds:               []string{
        "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
    },
}

resp, err := utilitiesLocation.SendSiteDevicesArbitraryBleBeacon(ctx, siteId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

