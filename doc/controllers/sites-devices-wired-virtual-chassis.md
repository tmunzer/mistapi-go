# Sites Devices-Wired-Virtual Chassis

```go
sitesDevicesWiredVirtualChassis := client.SitesDevicesWiredVirtualChassis()
```

## Class Name

`SitesDevicesWiredVirtualChassis`

## Methods

* [Create Site Virtual Chassis](../../doc/controllers/sites-devices-wired-virtual-chassis.md#create-site-virtual-chassis)
* [Delete Site Virtual Chassis](../../doc/controllers/sites-devices-wired-virtual-chassis.md#delete-site-virtual-chassis)
* [Get Site Device Virtual Chassis](../../doc/controllers/sites-devices-wired-virtual-chassis.md#get-site-device-virtual-chassis)
* [Set Site Vc Port](../../doc/controllers/sites-devices-wired-virtual-chassis.md#set-site-vc-port)
* [Update Site Virtual Chassis Member](../../doc/controllers/sites-devices-wired-virtual-chassis.md#update-site-virtual-chassis-member)


# Create Site Virtual Chassis

For models (e.g. EX3400 and up) having dedicated VC ports, it is easier to form a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new VC and update the inventory.

In case that the user would like to choose the dedicated switch as a VC master. Or for EX2300-C-12P and EX2300-C-12T which doesn’t have dedicated VC ports, below are procedures to automate the VC creation:

1. Power on the switch that is choosen as the VC master first. And the powering on the other member switches.
2. Claim or adopt all these switches under the same organization’s Inventory
3. Assign these switches into the same Site
4. Invoke vc command on the switch choosen to be the VC master. For EX2300-C-12P, VC ports will be created automatically.
5. Connect the cables to the VC ports for these switches
6. Wait for the VC to be formed. The Org’s inventory will be updated for the new VC.

```go
CreateSiteVirtualChassis(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.VirtualChassisConfig) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.VirtualChassisConfig`](../../doc/models/virtual-chassis-config.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseVirtualChassisConfig`](../../doc/models/response-virtual-chassis-config.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.VirtualChassisConfig{
    Members:        []models.VirtualChassisConfigMember{
        models.VirtualChassisConfigMember{
            Mac:      models.ToPointer("aff827549235"),
            VcPorts:  []string{
                "xe-0/1/0",
            },
            VcRole:   models.ToPointer(models.VirtualChassisConfigMemberVcRoleEnum("master")),
        },
        models.VirtualChassisConfigMember{
            Mac:      models.ToPointer("8396cd006c8c"),
            VcPorts:  []string{
                "xe-0/1/0",
                "xe-0/1/1",
            },
            VcRole:   models.ToPointer(models.VirtualChassisConfigMemberVcRoleEnum("backup")),
        },
        models.VirtualChassisConfigMember{
            Mac:      models.ToPointer("8396cd00888c"),
            VcPorts:  []string{
                "xe-0/1/0",
            },
            VcRole:   models.ToPointer(models.VirtualChassisConfigMemberVcRoleEnum("linecard")),
        },
    },
}

apiResponse, err := sitesDevicesWiredVirtualChassis.CreateSiteVirtualChassis(ctx, siteId, deviceId, &body)
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
  "id": {
    "members": [
      {
        "mac": "string",
        "member": 0,
        "vc_ports": [
          "string"
        ],
        "vc_role": "master"
      }
    ],
    "op": "add"
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


# Delete Site Virtual Chassis

When all the member switches of VC are removed and only member ID 0 is left, the cloud would detect this situation and automatically changes the single switch to non-VC role.

For some unexpected cases that the VC is gone and disconncted, the API below could be used to change the state of VC’s switches to be standalone. After it is executed, all the switches will be shown as standalone switches under Inventory.

```go
DeleteSiteVirtualChassis(
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

resp, err := sitesDevicesWiredVirtualChassis.DeleteSiteVirtualChassis(ctx, siteId, deviceId)
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


# Get Site Device Virtual Chassis

Get VC Status

The API returns a combined view of the VC status which includes topology and stats_

```go
GetSiteDeviceVirtualChassis(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseVirtualChassisConfig`](../../doc/models/response-virtual-chassis-config.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevicesWiredVirtualChassis.GetSiteDeviceVirtualChassis(ctx, siteId, deviceId)
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
  "id": {
    "members": [
      {
        "mac": "string",
        "member": 0,
        "vc_ports": [
          "string"
        ],
        "vc_role": "master"
      }
    ],
    "op": "add"
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


# Set Site Vc Port

Set VC port

```go
SetSiteVcPort(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.VirtualChassisPort) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.VirtualChassisPort`](../../doc/models/virtual-chassis-port.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.VirtualChassisPort{
    Members: []models.ConfigVcPortMember{
        models.ConfigVcPortMember{
            Member:  float64(0),
            VcPorts: []string{
                "xe-0/1/1",
            },
        },
        models.ConfigVcPortMember{
            Member:  float64(2),
            VcPorts: []string{
                "xe-0/1/1",
            },
        },
    },
    Op:      models.VirtualChassisPortOperationEnum("delete"),
}

resp, err := sitesDevicesWiredVirtualChassis.SetSiteVcPort(ctx, siteId, deviceId, &body)
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


# Update Site Virtual Chassis Member

The VC creation and adding member switch API will update the device’ s virtual chassis config which is applied after VC is formed to create JUNOS pre-provisioned virtual chassis configuration.

## Change to use preprovisioned VC

To switch the VC to use preprovisioned VC, enable preprovisioned in virtual_chassis config. Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config.

In this config, fpc0 has to be the same as the mac of device_id. Use renumber if you want to replace fpc0 which involves device_id change.

Notice: to configure preprovisioned VC, every member of the VC must be in the inventory.

## Add new members

For models (e.g. EX4300 and up) having dedicated VC ports, it is easier to add new member switches into a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new members and update the inventory.

For EX2300 VC, adding new members requires to follow the procedures below:

1. Powering on the new member switches and ensuring cables are not connected to any VC ports.
2. Claim or adopt all new member switches under the VC’s organization Inventory
3. Assign all new member switches to the same Site as the VC
4. Invoke vc command to add switches to the VC.
5. Connect the cables to the VC ports for these switches
6. After a while, the Org’s Inventory shows this new switches has been added into the VC.

## Removing member switch

To remove a member switch from the VC, following the procedures below:

1. Ensuring the VC is connected to the cloud first
2. Unplug the cable from the VC port of the switch
3. Waiting for the VC state (vc_state) of this switch is changed to not-present
4. Invoke update_vc with remove to remove this switch from the VC
5. The Org’s Inventory shows the switch is removed.

Please notice that member ID 0 (fpc0) cannot be removed. When a VC has two switches left, unpluging the cable may result in the situation that fpc0 becomes a line card (LC). When this situation is happened, please re-plug in the cable, wait for both switches becoming present (show virtual-chassis) and then removing the cable again.

## Renumber a member switch

When a member switch doesn’ t work properly and needed to be replaced, the renumber API could be used. The following two types of renumber are supported:

1. Replace a non-fpc0 member switch
2. Replace fpc0. When fpc0 is relaced, PAPI device config and JUNOS config will be both updated.

For renumber to work, the following procedures are needed:

1. Ensuring the VC is connected to the cloud and the state of the member switch to be replaced must be non present.
2. Adding the new member switch to the VC
3. Waiting for the VC state (vc_state) of this VC to be updated to API server
4. Invoke vc with renumber to replace the new member switch from fpc X to

## Perprovision VC members

By specifying “preprovision” op, you can convert the current VC to pre-provisioned mode, update VC members as well as specify vc_ports when adding new members for device models without dedicated vc ports. Use renumber for fpc0 replacement which involves device_id change.

Note:

1. vc_ports is used for adding new members and not needed if * the device model has dedicated vc ports, or * no new member is added
2. New VC members to be added should exist in the same Site as the VC

Update Device’s VC config can achieve similar purpose by directly modifying current virtual_chassis config. However, it cannot fulfill requests to enabling vc_ports on new members that are yet to belong to current VC.

```go
UpdateSiteVirtualChassisMember(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.VirtualChassisUpdate) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.VirtualChassisUpdate`](../../doc/models/virtual-chassis-update.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseVirtualChassisConfig`](../../doc/models/response-virtual-chassis-config.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.VirtualChassisUpdate{
    Members:   []models.VirtualChassisMemberUpdate{
        models.VirtualChassisMemberUpdate{
            Mac:     models.ToPointer("aff827549235"),
            Member:  models.ToPointer(0),
            VcPorts: []string{
                "xe-0/1/1",
            },
            VcRole:  models.ToPointer(models.VirtualChassisMemberUpdateVcRoleEnum("linecard")),
        },
        models.VirtualChassisMemberUpdate{
            Mac:     models.ToPointer("8396cd00777c"),
            VcPorts: []string{
                "xe-0/1/0",
            },
            VcRole:  models.ToPointer(models.VirtualChassisMemberUpdateVcRoleEnum("linecard")),
        },
    },
    Op:        models.ToPointer(models.VirtualChassisUpdateOpEnum("add")),
}

apiResponse, err := sitesDevicesWiredVirtualChassis.UpdateSiteVirtualChassisMember(ctx, siteId, deviceId, &body)
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
  "id": {
    "members": [
      {
        "mac": "string",
        "member": 0,
        "vc_ports": [
          "string"
        ],
        "vc_role": "master"
      }
    ],
    "op": "add"
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

