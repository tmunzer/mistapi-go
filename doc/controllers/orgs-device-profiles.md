# Orgs Device Profiles

```go
orgsDeviceProfiles := client.OrgsDeviceProfiles()
```

## Class Name

`OrgsDeviceProfiles`

## Methods

* [Assign Org Device Profile](../../doc/controllers/orgs-device-profiles.md#assign-org-device-profile)
* [Create Org Device Profiles](../../doc/controllers/orgs-device-profiles.md#create-org-device-profiles)
* [Delete Org Device Profile](../../doc/controllers/orgs-device-profiles.md#delete-org-device-profile)
* [Get Org Device Profile](../../doc/controllers/orgs-device-profiles.md#get-org-device-profile)
* [List Org Device Profiles](../../doc/controllers/orgs-device-profiles.md#list-org-device-profiles)
* [Unassign Org Device Profile](../../doc/controllers/orgs-device-profiles.md#unassign-org-device-profile)
* [Update Org Device Profile](../../doc/controllers/orgs-device-profiles.md#update-org-device-profile)


# Assign Org Device Profile

Assign Org Device Profile to Devices

```go
AssignOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseAssignSuccess`](../../doc/models/response-assign-success.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "5c5b350e0001",
        "5c5b350e0003",
    },
}

apiResponse, err := orgsDeviceProfiles.AssignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
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


# Create Org Device Profiles

Create Device Profile

```go
CreateOrgDeviceProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Deviceprofile) (
    models.ApiResponse[models.Deviceprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Deviceprofile`](../../doc/models/containers/deviceprofile.md) | Body, Optional | Request Body |

## Response Type

[`models.Deviceprofile`](../../doc/models/containers/deviceprofile.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.DeviceprofileContainer.FromDeviceprofileAp(models.DeviceprofileAp{
    Aeroscout:            models.ToPointer(models.ApAeroscout{
        Enabled:              models.ToPointer(false),
        Host:                 models.NewOptional(models.ToPointer("aero.pvt.net")),
        LocateConnected:      models.ToPointer(true),
    }),
    Led:                  models.ToPointer(models.ApLed{
        Brightness:           models.ToPointer(255),
        Enabled:              models.ToPointer(true),
    }),
    Name:                 models.ToPointer("string"),
    NtpServers:           []string{
        "10.10.10.10",
    },
    Type:                 "ap",
    UsbConfig:            models.ToPointer(models.ApUsb{
        Cacert:               models.NewOptional(models.ToPointer("string")),
        Channel:              models.ToPointer(3),
        Enabled:              models.ToPointer(true),
        Host:                 models.ToPointer("1.1.1.1"),
        Port:                 models.ToPointer(0),
        Type:                 models.ToPointer(models.ApUsbTypeEnum("imagotag")),
        VerifyCert:           models.ToPointer(true),
        VlanId:               models.ToPointer(1),
    }),
})

apiResponse, err := orgsDeviceProfiles.CreateOrgDeviceProfiles(ctx, orgId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceprofileAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *r)
    } else if r, ok := responseBody.AsDeviceprofileGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "aeroscout": {
    "enabled": true,
    "host": "string"
  },
  "created_time": 0,
  "disable_eth1": true,
  "disable_module": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ip_config": {
    "dns": [
      "string"
    ],
    "dns_suffix": [
      "string"
    ],
    "gateway": "192.168.0.1",
    "gateway6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "ip": "192.168.0.1",
    "ip6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "mtu": 0,
    "netmask": "192.168.0.1",
    "netmask6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "mesh": {
    "enabled": true,
    "group": 1,
    "role": "base"
  },
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "poe_passthrough": true,
  "switch_config": {
    "enabled": true,
    "eth0": {
      "port_vlan_id": 1,
      "vlan_ids": [
        1,
        10
      ]
    },
    "eth1": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "eth2": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "eth3": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "module": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "wds": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    }
  },
  "type": "ap",
  "usb_config": {
    "cacert": "string",
    "enabled": true,
    "host": "string",
    "type": "imagotag",
    "verify_cert": true
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


# Delete Org Device Profile

Delete Org Device Profile

```go
DeleteOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsDeviceProfiles.DeleteOrgDeviceProfile(ctx, orgId, deviceprofileId)
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


# Get Org Device Profile

Get Org device Profile Details

```go
GetOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID) (
    models.ApiResponse[models.Deviceprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Deviceprofile`](../../doc/models/containers/deviceprofile.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDeviceProfiles.GetOrgDeviceProfile(ctx, orgId, deviceprofileId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceprofileAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *r)
    } else if r, ok := responseBody.AsDeviceprofileGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "aeroscout": {
    "enabled": true,
    "host": "string"
  },
  "created_time": 0,
  "disable_eth1": true,
  "disable_module": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ip_config": {
    "dns": [
      "string"
    ],
    "dns_suffix": [
      "string"
    ],
    "gateway": "192.168.0.1",
    "gateway6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "ip": "192.168.0.1",
    "ip6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "mtu": 0,
    "netmask": "192.168.0.1",
    "netmask6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "mesh": {
    "enabled": true,
    "group": 1,
    "role": "base"
  },
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "poe_passthrough": true,
  "switch_config": {
    "enabled": true,
    "eth0": {
      "port_vlan_id": 1,
      "vlan_ids": [
        1,
        10
      ]
    },
    "eth1": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "eth2": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "eth3": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "module": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "wds": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    }
  },
  "type": "ap",
  "usb_config": {
    "cacert": "string",
    "enabled": true,
    "host": "string",
    "type": "imagotag",
    "verify_cert": true
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


# List Org Device Profiles

Get List of Org Device Profiles

```go
ListOrgDeviceProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Deviceprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | **Default**: `"ap"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Deviceprofile`](../../doc/models/containers/deviceprofile.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum("ap")

limit := 100

page := 1

apiResponse, err := orgsDeviceProfiles.ListOrgDeviceProfiles(ctx, orgId, &mType, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsDeviceprofileAp(); ok {
            fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *i)
        } else if i, ok := item.AsDeviceprofileGateway(); ok {
            fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *i)
        }
    }

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


# Unassign Org Device Profile

Unassign Org Device Profile from Devices

```go
UnassignOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseAssignSuccess`](../../doc/models/response-assign-success.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "5c5b350e0001",
        "5c5b350e0003",
    },
}

apiResponse, err := orgsDeviceProfiles.UnassignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
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


# Update Org Device Profile

Update Org Device Profile

```go
UpdateOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.Deviceprofile) (
    models.ApiResponse[models.Deviceprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Deviceprofile`](../../doc/models/containers/deviceprofile.md) | Body, Optional | Request Body |

## Response Type

[`models.Deviceprofile`](../../doc/models/containers/deviceprofile.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.DeviceprofileContainer.FromDeviceprofileAp(models.DeviceprofileAp{
    Aeroscout:            models.ToPointer(models.ApAeroscout{
        Enabled:              models.ToPointer(true),
        Host:                 models.NewOptional(models.ToPointer("string")),
    }),
    DisableEth1:          models.ToPointer(true),
    DisableModule:        models.ToPointer(true),
    Mesh:                 models.ToPointer(models.ApMesh{
        Enabled:              models.ToPointer(true),
        Group:                models.NewOptional(models.ToPointer(1)),
        Role:                 models.ToPointer(models.ApMeshRoleEnum("base")),
    }),
    Name:                 models.ToPointer("string"),
    PoePassthrough:       models.ToPointer(true),
    RadioConfig:          models.ToPointer(models.ApRadio{
        AntGain24:            models.ToPointer(0),
        AntGain5:             models.ToPointer(0),
        Band24:               models.ToPointer(models.ApRadioBand24{
            AllowRrmDisable:      models.ToPointer(true),
            AntennaMode:          models.ToPointer(models.RadioBandAntennaModeEnum("default")),
            Bandwidth:            models.ToPointer(models.Dot11Bandwidth24Enum(20)),
            Channel:              models.NewOptional(models.ToPointer(6)),
            Disabled:             models.ToPointer(true),
            Power:                models.NewOptional(models.ToPointer(8)),
            Preamble:             models.ToPointer(models.RadioBandPreambleEnum("auto")),
            AdditionalProperties: map[string]interface{}{
                "usage": interface{}("24"),
            },
        }),
        Band24Usage:          models.ToPointer(models.RadioBand24UsageEnum("24")),
        Band5:                models.ToPointer(models.ApRadioBand5{
            AllowRrmDisable:      models.ToPointer(true),
            AntennaMode:          models.ToPointer(models.RadioBandAntennaModeEnum("default")),
            Bandwidth:            models.ToPointer(models.Dot11Bandwidth5Enum(20)),
            Channel:              models.NewOptional(models.ToPointer(50)),
            Disabled:             models.ToPointer(true),
            PowerMax:             models.NewOptional(models.ToPointer(8)),
            PowerMin:             models.NewOptional(models.ToPointer(15)),
            Preamble:             models.ToPointer(models.RadioBandPreambleEnum("auto")),
            AdditionalProperties: map[string]interface{}{
                "usage": interface{}("24"),
            },
        }),
        ScanningEnabled:      models.ToPointer(true),
    }),
    Type:                 "ap",
})

apiResponse, err := orgsDeviceProfiles.UpdateOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceprofileAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *r)
    } else if r, ok := responseBody.AsDeviceprofileGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "aeroscout": {
    "enabled": true,
    "host": "string"
  },
  "created_time": 0,
  "disable_eth1": true,
  "disable_module": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ip_config": {
    "dns": [
      "string"
    ],
    "dns_suffix": [
      "string"
    ],
    "gateway": "192.168.0.1",
    "gateway6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "ip": "192.168.0.1",
    "ip6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "mtu": 0,
    "netmask": "192.168.0.1",
    "netmask6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "mesh": {
    "enabled": true,
    "group": 1,
    "role": "base"
  },
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "poe_passthrough": true,
  "switch_config": {
    "enabled": true,
    "eth0": {
      "port_vlan_id": 1,
      "vlan_ids": [
        1,
        10
      ]
    },
    "eth1": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "eth2": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "eth3": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "module": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    },
    "wds": {
      "port_vlan_id": 1,
      "vlan_ids": [
        10
      ]
    }
  },
  "type": "ap",
  "usb_config": {
    "cacert": "string",
    "enabled": true,
    "host": "string",
    "type": "imagotag",
    "verify_cert": true
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

