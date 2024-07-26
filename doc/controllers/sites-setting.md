# Sites Setting

```go
sitesSetting := client.SitesSetting()
```

## Class Name

`SitesSetting`

## Methods

* [Create Site Watched Stations](../../doc/controllers/sites-setting.md#create-site-watched-stations)
* [Create Site Wireless Clients Allowlist](../../doc/controllers/sites-setting.md#create-site-wireless-clients-allowlist)
* [Create Site Wireless Clients Blocklist](../../doc/controllers/sites-setting.md#create-site-wireless-clients-blocklist)
* [Delete Site Watched Stations](../../doc/controllers/sites-setting.md#delete-site-watched-stations)
* [Delete Site Wireless Clients Allowlist](../../doc/controllers/sites-setting.md#delete-site-wireless-clients-allowlist)
* [Delete Site Wireless Clients Blocklist](../../doc/controllers/sites-setting.md#delete-site-wireless-clients-blocklist)
* [Get Site Setting](../../doc/controllers/sites-setting.md#get-site-setting)
* [Update Site Settings](../../doc/controllers/sites-setting.md#update-site-settings)


# Create Site Watched Stations

This endpoint is to provide list of client macs for annotation as watched station.

Retrieve the current clients list from `watched_station_url` under Site:Setting

```go
CreateSiteWatchedStations(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "18-65-90-de-f4-c6",
        "84-89-ad-5d-69-0d",
    },
}

apiResponse, err := sitesSetting.CreateSiteWatchedStations(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
  ]
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


# Create Site Wireless Clients Allowlist

This endpoint is to provide list of client macs for annotation as whitelist.

Retrieve the current clients list from `whitelist_url` under Site:Setting

```go
CreateSiteWirelessClientsAllowlist(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "683b679ac024",
    },
}

apiResponse, err := sitesSetting.CreateSiteWirelessClientsAllowlist(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
  ]
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


# Create Site Wireless Clients Blocklist

This endpoint is to provide list of client macs for annotation blacklist.

Retrieve the current clients list `blacklist_url` under Site:Setting

```go
CreateSiteWirelessClientsBlocklist(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "18-65-90-de-f4-c6",
        "84-89-ad-5d-69-0d",
    },
}

apiResponse, err := sitesSetting.CreateSiteWirelessClientsBlocklist(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
  ]
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


# Delete Site Watched Stations

Delete Site Watched Station Clients

```go
DeleteSiteWatchedStations(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWatchedStations(ctx, siteId)
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


# Delete Site Wireless Clients Allowlist

Delete Site Whitelist Station Clients

```go
DeleteSiteWirelessClientsAllowlist(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWirelessClientsAllowlist(ctx, siteId)
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


# Delete Site Wireless Clients Blocklist

Delete Site Blacklist Station Clients

```go
DeleteSiteWirelessClientsBlocklist(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWirelessClientsBlocklist(ctx, siteId)
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


# Get Site Setting

Get Site Settings

```go
GetSiteSetting(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.SiteSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SiteSetting`](../../doc/models/site-setting.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesSetting.GetSiteSetting(ctx, siteId)
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


# Update Site Settings

Update Site Settings

```go
UpdateSiteSettings(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.SiteSetting) (
    models.ApiResponse[models.SiteSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SiteSetting`](../../doc/models/site-setting.md) | Body, Optional | Request Body |

## Response Type

[`models.SiteSetting`](../../doc/models/site-setting.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SiteSetting{
    BlacklistUrl:                    models.ToPointer("https://papi.s3.amazonaws.com/blacklist/xxx..."),
    ConfigAutoRevert:                models.ToPointer(false),
    DeviceUpdownThreshold:           models.ToPointer(0),
    ExtraRoutes6:                    map[string]models.ExtraRoute6{
        "2a02:1234:420a:10c9::/64": models.ExtraRoute6{
            Via:           models.ToPointer("2a02:1234:200a::100"),
        },
    },
    GatewayUpdownThreshold:          models.NewOptional(models.ToPointer(10)),
    OrgId:                           models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    PersistConfigOnDevice:           models.ToPointer(false),
    ReportGatt:                      models.ToPointer(false),
    SiteId:                          models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    SwitchUpdownThreshold:           models.NewOptional(models.ToPointer(0)),
    TrackAnonymousDevices:           models.ToPointer(false),
    TuntermMonitoringDisabled:       models.ToPointer(false),
    Vars:                            map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    VrfInstances:                    map[string]models.VrfInstance{
        "guest": models.VrfInstance{
            ExtraRoutes: map[string]models.VrfExtraRoute{
                "0.0.0.0/0": models.VrfExtraRoute{
                    Via: models.ToPointer("192.168.31.1"),
                },
            },
            Networks:    []string{
                "guest",
            },
        },
    },
    WatchedStationUrl:               models.ToPointer("https://papi.s3.amazonaws.com/watched_station/xxx..."),
    WhitelistUrl:                    models.ToPointer("https://papi.s3.amazonaws.com/whitelist/xxx..."),
}

apiResponse, err := sitesSetting.UpdateSiteSettings(ctx, siteId, &body)
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

