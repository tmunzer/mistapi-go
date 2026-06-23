# Orgs JSI

```go
orgsJSI := client.OrgsJSI()
```

## Class Name

`OrgsJSI`

## Methods

* [Adopt Org Jsi Device](../../doc/controllers/orgs-jsi.md#adopt-org-jsi-device)
* [Count Org Jsi Assets and Contracts](../../doc/controllers/orgs-jsi.md#count-org-jsi-assets-and-contracts)
* [Count Org Jsi Pbn](../../doc/controllers/orgs-jsi.md#count-org-jsi-pbn)
* [Count Org Jsi Sirt](../../doc/controllers/orgs-jsi.md#count-org-jsi-sirt)
* [Create Org Jsi Device Shell Session](../../doc/controllers/orgs-jsi.md#create-org-jsi-device-shell-session)
* [List Org Jsi Devices](../../doc/controllers/orgs-jsi.md#list-org-jsi-devices)
* [List Org Jsi Past Purchases](../../doc/controllers/orgs-jsi.md#list-org-jsi-past-purchases)
* [Search Org Jsi Assets and Contracts](../../doc/controllers/orgs-jsi.md#search-org-jsi-assets-and-contracts)
* [Search Org Jsi Pbn](../../doc/controllers/orgs-jsi.md#search-org-jsi-pbn)
* [Search Org Jsi Sirt](../../doc/controllers/orgs-jsi.md#search-org-jsi-sirt)


# Adopt Org Jsi Device

Return the outbound SSH registration command used to onboard Junos devices to Juniper Support Insights (JSI).

```go
AdoptOrgJsiDevice(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseDeviceConfigCmd],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceConfigCmd](../../doc/models/response-device-config-cmd.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsJSI.AdoptOrgJsiDevice(ctx, orgId)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.JsiInventoryCountDistinctEnum`](../../doc/models/jsi-inventory-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `account_id`, `claimed`, `has_support`, `end_of_sale_time`, `eos_time`, `version_time`, `model`, `sku`, `status`, `type`, `version`, `warranty_type` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

apiResponse, err := orgsJSI.CountOrgJsiAssetsAndContracts(ctx, orgId, nil, &limit)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseDetailString:
            log.Fatalln("ResponseDetailStringException: ", typedErr)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Count Org Jsi Pbn

Get count of PBN advisories grouped by specified field

```go
CountOrgJsiPbn(
    ctx context.Context,
    orgId uuid.UUID,
    distinct models.CountPbnDistinctEnum,
    limit *int,
    start *string,
    end *string) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`models.CountPbnDistinctEnum`](../../doc/models/count-pbn-distinct-enum.md) | Query, Required | Field to group by enum: `versions`, `models`, `customer_risk`, `bug_type` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.CountPbnDistinctEnum_VERSIONS

limit := 100

apiResponse, err := orgsJSI.CountOrgJsiPbn(ctx, orgId, distinct, &limit, nil, nil)
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
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# Count Org Jsi Sirt

Get count of SIRT advisories grouped by specified field

```go
CountOrgJsiSirt(
    ctx context.Context,
    orgId uuid.UUID,
    distinct models.CountSirtDistinctEnum,
    limit *int,
    start *string,
    end *string) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`models.CountSirtDistinctEnum`](../../doc/models/count-sirt-distinct-enum.md) | Query, Required | Field to group by. enum: `jsa_updated_date`, `models`, `severity`, `versions` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.CountSirtDistinctEnum_VERSIONS

limit := 100

apiResponse, err := orgsJSI.CountOrgJsiSirt(ctx, orgId, distinct, &limit, nil, nil)
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
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# Create Org Jsi Device Shell Session

Create a WebSocket-backed shell session for a JSI-connected device identified by MAC address.

```go
CreateOrgJsiDeviceShellSession(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSessionWithUrl](../../doc/models/websocket-session-with-url.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

apiResponse, err := orgsJSI.CreateOrgJsiDeviceShellSession(ctx, orgId, deviceMac)
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
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# List Org Jsi Devices

List organization devices connected to Juniper Support Insights (JSI), optionally filtered by model, serial number, or MAC address.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `serial` | `*string` | Query, Optional | Filter results by device serial number |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |

## Response Type

**200**: OK

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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `model` | `*string` | Query, Optional | Filter results by one or more device models. Supports comma-separated values |
| `serial` | `*string` | Query, Optional | Filter results by device serial number |

## Response Type

**200**: OK

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
    switch typedErr := err.(type) {
        case *errors.ResponseDetailString:
            log.Fatalln("ResponseDetailStringException: ", typedErr)
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
    "end_of_sale_time": 1561507200,
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


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
    endOfSaleAfter *string,
    endOfSaleBefore *string,
    eosAfter *string,
    eosBefore *string,
    versionEosAfter *string,
    versionEosBefore *string,
    hasSupport *bool,
    sirtId *string,
    pbnId *string,
    text *string,
    limit *int,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.JsInventorySearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `claimed` | `*bool` | Query, Optional | Device claim status, `true` for claimed devices, `false` for all devices. Accepts multiple comma-separated boolean values. |
| `model` | `*string` | Query, Optional | Filter results by device model. Accepts multiple comma-separated values. |
| `serial` | `*string` | Query, Optional | Filter results by device serial number. Accepts multiple comma-separated values. |
| `sku` | `*string` | Query, Optional | Filter results by SKU. Accepts multiple comma-separated values. |
| `status` | [`*models.DeviceStatusEnum`](../../doc/models/device-status-enum.md) | Query, Optional | Device status. enum: `all`, `connected`, `disconnected`<br><br>**Default**: `"all"` |
| `warrantyType` | [`*models.JsiWarrantyTypeEnum`](../../doc/models/jsi-warranty-type-enum.md) | Query, Optional | Device warranty type used to filter Juniper Support Insight inventory. enum: `Standard Hardware Warranty`, `Enhanced Hardware Warranty`, `Dead On Arrival Warranty`, `Limited Lifetime Warranty`, `Software Warranty`, `Limited Lifetime Warranty for WLA`, `Warranty-JCPO EOL (DOA Not Included)`, `MIST Enhanced Hardware Warranty`, `MIST Standard Warranty`, `Determine Lifetime warranty` |
| `endOfSaleAfter` | `*string` | Query, Optional | Filter devices with End Of Sale date after this date |
| `endOfSaleBefore` | `*string` | Query, Optional | Filter devices with End Of Sale date before this date |
| `eosAfter` | `*string` | Query, Optional | Filter devices with End Of Support date after this date |
| `eosBefore` | `*string` | Query, Optional | Filter devices with End Of Support date before this date |
| `versionEosAfter` | `*string` | Query, Optional | Filter devices with OS Version End Of Support date after this date |
| `versionEosBefore` | `*string` | Query, Optional | Filter devices with OS Version End Of Support date before this date |
| `hasSupport` | `*bool` | Query, Optional | Indicates if the device is covered under active support contract. Accepts multiple comma-separated boolean values. |
| `sirtId` | `*string` | Query, Optional | To get the onboarded devices that are affected by the SIRT ID |
| `pbnId` | `*string` | Query, Optional | To get the onboarded devices that are affected by the PBN ID |
| `text` | `*string` | Query, Optional | Wildcards for `serial`, `model`, `account_id` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.JsInventorySearch](../../doc/models/js-inventory-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

claimed := true

model := "QFX10000-36Q,QFX10000-30C"

serial := "ACNP4666,ACNP6969"

sku := "QFX10000-36Q,QFX10000-30C"

status := models.DeviceStatusEnum_ALL

endOfSaleAfter := "2024-01-01"

endOfSaleBefore := "2025-12-31"

eosAfter := "2024-01-01"

eosBefore := "2025-12-31"

versionEosAfter := "2024-01-01"

versionEosBefore := "2025-12-31"

hasSupport := true

sirtId := "JSA12345"

pbnId := "PBN67890"

limit := 100

sort := "-site_id"

apiResponse, err := orgsJSI.SearchOrgJsiAssetsAndContracts(ctx, orgId, &claimed, &model, &serial, &sku, &status, nil, &endOfSaleAfter, &endOfSaleBefore, &eosAfter, &eosBefore, &versionEosAfter, &versionEosBefore, &hasSupport, &sirtId, &pbnId, nil, &limit, &sort, nil)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseDetailString:
            log.Fatalln("ResponseDetailStringException: ", typedErr)
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
      "end_of_sale_time": 1561507200,
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Jsi Pbn

Text search for PBN (Problem Bug Notification) advisories. Search can be done on versions, models, customer_risk, id, and bug_type fields.

```go
SearchOrgJsiPbn(
    ctx context.Context,
    orgId uuid.UUID,
    versions *string,
    mModels *string,
    customerRisk *string,
    id *string,
    bugType *string,
    limit *int,
    page *int,
    searchAfter *string,
    start *string,
    end *string) (
    models.ApiResponse[models.JsiPbnSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `versions` | `*string` | Query, Optional | OS versions to search for |
| `mModels` | `*string` | Query, Optional | Device models to search for |
| `customerRisk` | `*string` | Query, Optional | Customer risk level to filter by |
| `id` | `*string` | Query, Optional | PBN ID to search for |
| `bugType` | `*string` | Query, Optional | Bug type to filter by |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.JsiPbnSearch](../../doc/models/jsi-pbn-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

versions := "20.4R3"

id := "1403338"

limit := 100

page := 1

apiResponse, err := orgsJSI.SearchOrgJsiPbn(ctx, orgId, &versions, nil, nil, &id, nil, &limit, &page, nil, nil, nil)
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
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# Search Org Jsi Sirt

Search and get all the SIRT for the onboarded devices. Search can be done on severity, id, updated_after, updated_before, published_after, published_before, models, versions, and text fields.

```go
SearchOrgJsiSirt(
    ctx context.Context,
    orgId uuid.UUID,
    severity *string,
    id *string,
    updatedAfter *string,
    updatedBefore *string,
    publishedAfter *string,
    publishedBefore *string,
    mModels *string,
    versions *string,
    text *string,
    limit *int,
    page *int,
    sort *string,
    searchAfter *string,
    start *string,
    end *string) (
    models.ApiResponse[models.JsiSirtSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `severity` | `*string` | Query, Optional | Filter results by severity |
| `id` | `*string` | Query, Optional | Filter results by identifier |
| `updatedAfter` | `*string` | Query, Optional | JSA Updated date to be filtered after this date |
| `updatedBefore` | `*string` | Query, Optional | JSA Updated date to be filtered before this date |
| `publishedAfter` | `*string` | Query, Optional | JSA Published date to be filtered after this date |
| `publishedBefore` | `*string` | Query, Optional | JSA Published date to be filtered before this date |
| `mModels` | `*string` | Query, Optional | Filter results by models |
| `versions` | `*string` | Query, Optional | Software version affected by the SIRT |
| `text` | `*string` | Query, Optional | Wildcards search on os_version_affected, affected_models, severity, jsa_id |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.JsiSirtSearch](../../doc/models/jsi-sirt-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

id := "JSA100053"

versions := "20.4R3"

limit := 100

page := 1

sort := "-site_id"

apiResponse, err := orgsJSI.SearchOrgJsiSirt(ctx, orgId, nil, &id, nil, nil, nil, nil, nil, &versions, nil, &limit, &page, &sort, nil, nil, nil)
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
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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

