# Orgs Stats-Other Devices

```go
orgsStatsOtherDevices := client.OrgsStatsOtherDevices()
```

## Class Name

`OrgsStatsOtherDevices`


# Get Org Other Device Stats

Get Otherdevice Stats

```go
GetOrgOtherDeviceStats(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.StatsDeviceOther],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | - |

## Response Type

[`models.StatsDeviceOther`](../../doc/models/stats-device-other.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

apiResponse, err := orgsStatsOtherDevices.GetOrgOtherDeviceStats(ctx, orgId, deviceMac)
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
  "config_status": "synced",
  "last_config": 1675392788,
  "last_seen": 1675843629,
  "mac": "5c5b35000018",
  "status": "online",
  "uptime": 20296,
  "vendor": "cradlepoint",
  "vendor_specific": {
    "ports": {
      "mdm-4d0e073b": {
        "bytes_in": 33004879,
        "bytes_out": 41103393,
        "health_category": "",
        "health_score": 0,
        "id": "101027967",
        "mode": "wan",
        "model": "Internal 5GB (SIM1)",
        "state": "READY",
        "type": "5G",
        "uptime": 252371.34149021498
      }
    },
    "router_id": null,
    "target_version": "7.23.40"
  },
  "version": "7.22.70"
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

