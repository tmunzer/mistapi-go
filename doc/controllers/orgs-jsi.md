# Orgs JSI

```go
orgsJSI := client.OrgsJSI()
```

## Class Name

`OrgsJSI`

## Methods

* [Adopt Org Jsi Device](../../doc/controllers/orgs-jsi.md#adopt-org-jsi-device)
* [Create Org Jsi Device Shell Session](../../doc/controllers/orgs-jsi.md#create-org-jsi-device-shell-session)
* [List Org Jsi Devices](../../doc/controllers/orgs-jsi.md#list-org-jsi-devices)
* [List Org Jsi Past Purchases](../../doc/controllers/orgs-jsi.md#list-org-jsi-past-purchases)


# Adopt Org Jsi Device

Adopt JSI devices

```go
AdoptOrgJsiDevice(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseDeviceConfigCmd],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseDeviceConfigCmd`](../../doc/models/response-device-config-cmd.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsJSI.AdoptOrgJsiDevice(ctx, orgId)
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
  "cmd": "set system services ssh...\n...\nset system services outbound-ssh client mist ..."
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


# Create Org Jsi Device Shell Session

Create Shell Session

```go
CreateOrgJsiDeviceShellSession(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | - |

## Response Type

[`models.WebsocketSessionWithUrl`](../../doc/models/websocket-session-with-url.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

apiResponse, err := orgsJSI.CreateOrgJsiDeviceShellSession(ctx, orgId, deviceMac)
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


# List Org Jsi Devices

Get List of Org devices that connected to JSI

```go
ListOrgJsiDevices(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int,
    model *string,
    serial *string,
    mac *string) (
    models.ApiResponse[[]models.JseDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `model` | `*string` | Query, Optional | Device model |
| `serial` | `*string` | Query, Optional | Device serial |
| `mac` | `*string` | Query, Optional | Device MAC Address |

## Response Type

[`[]models.JseDevice`](../../doc/models/jse-device.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1







apiResponse, err := orgsJSI.ListOrgJsiDevices(ctx, orgId, &limit, &page, nil, nil, nil)
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
    "ext_ip": "73.92.124.103",
    "last_seen": 1654636867,
    "mac": "c15353123096",
    "model": "EX2300-C-12P",
    "serial": "DGCOO0015"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Jsi Past Purchases

Get List of all devices purchased from the accounts associated with the Org

```go
ListOrgJsiPastPurchases(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int,
    model *string,
    serial *string) (
    models.ApiResponse[[]models.JseInventoryItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `model` | `*string` | Query, Optional | - |
| `serial` | `*string` | Query, Optional | - |

## Response Type

[`[]models.JseInventoryItem`](../../doc/models/jse-inventory-item.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1





apiResponse, err := orgsJSI.ListOrgJsiPastPurchases(ctx, orgId, &limit, &page, nil, nil)
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
    "contract_end_time": 1708214400,
    "contract_id": "68247710",
    "contract_sku": "PAR-SDCE-EX4300T48",
    "contract_start_time": 1629331200,
    "customer_po": "4501942932",
    "distributor": "dist1",
    "eol_time": 1671062400,
    "eos_time": 1828828800,
    "installed_address": "Address",
    "model": "EX4300-48T",
    "order_id": "15909647",
    "reseller": "reseller1",
    "serial": "PE3721050223",
    "sku": "EX4300-48T-AFI",
    "type": "switch",
    "warranty_end_time": 2258236800,
    "warranty_start_time": 1619395200,
    "warranty_type": "Enhanced Hardware Warranty"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Request - no Juniper Account Linked | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

