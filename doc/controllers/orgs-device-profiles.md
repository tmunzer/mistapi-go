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
    Macs: []string{
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
    Aeroscout:        models.ToPointer(models.ApAeroscout{
        Enabled:         models.ToPointer(false),
        Host:            models.NewOptional(models.ToPointer("aero.pvt.net")),
        LocateConnected: models.ToPointer(true),
    }),
    Led:              models.ToPointer(models.ApLed{
        Brightness: models.ToPointer(255),
        Enabled:    models.ToPointer(true),
    }),
    Name:             models.ToPointer("string"),
    NtpServers:       []string{
        "10.10.10.10",
    },
    Type:             "ap",
    UsbConfig:        models.ToPointer(models.ApUsb{
        Cacert:     models.NewOptional(models.ToPointer("string")),
        Channel:    models.ToPointer(3),
        Enabled:    models.ToPointer(true),
        Host:       models.ToPointer("1.1.1.1"),
        Port:       models.ToPointer(0),
        Type:       models.ToPointer(models.ApUsbTypeEnum("imagotag")),
        VerifyCert: models.ToPointer(true),
        VlanId:     models.ToPointer(1),
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

## Example Response

```
[
  {
    "aeroscout": {
      "enabled": false,
      "host": "aero.pvt.net",
      "locate_connected": true
    },
    "ble_config": {
      "beacon_enabled": false,
      "beacon_rate": 3,
      "beacon_rate_mode": "custom",
      "beam_disabled": [
        1,
        3,
        6
      ],
      "custom_ble_packet_enabled": false,
      "custom_ble_packet_frame": "0x........",
      "custom_ble_packet_freq_msec": 300,
      "eddystone_uid_adv_power": -65,
      "eddystone_uid_beams": "2-4,7",
      "eddystone_uid_enabled": false,
      "eddystone_uid_freq_msec": 200,
      "eddystone_uid_instance": "5c5b35000001",
      "eddystone_uid_namespace": "2818e3868dec25629ede",
      "eddystone_url_adv_power": -65,
      "eddystone_url_beams": "2-4,7",
      "eddystone_url_enabled": true,
      "eddystone_url_freq_msec": 1000,
      "eddystone_url_url": "https://www.abc.com",
      "ibeacon_adv_power": -65,
      "ibeacon_beams": "2-4,7",
      "ibeacon_enabled": false,
      "ibeacon_freq_msec": 0,
      "ibeacon_major": 13,
      "ibeacon_minor": 138,
      "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
      "power": 10,
      "power_mode": "custom"
    },
    "created_time": 0,
    "disable_eth1": false,
    "disable_eth2": false,
    "disable_eth3": false,
    "disable_module": false,
    "for_site": true,
    "height": 0,
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6108",
    "ip_config": {
      "dns": [
        "8.8.8.8",
        "4.4.4.4"
      ],
      "dns_suffix": [
        ".mist.local",
        ".mist.com"
      ],
      "gateway": "10.2.1.254",
      "gateway6": "2607:f8b0:4005:808::1",
      "ip": "10.2.1.1",
      "ip6": "2607:f8b0:4005:808::2004",
      "mtu": 0,
      "netmask": "255.255.255.0",
      "netmask6": "/32",
      "type": "static",
      "type6": "static",
      "vlan_id": 1
    },
    "led": {
      "brightness": 255,
      "enabled": true
    },
    "map_id": "09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1",
    "mesh": {
      "enabled": false,
      "group": 1,
      "role": "base"
    },
    "modified_time": 0,
    "name": "string",
    "notes": "string",
    "ntp_servers": [
      "string"
    ],
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "orientation": 0,
    "orientation_overwrite": true,
    "poe_passthrough": false,
    "port_config": {
      "property1": {
        "additional_vlan_ids": [
          55,
          66
        ],
        "authentication_protocol": "pap",
        "disabled": true,
        "dynamic_vlan": {
          "default_vlan_id": 999,
          "enabled": true,
          "type": "string",
          "vlans": {
            "1-10": null,
            "user": null
          }
        },
        "enable_mac_auth": false,
        "forwarding": "all",
        "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
        "mxtunnel_name": "string",
        "port_auth": "none",
        "port_vlan_id": 1,
        "radius_config": {
          "acct_interim_interval": 0,
          "acct_servers": [
            {
              "host": "1.2.3.4",
              "keywrap_enabled": true,
              "keywrap_format": "hex",
              "keywrap_kek": "1122334455",
              "keywrap_mack": "1122334455",
              "port": 1813,
              "secret": "testing123"
            }
          ],
          "auth_servers": [
            {
              "host": "1.2.3.4",
              "keywrap_enabled": true,
              "keywrap_format": "hex",
              "keywrap_kek": "1122334455",
              "keywrap_mack": "1122334455",
              "port": 1812,
              "secret": "testing123"
            }
          ],
          "auth_servers_retries": 3,
          "auth_servers_timeout": 5,
          "coa_enabled": false,
          "coa_port": 3799,
          "network": "string",
          "source_ip": "string"
        },
        "radsec": {
          "enabled": true,
          "idle_timeout": 60,
          "mxcluster_ids": [
            "572586b7-f97b-a22b-526c-8b97a3f609c4"
          ],
          "proxy_hosts": [
            "mxedge1.local"
          ],
          "server_name": "radsec.abc.com",
          "servers": [
            {
              "host": "1.1.1.1",
              "port": 1812
            }
          ],
          "use_mxedge": true,
          "use_site_mxedge": false
        },
        "vlan_id": 9,
        "vland_ids": [
          1,
          10,
          50
        ],
        "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
        "wxtunnel_remote_id": "wifiguest"
      },
      "property2": {
        "additional_vlan_ids": [
          55,
          66
        ],
        "authentication_protocol": "pap",
        "disabled": true,
        "dynamic_vlan": {
          "default_vlan_id": 999,
          "enabled": true,
          "type": "string",
          "vlans": {
            "1-10": null,
            "user": null
          }
        },
        "enable_mac_auth": false,
        "forwarding": "all",
        "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
        "mxtunnel_name": "string",
        "port_auth": "none",
        "port_vlan_id": 1,
        "radius_config": {
          "acct_interim_interval": 0,
          "acct_servers": [
            {
              "host": "1.2.3.4",
              "keywrap_enabled": true,
              "keywrap_format": "hex",
              "keywrap_kek": "1122334455",
              "keywrap_mack": "1122334455",
              "port": 1813,
              "secret": "testing123"
            }
          ],
          "auth_servers": [
            {
              "host": "1.2.3.4",
              "keywrap_enabled": true,
              "keywrap_format": "hex",
              "keywrap_kek": "1122334455",
              "keywrap_mack": "1122334455",
              "port": 1812,
              "secret": "testing123"
            }
          ],
          "auth_servers_retries": 3,
          "auth_servers_timeout": 5,
          "coa_enabled": false,
          "coa_port": 3799,
          "network": "string",
          "source_ip": "string"
        },
        "radsec": {
          "enabled": true,
          "idle_timeout": 60,
          "mxcluster_ids": [
            "572586b7-f97b-a22b-526c-8b97a3f609c4"
          ],
          "proxy_hosts": [
            "mxedge1.local"
          ],
          "server_name": "radsec.abc.com",
          "servers": [
            {
              "host": "1.1.1.1",
              "port": 1812
            }
          ],
          "use_mxedge": true,
          "use_site_mxedge": false
        },
        "vlan_id": 9,
        "vland_ids": [
          1,
          10,
          50
        ],
        "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
        "wxtunnel_remote_id": "wifiguest"
      }
    },
    "pwr_config": {
      "base": 0
    },
    "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
    "type": "ap",
    "x": 0,
    "y": 0
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
    Macs: []string{
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

body := models.DeviceprofileContainer.FromDeviceprofileGateway(models.DeviceprofileGateway{
    DhcpdConfig:           models.ToPointer(models.DhcpdConfig{
    }),
    DnsOverride:           models.ToPointer(true),
    DnsServers:            []string{
        "10.3.20.201",
        "10.3.51.222",
        "1.1.1.1",
    },
    DnsSuffix:             []string{
        "example.com",
    },
    ExtraRoutes:           map[string]models.GatewayExtraRoute{
        "10.101.0.0/16": models.GatewayExtraRoute{
            Via: models.ToPointer("10.3.100.10"),
        },
    },
    IpConfigs:             map[string]models.GatewayIpConfigProperty{
        "Corp-Core": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.100.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
        "Corp-Mgmt": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.172.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
        "Corp-lan": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.171.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
    },
    Name:                  "ITParis",
    NtpOverride:           models.ToPointer(true),
    NtpServers:            []string{
        "10.3.51.222",
    },
    PathPreferences:       map[string]models.GatewayPathPreferences{
        "core": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-Core",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "lab": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-lan",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "mgmt": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-Mgmt",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "untrust": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Name:           models.ToPointer("wan"),
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("wan")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
    },
    PortConfig:            map[string]models.GatewayPortConfig{
        "ge-0/0/0": models.GatewayPortConfig{
            Aggregated:      models.ToPointer(false),
            IpConfig:        models.ToPointer(models.GatewayPortConfigIpConfig{
                Gateway:       models.ToPointer("192.168.1.1"),
                Ip:            models.ToPointer("192.168.1.9"),
                Netmask:       models.ToPointer("/24"),
                Type:          models.ToPointer(models.GatewayWanTypeEnum("static")),
            }),
            Name:            models.ToPointer("wan"),
            Redundant:       models.ToPointer(false),
            TrafficShaping:  models.ToPointer(models.GatewayTrafficShaping{
                Enabled:          models.ToPointer(false),
            }),
            Usage:           models.GatewayPortUsageEnum("wan"),
            WanType:         models.ToPointer(models.GatewayPortWanTypeEnum("broadband")),
        },
        "ge-0/0/6-7": models.GatewayPortConfig{
            AeDisableLacp:   models.ToPointer(false),
            AeIdx:           models.NewOptional(models.ToPointer("0")),
            AeLacpForceUp:   models.ToPointer(true),
            Aggregated:      models.ToPointer(true),
            Networks:        []string{
                "Corp-lan",
                "Corp-Mgmt",
                "Corp-Core",
            },
            Usage:           models.GatewayPortUsageEnum("lan"),
        },
    },
    ServicePolicies:       []models.ServicePolicy{
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("ITParis-Internal"),
            PathPreference:  models.ToPointer("core"),
            Services:        []string{
                "internal_dns",
                "drive",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("ITParis-internet"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-core"),
            PathPreference:  models.ToPointer("core"),
            Services:        []string{
                "internal_dns",
                "internal_ntp",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-mxe-tt-in"),
            PathPreference:  models.ToPointer("mxe-in"),
            Services:        []string{
                "internal_any",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "mxedge-updates",
                "radsec",
                "icmp",
                "internet_dns",
                "internet_ntp",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mxe-data-0-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-mgt"),
            PathPreference:  models.ToPointer("mgmt"),
            Services:        []string{
                "mgmt",
            },
            Tenants:         []string{
                "domain.Corp-Core",
                "lan.Corp-Core",
                "servers.Corp-Core",
                "Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-edge-in"),
            PathPreference:  models.ToPointer("mxe-in"),
            Services:        []string{
                "internal_any",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-iot"),
            PathPreference:  models.ToPointer("iot"),
            Services:        []string{
                "iot",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "servers-hassio.Corp-Core",
                "servers-kubes.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("tanker-to-cctv"),
            PathPreference:  models.ToPointer("iot"),
            Services:        []string{
                "rtsp",
            },
            Tenants:         []string{
                "servers-tanker.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("core-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "domain.Corp-Core",
                "servers.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("deny")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("iot-upgrade-cctv"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "motioneye",
                "nodejs",
                "raspbian",
            },
            Tenants:         []string{
                "printer",
            },
        },
    },
    Type:                  "gateway",
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

