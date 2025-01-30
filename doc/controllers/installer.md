# Installer

In a typical enterprise, a separate group of people, Installers, are responsible for install new devices. May it be a new installation (e.g. new stores), a replacement installation (e.g. replacing Cisco APs with Mist APs), or addition (e.g. adding new APs for better coverage). Instead of granting them Admin/Write privilege, it's more desirable to grant them minimum privileges to do the initial provisioning so they cannot read sensible information (e.g. PSK of a WLAN), or change configs of running APs.
At a high level, Installer APs try to achieve the following:

1. identifying a device by MAC (that\u2019\
   s what they see)
2. they can only touch configurations of the devices they\u2019\
   re installing
3. allow the following configurations:

* name * site assignment
* device profile assignment
* map and location (x/y) assignment
* claim (if not already in the inventory)
* replace existing device with the device being installed

**Grace Period**

Grace period provides a dynamic way to limit what devices / sites installer can work on. Generally installers work on recent deployments - bringing up new sites, add newly claimed devices to new / existing sites. They make mistakes, too, and may need to further tweak some of the parameters. Default grace period is 7 days and can be set from 1 day to 365 days.

```go
installer := client.Installer()
```

## Class Name

`Installer`

## Methods

* [Add Installer Device Image](../../doc/controllers/installer.md#add-installer-device-image)
* [Claim Installer Devices](../../doc/controllers/installer.md#claim-installer-devices)
* [Create Installer Map](../../doc/controllers/installer.md#create-installer-map)
* [Create Installer Virtual Chassis](../../doc/controllers/installer.md#create-installer-virtual-chassis)
* [Create or Update Installer Sites](../../doc/controllers/installer.md#create-or-update-installer-sites)
* [Delete Installer Device Image](../../doc/controllers/installer.md#delete-installer-device-image)
* [Delete Installer Map](../../doc/controllers/installer.md#delete-installer-map)
* [Get Installer Device Virtual Chassis](../../doc/controllers/installer.md#get-installer-device-virtual-chassis)
* [Import Installer Map](../../doc/controllers/installer.md#import-installer-map)
* [List Installer Alarm Templates](../../doc/controllers/installer.md#list-installer-alarm-templates)
* [List Installer Device Profiles](../../doc/controllers/installer.md#list-installer-device-profiles)
* [List Installer List of Renctly Claimed Devices](../../doc/controllers/installer.md#list-installer-list-of-renctly-claimed-devices)
* [List Installer Maps](../../doc/controllers/installer.md#list-installer-maps)
* [List Installer Rf Templates Names](../../doc/controllers/installer.md#list-installer-rf-templates-names)
* [List Installer Site Groups](../../doc/controllers/installer.md#list-installer-site-groups)
* [List Installer Sites](../../doc/controllers/installer.md#list-installer-sites)
* [Optimize Installer Rrm](../../doc/controllers/installer.md#optimize-installer-rrm)
* [Provision Installer Devices](../../doc/controllers/installer.md#provision-installer-devices)
* [Start Installer Locate Device](../../doc/controllers/installer.md#start-installer-locate-device)
* [Stop Installer Locate Device](../../doc/controllers/installer.md#stop-installer-locate-device)
* [Unassign Installer Recently Claimed Device](../../doc/controllers/installer.md#unassign-installer-recently-claimed-device)
* [Update Installer Map](../../doc/controllers/installer.md#update-installer-map)
* [Update Installer Virtual Chassis Member](../../doc/controllers/installer.md#update-installer-virtual-chassis-member)


# Add Installer Device Image

Add image

```go
AddInstallerDeviceImage(
    ctx context.Context,
    orgId uuid.UUID,
    imageName string,
    deviceMac string,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapImportJson) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `imageName` | `string` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `autoDeviceprofileAssignment` | `*bool` | Form, Optional | whether to auto assign device to deviceprofile by name |
| `csv` | `*models.FileWrapper` | Form, Optional | csv file for ap name mapping, optional |
| `file` | `*models.FileWrapper` | Form, Optional | ekahau or ibwave file |
| `json` | [`*models.MapImportJson`](../../doc/models/map-import-json.md) | Form, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageName := "image_name6"

deviceMac := "0000000000ab"

autoDeviceprofileAssignment := true





json := models.MapImportJson{
    ImportAllFloorplans:  models.ToPointer(false),
    ImportHeight:         models.ToPointer(true),
    ImportOrientation:    models.ToPointer(true),
    VendorName:           models.MapImportJsonVendorNameEnum_EKAHAU,
}

resp, err := installer.AddInstallerDeviceImage(ctx, orgId, imageName, deviceMac, &autoDeviceprofileAssignment, nil, nil, &json)
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


# Claim Installer Devices

This mirrors `POST /api/v1/orgs/{org_id}/inventory` (see Inventory API)

```go
ClaimInstallerDevices(
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

apiResponse, err := installer.ClaimInstallerDevices(ctx, orgId, body)
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


# Create Installer Map

Create a MAP

```go
CreateInstallerMap(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    mapId uuid.UUID,
    body *models.Map) (
    models.ApiResponse[models.Map],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteName` | `string` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Map`](../../doc/models/map.md) | Body, Optional | Request Body |

## Response Type

[`models.Map`](../../doc/models/map.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteName := "site_name8"

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Map{
    Height:               models.ToPointer(0),
    LatlngBr:             models.ToPointer(models.LatlngBr{
        Lat:                  models.ToPointer("string"),
        Lng:                  models.ToPointer("string"),
    }),
    LatlngTl:             models.ToPointer(models.LatlngTl{
        Lat:                  models.ToPointer("string"),
        Lng:                  models.ToPointer("string"),
    }),
    Locked:               models.ToPointer(true),
    Name:                 models.ToPointer("string"),
    Orientation:          models.ToPointer(0),
    OriginX:              models.ToPointer(0),
    OriginY:              models.ToPointer(0),
    Ppm:                  models.ToPointer(float64(0)),
    SitesurveyPath:       []models.MapSitesurveyPathItems{
        models.MapSitesurveyPathItems{
            Coordinate:           models.ToPointer("string"),
            Name:                 models.ToPointer("string"),
            Nodes:                []models.MapNode{
                models.MapNode{
                    Edges:                map[string]string{
                        "N2": "string",
                    },
                    Name:                 "string",
                    Position:             models.ToPointer(models.MapNodePosition{
                        X:                    float64(0),
                        Y:                    float64(0),
                    }),
                },
            },
        },
    },
    Type:                 models.ToPointer(models.MapTypeEnum_IMAGE),
    View:                 models.NewOptional(models.ToPointer(models.MapViewEnum_ROADMAP)),
    WallPath:             models.ToPointer(models.MapWallPath{
        Coordinate:           models.ToPointer("string"),
        Nodes:                []models.MapNode{
            models.MapNode{
                Edges:                map[string]string{
                    "N2": "string",
                },
                Name:                 "string",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(0),
                    Y:                    float64(0),
                }),
            },
        },
    }),
    Wayfinding:           models.ToPointer(models.MapWayfinding{
        Micello:              models.ToPointer(models.MapWayfindingMicello{
            AccountKey:           models.ToPointer("string"),
            DefaultLevelId:       models.ToPointer(0),
        }),
        SnapToPath:           models.ToPointer(true),
    }),
    WayfindingPath:       models.ToPointer(models.MapWayfindingPath{
        Coordinate:           models.ToPointer("string"),
        Nodes:                []models.MapNode{
            models.MapNode{
                Edges:                map[string]string{
                    "N2": "string",
                },
                Name:                 "string",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(0),
                    Y:                    float64(0),
                }),
            },
        },
    }),
    Width:                models.ToPointer(0),
}

apiResponse, err := installer.CreateInstallerMap(ctx, orgId, siteName, mapId, &body)
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


# Create Installer Virtual Chassis

For models (e.g. EX3400 and up) having dedicated VC ports, it is easier to form a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new VC and update the inventory.

In case that the user would like to choose the dedicated switch as a VC master or for EX2300-C-12P and EX2300-C-12T which doesn’t have dedicated VC ports, below are procedures to automate the VC creation:

1. Power on the switch that is choosen as the VC master first. And then powering on the other member switches.
2. Claim or adopt all these switches under the same organization’s Inventory
3. Assign these switches into the same Site
4. Invoke vc command on the switch choosen to be the VC master. For EX2300-C-12P, VC ports will be created automatically.
5. Connect the cables to the VC ports for these switches
6. Wait for the VC to be formed. The Org’s inventory will be updated for the new VC.

```go
CreateInstallerVirtualChassis(
    ctx context.Context,
    orgId uuid.UUID,
    fpc0Mac string,
    body *models.VirtualChassisConfig) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `fpc0Mac` | `string` | Template, Required | FPC0 MAC Address |
| `body` | [`*models.VirtualChassisConfig`](../../doc/models/virtual-chassis-config.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseVirtualChassisConfig`](../../doc/models/response-virtual-chassis-config.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

fpc0Mac := "fpc0_mac6"

body := models.VirtualChassisConfig{
    Members:              []models.VirtualChassisConfigMember{
        models.VirtualChassisConfigMember{
            Mac:                  "aff827549235",
            VcPorts:              []string{
                "xe-0/1/0",
                "xe-0/1/1",
            },
            VcRole:               models.VirtualChassisConfigMemberVcRoleEnum_MASTER,
        },
        models.VirtualChassisConfigMember{
            Mac:                  "8396cd006c8c",
            VcPorts:              []string{
                "xe-0/1/0",
                "xe-0/1/1",
            },
            VcRole:               models.VirtualChassisConfigMemberVcRoleEnum_BACKUP,
        },
        models.VirtualChassisConfigMember{
            Mac:                  "8396cd00888c",
            VcPorts:              []string{
                "xe-0/1/0",
                "xe-0/1/1",
            },
            VcRole:               models.VirtualChassisConfigMemberVcRoleEnum_LINECARD,
        },
    },
}

apiResponse, err := installer.CreateInstallerVirtualChassis(ctx, orgId, fpc0Mac, &body)
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
  "config_type": "nonprovisioned",
  "id": "00000000-0000-0000-1000-52d9107af289",
  "mac": "52d9107af289",
  "members": [
    {
      "_idx": 0,
      "boot_partition": "junos",
      "cpld_version": "6",
      "cpu_stat": {
        "idle": 69,
        "interrupt": 1,
        "load_avg": [
          0.8,
          1.09,
          1.07
        ],
        "system": 11,
        "user": 19
      },
      "fans": [
        {
          "airflow": "out",
          "name": "Fan Tray 0 Fan 0",
          "rpm": 0,
          "status": "ok"
        },
        {
          "airflow": "out",
          "name": "Fan Tray 1 Fan 0",
          "rpm": 0,
          "status": "ok"
        }
      ],
      "fpc_idx": 0,
      "mac": "52d9107af289",
      "memory_stat": {
        "usage": 42
      },
      "model": "EX2300-48P",
      "pics": [
        {
          "index": 0,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 48,
              "type": "GE"
            }
          ]
        },
        {
          "index": 1,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 4,
              "type": "SFP/SFP+"
            }
          ]
        }
      ],
      "poe": {
        "max_power": 750,
        "power_draw": 40.4,
        "status": "AT_MODE"
      },
      "poe_version": "2.1.1.19.3 (type1)",
      "psus": [
        {
          "name": "Power Supply 0",
          "status": "ok"
        }
      ],
      "recovery_version": "21.4R3-S4.18",
      "serial": "JW0000000000",
      "temperatures": [
        {
          "celsius": 33,
          "name": "CPU Sensor",
          "status": "ok"
        },
        {
          "celsius": 29,
          "name": "PSU Sensor",
          "status": "ok"
        }
      ],
      "type": "fpc",
      "uboot_version": "U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2",
      "uptime": 27636720,
      "vc_links": [
        {
          "neighbor_module_idx": 1,
          "neighbor_port_id": "vcp-1/1/0",
          "port_id": "vcp-0/1/0"
        }
      ],
      "vc_mode": "HiGiG",
      "vc_role": "master",
      "vc_state": "present",
      "version": "21.4R3-S4.18"
    },
    {
      "_idx": 1,
      "boot_partition": "junos",
      "cpld_version": "6",
      "cpu_stat": {
        "idle": 76,
        "interrupt": 0,
        "load_avg": [
          0.96,
          0.87,
          0.76
        ],
        "system": 6,
        "user": 17
      },
      "fans": [
        {
          "airflow": "out",
          "name": "Fan Tray 0 Fan 0",
          "rpm": 0,
          "status": "ok"
        },
        {
          "airflow": "out",
          "name": "Fan Tray 1 Fan 0",
          "rpm": 0,
          "status": "ok"
        }
      ],
      "fpc_idx": 1,
      "mac": "d0dd4991652d",
      "memory_stat": {
        "usage": 18
      },
      "model": "EX2300-48P",
      "pics": [
        {
          "index": 0,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 48,
              "type": "GE"
            }
          ]
        },
        {
          "index": 1,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 4,
              "type": "SFP/SFP+"
            }
          ]
        }
      ],
      "poe": {
        "max_power": 750,
        "power_draw": 21.2,
        "status": "AT_MODE"
      },
      "poe_version": "2.1.1.19.3 (type1)",
      "psus": [
        {
          "name": "Power Supply 0",
          "status": "ok"
        }
      ],
      "recovery_version": "21.4R3-S4.18",
      "serial": "JW3619300922",
      "temperatures": [
        {
          "celsius": 32,
          "name": "CPU Sensor",
          "status": "ok"
        },
        {
          "celsius": 29,
          "name": "PSU Sensor",
          "status": "ok"
        }
      ],
      "type": "fpc",
      "uboot_version": "U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2",
      "uptime": 27636720,
      "vc_links": [
        {
          "neighbor_module_idx": 0,
          "neighbor_port_id": "vcp-0/1/0",
          "port_id": "vcp-1/1/0"
        }
      ],
      "vc_mode": "HiGiG",
      "vc_role": "backup",
      "vc_state": "present",
      "version": "21.4R3-S4.18"
    }
  ],
  "model": "EX2300-48P",
  "org_id": "1e9a61a9-bc42-42ca-bf67-1ad87443d9b8",
  "serial": "JW3619300157",
  "site_id": "ab0aca7a-d45c-469b-b3bb-4fe240642d0b",
  "status": "connected",
  "type": "switch",
  "vc_mac": "52d9107af289"
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


# Create or Update Installer Sites

Often the Installers are asked to assign Devices to Sites. The Sites can either be pre-created or created/modified by the Installer. If this is an update, the same grace period also applies.

```go
CreateOrUpdateInstallerSites(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    body *models.InstallerSite) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteName` | `string` | Template, Required | - |
| `body` | [`*models.InstallerSite`](../../doc/models/installer-site.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteName := "site_name8"

body := models.InstallerSite{
    Address:              "1601 S. Deanza Blvd., Cupertino, CA, 95014",
    CountryCode:          "US",
    Latlng:               models.LatLng{
        Lat:                  float64(37.295833),
        Lng:                  float64(-122.032946),
    },
    Name:                 "string",
    RftemplateName:       models.ToPointer("rftemplate1"),
    SitegroupNames:       []string{
        "sg1",
        "sg2",
    },
    Timezone:             models.ToPointer("America/Los_Angeles"),
}

resp, err := installer.CreateOrUpdateInstallerSites(ctx, orgId, siteName, &body)
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


# Delete Installer Device Image

delete image

```go
DeleteInstallerDeviceImage(
    ctx context.Context,
    orgId uuid.UUID,
    imageName string,
    deviceMac string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `imageName` | `string` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageName := "image_name6"

deviceMac := "0000000000ab"

resp, err := installer.DeleteInstallerDeviceImage(ctx, orgId, imageName, deviceMac)
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


# Delete Installer Map

Delete Map

```go
DeleteInstallerMap(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    mapId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteName` | `string` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteName := "site_name8"

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := installer.DeleteInstallerMap(ctx, orgId, siteName, mapId)
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


# Get Installer Device Virtual Chassis

Get VC Status

The API returns a combined view of the VC status which includes topology and stats

```go
GetInstallerDeviceVirtualChassis(
    ctx context.Context,
    orgId uuid.UUID,
    fpc0Mac string) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `fpc0Mac` | `string` | Template, Required | FPC0 MAC Address |

## Response Type

[`models.ResponseVirtualChassisConfig`](../../doc/models/response-virtual-chassis-config.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

fpc0Mac := "fpc0_mac6"

apiResponse, err := installer.GetInstallerDeviceVirtualChassis(ctx, orgId, fpc0Mac)
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
  "config_type": "nonprovisioned",
  "id": "00000000-0000-0000-1000-52d9107af289",
  "mac": "52d9107af289",
  "members": [
    {
      "_idx": 0,
      "boot_partition": "junos",
      "cpld_version": "6",
      "cpu_stat": {
        "idle": 69,
        "interrupt": 1,
        "load_avg": [
          0.8,
          1.09,
          1.07
        ],
        "system": 11,
        "user": 19
      },
      "fans": [
        {
          "airflow": "out",
          "name": "Fan Tray 0 Fan 0",
          "rpm": 0,
          "status": "ok"
        },
        {
          "airflow": "out",
          "name": "Fan Tray 1 Fan 0",
          "rpm": 0,
          "status": "ok"
        }
      ],
      "fpc_idx": 0,
      "mac": "52d9107af289",
      "memory_stat": {
        "usage": 42
      },
      "model": "EX2300-48P",
      "pics": [
        {
          "index": 0,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 48,
              "type": "GE"
            }
          ]
        },
        {
          "index": 1,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 4,
              "type": "SFP/SFP+"
            }
          ]
        }
      ],
      "poe": {
        "max_power": 750,
        "power_draw": 40.4,
        "status": "AT_MODE"
      },
      "poe_version": "2.1.1.19.3 (type1)",
      "psus": [
        {
          "name": "Power Supply 0",
          "status": "ok"
        }
      ],
      "recovery_version": "21.4R3-S4.18",
      "serial": "JW0000000000",
      "temperatures": [
        {
          "celsius": 33,
          "name": "CPU Sensor",
          "status": "ok"
        },
        {
          "celsius": 29,
          "name": "PSU Sensor",
          "status": "ok"
        }
      ],
      "type": "fpc",
      "uboot_version": "U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2",
      "uptime": 27636720,
      "vc_links": [
        {
          "neighbor_module_idx": 1,
          "neighbor_port_id": "vcp-1/1/0",
          "port_id": "vcp-0/1/0"
        }
      ],
      "vc_mode": "HiGiG",
      "vc_role": "master",
      "vc_state": "present",
      "version": "21.4R3-S4.18"
    },
    {
      "_idx": 1,
      "boot_partition": "junos",
      "cpld_version": "6",
      "cpu_stat": {
        "idle": 76,
        "interrupt": 0,
        "load_avg": [
          0.96,
          0.87,
          0.76
        ],
        "system": 6,
        "user": 17
      },
      "fans": [
        {
          "airflow": "out",
          "name": "Fan Tray 0 Fan 0",
          "rpm": 0,
          "status": "ok"
        },
        {
          "airflow": "out",
          "name": "Fan Tray 1 Fan 0",
          "rpm": 0,
          "status": "ok"
        }
      ],
      "fpc_idx": 1,
      "mac": "d0dd4991652d",
      "memory_stat": {
        "usage": 18
      },
      "model": "EX2300-48P",
      "pics": [
        {
          "index": 0,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 48,
              "type": "GE"
            }
          ]
        },
        {
          "index": 1,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 4,
              "type": "SFP/SFP+"
            }
          ]
        }
      ],
      "poe": {
        "max_power": 750,
        "power_draw": 21.2,
        "status": "AT_MODE"
      },
      "poe_version": "2.1.1.19.3 (type1)",
      "psus": [
        {
          "name": "Power Supply 0",
          "status": "ok"
        }
      ],
      "recovery_version": "21.4R3-S4.18",
      "serial": "JW3619300922",
      "temperatures": [
        {
          "celsius": 32,
          "name": "CPU Sensor",
          "status": "ok"
        },
        {
          "celsius": 29,
          "name": "PSU Sensor",
          "status": "ok"
        }
      ],
      "type": "fpc",
      "uboot_version": "U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2",
      "uptime": 27636720,
      "vc_links": [
        {
          "neighbor_module_idx": 0,
          "neighbor_port_id": "vcp-0/1/0",
          "port_id": "vcp-1/1/0"
        }
      ],
      "vc_mode": "HiGiG",
      "vc_role": "backup",
      "vc_state": "present",
      "version": "21.4R3-S4.18"
    }
  ],
  "model": "EX2300-48P",
  "org_id": "1e9a61a9-bc42-42ca-bf67-1ad87443d9b8",
  "serial": "JW3619300157",
  "site_id": "ab0aca7a-d45c-469b-b3bb-4fe240642d0b",
  "status": "connected",
  "type": "switch",
  "vc_mac": "52d9107af289"
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


# Import Installer Map

Import data from files is a multipart POST which has an file, an optional json, and an optional csv, to create floorplan, assign & place ap if name or mac matches

```go
ImportInstallerMap(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapImportJson) (
    models.ApiResponse[models.ResponseMapImport],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteName` | `string` | Template, Required | - |
| `autoDeviceprofileAssignment` | `*bool` | Form, Optional | whether to auto assign device to deviceprofile by name |
| `csv` | `*models.FileWrapper` | Form, Optional | csv file for ap name mapping, optional |
| `file` | `*models.FileWrapper` | Form, Optional | ekahau or ibwave file |
| `json` | [`*models.MapImportJson`](../../doc/models/map-import-json.md) | Form, Optional | - |

## Response Type

[`models.ResponseMapImport`](../../doc/models/response-map-import.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteName := "site_name8"

autoDeviceprofileAssignment := true





json := models.MapImportJson{
    ImportAllFloorplans:  models.ToPointer(false),
    ImportHeight:         models.ToPointer(true),
    ImportOrientation:    models.ToPointer(true),
    VendorName:           models.MapImportJsonVendorNameEnum_EKAHAU,
}

apiResponse, err := installer.ImportInstallerMap(ctx, orgId, siteName, &autoDeviceprofileAssignment, nil, nil, &json)
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
  "aps": [
    {
      "action": "assigned-placed",
      "floorplan_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
      "height": 3,
      "mac": "5c5b35000001",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "orientation": 45
    }
  ],
  "floorplans": [
    {
      "action": "imported",
      "id": "cbdb7f0b-3be0-4872-88f9-58790b509c23",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "name": "map1"
    }
  ],
  "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
  "summary": {
    "num_ap_assigned": 1,
    "num_inv_assigned": 1,
    "num_map_assigned": 1
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


# List Installer Alarm Templates

Get List of alarm templates

```go
ListInstallerAlarmTemplates(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.InstallersItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.InstallersItem`](../../doc/models/installers-item.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := installer.ListInstallerAlarmTemplates(ctx, orgId)
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
    "id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
    "name": "AlarmTemplate 1"
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


# List Installer Device Profiles

Get List of Device Profiles

```go
ListInstallerDeviceProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum) (
    models.ApiResponse[[]models.InstallersItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | **Default**: `"ap"` |

## Response Type

[`[]models.InstallersItem`](../../doc/models/installers-item.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum_AP

apiResponse, err := installer.ListInstallerDeviceProfiles(ctx, orgId, &mType)
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
    "id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
    "name": "DeviceProfile 1"
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


# List Installer List of Renctly Claimed Devices

Get List of recently claimed devices

```go
ListInstallerListOfRenctlyClaimedDevices(
    ctx context.Context,
    orgId uuid.UUID,
    model *string,
    siteName *string,
    siteId *uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.InstallerDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `model` | `*string` | Query, Optional | Device Model |
| `siteName` | `*string` | Query, Optional | Site Name |
| `siteId` | `*uuid.UUID` | Query, Optional | Site ID |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.InstallerDevice`](../../doc/models/installer-device.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")







limit := 100

page := 1

apiResponse, err := installer.ListInstallerListOfRenctlyClaimedDevices(ctx, orgId, nil, nil, nil, &limit, &page)
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
    "deviceprofile_name": "SJ1",
    "height": 2.7,
    "mac": "5c5b35000018",
    "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
    "model": "AP41",
    "name": "hallway",
    "orientation": 90,
    "serial": "FXLH2015150025",
    "site_name": "SJ1",
    "x": 150,
    "y": 300
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


# List Installer Maps

Get List of Maps

```go
ListInstallerMaps(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string) (
    models.ApiResponse[[]models.Map],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteName` | `string` | Template, Required | - |

## Response Type

[`[]models.Map`](../../doc/models/map.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteName := "site_name8"

apiResponse, err := installer.ListInstallerMaps(ctx, orgId, siteName)
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
    "created_time": 0,
    "flags": {},
    "height": 0,
    "height_m": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "latlng_br": {
      "lat": "string",
      "lng": "string"
    },
    "latlng_tl": {
      "lat": "string",
      "lng": "string"
    },
    "locked": true,
    "modified_time": 0,
    "name": "string",
    "occupancy_limit": 0,
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "orientation": 0,
    "origin_x": 0,
    "origin_y": 0,
    "ppm": 0,
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "sitesurvey_path": [
      {
        "coordinate": "string",
        "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
        "name": "string",
        "nodes": [
          {
            "edges": {
              "N2": "string"
            },
            "name": "string",
            "position": {
              "x": 0,
              "y": 0
            }
          }
        ]
      }
    ],
    "thumbnail_url": "string",
    "type": "image",
    "url": "string",
    "view": "roadmap",
    "wall_path": {
      "coordinate": "string",
      "nodes": [
        {
          "edges": {
            "N2": "string"
          },
          "name": "string",
          "position": {
            "x": 0,
            "y": 0
          }
        }
      ]
    },
    "wayfinding": {
      "micello": {
        "account_key": "string",
        "default_level_id": 0,
        "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
      },
      "snap_to_path": true
    },
    "wayfinding_path": {
      "coordinate": "string",
      "nodes": [
        {
          "edges": {
            "N2": "string"
          },
          "name": "string",
          "position": {
            "x": 0,
            "y": 0
          }
        }
      ]
    },
    "width": 0,
    "width_m": 0
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


# List Installer Rf Templates Names

Get List of RF Templates

```go
ListInstallerRfTemplatesNames(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.InstallersItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.InstallersItem`](../../doc/models/installers-item.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := installer.ListInstallerRfTemplatesNames(ctx, orgId)
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
    "id": "bb8a9017-1e36-5d6c-6f2b-551abe8a76a2",
    "name": "RFTemplate 1"
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


# List Installer Site Groups

Get List of Site Groups

```go
ListInstallerSiteGroups(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.InstallersItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.InstallersItem`](../../doc/models/installers-item.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := installer.ListInstallerSiteGroups(ctx, orgId)
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
    "id": "581328b6-e382-f54e-c9dc-999983183a34",
    "name": "SiteGroup 1"
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


# List Installer Sites

Get List of Sites

```go
ListInstallerSites(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.InstallerSite],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.InstallerSite`](../../doc/models/installer-site.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := installer.ListInstallerSites(ctx, orgId)
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
    "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
    "country_code": "US",
    "id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
    "latlng": {
      "lat": 37.295833,
      "lng": -122.032946
    },
    "name": "Mist Office",
    "rftemplate_name": "rftemplate1",
    "sitegroup_names": [
      "sg1",
      "sg2"
    ],
    "timezone": "America/Los_Angeles"
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


# Optimize Installer Rrm

After installation is considered complete (APs are placed on maps, all powered up), you can trigger an optimize operation where RRM will kick in (and maybe other things in the future) before it’s automatically scheduled.

```go
OptimizeInstallerRrm(
    ctx context.Context,
    siteName string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteName` | `string` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteName := "site_name8"

resp, err := installer.OptimizeInstallerRrm(ctx, siteName)
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


# Provision Installer Devices

Provision or Replace a device

If replacing_mac is in the request payload, other attributes are ignored, we attempt to replace existing device (with mac replacing_mac) with the inventory device being configured. The replacement device must be in the inventory but not assigned, and the replacing_mac device must be assigned to a site, and satisfy grace period requirements. The Device replaced will become unassigned.

```go
ProvisionInstallerDevices(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string,
    body *models.InstallerProvisionDevice) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `body` | [`*models.InstallerProvisionDevice`](../../doc/models/installer-provision-device.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

body := models.InstallerProvisionDevice{
    DeviceprofileName:    models.ToPointer("SJ1"),
    Height:               models.ToPointer(float64(2.7)),
    Name:                 "SJ1-AP1",
    Orientation:          models.ToPointer(90),
    SiteId:               models.ToPointer(uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")),
    SiteName:             models.ToPointer("SJ1"),
    X:                    models.ToPointer(float64(150)),
    Y:                    models.ToPointer(float64(300)),
}

resp, err := installer.ProvisionInstallerDevices(ctx, orgId, deviceMac, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Request | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not Found | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Start Installer Locate Device

Locate a Device by blinking it’s LED, it’s a persisted state that has to be stopped by calling Stop Locating API

```go
StartInstallerLocateDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

resp, err := installer.StartInstallerLocateDevice(ctx, orgId, deviceMac)
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


# Stop Installer Locate Device

Stop it

```go
StopInstallerLocateDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

resp, err := installer.StopInstallerLocateDevice(ctx, orgId, deviceMac)
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


# Unassign Installer Recently Claimed Device

Unassign recently claimed devices

```go
UnassignInstallerRecentlyClaimedDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

resp, err := installer.UnassignInstallerRecentlyClaimedDevice(ctx, orgId, deviceMac)
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


# Update Installer Map

Update map

```go
UpdateInstallerMap(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    mapId uuid.UUID,
    body *models.Map) (
    models.ApiResponse[models.Map],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteName` | `string` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Map`](../../doc/models/map.md) | Body, Optional | Request Body |

## Response Type

[`models.Map`](../../doc/models/map.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteName := "site_name8"

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Map{
    Height:               models.ToPointer(1500),
    Locked:               models.ToPointer(false),
    Name:                 models.ToPointer("Mist Office"),
    Orientation:          models.ToPointer(30),
    OriginX:              models.ToPointer(35),
    OriginY:              models.ToPointer(60),
    Ppm:                  models.ToPointer(float64(40.94)),
    Type:                 models.ToPointer(models.MapTypeEnum_IMAGE),
    UseAutoOrientation:   models.ToPointer(false),
    UseAutoPlacement:     models.ToPointer(false),
    Width:                models.ToPointer(1250),
}

apiResponse, err := installer.UpdateInstallerMap(ctx, orgId, siteName, mapId, &body)
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


# Update Installer Virtual Chassis Member

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

Please notice that member ID 0 (fpc0) cannot be removed. When a VC has two switches left, unplugging the cable may result in the situation that fpc0 becomes a line card (LC). When this situation is happening, please re-plug in the cable, wait for both switches becoming present (show virtual-chassis) and then removing the cable again.

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
UpdateInstallerVirtualChassisMember(
    ctx context.Context,
    orgId uuid.UUID,
    fpc0Mac string,
    body *models.VirtualChassisUpdate) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `fpc0Mac` | `string` | Template, Required | FPC0 MAC Address |
| `body` | [`*models.VirtualChassisUpdate`](../../doc/models/virtual-chassis-update.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseVirtualChassisConfig`](../../doc/models/response-virtual-chassis-config.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

fpc0Mac := "fpc0_mac6"

body := models.VirtualChassisUpdate{
    Members:              []models.VirtualChassisMemberUpdate{
        models.VirtualChassisMemberUpdate{
            Mac:                  models.ToPointer("aff827549235"),
            MemberId:             models.ToPointer(2),
            VcPorts:              []string{
                "xe-0/1/0",
                "xe-0/1/1",
            },
            VcRole:               models.ToPointer(models.VirtualChassisMemberUpdateVcRoleEnum_LINECARD),
        },
        models.VirtualChassisMemberUpdate{
            Mac:                  models.ToPointer("8396cd00777c"),
            MemberId:             models.ToPointer(3),
            VcPorts:              []string{
                "xe-0/1/0",
                "xe-0/1/1",
            },
            VcRole:               models.ToPointer(models.VirtualChassisMemberUpdateVcRoleEnum_LINECARD),
        },
    },
    Op:                   models.ToPointer(models.VirtualChassisUpdateOpEnum_ADD),
}

apiResponse, err := installer.UpdateInstallerVirtualChassisMember(ctx, orgId, fpc0Mac, &body)
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
  "config_type": "nonprovisioned",
  "id": "00000000-0000-0000-1000-52d9107af289",
  "mac": "52d9107af289",
  "members": [
    {
      "_idx": 0,
      "boot_partition": "junos",
      "cpld_version": "6",
      "cpu_stat": {
        "idle": 69,
        "interrupt": 1,
        "load_avg": [
          0.8,
          1.09,
          1.07
        ],
        "system": 11,
        "user": 19
      },
      "fans": [
        {
          "airflow": "out",
          "name": "Fan Tray 0 Fan 0",
          "rpm": 0,
          "status": "ok"
        },
        {
          "airflow": "out",
          "name": "Fan Tray 1 Fan 0",
          "rpm": 0,
          "status": "ok"
        }
      ],
      "fpc_idx": 0,
      "mac": "52d9107af289",
      "memory_stat": {
        "usage": 42
      },
      "model": "EX2300-48P",
      "pics": [
        {
          "index": 0,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 48,
              "type": "GE"
            }
          ]
        },
        {
          "index": 1,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 4,
              "type": "SFP/SFP+"
            }
          ]
        }
      ],
      "poe": {
        "max_power": 750,
        "power_draw": 40.4,
        "status": "AT_MODE"
      },
      "poe_version": "2.1.1.19.3 (type1)",
      "psus": [
        {
          "name": "Power Supply 0",
          "status": "ok"
        }
      ],
      "recovery_version": "21.4R3-S4.18",
      "serial": "JW0000000000",
      "temperatures": [
        {
          "celsius": 33,
          "name": "CPU Sensor",
          "status": "ok"
        },
        {
          "celsius": 29,
          "name": "PSU Sensor",
          "status": "ok"
        }
      ],
      "type": "fpc",
      "uboot_version": "U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2",
      "uptime": 27636720,
      "vc_links": [
        {
          "neighbor_module_idx": 1,
          "neighbor_port_id": "vcp-1/1/0",
          "port_id": "vcp-0/1/0"
        }
      ],
      "vc_mode": "HiGiG",
      "vc_role": "master",
      "vc_state": "present",
      "version": "21.4R3-S4.18"
    },
    {
      "_idx": 1,
      "boot_partition": "junos",
      "cpld_version": "6",
      "cpu_stat": {
        "idle": 76,
        "interrupt": 0,
        "load_avg": [
          0.96,
          0.87,
          0.76
        ],
        "system": 6,
        "user": 17
      },
      "fans": [
        {
          "airflow": "out",
          "name": "Fan Tray 0 Fan 0",
          "rpm": 0,
          "status": "ok"
        },
        {
          "airflow": "out",
          "name": "Fan Tray 1 Fan 0",
          "rpm": 0,
          "status": "ok"
        }
      ],
      "fpc_idx": 1,
      "mac": "d0dd4991652d",
      "memory_stat": {
        "usage": 18
      },
      "model": "EX2300-48P",
      "pics": [
        {
          "index": 0,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 48,
              "type": "GE"
            }
          ]
        },
        {
          "index": 1,
          "model_number": "EX2300-48P",
          "port_groups": [
            {
              "count": 4,
              "type": "SFP/SFP+"
            }
          ]
        }
      ],
      "poe": {
        "max_power": 750,
        "power_draw": 21.2,
        "status": "AT_MODE"
      },
      "poe_version": "2.1.1.19.3 (type1)",
      "psus": [
        {
          "name": "Power Supply 0",
          "status": "ok"
        }
      ],
      "recovery_version": "21.4R3-S4.18",
      "serial": "JW3619300922",
      "temperatures": [
        {
          "celsius": 32,
          "name": "CPU Sensor",
          "status": "ok"
        },
        {
          "celsius": 29,
          "name": "PSU Sensor",
          "status": "ok"
        }
      ],
      "type": "fpc",
      "uboot_version": "U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2",
      "uptime": 27636720,
      "vc_links": [
        {
          "neighbor_module_idx": 0,
          "neighbor_port_id": "vcp-0/1/0",
          "port_id": "vcp-1/1/0"
        }
      ],
      "vc_mode": "HiGiG",
      "vc_role": "backup",
      "vc_state": "present",
      "version": "21.4R3-S4.18"
    }
  ],
  "model": "EX2300-48P",
  "org_id": "1e9a61a9-bc42-42ca-bf67-1ad87443d9b8",
  "serial": "JW3619300157",
  "site_id": "ab0aca7a-d45c-469b-b3bb-4fe240642d0b",
  "status": "connected",
  "type": "switch",
  "vc_mac": "52d9107af289"
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

