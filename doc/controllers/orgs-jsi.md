# Orgs JSI

```go
orgsJSI := client.OrgsJSI()
```

## Class Name

`OrgsJSI`

## Methods

* [Adopt Org Jsi Device](../../doc/controllers/orgs-jsi.md#adopt-org-jsi-device)
* [Count Org Jsi Assets and Contracts](../../doc/controllers/orgs-jsi.md#count-org-jsi-assets-and-contracts)
* [Create Org Jsi Device Shell Session](../../doc/controllers/orgs-jsi.md#create-org-jsi-device-shell-session)
* [List Org Jsi Devices](../../doc/controllers/orgs-jsi.md#list-org-jsi-devices)
* [List Org Jsi Past Purchases](../../doc/controllers/orgs-jsi.md#list-org-jsi-past-purchases)
* [Search Org Jsi Assets and Contracts](../../doc/controllers/orgs-jsi.md#search-org-jsi-assets-and-contracts)


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceConfigCmd](../../doc/models/response-device-config-cmd.md).

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Count Org Jsi Assets and Contracts

Count devices purchased from the accounts associated with the Org

```go
CountOrgJsiAssetsAndContracts(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.JsiInventoryCountDistinctEnum,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.JsiInventoryCountDistinctEnum`](../../doc/models/jsi-inventory-count-distinct-enum.md) | Query, Optional | Distinct attributes to count |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

apiResponse, err := orgsJSI.CountOrgJsiAssetsAndContracts(ctx, orgId, nil, &limit)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Request - no Juniper Account Linked | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSessionWithUrl](../../doc/models/websocket-session-with-url.md).

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |
| `model` | `*string` | Query, Optional | Device model |
| `serial` | `*string` | Query, Optional | Device serial |
| `mac` | `*string` | Query, Optional | Device MAC Address |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.JseDevice](../../doc/models/jse-device.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

model := "AP43"

serial := "FXLH2015150025"

mac := "5c5b350e0001"

apiResponse, err := orgsJSI.ListOrgJsiDevices(ctx, orgId, &limit, &page, &model, &serial, &mac)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Jsi Past Purchases

This gets all devices purchased from the accounts associated with the Org

* Fetch Install base devices for all linked accounts and associated account of the linked accounts.
* The primary and the associated account ids will be queries from SFDC by passing the linked account
* Returns only the device centric details of the Install base device. No customer specific information will be returned.

```go
ListOrgJsiPastPurchases(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int,
    model *string,
    serial *string) (
    models.ApiResponse[[]models.JsInventoryItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |
| `model` | `*string` | Query, Optional | - |
| `serial` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.JsInventoryItem](../../doc/models/js-inventory-item.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

model := "AP43"

serial := "FXLH2015150025"

apiResponse, err := orgsJSI.ListOrgJsiPastPurchases(ctx, orgId, &limit, &page, &model, &serial)
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
    "device_name": "name1",
    "eol_time": 1561507200,
    "eos_time": 1672012800,
    "master": true,
    "model": "EX2300-24MP",
    "org_id": "6e843b41-f953-4af9-80e5-e1a70f65754a",
    "serial": "XN3123300095",
    "sku": "EX2300",
    "status": "connected",
    "suggested_version": "Latest 21.4R3-Sx",
    "type": "switch",
    "version": "23.4R2-S4.11",
    "version_eos_time": 1672012800,
    "version_time": 1561507200,
    "warranty": "Enhanced Hardware Warranty",
    "warranty_time": 1672012800,
    "warranty_type": "Enhanced Hardware Warranty"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Request - no Juniper Account Linked | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Search Org Jsi Assets and Contracts

This gets all devices purchased from the accounts associated with the Org

* Fetch Install base devices for all linked accounts and associated account of the linked accounts.
* The primary and the associated account ids will be queries from SFDC by passing the linked account
* Returns only the device centric details of the Install base device. No customer specific information will be returned.

```go
SearchOrgJsiAssetsAndContracts(
    ctx context.Context,
    orgId uuid.UUID,
    claimed *bool,
    model *string,
    serial *string,
    sku *string,
    status *models.DeviceStatusEnum,
    warrantyType *models.JsiWarrantyTypeEnum,
    eolDuration *string,
    eosDuration *string,
    hasSupport *bool,
    text *string,
    limit *int,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.JsInventorySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `claimed` | `*bool` | Query, Optional | Device claim status, `true` for claimed devices, `false` for all devices |
| `model` | `*string` | Query, Optional | Device model |
| `serial` | `*string` | Query, Optional | Device serial |
| `sku` | `*string` | Query, Optional | SKU name of the device |
| `status` | [`*models.DeviceStatusEnum`](../../doc/models/device-status-enum.md) | Query, Optional | Device status<br><br>**Default**: `"all"` |
| `warrantyType` | [`*models.JsiWarrantyTypeEnum`](../../doc/models/jsi-warranty-type-enum.md) | Query, Optional | Device warranty type |
| `eolDuration` | `*string` | Query, Optional | Device EOL duration in days |
| `eosDuration` | `*string` | Query, Optional | Device EOS duration in days |
| `hasSupport` | `*bool` | Query, Optional | Indicates if the device is covered under active support contract |
| `text` | `*string` | Query, Optional | Wildcards for `model`, `serial`, `account_id`, `status`, `warranty_type` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.JsInventorySearch](../../doc/models/js-inventory-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

claimed := true

model := "AP43"

serial := "FXLH2015150025"

sku := "EX2300"

status := models.DeviceStatusEnum_ALL

eolDuration := "30d"

eosDuration := "30d"

hasSupport := true

limit := 100

sort := "-site_id"

apiResponse, err := orgsJSI.SearchOrgJsiAssetsAndContracts(ctx, orgId, &claimed, &model, &serial, &sku, &status, nil, &eolDuration, &eosDuration, &hasSupport, nil, &limit, &sort, nil)
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
  "end": 1748023308,
  "limit": 1000,
  "results": [
    {
      "claimed": true,
      "device_name": "name1",
      "eol_time": 1561507200,
      "eos_time": 1672012800,
      "has_support": true,
      "master": true,
      "model": "EX2300-24MP",
      "org_id": "6e843b41-f953-4af9-80e5-e1a70f65754a",
      "serial": "XN3123300095",
      "sku": "EX2300",
      "status": "connected",
      "suggested_version": "Latest 21.4R3-Sx",
      "type": "switch",
      "version": "23.4R2-S4.11",
      "version_eos_time": 1672012800,
      "version_time": 1561507200,
      "warranty": "Enhanced Hardware Warranty",
      "warranty_time": 1672012800,
      "warranty_type": "Enhanced Hardware Warranty"
    }
  ],
  "start": 1748019708,
  "total": 1
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Request - no Juniper Account Linked | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

