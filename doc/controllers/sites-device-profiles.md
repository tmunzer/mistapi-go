# Sites Device Profiles

```go
sitesDeviceProfiles := client.SitesDeviceProfiles()
```

## Class Name

`SitesDeviceProfiles`


# List Site Device Profiles Derived

Retrieves the list of Device Profiles available for the Site

```go
ListSiteDeviceProfilesDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[[]models.ListSiteDeviceProfilesDerivedResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | whether resolve the site variables |

## Response Type

[`[]models.ListSiteDeviceProfilesDerivedResponse`](../../doc/models/containers/list-site-device-profiles-derived-response.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resolve := false

apiResponse, err := sitesDeviceProfiles.ListSiteDeviceProfilesDerived(ctx, siteId, &resolve)
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

## Example Response *(as JSON)*

```json
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

