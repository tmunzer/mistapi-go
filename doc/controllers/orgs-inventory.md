# Orgs Inventory

```go
orgsInventory := client.OrgsInventory()
```

## Class Name

`OrgsInventory`

## Methods

* [Add Org Inventory](../../doc/controllers/orgs-inventory.md#add-org-inventory)
* [Get Org Inventory](../../doc/controllers/orgs-inventory.md#get-org-inventory)
* [Reevaluate Org Auto Assignment](../../doc/controllers/orgs-inventory.md#reevaluate-org-auto-assignment)
* [Replace Org Devices](../../doc/controllers/orgs-inventory.md#replace-org-devices)
* [Update Org Inventory Assignment](../../doc/controllers/orgs-inventory.md#update-org-inventory-assignment)


# Add Org Inventory

Add Device to Org Inventory with the device claim codes

```go
AddOrgInventory(
    ctx context.Context,
    orgId uuid.UUID,
    body []string) (
    models.ApiResponse[models.ResponseInventory],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | `[]string` | Body, Optional | Request Body |

## Response Type

[`models.ResponseInventory`](../../doc/models/response-inventory.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := []string{
    "6JG8E-PTFV2-A9Z2N",
    "DVH4V-SNMSZ-PDXBR",
}

apiResponse, err := orgsInventory.AddOrgInventory(ctx, orgId, body)
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
  "added": [
    "6JG8E-PTFV2-A9Z2N"
  ],
  "duplicated": [
    "DVH4V-SNMSZ-PDXBR"
  ],
  "error": [
    "PO1025335ohoh"
  ],
  "inventory_added": [
    {
      "mac": "5c5b35000018",
      "magic": "6JG8EPTFV2A9Z2N",
      "model": "AP41",
      "serial": "FXLH2015150025",
      "type": "ap"
    }
  ],
  "inventory_duplicated": [
    {
      "mac": "5c5b35000012",
      "magic": "DVH4VSNMSZPDXBR",
      "model": "AP41",
      "serial": "FXLH2015150027",
      "type": "ap"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | if none of the entries are valid | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Org Inventory

Get Org Inventory

### VC (Virtual-Chassis) Management

Ideally VC should be managed as a single device - where - one managed entity where config / monitoring is anchored against (with a stable identify MAC) - all members appears in the inventory

In our implementation, we strive to achieve that without manual user configurations by

1. during claim or adoption a VC, we require FPC0 to exist and will use its MAC as identify for the entire chassis
2. other VC members will be automatically populated when they’re all present

The perceivable result is

1. from `/sites/:site_id/stats/devices/:fpc0_mac` API, you’d see the VC where module_stat contains the VC members
2. from `/orgs/:org_id/inventory?vc=true` API, you’d see all VC members with vc_mac pointing to the FPC0

```go
GetOrgInventory(
    ctx context.Context,
    orgId uuid.UUID,
    serial *string,
    model *string,
    mType *models.DeviceTypeEnum,
    mac *string,
    siteId *string,
    vcMac *string,
    vc *bool,
    unassigned *bool,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Inventory],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `serial` | `*string` | Query, Optional | device serial |
| `model` | `*string` | Query, Optional | device model |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | - |
| `mac` | `*string` | Query, Optional | MAC address |
| `siteId` | `*string` | Query, Optional | site id if assigned, null if not assigned |
| `vcMac` | `*string` | Query, Optional | Virtual Chassis MAC Address |
| `vc` | `*bool` | Query, Optional | To display Virtual Chassis members |
| `unassigned` | `*bool` | Query, Optional | to display Unassigned devices |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Inventory`](../../doc/models/inventory.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





mType := models.DeviceTypeEnum("ap")







vc := false

unassigned := true

limit := 100

page := 1

apiResponse, err := orgsInventory.GetOrgInventory(ctx, orgId, nil, nil, &mType, nil, nil, nil, &vc, &unassigned, &limit, &page)
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
    "connected": true,
    "created_time": 1542328276,
    "deviceprofile_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
    "id": "00000000-0000-0000-0000-5c5b35000018",
    "mac": "5c5b35000018",
    "model": "AP41",
    "modified_time": 1542829778,
    "name": "hallway",
    "serial": "FXLH2015150025",
    "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
    "type": "ap"
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


# Reevaluate Org Auto Assignment

Reevaluate Auto Assignment

```go
ReevaluateOrgAutoAssignment(
    ctx context.Context,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsInventory.ReevaluateOrgAutoAssignment(ctx, orgId)
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


# Replace Org Devices

It’s a common request we get from the customers. When a AP HW has problem and need a replacement, they would want to copy the existing attributes (Device Config) of this old AP to the new one. It can be done by providing the MAC of a device that’s currently in the inventory but not assigned. The Device replaced will become unassigned.

This API also supports replacement of Mist Edges. This API copies device agnostic attributes from old Mist edge to new one.
Mist manufactured Mist Edges will be reset to factory settings but will still be in Inventory.Brownfield or VM’s will be
deleted from Inventory

**Note:** For Gateway devices only like-for-like replacements (can only replace a SRX320 with a SRX320 and not some otehr model) are allowed.

```go
ReplaceOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.ReplaceDevice) (
    models.ApiResponse[models.ResponseOrgInventoryChange],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.ReplaceDevice`](../../doc/models/replace-device.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseOrgInventoryChange`](../../doc/models/response-org-inventory-change.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.ReplaceDevice{
    Discard:           []string{
    },
    InventoryMac:      models.ToPointer("5c5b35000301"),
    Mac:               models.ToPointer("5c5b35000101"),
    SiteId:            models.ToPointer("4ac1dcf4-9d8b-7211-65c4-057819f0862b"),
}

apiResponse, err := orgsInventory.ReplaceOrgDevices(ctx, orgId, &body)
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
  "error": [],
  "op": "assign",
  "reason": [],
  "success": [
    "5c5b350e0001"
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


# Update Org Inventory Assignment

Update Org Inventory

```go
UpdateOrgInventoryAssignment(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.InventoryUpdate) (
    models.ApiResponse[models.ResponseOrgInventoryChange],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.InventoryUpdate`](../../doc/models/inventory-update.md) | Body, Optional | - |

## Response Type

[`models.ResponseOrgInventoryChange`](../../doc/models/response-org-inventory-change.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.InventoryUpdate{
    DisableAutoConfig: models.ToPointer(false),
    Macs:              []string{
        "5c5b350e0001",
    },
    Managed:           models.ToPointer(false),
    NoReassign:        models.ToPointer(false),
    Op:                models.InventoryUpdateOperationEnum("assign"),
    SiteId:            models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
}

apiResponse, err := orgsInventory.UpdateOrgInventoryAssignment(ctx, orgId, &body)
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
  "error": [],
  "op": "assign",
  "reason": [],
  "success": [
    "5c5b350e0001"
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
