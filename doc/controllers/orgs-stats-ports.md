# Orgs Stats-Ports

```go
orgsStatsPorts := client.OrgsStatsPorts()
```

## Class Name

`OrgsStatsPorts`


# Search Org Sw or Gw Ports

Search Switch / Gateway Ports

```go
SearchOrgSwOrGwPorts(
    ctx context.Context,
    orgId uuid.UUID,
    fullDuplex *bool,
    mac *string,
    neighborMac *string,
    neighborPortDesc *string,
    neighborSystemName *string,
    poeDisabled *bool,
    poeMode *string,
    poeOn *bool,
    portId *string,
    portMac *string,
    powerDraw *float64,
    txPkts *int,
    rxPkts *int,
    rxBytes *int,
    txBps *int,
    rxBps *int,
    txErrors *int,
    rxErrors *int,
    txMcastPkts *int,
    txBcastPkts *int,
    rxMcastPkts *int,
    rxBcastPkts *int,
    speed *int,
    macLimit *int,
    macCount *int,
    up *bool,
    stpState *models.SearchOrgSwOrGwPortsStpStateEnum,
    stpRole *models.SearchOrgSwOrGwPortsStpRoleEnum,
    authState *models.SearchOrgSwOrGwPortsAuthStateEnum,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponsePortStatsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `fullDuplex` | `*bool` | Query, Optional | Indicates full or half duplex |
| `mac` | `*string` | Query, Optional | Device identifier |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” |
| `poeDisabled` | `*bool` | Query, Optional | Is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | POE mode depending on class E.g. “802.3at” |
| `poeOn` | `*bool` | Query, Optional | Is the device attached to POE |
| `portId` | `*string` | Query, Optional | Interface name |
| `portMac` | `*string` | Query, Optional | Interface mac address |
| `powerDraw` | `*float64` | Query, Optional | Amount of power being used by the interface at the time the command is executed. Unit in watts. |
| `txPkts` | `*int` | Query, Optional | Output packets |
| `rxPkts` | `*int` | Query, Optional | Input packets |
| `rxBytes` | `*int` | Query, Optional | Input bytes |
| `txBps` | `*int` | Query, Optional | Output rate |
| `rxBps` | `*int` | Query, Optional | Input rate |
| `txErrors` | `*int` | Query, Optional | Output errors |
| `rxErrors` | `*int` | Query, Optional | Input errors |
| `txMcastPkts` | `*int` | Query, Optional | Multicast output packets |
| `txBcastPkts` | `*int` | Query, Optional | Broadcast output packets |
| `rxMcastPkts` | `*int` | Query, Optional | Multicast input packets |
| `rxBcastPkts` | `*int` | Query, Optional | Broadcast input packets |
| `speed` | `*int` | Query, Optional | Port speed |
| `macLimit` | `*int` | Query, Optional | Limit on number of dynamically learned macs |
| `macCount` | `*int` | Query, Optional | Number of mac addresses in the forwarding table |
| `up` | `*bool` | Query, Optional | Indicates if interface is up |
| `stpState` | [`*models.SearchOrgSwOrGwPortsStpStateEnum`](../../doc/models/search-org-sw-or-gw-ports-stp-state-enum.md) | Query, Optional | If `up`==`true` |
| `stpRole` | [`*models.SearchOrgSwOrGwPortsStpRoleEnum`](../../doc/models/search-org-sw-or-gw-ports-stp-role-enum.md) | Query, Optional | If `up`==`true` |
| `authState` | [`*models.SearchOrgSwOrGwPortsAuthStateEnum`](../../doc/models/search-org-sw-or-gw-ports-auth-state-enum.md) | Query, Optional | If `up`==`true` && has Authenticator role |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

[`models.ResponsePortStatsSearch`](../../doc/models/response-port-stats-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



























































limit := 100





duration := "10m"

apiResponse, err := orgsStatsPorts.SearchOrgSwOrGwPorts(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

