# Sites Devices-Wired

```go
sitesDevicesWired := client.SitesDevicesWired()
```

## Class Name

`SitesDevicesWired`

## Methods

* [Delete Site Local Switch Port Config](../../doc/controllers/sites-devices-wired.md#delete-site-local-switch-port-config)
* [Update Site Local Switch Port Config](../../doc/controllers/sites-devices-wired.md#update-site-local-switch-port-config)


# Delete Site Local Switch Port Config

Sometimes HelpDesk Admin needs to change port configs

```go
DeleteSiteLocalSwitchPortConfig(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesDevicesWired.DeleteSiteLocalSwitchPortConfig(ctx, siteId, deviceId)
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


# Update Site Local Switch Port Config

Sometimes HelpDesk Admin needs to change port configs

```go
UpdateSiteLocalSwitchPortConfig(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.JunosPortConfig) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.JunosPortConfig{
    AeDisableLacp:    models.ToPointer(true),
    AeIdx:            models.ToPointer(0),
    Aggregated:       models.ToPointer(false),
    Description:      models.ToPointer("string"),
    DisableAutoneg:   models.ToPointer(true),
    Duplex:           models.ToPointer(models.JunosPortConfigDuplexEnum("auto")),
    DynamicUsage:     models.NewOptional(models.ToPointer("string")),
    Esilag:           models.ToPointer(true),
    PoeDisabled:      models.ToPointer(true),
    Speed:            models.ToPointer(models.JunosPortConfigSpeedEnum("auto")),
    Usage:            "string",
}

resp, err := sitesDevicesWired.UpdateSiteLocalSwitchPortConfig(ctx, siteId, deviceId, &body)
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

