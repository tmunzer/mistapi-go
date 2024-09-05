# Orgs Stats-Devices

```go
orgsStatsDevices := client.OrgsStatsDevices()
```

## Class Name

`OrgsStatsDevices`


# List Org Devices Stats

Get List of Org Devices stats
This API renders some high-level device stats, pagination is assumed and returned in response header (as the response is an array)

```go
ListOrgDevicesStats(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeWithAllEnum,
    status *models.DeviceStatusEnum,
    siteId *string,
    mac *string,
    evpntopoId *string,
    evpnUnused *string,
    fields *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeWithAllEnum`](../../doc/models/device-type-with-all-enum.md) | Query, Optional | - |
| `status` | [`*models.DeviceStatusEnum`](../../doc/models/device-status-enum.md) | Query, Optional | - |
| `siteId` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `evpntopoId` | `*string` | Query, Optional | EVPN Topology ID |
| `evpnUnused` | `*string` | Query, Optional | if `evpn_unused`==`true`, find EVPN eligible switches which don’t belong to any EVPN Topology yet |
| `fields` | `*string` | Query, Optional | list of additional fields requests, comma separeted, or `fields=*` for all of them |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`[]models.StatsDevice`](../../doc/models/containers/stats-device.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeWithAllEnum("ap")

status := models.DeviceStatusEnum("all")









fields := "field1,field2"





duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsStatsDevices.ListOrgDevicesStats(ctx, orgId, &mType, &status, nil, nil, nil, nil, &fields, nil, nil, &duration, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsStatsAp(); ok {
            fmt.Println("Value narrowed down to models.StatsAp: ", *i)
        } else if i, ok := item.AsStatsSwitch(); ok {
            fmt.Println("Value narrowed down to models.StatsSwitch: ", *i)
        } else if i, ok := item.AsStatsGateway(); ok {
            fmt.Println("Value narrowed down to models.StatsGateway: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
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

