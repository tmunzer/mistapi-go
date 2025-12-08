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

API Calls delete all the existing port config local overrides, and reapply the configured planed at the device level
(with site / template heritance).

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Site Local Switch Port Config

API Calls to add port config local overrides. This can be used by Switch Port Operators or Helpdesk administrators
to change a Switch Port configuration without having to change the switch configuration.

The local overrides configured for the switchports with `no_local_overwrite`==`true` won't be applied to the switch configuration.

> NOTE:
> 
> When using the API Call, it is required to put send all overrides in the PUT request Payload, even the existing once.
> 
> The current overrides can be retrieved with the API Call [Get Site Device](../../doc/controllers/sites-devices.md#get-site-device). The local overrides will show up separately from the `port_config` in the `local_port_config` so it can be easily identified (and cleared)

```go
UpdateSiteLocalSwitchPortConfig(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body map[string]models.JunosLocalPortConfig) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`map[string]models.JunosLocalPortConfig`](../../doc/models/junos-local-port-config.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := map[string]models.JunosLocalPortConfig{
    "ge-0/0/0-1": models.JunosLocalPortConfig{
        PoeDisabled:                              models.ToPointer(true),
        Usage:                                    "iot",
    },
}

resp, err := sitesDevicesWired.UpdateSiteLocalSwitchPortConfig(ctx, siteId, deviceId, body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

