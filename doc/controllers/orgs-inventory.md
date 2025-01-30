# Orgs Inventory

```go
orgsInventory := client.OrgsInventory()
```

## Class Name

`OrgsInventory`

## Methods

* [Add Org Inventory](../../doc/controllers/orgs-inventory.md#add-org-inventory)
* [Count Org Inventory](../../doc/controllers/orgs-inventory.md#count-org-inventory)
* [Create Org Gateway Ha Cluster](../../doc/controllers/orgs-inventory.md#create-org-gateway-ha-cluster)
* [Delete Org Gateway Ha Cluster](../../doc/controllers/orgs-inventory.md#delete-org-gateway-ha-cluster)
* [Get Org Inventory](../../doc/controllers/orgs-inventory.md#get-org-inventory)
* [Reevaluate Org Auto Assignment](../../doc/controllers/orgs-inventory.md#reevaluate-org-auto-assignment)
* [Replace Org Devices](../../doc/controllers/orgs-inventory.md#replace-org-devices)
* [Search Org Inventory](../../doc/controllers/orgs-inventory.md#search-org-inventory)
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
| `body` | `[]string` | Body, Optional | Request Body<br>**Constraints**: *Unique Items Required* |

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
  ],
  "reason": [
    "belongs to another org ('e2f543f7-d6e1-409f-a565-e77a1f098d3b' (other) != '0de5d6fc-219a-414d-a840-67d6b919ad8f' (you))"
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | OK - if any of entries are valid or there’s no errors | [`ResponseInventoryErrorException`](../../doc/models/response-inventory-error-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Count Org Inventory

Count in the Org Inventory

```go
CountOrgInventory(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    distinct *models.InventoryCountDistinctEnum,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | **Default**: `"ap"` |
| `distinct` | [`*models.InventoryCountDistinctEnum`](../../doc/models/inventory-count-distinct-enum.md) | Query, Optional | **Default**: `"model"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum_AP

distinct := models.InventoryCountDistinctEnum_MODEL

limit := 100

page := 1

apiResponse, err := orgsInventory.CountOrgInventory(ctx, orgId, &mType, &distinct, &limit, &page)
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
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Create Org Gateway Ha Cluster

Create HA Cluster from unassigned Gateways

```go
CreateOrgGatewayHaCluster(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.HaClusterConfig) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.HaClusterConfig`](../../doc/models/ha-cluster-config.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.HaClusterConfig{
    DisableAutoConfig:    models.ToPointer(true),
    Managed:              models.ToPointer(true),
    Nodes:                []models.HaClusterConfigNode{
        models.HaClusterConfigNode{
            Mac:                  models.ToPointer("aff827549235"),
        },
        models.HaClusterConfigNode{
            Mac:                  models.ToPointer("8396cd006c8c"),
        },
    },
    SiteId:               models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
}

resp, err := orgsInventory.CreateOrgGatewayHaCluster(ctx, orgId, &body)
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


# Delete Org Gateway Ha Cluster

Delete HA Cluster

After HA cluster deleted, both of the nodes will be unassigned.

```go
DeleteOrgGatewayHaCluster(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.HaClusterDelete) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.HaClusterDelete`](../../doc/models/ha-cluster-delete.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.HaClusterDelete{
    Mac:                  models.ToPointer("aff827549235"),
}

resp, err := orgsInventory.DeleteOrgGatewayHaCluster(ctx, orgId, &body)
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
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | **Default**: `"ap"` |
| `mac` | `*string` | Query, Optional | MAC address |
| `siteId` | `*string` | Query, Optional | site id if assigned, null if not assigned |
| `vcMac` | `*string` | Query, Optional | Virtual Chassis MAC Address |
| `vc` | `*bool` | Query, Optional | To display Virtual Chassis members<br>**Default**: `false` |
| `unassigned` | `*bool` | Query, Optional | to display Unassigned devices<br>**Default**: `true` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Inventory`](../../doc/models/inventory.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





mType := models.DeviceTypeEnum_AP







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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
    Discard:              []string{
    },
    InventoryMac:         models.ToPointer("5c5b35000301"),
    Mac:                  models.ToPointer("5c5b35000101"),
    SiteId:               models.ToPointer("4ac1dcf4-9d8b-7211-65c4-057819f0862b"),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Search Org Inventory

Search in the Org Inventory

```go
SearchOrgInventory(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    mac *string,
    vcMac *string,
    masterMac *string,
    siteId *string,
    serial *string,
    master *string,
    sku *string,
    version *string,
    status *string,
    text *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.InventorySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | **Default**: `"ap"` |
| `mac` | `*string` | Query, Optional | MAC address |
| `vcMac` | `*string` | Query, Optional | Virtual Chassis MAC Address |
| `masterMac` | `*string` | Query, Optional | master device mac for virtual mac cluster |
| `siteId` | `*string` | Query, Optional | site id if assigned, null if not assigned |
| `serial` | `*string` | Query, Optional | device serial |
| `master` | `*string` | Query, Optional | true / false |
| `sku` | `*string` | Query, Optional | device sku |
| `version` | `*string` | Query, Optional | device version |
| `status` | `*string` | Query, Optional | device status |
| `text` | `*string` | Query, Optional | wildcards for name, mac, serial |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.InventorySearch`](../../doc/models/inventory-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum_AP





















limit := 100

page := 1

apiResponse, err := orgsInventory.SearchOrgInventory(ctx, orgId, &mType, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, &page)
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
  "limit": 1000,
  "results": [
    {
      "mac": "f01c2df166e0",
      "master": true,
      "members": [
        {
          "mac": "f01c2df166e0",
          "model": "EX4300-48P",
          "serial": "PD3714460200"
        }
      ],
      "model": "EX4300-48P",
      "name": "mist-wa-ex4300-VC",
      "org_id": "9b853544-51e4-45fb-81ac-a442e4a111d0",
      "serial": "PD3714460200",
      "site_id": "01dc141d-b6af-4baa-b00f-0e31ef954c4f",
      "sku": "EX4300-48P",
      "status": "disconnected",
      "type": "switch",
      "vc_mac": "f01c2df166e0",
      "version": "21.4R3.5"
    }
  ],
  "total": 1
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
    DisableAutoConfig:    models.ToPointer(false),
    Macs:                 []string{
        "5c5b350e0001",
    },
    Managed:              models.ToPointer(false),
    NoReassign:           models.ToPointer(false),
    Op:                   models.InventoryUpdateOperationEnum_ASSIGN,
    SiteId:               models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

