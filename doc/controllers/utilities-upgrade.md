# Utilities Upgrade

```go
utilitiesUpgrade := client.UtilitiesUpgrade()
```

## Class Name

`UtilitiesUpgrade`

## Methods

* [Cancel Org Device Upgrade](../../doc/controllers/utilities-upgrade.md#cancel-org-device-upgrade)
* [Cancel Org Ssr Upgrade](../../doc/controllers/utilities-upgrade.md#cancel-org-ssr-upgrade)
* [Cancel Site Device Upgrade](../../doc/controllers/utilities-upgrade.md#cancel-site-device-upgrade)
* [Get Org Device Upgrade](../../doc/controllers/utilities-upgrade.md#get-org-device-upgrade)
* [Get Org Mx Edge Upgrade](../../doc/controllers/utilities-upgrade.md#get-org-mx-edge-upgrade)
* [Get Org Ssr Upgrade](../../doc/controllers/utilities-upgrade.md#get-org-ssr-upgrade)
* [Get Site Device Upgrade](../../doc/controllers/utilities-upgrade.md#get-site-device-upgrade)
* [Get Site Ssr Upgrade](../../doc/controllers/utilities-upgrade.md#get-site-ssr-upgrade)
* [List Org Available Device Versions](../../doc/controllers/utilities-upgrade.md#list-org-available-device-versions)
* [List Org Available Ssr Versions](../../doc/controllers/utilities-upgrade.md#list-org-available-ssr-versions)
* [List Org Device Upgrades](../../doc/controllers/utilities-upgrade.md#list-org-device-upgrades)
* [List Org Mx Edge Upgrades](../../doc/controllers/utilities-upgrade.md#list-org-mx-edge-upgrades)
* [List Org Ssr Upgrades](../../doc/controllers/utilities-upgrade.md#list-org-ssr-upgrades)
* [List Site Available Device Versions](../../doc/controllers/utilities-upgrade.md#list-site-available-device-versions)
* [List Site Device Upgrades](../../doc/controllers/utilities-upgrade.md#list-site-device-upgrades)
* [Upgrade Device](../../doc/controllers/utilities-upgrade.md#upgrade-device)
* [Upgrade Org Devices](../../doc/controllers/utilities-upgrade.md#upgrade-org-devices)
* [Upgrade Org Jsi Device](../../doc/controllers/utilities-upgrade.md#upgrade-org-jsi-device)
* [Upgrade Org Mx Edges](../../doc/controllers/utilities-upgrade.md#upgrade-org-mx-edges)
* [Upgrade Org Ssrs](../../doc/controllers/utilities-upgrade.md#upgrade-org-ssrs)
* [Upgrade Site Devices](../../doc/controllers/utilities-upgrade.md#upgrade-site-devices)
* [Upgrade Ssr](../../doc/controllers/utilities-upgrade.md#upgrade-ssr)


# Cancel Org Device Upgrade

Best effort to cancel an upgrade. Devices which are already upgraded wont be touched

```go
CancelOrgDeviceUpgrade(
    ctx context.Context,
    orgId uuid.UUID,
    upgradeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `upgradeId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upgradeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := utilitiesUpgrade.CancelOrgDeviceUpgrade(ctx, orgId, upgradeId)
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


# Cancel Org Ssr Upgrade

Best effort to cancel an upgrade. Devices which are already upgraded wont be touched↵

```go
CancelOrgSsrUpgrade(
    ctx context.Context,
    orgId uuid.UUID,
    upgradeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `upgradeId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upgradeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := utilitiesUpgrade.CancelOrgSsrUpgrade(ctx, orgId, upgradeId)
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


# Cancel Site Device Upgrade

Best effort to cancel an upgrade. Devices which are already upgraded wont be touched

```go
CancelSiteDeviceUpgrade(
    ctx context.Context,
    siteId uuid.UUID,
    upgradeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `upgradeId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upgradeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := utilitiesUpgrade.CancelSiteDeviceUpgrade(ctx, siteId, upgradeId)
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


# Get Org Device Upgrade

Get Multiple Devices Upgrade

```go
GetOrgDeviceUpgrade(
    ctx context.Context,
    orgId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseUpgradeOrgDevices],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `upgradeId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseUpgradeOrgDevices](../../doc/models/response-upgrade-org-devices.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upgradeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.GetOrgDeviceUpgrade(ctx, orgId, upgradeId)
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
  "enable_p2p": true,
  "force": true,
  "id": "31223085-405d-4b64-8aea-9c5b98098b4b",
  "strategy": "big_bang",
  "target_version": "0.14.29411",
  "upgrades": [
    {
      "site_id": "1bbe6e79-2583-403c-be1a-9881b4691ab6",
      "upgrade": {
        "id": "473f6eca-6276-4993-bfeb-53cbbbba6f18",
        "start_time": 1717658765,
        "status": "upgrading",
        "targets": {
          "download_requested": [
            "5c5b3550bd2e"
          ],
          "downloaded": [
            "003e7316ff9e"
          ],
          "total": 4,
          "upgraded": [
            "5c5b3550bd2e",
            "003e7316ff9e"
          ]
        }
      }
    }
  ]
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


# Get Org Mx Edge Upgrade

Get Mist Edge Upgrade

```go
GetOrgMxEdgeUpgrade(
    ctx context.Context,
    orgId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseMxedgeUpgrade],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `upgradeId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseMxedgeUpgrade](../../doc/models/response-mxedge-upgrade.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upgradeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.GetOrgMxEdgeUpgrade(ctx, orgId, upgradeId)
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


# Get Org Ssr Upgrade

Get Specific Org SSR Upgrade

```go
GetOrgSsrUpgrade(
    ctx context.Context,
    orgId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseSsrUpgradeStatus],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `upgradeId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSsrUpgradeStatus](../../doc/models/response-ssr-upgrade-status.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upgradeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.GetOrgSsrUpgrade(ctx, orgId, upgradeId)
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
  "channel": "stable",
  "device_type": "gateway",
  "id": "5cbcee0a-c620-4bb4-a25e-15000934e9d8",
  "status": "upgrading",
  "targets": {
    "failed": [],
    "queued": [],
    "success": [],
    "upgrading": [
      "8e525f1d-4178-4ae1-a988-2b0176855e55"
    ]
  },
  "versions": {}
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


# Get Site Device Upgrade

Get Site Device Upgrade

```go
GetSiteDeviceUpgrade(
    ctx context.Context,
    siteId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseSiteDeviceUpgrade],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `upgradeId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSiteDeviceUpgrade](../../doc/models/response-site-device-upgrade.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upgradeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.GetSiteDeviceUpgrade(ctx, siteId, upgradeId)
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
  "canary_phases": [
    1,
    10,
    50,
    100
  ],
  "current_phase": 0,
  "enable_p2p": true,
  "force": false,
  "id": "b910de88-0e0e-43a0-b20e-a43aa8788917",
  "max_failure_percentage": 5,
  "max_failures": [
    5,
    5,
    5,
    5
  ],
  "p2p_cluster_size": 10,
  "p2p_parallelism": 2,
  "reboot_at": 1738233315,
  "start_time": 1738232698,
  "status": "created",
  "strategy": "canary",
  "target_version": "0.14.29633",
  "targets": {
    "downloaded": [
      "30783a619182"
    ],
    "total": 2,
    "upgraded": [
      "a8378280da5e"
    ]
  },
  "upgrade_plan": {
    "1": [
      "30783a619182"
    ],
    "2": [
      "a8378280da5e"
    ]
  }
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


# Get Site Ssr Upgrade

Get Specific Site SSR Upgrade

```go
GetSiteSsrUpgrade(
    ctx context.Context,
    siteId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseSsrUpgradeStatus],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `upgradeId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSsrUpgradeStatus](../../doc/models/response-ssr-upgrade-status.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

upgradeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.GetSiteSsrUpgrade(ctx, siteId, upgradeId)
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
  "channel": "stable",
  "device_type": "gateway",
  "id": "5cbcee0a-c620-4bb4-a25e-15000934e9d8",
  "status": "upgrading",
  "targets": {
    "failed": [],
    "queued": [],
    "success": [],
    "upgrading": [
      "8e525f1d-4178-4ae1-a988-2b0176855e55"
    ]
  },
  "versions": {}
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


# List Org Available Device Versions

Get List of Available Device Versions

```go
ListOrgAvailableDeviceVersions(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeDefaultApEnum,
    model *string) (
    models.ApiResponse[[]models.DeviceVersionItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | **Default**: `"ap"` |
| `model` | `*string` | Query, Optional | Fetch version for device model, use/combine with `type` as needed (for switch and gateway devices) |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.DeviceVersionItem](../../doc/models/device-version-item.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeDefaultApEnum_AP

model := "AP43"

apiResponse, err := utilitiesUpgrade.ListOrgAvailableDeviceVersions(ctx, orgId, &mType, &model)
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
[
  {
    "model": "AP41",
    "tag": "stable",
    "version": "v0.1.543"
  },
  {
    "model": "AP21",
    "version": "v0.1.545"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Available Ssr Versions

Get available version for SSR

```go
ListOrgAvailableSsrVersions(
    ctx context.Context,
    orgId uuid.UUID,
    channel *models.SsrVersionChannelEnum,
    mac *string) (
    models.ApiResponse[[]models.SsrVersion],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `channel` | [`*models.SsrVersionChannelEnum`](../../doc/models/ssr-version-channel-enum.md) | Query, Optional | SSR version channel<br><br>**Default**: `"stable"` |
| `mac` | `*string` | Query, Optional | Optional. MAC address, or comma separated MAC address list. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.SsrVersion](../../doc/models/ssr-version.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

channel := models.SsrVersionChannelEnum_STABLE

apiResponse, err := utilitiesUpgrade.ListOrgAvailableSsrVersions(ctx, orgId, &channel, nil)
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
[
  {
    "default": true,
    "package": "SSR",
    "version": "5.3.1-17"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Device Upgrades

Get List of Org multiple devices upgrades

```go
ListOrgDeviceUpgrades(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.UpgradeOrgDevicesItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.UpgradeOrgDevicesItem](../../doc/models/upgrade-org-devices-item.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.ListOrgDeviceUpgrades(ctx, orgId)
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
[
  {
    "id": "466f6eca-6276-4993-bfeb-53cbbbba6f88",
    "site_upgrades": [
      {
        "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
        "upgrade_id": "174bda0-06a3-40ee-b918-d9cbde303690"
      }
    ]
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Mx Edge Upgrades

Get List of Org Mist Edge Upgrades

```go
ListOrgMxEdgeUpgrades(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.ResponseMxedgeUpgrade],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ResponseMxedgeUpgrade](../../doc/models/response-mxedge-upgrade.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.ListOrgMxEdgeUpgrades(ctx, orgId)
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


# List Org Ssr Upgrades

Get List of Org SSR Upgrades

```go
ListOrgSsrUpgrades(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.ResponseSsrUpgrade],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ResponseSsrUpgrade](../../doc/models/response-ssr-upgrade.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.ListOrgSsrUpgrades(ctx, orgId)
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
[
  {
    "channel": "stable",
    "counts": {
      "failed": 0,
      "queued": 1,
      "success": 0,
      "upgrading": 1
    },
    "device_type": "gateway",
    "id": "ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae",
    "status": "upgrading",
    "strategy": "serial",
    "versions": {}
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Site Available Device Versions

Get List of Available Device Versions

```go
ListSiteAvailableDeviceVersions(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeDefaultApEnum,
    model *string) (
    models.ApiResponse[[]models.DeviceVersionItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | **Default**: `"ap"` |
| `model` | `*string` | Query, Optional | Fetch version for device model, use/combine with `type` as needed (for switch and gateway devices) |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.DeviceVersionItem](../../doc/models/device-version-item.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeDefaultApEnum_AP

apiResponse, err := utilitiesUpgrade.ListSiteAvailableDeviceVersions(ctx, siteId, &mType, nil)
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
[
  {
    "model": "AP41",
    "tag": "stable",
    "version": "v0.1.543"
  },
  {
    "model": "AP21",
    "version": "v0.1.545"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Site Device Upgrades

Get all upgrades for site

```go
ListSiteDeviceUpgrades(
    ctx context.Context,
    siteId uuid.UUID,
    status *models.UpgradeDeviceStatusEnum) (
    models.ApiResponse[[]models.ResponseSiteDeviceUpgradesItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `status` | [`*models.UpgradeDeviceStatusEnum`](../../doc/models/upgrade-device-status-enum.md) | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ResponseSiteDeviceUpgradesItem](../../doc/models/response-site-device-upgrades-item.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesUpgrade.ListSiteDeviceUpgrades(ctx, siteId, nil)
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
[
  {
    "counts": {
      "download_requested": 0,
      "downloaded": 0,
      "failed": 0,
      "reboot_in_progress": 0,
      "rebooted": 0,
      "skipped": 0,
      "total": 0
    },
    "enable_p2p": true,
    "force": true,
    "id": "472f6eca-6276-4993-bfeb-53cbbbba6f28",
    "start_time": 0,
    "status": "created",
    "strategy": "big_bang",
    "target_version": "string"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Upgrade Device

Device Upgrade

```go
UpgradeDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.DeviceUpgrade) (
    models.ApiResponse[models.ResponseDeviceUpgrade],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.DeviceUpgrade`](../../doc/models/device-upgrade.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceUpgrade](../../doc/models/response-device-upgrade.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.DeviceUpgrade{
    Version:              "3.1.5",
}

apiResponse, err := utilitiesUpgrade.UpgradeDevice(ctx, siteId, deviceId, &body)
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
  "status": "inprogress",
  "timestamp": 1428949501
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


# Upgrade Org Devices

Upgrade Multiple Sites (Only supported for Access Points upgrades)

```go
UpgradeOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.UpgradeOrgDevices) (
    models.ApiResponse[models.ResponseUpgradeOrgDevices],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UpgradeOrgDevices`](../../doc/models/upgrade-org-devices.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseUpgradeOrgDevices](../../doc/models/response-upgrade-org-devices.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UpgradeOrgDevices{
    AllSites:                models.ToPointer(false),
    CanaryPhases:            nil,
    DownloadStrategy:        models.ToPointer(models.UpgradeOrgDevicesDownloadStrategyEnum_BIGBANG),
    MaxFailurePercentage:    models.ToPointer(5),
    P2pClusterSize:          models.ToPointer(0),
    RebootAt:                models.ToPointer(1624399840),
    RebootDatetime:          models.ToPointer("2024-06-13 15:00:00-07:00"),
    RebootStrategy:          models.ToPointer(models.UpgradeDeviceStrategyEnum_BIGBANG),
    RrmFirstBatchPercentage: models.ToPointer(2),
    RrmMaxBatchPercentage:   models.ToPointer(10),
    RrmMeshUpgrade:          models.ToPointer(models.UpgradeDeviceRrmMeshUpgradeEnum_SEQUENTIAL),
    RrmNodeOrder:            models.ToPointer(models.UpgradeDeviceRrmNodeOrderEnum_FRINGETOCENTER),
    Rules:                   []map[string]string{
        []map[string]string{
            []map[string]string{
            },
            []map[string]string{
            },
        },
        []map[string]string{
            []map[string]string{
            },
        },
    },
    Snapshot:                models.ToPointer(false),
    StartDatetime:           models.ToPointer("2024-06-13 15:00:00-07:00"),
    StartTime:               models.ToPointer(1624399840),
    Strategy:                models.ToPointer(models.UpgradeDeviceStrategyEnum_BIGBANG),
}

apiResponse, err := utilitiesUpgrade.UpgradeOrgDevices(ctx, orgId, &body)
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
  "enable_p2p": true,
  "force": true,
  "id": "31223085-405d-4b64-8aea-9c5b98098b4b",
  "strategy": "big_bang",
  "target_version": "0.14.29411",
  "upgrades": [
    {
      "site_id": "1bbe6e79-2583-403c-be1a-9881b4691ab6",
      "upgrade": {
        "id": "473f6eca-6276-4993-bfeb-53cbbbba6f18",
        "start_time": 1717658765,
        "status": "upgrading",
        "targets": {
          "download_requested": [
            "5c5b3550bd2e"
          ],
          "downloaded": [
            "003e7316ff9e"
          ],
          "total": 4,
          "upgraded": [
            "5c5b3550bd2e",
            "003e7316ff9e"
          ]
        }
      }
    }
  ]
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


# Upgrade Org Jsi Device

Upgrade

```go
UpgradeOrgJsiDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string,
    body *models.VersionString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `body` | [`*models.VersionString`](../../doc/models/version-string.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

body := models.VersionString{
    Version:              models.ToPointer("3.1.5"),
}

resp, err := utilitiesUpgrade.UpgradeOrgJsiDevice(ctx, orgId, deviceMac, &body)
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


# Upgrade Org Mx Edges

Upgrade Mist Edges

```go
UpgradeOrgMxEdges(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MxedgeUpgradeMulti) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MxedgeUpgradeMulti`](../../doc/models/mxedge-upgrade-multi.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MxedgeUpgradeMulti{
    AllowDowngrades:      models.ToPointer(models.MxedgeUpgradeMultiAllowDowngrades{
        Mxagent:              models.ToPointer(false),
        Tunterm:              models.ToPointer(true),
    }),
    CanaryPhases:         []int{
        5,
        25,
        50,
        100,
    },
    Channel:              models.ToPointer(models.MxedgeUpgradeChannelEnum_STABLE),
    MaxFailurePercentage: models.ToPointer(5),
    MxedgeIds:            []uuid.UUID{
        uuid.MustParse("387804a7-3474-85ce-15a2-f9a9684c9c90"),
    },
    StartTime:            models.ToPointer(1624399840),
    Strategy:             models.ToPointer(models.MxedgeUpgradeStrategyEnum_SERIAL),
    Versions:             models.ToPointer(models.MxedgeUpgradeVersion{
        Mxagent:              "latest",
        Radsecproxy:          models.ToPointer("<version>"),
        Tunterm:              "default",
    }),
    AdditionalProperties: map[string]interface{}{
        "max_failures": interface{}("1155"),
    },
}

resp, err := utilitiesUpgrade.UpgradeOrgMxEdges(ctx, orgId, &body)
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


# Upgrade Org Ssrs

Upgrade Org SSRs

```go
UpgradeOrgSsrs(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SsrUpgradeMulti) (
    models.ApiResponse[models.ResponseSsrUpgrade],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SsrUpgradeMulti`](../../doc/models/ssr-upgrade-multi.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSsrUpgrade](../../doc/models/response-ssr-upgrade.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsrUpgradeMulti{
    Channel:              models.ToPointer(models.SsrUpgradeChannelEnum_STABLE),
    DeviceIds:            []uuid.UUID{
        uuid.MustParse("00000000-0000-0000-1000-5c5b3500001f"),
        uuid.MustParse("00000000-0000-0000-1000-5c5b35000020"),
    },
    Strategy:             models.ToPointer(models.SsrUpgradeStrategyEnum_BIGBANG),
    Version:              models.ToPointer("5.3.0-93"),
}

apiResponse, err := utilitiesUpgrade.UpgradeOrgSsrs(ctx, orgId, &body)
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
  "channel": "stable",
  "counts": {
    "failed": 0,
    "queued": 1,
    "success": 0,
    "upgrading": 1
  },
  "device_type": "gateway",
  "id": "ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae",
  "status": "upgrading",
  "strategy": "serial",
  "versions": {}
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


# Upgrade Site Devices

Upgrade Site Device

**Note**: this call doesn’t guarantee the devices to be upgraded right away (they may be offline)

```go
UpgradeSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UpgradeSiteDevices) (
    models.ApiResponse[models.ResponseUpgradeId],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UpgradeSiteDevices`](../../doc/models/upgrade-site-devices.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseUpgradeId](../../doc/models/response-upgrade-id.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UpgradeSiteDevices{
    CanaryPhases:            nil,
    Force:                   models.ToPointer(false),
    MaxFailurePercentage:    models.ToPointer(5),
    P2pClusterSize:          models.ToPointer(0),
    Reboot:                  models.ToPointer(false),
    RebootAt:                models.ToPointer(1624399840),
    RrmFirstBatchPercentage: models.ToPointer(2),
    RrmMaxBatchPercentage:   models.ToPointer(10),
    RrmMeshUpgrade:          models.ToPointer(models.UpgradeDeviceRrmMeshUpgradeEnum_SEQUENTIAL),
    RrmNodeOrder:            models.ToPointer(models.UpgradeDeviceRrmNodeOrderEnum_FRINGETOCENTER),
    Rules:                   []map[string]string{
        []map[string]string{
            []map[string]string{
            },
            []map[string]string{
            },
        },
        []map[string]string{
            []map[string]string{
            },
        },
    },
    Snapshot:                models.ToPointer(false),
    StartTime:               models.ToPointer(1624399840),
    Strategy:                models.ToPointer(models.UpgradeDeviceStrategyEnum_BIGBANG),
    Version:                 models.ToPointer("3.1.5"),
}

apiResponse, err := utilitiesUpgrade.UpgradeSiteDevices(ctx, siteId, &body)
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
  "upgrade_id": "4316c116-0acb-4c43-8f06-6723154e741e"
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


# Upgrade Ssr

Upgrade Site SSR device

```go
UpgradeSsr(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.SsrUpgrade) (
    models.ApiResponse[models.ResponseSsrUpgrade],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SsrUpgrade`](../../doc/models/ssr-upgrade.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSsrUpgrade](../../doc/models/response-ssr-upgrade.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsrUpgrade{
    Channel:              models.ToPointer(models.SsrUpgradeChannelEnum_STABLE),
    Version:              "5.3.1-170-93",
}

apiResponse, err := utilitiesUpgrade.UpgradeSsr(ctx, siteId, deviceId, &body)
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
  "channel": "stable",
  "counts": {
    "failed": 0,
    "queued": 1,
    "success": 0,
    "upgrading": 1
  },
  "device_type": "gateway",
  "id": "ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae",
  "status": "upgrading",
  "strategy": "serial",
  "versions": {}
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

