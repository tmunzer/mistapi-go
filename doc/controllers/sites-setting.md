# Sites Setting

```go
sitesSetting := client.SitesSetting()
```

## Class Name

`SitesSetting`

## Methods

* [Create Site Watched Stations](../../doc/controllers/sites-setting.md#create-site-watched-stations)
* [Create Site Wireless Clients Allowlist](../../doc/controllers/sites-setting.md#create-site-wireless-clients-allowlist)
* [Create Site Wireless Clients Blocklist](../../doc/controllers/sites-setting.md#create-site-wireless-clients-blocklist)
* [Delete Site Watched Stations](../../doc/controllers/sites-setting.md#delete-site-watched-stations)
* [Delete Site Wireless Clients Allowlist](../../doc/controllers/sites-setting.md#delete-site-wireless-clients-allowlist)
* [Delete Site Wireless Clients Blocklist](../../doc/controllers/sites-setting.md#delete-site-wireless-clients-blocklist)
* [Get Site Setting](../../doc/controllers/sites-setting.md#get-site-setting)
* [Get Site Setting Derived](../../doc/controllers/sites-setting.md#get-site-setting-derived)
* [Update Site Settings](../../doc/controllers/sites-setting.md#update-site-settings)


# Create Site Watched Stations

This endpoint is to provide list of client macs for annotation as watched station.

Retrieve the current clients list from `watched_station_url` under Site:Setting

```go
CreateSiteWatchedStations(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "18-65-90-de-f4-c6",
        "84-89-ad-5d-69-0d",
    },
}

apiResponse, err := sitesSetting.CreateSiteWatchedStations(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
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


# Create Site Wireless Clients Allowlist

This endpoint is to provide list of client macs for annotation as whitelist.

Retrieve the current clients list from `whitelist_url` under Site:Setting

```go
CreateSiteWirelessClientsAllowlist(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "683b679ac024",
    },
}

apiResponse, err := sitesSetting.CreateSiteWirelessClientsAllowlist(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
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


# Create Site Wireless Clients Blocklist

This endpoint is to provide list of client macs for annotation blacklist.

Retrieve the current clients list `blacklist_url` under Site:Setting

```go
CreateSiteWirelessClientsBlocklist(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "18-65-90-de-f4-c6",
        "84-89-ad-5d-69-0d",
    },
}

apiResponse, err := sitesSetting.CreateSiteWirelessClientsBlocklist(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
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


# Delete Site Watched Stations

Delete Site Watched Station Clients

```go
DeleteSiteWatchedStations(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWatchedStations(ctx, siteId)
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


# Delete Site Wireless Clients Allowlist

Delete Site Whitelist Station Clients

```go
DeleteSiteWirelessClientsAllowlist(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWirelessClientsAllowlist(ctx, siteId)
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


# Delete Site Wireless Clients Blocklist

Delete Site Blacklist Station Clients

```go
DeleteSiteWirelessClientsBlocklist(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWirelessClientsBlocklist(ctx, siteId)
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


# Get Site Setting

Get the Derived Site Settings, generated by merging the Org level templates (network templates, gateway templates) and the Site level configuration. If the same parameter is defined in both scopes, the Site level one is used.

```go
GetSiteSetting(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.SiteSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SiteSetting`](../../doc/models/site-setting.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesSetting.GetSiteSetting(ctx, siteId)
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
  "additional_config_cmds": [
    "set snmp community public"
  ],
  "analytic": {
    "enabled": false
  },
  "ap_matching": {
    "enabled": true,
    "rules": [
      {
        "eth1,eth2": {
          "port_vlan_id": 1,
          "vlan_ids": [
            1,
            10,
            50
          ]
        }
      }
    ]
  },
  "ap_port_config": {
    "model_specific": {
      "AP32": {
        "eth1,eth2": {
          "port_vlan_id": 1,
          "vlan_ids": [
            1,
            10,
            50
          ]
        }
      }
    }
  },
  "auto_placement": {
    "orientation": 45,
    "x": 30,
    "y": 60
  },
  "auto_upgrade": {
    "custom_versions": {
      "AP21": "stable",
      "AP41": "0.1.5135",
      "AP61": "0.1.7215"
    },
    "day_of_week": "sun",
    "enabled": false,
    "time_of_day": "12:00",
    "version": "beta"
  },
  "blacklist_url": "https://papi.s3.amazonaws.com/blacklist/xxx...",
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
    "power": 6,
    "power_mode": "custom"
  },
  "config_auto_revert": false,
  "created_time": 0,
  "device_updown_threshold": 0,
  "dns_servers": [
    "string"
  ],
  "dns_suffix": [
    "string"
  ],
  "engagement": {
    "dwell_tag_names": {
      "bounce": "Bounce",
      "engaged": "Engaged",
      "passerby": "Passer By",
      "stationed": "Stationed"
    },
    "dwell_tags": {
      "bounce": null,
      "engaged": "300-14400",
      "passerby": null,
      "stationed": "14400-43200"
    },
    "hours": {
      "fri": "09:00-17:00",
      "mon": "09:00-17:00",
      "sat": "09:00-12:00",
      "sun": "09:00-12:00",
      "thu": "09:00-17:00",
      "tue": "09:00-17:00",
      "wed": "09:00-17:00"
    },
    "max_dwell": 43200,
    "min_dwell": 0
  },
  "evpn_options": {
    "auto_loopback_subnet": "100.101.0.0/16",
    "auto_router_id_subnet": "100.100.0.0/24",
    "core_as_border": false,
    "overlay": {
      "as": 65000
    },
    "per_vlan_vga_v4_mac": false,
    "routed_at": "edge",
    "underlay": {
      "as_base": 65001,
      "routed_id_prefix": "/24",
      "subnet": "10.255.240.0/20"
    }
  },
  "flags": {
    "property1": "string",
    "property2": "string"
  },
  "for_site": true,
  "gateway_additional_config_cmds": [
    "set snmp community public"
  ],
  "gateway_mgmt": {
    "admin_sshkeys": [
      "string"
    ],
    "app_probing": {
      "apps": [
        "string"
      ],
      "custom_apps": [
        {
          "app_type": "string",
          "hostname": [
            "string"
          ],
          "name": "string",
          "protocol": "http"
        }
      ],
      "enabled": true
    },
    "app_usage": true,
    "auto_signature_update": {
      "day_of_week": "mon",
      "enable": true,
      "time_of_day": "string"
    },
    "config_revert_timer": 10,
    "probe_hosts": [
      "string"
    ],
    "root_password": "string",
    "security_log_source_address": "192.168.1.1",
    "security_log_source_interface": "string"
  },
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f09",
  "led": {
    "brightness": 255,
    "enabled": true
  },
  "modified_time": 0,
  "mxedge": {
    "mist_das": {
      "coa_servers": [
        {
          "disable_event_timestamp_check": false,
          "enabled": true,
          "host": "string",
          "port": 3799,
          "secret": "string"
        }
      ],
      "enabled": false
    },
    "radsec": {
      "acct_servers": [
        {
          "host": "string",
          "port": 1813,
          "secret": "string",
          "ssids": [
            "string"
          ]
        }
      ],
      "auth_servers": [
        {
          "host": "string",
          "keywrap_enabled": true,
          "keywrap_format": "hex",
          "keywrap_kek": "string",
          "keywrap_mack": "string",
          "port": 1812,
          "secret": "string",
          "ssids": [
            "string"
          ]
        }
      ],
      "enabled": true,
      "match_ssid": true,
      "proxy_hosts": [
        "string"
      ],
      "server_selection": "ordered",
      "source": "any"
    }
  },
  "mxedge_mgmt": {
    "mist_password": "MIST_PASSWORD",
    "root_password": "ROOT_PASSWORD"
  },
  "ntp_servers": [
    "pool.ntp.org"
  ],
  "occupancy": {
    "assets_enabled": false,
    "clients_enabled": true,
    "min_duration": 3000,
    "sdkclients_enabled": false,
    "unconnected_clients_enabled": false
  },
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "ospf_areas": {
    "property1": {
      "include_loopback": false,
      "networks": {
        "corp": {
          "auth_keys": {
            "1": "auth-key-1"
          },
          "auth_type": "md5",
          "bfd_minimum_interval": 500,
          "dead_interval": 40,
          "hello_interval": 10,
          "interface_type": "nbma",
          "metric": 10000
        },
        "guest": {
          "passive": true
        }
      },
      "type": "default"
    },
    "property2": {
      "include_loopback": false,
      "networks": {
        "corp": {
          "auth_keys": {
            "1": "auth-key-1"
          },
          "auth_type": "md5",
          "bfd_minimum_interval": 500,
          "dead_interval": 40,
          "hello_interval": 10,
          "interface_type": "nbma",
          "metric": 10000
        },
        "guest": {
          "passive": true
        }
      },
      "type": "default"
    }
  },
  "persist_config_on_device": false,
  "port_mirroring": {
    "property1": {
      "input_networks_ingress": [
        "corp"
      ],
      "input_port_ids_egress": [
        "ge-0/0/3"
      ],
      "input_port_ids_ingress": [
        "ge-0/0/3"
      ],
      "output_network": "analyze",
      "output_port_id": "ge-0/0/5"
    },
    "property2": {
      "input_networks_ingress": [
        "corp"
      ],
      "input_port_ids_egress": [
        "ge-0/0/3"
      ],
      "input_port_ids_ingress": [
        "ge-0/0/3"
      ],
      "output_network": "analyze",
      "output_port_id": "ge-0/0/5"
    }
  },
  "port_usages": {
    "dynamic": {
      "mode": "dynamic",
      "reset_default_when": "link_down",
      "rules": [
        {
          "equals": "string",
          "equals_any": [
            "string"
          ],
          "expression": "string",
          "src": "lldp_chassis_id",
          "usage": "string"
        }
      ]
    },
    "property1": {
      "all_networks": false,
      "allow_dhcpd": true,
      "authentication_protocol": "pap",
      "bypass_auth_when_server_down": true,
      "description": "string",
      "disable_autoneg": false,
      "disabled": false,
      "duplex": "auto",
      "enable_mac_auth": true,
      "enable_qos": true,
      "guest_network": "string",
      "mac_auth_only": true,
      "mac_limit": 0,
      "mode": "access",
      "mtu": 0,
      "networks": [
        "string"
      ],
      "persist_mac": false,
      "poe_disabled": false,
      "port_auth": "dot1x",
      "port_network": "string",
      "rejected_network": null,
      "speed": "auto",
      "storm_control": {
        "no_broadcast": false,
        "no_multicast": false,
        "no_registered_multicast": false,
        "no_unknown_unicast": false,
        "percentage": 80
      },
      "stp_edge": true,
      "voip_network": "string"
    },
    "property2": {
      "all_networks": false,
      "allow_dhcpd": true,
      "authentication_protocol": "pap",
      "bypass_auth_when_server_down": true,
      "description": "string",
      "disable_autoneg": false,
      "disabled": false,
      "duplex": "auto",
      "enable_mac_auth": true,
      "enable_qos": true,
      "guest_network": "string",
      "mac_auth_only": true,
      "mac_limit": 0,
      "mode": "access",
      "mtu": 0,
      "networks": [
        "string"
      ],
      "persist_mac": false,
      "poe_disabled": false,
      "port_network": "string",
      "rejected_network": null,
      "speed": "auto",
      "storm_control": {
        "no_broadcast": false,
        "no_multicast": false,
        "no_registered_multicast": false,
        "no_unknown_unicast": false,
        "percentage": 80
      },
      "stp_edge": true,
      "voip_network": "string"
    }
  },
  "proxy": {
    "url": "http://proxy.internal:8080/"
  },
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
  "remote_syslog": {
    "archive": {
      "files": 20,
      "size": "5m"
    },
    "console": {
      "contents": [
        {
          "facility": "config",
          "severity": "warning"
        }
      ]
    },
    "enabled": false,
    "files": [
      {
        "archive": {
          "files": 10,
          "size": "5m"
        },
        "contents": [
          {
            "facility": "config",
            "severity": "warning"
          }
        ],
        "explicit_priority": true,
        "file": "file-name",
        "match": "!alarm|ntp|errors.crc_error[chan]",
        "structured_data": true
      }
    ],
    "network": "default",
    "send_to_all_servers": false,
    "servers": [
      {
        "facility": "config",
        "host": "syslogd.internal",
        "port": 514,
        "protocol": "udp",
        "severity": "info",
        "tag": ""
      }
    ],
    "time_format": "millisecond",
    "users": [
      {
        "contents": [
          {
            "facility": "config",
            "severity": "warning"
          }
        ],
        "match": "\"!alarm|ntp|errors.crc_error[chan]\"",
        "user": "*"
      }
    ]
  },
  "report_gatt": false,
  "rogue": {
    "enabled": false,
    "honeypot_enabled": false,
    "min_duration": 10,
    "min_rssi": -80,
    "whitelisted_bssids": [
      "NeighborSSID"
    ],
    "whitelisted_ssids": [
      "cc:8e:6f:d4:bf:16",
      "cc-8e-6f-d4-bf-16",
      "cc-73-*",
      "cc:82:*"
    ]
  },
  "rtsa": {
    "app_waking": false,
    "disable_dead_reckoning": true,
    "disable_pressure_sensor": false,
    "enabled": true,
    "track_asset": false
  },
  "simple_alert": {
    "arp_failure": {
      "client_count": 10,
      "duration": 20,
      "incident_count": 10
    },
    "dhcp_failure": {
      "client_count": 10,
      "duration": 10,
      "incident_count": 20
    },
    "dns_failure": {
      "client_count": 20,
      "duration": 10,
      "incident_count": 30
    }
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "skyatp": {
    "enabled": true,
    "send_ip_mac_mapping": true
  },
  "srx_app": {
    "enabled": false
  },
  "ssh_keys": [
    "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host"
  ],
  "ssr": {
    "conductor_hosts": [
      "\"1.1.1.1\", \"2.2.2.2\""
    ],
    "disable_stats": true
  },
  "status_portal": {
    "enabled": false,
    "hostnames": [
      "my.misty.com"
    ]
  },
  "switch_mgmt": {
    "ap_affinity_threshold": 10,
    "config_revert_timer": 10,
    "dhcp_option_fqdn": false,
    "mxedge_proxy_host": "string",
    "mxedge_proxy_port": 2222,
    "root_password": "string",
    "tacacs": {
      "acct_servers": [
        {
          "host": "198.51.100.1",
          "port": "49",
          "secret": "string",
          "timeout": 10
        }
      ],
      "enabled": true,
      "network": "string",
      "tacplus_servers": [
        {
          "host": "198.51.100.1",
          "port": "49",
          "secret": "string",
          "timeout": 10
        }
      ]
    },
    "use_mxedge_proxy": true
  },
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "vna": {
    "enabled": false
  },
  "vrf_instances": {
    "guest": {
      "extra_routes": {
        "0.0.0.0/0": {
          "via": "192.168.31.1"
        }
      },
      "networks": [
        "guest"
      ]
    }
  },
  "vrrp_groups": {
    "property1": {
      "auth_key": "auth-key-1",
      "auth_password": "string",
      "auth_type": "md5",
      "networks": {
        "data": {
          "ip": "10.182.96.1"
        },
        "mgmt": {
          "ip": "10.182.104.1"
        },
        "v10": {
          "ip": "10.182.104.129"
        },
        "wap": {
          "ip": "10.182.102.1"
        }
      }
    },
    "property2": {
      "auth_key": "auth-key-1",
      "auth_password": "string",
      "auth_type": "md5",
      "networks": {
        "data": {
          "ip": "10.182.96.1"
        },
        "mgmt": {
          "ip": "10.182.104.1"
        },
        "v10": {
          "ip": "10.182.104.129"
        },
        "wap": {
          "ip": "10.182.102.1"
        }
      }
    }
  },
  "wan_vna": {
    "enabled": false
  },
  "watched_station_url": "https://papi.s3.amazonaws.com/watched_station/xxx...",
  "whitelist_url": "https://papi.s3.amazonaws.com/whitelist/xxx...",
  "wids": {
    "repeated_auth_failures": {
      "duration": 60,
      "threshold": 0
    }
  },
  "wifi": {
    "cisco_enabled": true,
    "disable_11k": false,
    "disable_radios_when_power_constrained": false,
    "enable_arp_spoof_check": false,
    "enable_shared_radio_scanning": true,
    "enabled": true,
    "locate_connected": true,
    "locate_unconnected": false,
    "mesh_allow_dfs": false,
    "mesh_enable_crm": false,
    "mesh_enabled": false,
    "mesh_psk": "string",
    "mesh_ssid": "string",
    "proxy_arp": "default"
  },
  "wired_vna": {
    "enabled": false
  },
  "zone_occupancy_alert": {
    "email_notifiers": [
      "foo@juniper.net",
      "bar@juniper.net"
    ],
    "enabled": false,
    "threshold": 5
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


# Get Site Setting Derived

Get Site Settings

```go
GetSiteSettingDerived(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.SiteSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SiteSetting`](../../doc/models/site-setting.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesSetting.GetSiteSettingDerived(ctx, siteId)
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
  "additional_config_cmds": [
    "set snmp community public"
  ],
  "analytic": {
    "enabled": false
  },
  "ap_matching": {
    "enabled": true,
    "rules": [
      {
        "eth1,eth2": {
          "port_vlan_id": 1,
          "vlan_ids": [
            1,
            10,
            50
          ]
        }
      }
    ]
  },
  "ap_port_config": {
    "model_specific": {
      "AP32": {
        "eth1,eth2": {
          "port_vlan_id": 1,
          "vlan_ids": [
            1,
            10,
            50
          ]
        }
      }
    }
  },
  "auto_placement": {
    "orientation": 45,
    "x": 30,
    "y": 60
  },
  "auto_upgrade": {
    "custom_versions": {
      "AP21": "stable",
      "AP41": "0.1.5135",
      "AP61": "0.1.7215"
    },
    "day_of_week": "sun",
    "enabled": false,
    "time_of_day": "12:00",
    "version": "beta"
  },
  "blacklist_url": "https://papi.s3.amazonaws.com/blacklist/xxx...",
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
    "power": 6,
    "power_mode": "custom"
  },
  "config_auto_revert": false,
  "created_time": 0,
  "device_updown_threshold": 0,
  "dns_servers": [
    "string"
  ],
  "dns_suffix": [
    "string"
  ],
  "engagement": {
    "dwell_tag_names": {
      "bounce": "Bounce",
      "engaged": "Engaged",
      "passerby": "Passer By",
      "stationed": "Stationed"
    },
    "dwell_tags": {
      "bounce": null,
      "engaged": "300-14400",
      "passerby": null,
      "stationed": "14400-43200"
    },
    "hours": {
      "fri": "09:00-17:00",
      "mon": "09:00-17:00",
      "sat": "09:00-12:00",
      "sun": "09:00-12:00",
      "thu": "09:00-17:00",
      "tue": "09:00-17:00",
      "wed": "09:00-17:00"
    },
    "max_dwell": 43200,
    "min_dwell": 0
  },
  "evpn_options": {
    "auto_loopback_subnet": "100.101.0.0/16",
    "auto_router_id_subnet": "100.100.0.0/24",
    "core_as_border": false,
    "overlay": {
      "as": 65000
    },
    "per_vlan_vga_v4_mac": false,
    "routed_at": "edge",
    "underlay": {
      "as_base": 65001,
      "routed_id_prefix": "/24",
      "subnet": "10.255.240.0/20"
    }
  },
  "flags": {
    "property1": "string",
    "property2": "string"
  },
  "for_site": true,
  "gateway_additional_config_cmds": [
    "set snmp community public"
  ],
  "gateway_mgmt": {
    "admin_sshkeys": [
      "string"
    ],
    "app_probing": {
      "apps": [
        "string"
      ],
      "custom_apps": [
        {
          "app_type": "string",
          "hostname": [
            "string"
          ],
          "name": "string",
          "protocol": "http"
        }
      ],
      "enabled": true
    },
    "app_usage": true,
    "auto_signature_update": {
      "day_of_week": "mon",
      "enable": true,
      "time_of_day": "string"
    },
    "config_revert_timer": 10,
    "probe_hosts": [
      "string"
    ],
    "root_password": "string",
    "security_log_source_address": "192.168.1.1",
    "security_log_source_interface": "string"
  },
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f09",
  "led": {
    "brightness": 255,
    "enabled": true
  },
  "modified_time": 0,
  "mxedge": {
    "mist_das": {
      "coa_servers": [
        {
          "disable_event_timestamp_check": false,
          "enabled": true,
          "host": "string",
          "port": 3799,
          "secret": "string"
        }
      ],
      "enabled": false
    },
    "radsec": {
      "acct_servers": [
        {
          "host": "string",
          "port": 1813,
          "secret": "string",
          "ssids": [
            "string"
          ]
        }
      ],
      "auth_servers": [
        {
          "host": "string",
          "keywrap_enabled": true,
          "keywrap_format": "hex",
          "keywrap_kek": "string",
          "keywrap_mack": "string",
          "port": 1812,
          "secret": "string",
          "ssids": [
            "string"
          ]
        }
      ],
      "enabled": true,
      "match_ssid": true,
      "proxy_hosts": [
        "string"
      ],
      "server_selection": "ordered",
      "source": "any"
    }
  },
  "mxedge_mgmt": {
    "mist_password": "MIST_PASSWORD",
    "root_password": "ROOT_PASSWORD"
  },
  "ntp_servers": [
    "pool.ntp.org"
  ],
  "occupancy": {
    "assets_enabled": false,
    "clients_enabled": true,
    "min_duration": 3000,
    "sdkclients_enabled": false,
    "unconnected_clients_enabled": false
  },
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "ospf_areas": {
    "property1": {
      "include_loopback": false,
      "networks": {
        "corp": {
          "auth_keys": {
            "1": "auth-key-1"
          },
          "auth_type": "md5",
          "bfd_minimum_interval": 500,
          "dead_interval": 40,
          "hello_interval": 10,
          "interface_type": "nbma",
          "metric": 10000
        },
        "guest": {
          "passive": true
        }
      },
      "type": "default"
    },
    "property2": {
      "include_loopback": false,
      "networks": {
        "corp": {
          "auth_keys": {
            "1": "auth-key-1"
          },
          "auth_type": "md5",
          "bfd_minimum_interval": 500,
          "dead_interval": 40,
          "hello_interval": 10,
          "interface_type": "nbma",
          "metric": 10000
        },
        "guest": {
          "passive": true
        }
      },
      "type": "default"
    }
  },
  "persist_config_on_device": false,
  "port_mirroring": {
    "property1": {
      "input_networks_ingress": [
        "corp"
      ],
      "input_port_ids_egress": [
        "ge-0/0/3"
      ],
      "input_port_ids_ingress": [
        "ge-0/0/3"
      ],
      "output_network": "analyze",
      "output_port_id": "ge-0/0/5"
    },
    "property2": {
      "input_networks_ingress": [
        "corp"
      ],
      "input_port_ids_egress": [
        "ge-0/0/3"
      ],
      "input_port_ids_ingress": [
        "ge-0/0/3"
      ],
      "output_network": "analyze",
      "output_port_id": "ge-0/0/5"
    }
  },
  "port_usages": {
    "dynamic": {
      "mode": "dynamic",
      "reset_default_when": "link_down",
      "rules": [
        {
          "equals": "string",
          "equals_any": [
            "string"
          ],
          "expression": "string",
          "src": "lldp_chassis_id",
          "usage": "string"
        }
      ]
    },
    "property1": {
      "all_networks": false,
      "allow_dhcpd": true,
      "authentication_protocol": "pap",
      "bypass_auth_when_server_down": true,
      "description": "string",
      "disable_autoneg": false,
      "disabled": false,
      "duplex": "auto",
      "enable_mac_auth": true,
      "enable_qos": true,
      "guest_network": "string",
      "mac_auth_only": true,
      "mac_limit": 0,
      "mode": "access",
      "mtu": 0,
      "networks": [
        "string"
      ],
      "persist_mac": false,
      "poe_disabled": false,
      "port_auth": "dot1x",
      "port_network": "string",
      "rejected_network": null,
      "speed": "auto",
      "storm_control": {
        "no_broadcast": false,
        "no_multicast": false,
        "no_registered_multicast": false,
        "no_unknown_unicast": false,
        "percentage": 80
      },
      "stp_edge": true,
      "voip_network": "string"
    },
    "property2": {
      "all_networks": false,
      "allow_dhcpd": true,
      "authentication_protocol": "pap",
      "bypass_auth_when_server_down": true,
      "description": "string",
      "disable_autoneg": false,
      "disabled": false,
      "duplex": "auto",
      "enable_mac_auth": true,
      "enable_qos": true,
      "guest_network": "string",
      "mac_auth_only": true,
      "mac_limit": 0,
      "mode": "access",
      "mtu": 0,
      "networks": [
        "string"
      ],
      "persist_mac": false,
      "poe_disabled": false,
      "port_network": "string",
      "rejected_network": null,
      "speed": "auto",
      "storm_control": {
        "no_broadcast": false,
        "no_multicast": false,
        "no_registered_multicast": false,
        "no_unknown_unicast": false,
        "percentage": 80
      },
      "stp_edge": true,
      "voip_network": "string"
    }
  },
  "proxy": {
    "url": "http://proxy.internal:8080/"
  },
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
  "remote_syslog": {
    "archive": {
      "files": 20,
      "size": "5m"
    },
    "console": {
      "contents": [
        {
          "facility": "config",
          "severity": "warning"
        }
      ]
    },
    "enabled": false,
    "files": [
      {
        "archive": {
          "files": 10,
          "size": "5m"
        },
        "contents": [
          {
            "facility": "config",
            "severity": "warning"
          }
        ],
        "explicit_priority": true,
        "file": "file-name",
        "match": "!alarm|ntp|errors.crc_error[chan]",
        "structured_data": true
      }
    ],
    "network": "default",
    "send_to_all_servers": false,
    "servers": [
      {
        "facility": "config",
        "host": "syslogd.internal",
        "port": 514,
        "protocol": "udp",
        "severity": "info",
        "tag": ""
      }
    ],
    "time_format": "millisecond",
    "users": [
      {
        "contents": [
          {
            "facility": "config",
            "severity": "warning"
          }
        ],
        "match": "\"!alarm|ntp|errors.crc_error[chan]\"",
        "user": "*"
      }
    ]
  },
  "report_gatt": false,
  "rogue": {
    "enabled": false,
    "honeypot_enabled": false,
    "min_duration": 10,
    "min_rssi": -80,
    "whitelisted_bssids": [
      "NeighborSSID"
    ],
    "whitelisted_ssids": [
      "cc:8e:6f:d4:bf:16",
      "cc-8e-6f-d4-bf-16",
      "cc-73-*",
      "cc:82:*"
    ]
  },
  "rtsa": {
    "app_waking": false,
    "disable_dead_reckoning": true,
    "disable_pressure_sensor": false,
    "enabled": true,
    "track_asset": false
  },
  "simple_alert": {
    "arp_failure": {
      "client_count": 10,
      "duration": 20,
      "incident_count": 10
    },
    "dhcp_failure": {
      "client_count": 10,
      "duration": 10,
      "incident_count": 20
    },
    "dns_failure": {
      "client_count": 20,
      "duration": 10,
      "incident_count": 30
    }
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "skyatp": {
    "enabled": true,
    "send_ip_mac_mapping": true
  },
  "srx_app": {
    "enabled": false
  },
  "ssh_keys": [
    "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host"
  ],
  "ssr": {
    "conductor_hosts": [
      "\"1.1.1.1\", \"2.2.2.2\""
    ],
    "disable_stats": true
  },
  "status_portal": {
    "enabled": false,
    "hostnames": [
      "my.misty.com"
    ]
  },
  "switch_mgmt": {
    "ap_affinity_threshold": 10,
    "config_revert_timer": 10,
    "dhcp_option_fqdn": false,
    "mxedge_proxy_host": "string",
    "mxedge_proxy_port": 2222,
    "root_password": "string",
    "tacacs": {
      "acct_servers": [
        {
          "host": "198.51.100.1",
          "port": "49",
          "secret": "string",
          "timeout": 10
        }
      ],
      "enabled": true,
      "network": "string",
      "tacplus_servers": [
        {
          "host": "198.51.100.1",
          "port": "49",
          "secret": "string",
          "timeout": 10
        }
      ]
    },
    "use_mxedge_proxy": true
  },
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "vna": {
    "enabled": false
  },
  "vrf_instances": {
    "guest": {
      "extra_routes": {
        "0.0.0.0/0": {
          "via": "192.168.31.1"
        }
      },
      "networks": [
        "guest"
      ]
    }
  },
  "vrrp_groups": {
    "property1": {
      "auth_key": "auth-key-1",
      "auth_password": "string",
      "auth_type": "md5",
      "networks": {
        "data": {
          "ip": "10.182.96.1"
        },
        "mgmt": {
          "ip": "10.182.104.1"
        },
        "v10": {
          "ip": "10.182.104.129"
        },
        "wap": {
          "ip": "10.182.102.1"
        }
      }
    },
    "property2": {
      "auth_key": "auth-key-1",
      "auth_password": "string",
      "auth_type": "md5",
      "networks": {
        "data": {
          "ip": "10.182.96.1"
        },
        "mgmt": {
          "ip": "10.182.104.1"
        },
        "v10": {
          "ip": "10.182.104.129"
        },
        "wap": {
          "ip": "10.182.102.1"
        }
      }
    }
  },
  "wan_vna": {
    "enabled": false
  },
  "watched_station_url": "https://papi.s3.amazonaws.com/watched_station/xxx...",
  "whitelist_url": "https://papi.s3.amazonaws.com/whitelist/xxx...",
  "wids": {
    "repeated_auth_failures": {
      "duration": 60,
      "threshold": 0
    }
  },
  "wifi": {
    "cisco_enabled": true,
    "disable_11k": false,
    "disable_radios_when_power_constrained": false,
    "enable_arp_spoof_check": false,
    "enable_shared_radio_scanning": true,
    "enabled": true,
    "locate_connected": true,
    "locate_unconnected": false,
    "mesh_allow_dfs": false,
    "mesh_enable_crm": false,
    "mesh_enabled": false,
    "mesh_psk": "string",
    "mesh_ssid": "string",
    "proxy_arp": "default"
  },
  "wired_vna": {
    "enabled": false
  },
  "zone_occupancy_alert": {
    "email_notifiers": [
      "foo@juniper.net",
      "bar@juniper.net"
    ],
    "enabled": false,
    "threshold": 5
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


# Update Site Settings

Update Site Settings

```go
UpdateSiteSettings(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.SiteSetting) (
    models.ApiResponse[models.SiteSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SiteSetting`](../../doc/models/site-setting.md) | Body, Optional | Request Body |

## Response Type

[`models.SiteSetting`](../../doc/models/site-setting.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SiteSetting{
    AdditionalConfigCmds:            []string{
        "set snmp community public",
    },
    Analytic:                        models.ToPointer(models.SiteSettingAnalytic{
        Enabled:              models.ToPointer(false),
    }),
    ApMatching:                      models.ToPointer(models.SiteSettingApMatching{
        Enabled:              models.ToPointer(true),
        Rules:                []models.SiteSettingApMatchingRule{
            models.SiteSettingApMatchingRule{
                AdditionalProperties: map[string]interface{}{
                    "eth1,eth2": interface{}("[port_vlan_id, 1][vlan_ids, System.Object[]]"),
                },
            },
        },
    }),
    ApPortConfig:                    models.ToPointer(models.SiteSettingApPortConfig{
        ModelSpecific:        map[string]models.ApPortConfig{
            "AP32": models.ApPortConfig{
                AdditionalProperties: map[string]interface{}{
                    "eth1,eth2": interface{}("[port_vlan_id, 1][vlan_ids, System.Object[]]"),
                },
            },
        },
    }),
    AutoUpgrade:                     models.ToPointer(models.SiteSettingAutoUpgrade{
        CustomVersions:       map[string]string{
            "AP21": "stable",
            "AP41": "0.1.5135",
            "AP61": "0.1.7215",
        },
        DayOfWeek:            models.ToPointer(models.DayOfWeekEnum("sun")),
        Enabled:              models.ToPointer(false),
        TimeOfDay:            models.ToPointer("12:00"),
        Version:              models.ToPointer(models.SiteAutoUpgradeVersionEnum("beta")),
    }),
    ConfigAutoRevert:                models.ToPointer(false),
    DeviceUpdownThreshold:           models.NewOptional(models.ToPointer(0)),
    DnsServers:                      []string{
        "string",
    },
    DnsSuffix:                       []string{
        "string",
    },
    Engagement:                      models.ToPointer(models.SiteEngagement{
        DwellTagNames:        models.ToPointer(models.SiteEngagementDwellTagNames{
            Bounce:               models.ToPointer("Bounce"),
            Engaged:              models.ToPointer("Engaged"),
            Passerby:             models.ToPointer("Passer By"),
            Stationed:            models.ToPointer("Stationed"),
        }),
        DwellTags:            models.ToPointer(models.SiteEngagementDwellTags{
            Engaged:              models.NewOptional(models.ToPointer("300-14400")),
            Stationed:            models.NewOptional(models.ToPointer("14400-43200")),
        }),
        Hours:                models.ToPointer(models.Hours{
            Fri:                  models.ToPointer("09:00-17:00"),
            Mon:                  models.ToPointer("09:00-17:00"),
            Sat:                  models.ToPointer("09:00-12:00"),
            Sun:                  models.ToPointer("09:00-12:00"),
            Thu:                  models.ToPointer("09:00-17:00"),
            Tue:                  models.ToPointer("09:00-17:00"),
            Wed:                  models.ToPointer("09:00-17:00"),
        }),
        MaxDwell:             models.ToPointer(43200),
        MinDwell:             models.ToPointer(0),
    }),
    EvpnOptions:                     models.ToPointer(models.EvpnOptions{
        AutoLoopbackSubnet:   models.ToPointer("100.101.0.0/16"),
        AutoRouterIdSubnet:   models.ToPointer("100.100.0.0/24"),
        CoreAsBorder:         models.ToPointer(false),
        Overlay:              models.ToPointer(models.EvpnOptionsOverlay{
            As:                   models.ToPointer(65000),
        }),
        PerVlanVgaV4Mac:      models.ToPointer(false),
        RoutedAt:             models.ToPointer(models.EvpnOptionsRoutedAtEnum("edge")),
        Underlay:             models.ToPointer(models.EvpnOptionsUnderlay{
            AsBase:               models.ToPointer(65001),
            RoutedIdPrefix:       models.ToPointer("/24"),
            Subnet:               models.ToPointer("10.255.240.0/20"),
        }),
    }),
    GatewayAdditionalConfigCmds:     []string{
        "set snmp community public",
    },
    GatewayMgmt:                     models.ToPointer(models.SiteSettingGatewayMgmt{
        AdminSshkeys:               []string{
            "string",
        },
        AppProbing:                 models.ToPointer(models.AppProbing{
            Apps:                 []string{
                "string",
            },
            CustomApps:           []models.AppProbingCustomApp{
                models.AppProbingCustomApp{
                    AppType:              models.ToPointer("string"),
                    Name:                 models.ToPointer("string"),
                    Protocol:             models.ToPointer(models.AppProbingCustomAppProtocolEnum("http")),
                    AdditionalProperties: map[string]interface{}{
                        "hostname": interface{}("string"),
                    },
                },
            },
            Enabled:              models.ToPointer(true),
        }),
        AppUsage:                   models.ToPointer(true),
        AutoSignatureUpdate:        models.ToPointer(models.SiteSettingGatewayMgmtAutoSignatureUpdate{
            DayOfWeek:            models.ToPointer(models.DayOfWeekEnum("any")),
            Enable:               models.ToPointer(true),
            TimeOfDay:            models.ToPointer("string"),
        }),
        ConfigRevertTimer:          models.ToPointer(10),
        ProbeHosts:                 []string{
            "string",
        },
        RootPassword:               models.ToPointer("string"),
        SecurityLogSourceAddress:   models.ToPointer("192.168.1.1"),
        SecurityLogSourceInterface: models.ToPointer("string"),
    }),
    Led:                             models.ToPointer(models.ApLed{
        Brightness:           models.ToPointer(255),
        Enabled:              models.ToPointer(true),
    }),
    MxedgeMgmt:                      models.ToPointer(models.MxedgeMgmt{
        MistPassword:         models.ToPointer("MIST_PASSWORD"),
        RootPassword:         models.ToPointer("ROOT_PASSWORD"),
    }),
    Networks:                        map[string]models.SwitchNetwork{
        "property1": models.SwitchNetwork{
            Gateway:              models.ToPointer("string"),
            Subnet:               models.ToPointer("string"),
            VlanId:               models.VlanIdWithVariableContainer.FromNumber(10),
            AdditionalProperties: map[string]interface{}{
                "dns": interface{}("string"),
                "dns_suffix": interface{}("string"),
                "ospf_interface_type": interface{}("string"),
                "zone": interface{}("string"),
            },
        },
        "property2": models.SwitchNetwork{
            Gateway:              models.ToPointer("string"),
            Subnet:               models.ToPointer("string"),
            VlanId:               models.VlanIdWithVariableContainer.FromNumber(10),
            AdditionalProperties: map[string]interface{}{
                "dns": interface{}("string"),
                "dns_suffix": interface{}("string"),
                "ospf_interface_type": interface{}("string"),
                "zone": interface{}("string"),
            },
        },
    },
    NtpServers:                      []string{
        "string",
    },
    Occupancy:                       models.ToPointer(models.SiteOccupancyAnalytics{
        AssetsEnabled:             models.ToPointer(false),
        ClientsEnabled:            models.ToPointer(true),
        MinDuration:               models.ToPointer(3000),
        SdkclientsEnabled:         models.ToPointer(false),
        UnconnectedClientsEnabled: models.ToPointer(false),
    }),
    OspfAreas:                       map[string]models.OspfArea{
        "property1": models.OspfArea{
            IncludeLoopback:      models.ToPointer(false),
            Networks:             map[string]models.OspfAreasNetwork{
                "corp": models.OspfAreasNetwork{
                    AuthKeys:               map[string]string{
                        "1": "auth-key-1",
                    },
                    AuthType:               models.ToPointer(models.OspfAreaNetworkAuthTypeEnum("md5")),
                    BfdMinimumInterval:     models.ToPointer(500),
                    DeadInterval:           models.ToPointer(40),
                    HelloInterval:          models.ToPointer(10),
                    InterfaceType:          models.ToPointer(models.OspfAreaNetworkInterfaceTypeEnum("nbma")),
                    Metric:                 models.NewOptional(models.ToPointer(10000)),
                },
                "guest": models.OspfAreasNetwork{
                    Passive:                models.ToPointer(true),
                },
            },
            Type:                 models.ToPointer(models.OspfAreaTypeEnum("default")),
        },
        "property2": models.OspfArea{
            IncludeLoopback:      models.ToPointer(false),
            Networks:             map[string]models.OspfAreasNetwork{
                "corp": models.OspfAreasNetwork{
                    AuthKeys:               map[string]string{
                        "1": "auth-key-1",
                    },
                    AuthType:               models.ToPointer(models.OspfAreaNetworkAuthTypeEnum("md5")),
                    BfdMinimumInterval:     models.ToPointer(500),
                    DeadInterval:           models.ToPointer(40),
                    HelloInterval:          models.ToPointer(10),
                    InterfaceType:          models.ToPointer(models.OspfAreaNetworkInterfaceTypeEnum("nbma")),
                    Metric:                 models.NewOptional(models.ToPointer(10000)),
                },
                "guest": models.OspfAreasNetwork{
                    Passive:                models.ToPointer(true),
                },
            },
            Type:                 models.ToPointer(models.OspfAreaTypeEnum("default")),
        },
    },
    PersistConfigOnDevice:           models.ToPointer(false),
    PortMirroring:                   map[string]models.SwitchPortMirroringProperty{
        "property1": models.SwitchPortMirroringProperty{
            InputNetworksIngress: []string{
                "corp",
            },
            InputPortIdsEgress:   []string{
                "ge-0/0/3",
            },
            InputPortIdsIngress:  []string{
                "ge-0/0/3",
            },
            OutputNetwork:        models.ToPointer("analyze"),
            OutputPortId:         models.ToPointer("ge-0/0/5"),
        },
        "property2": models.SwitchPortMirroringProperty{
            InputNetworksIngress: []string{
                "corp",
            },
            InputPortIdsEgress:   []string{
                "ge-0/0/3",
            },
            InputPortIdsIngress:  []string{
                "ge-0/0/3",
            },
            OutputNetwork:        models.ToPointer("analyze"),
            OutputPortId:         models.ToPointer("ge-0/0/5"),
        },
    },
    PortUsages:                      map[string]models.SwitchPortUsage{
        "dynamic": models.SwitchPortUsage{
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum("dynamic")),
            ResetDefaultWhen:                         models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum("link_down")),
            Rules:                                    []models.SwitchPortUsageDynamicRule{
                models.SwitchPortUsageDynamicRule{
                    Equals:               models.ToPointer("string"),
                    EqualsAny:            []string{
                        "string",
                    },
                    Expression:           models.ToPointer("string"),
                    Src:                  models.SwitchPortUsageDynamicRuleSrcEnum("lldp_chassis_id"),
                    Usage:                models.ToPointer("string"),
                },
            },
        },
        "property1": models.SwitchPortUsage{
            AllNetworks:                              models.ToPointer(false),
            AllowDhcpd:                               models.ToPointer(true),
            BypassAuthWhenServerDown:                 models.ToPointer(true),
            Description:                              models.ToPointer("string"),
            DisableAutoneg:                           models.ToPointer(false),
            Disabled:                                 models.ToPointer(false),
            Duplex:                                   models.ToPointer(models.SwitchPortUsageDuplexEnum("auto")),
            EnableMacAuth:                            models.ToPointer(true),
            EnableQos:                                models.ToPointer(true),
            GuestNetwork:                             models.NewOptional(models.ToPointer("string")),
            MacAuthOnly:                              models.ToPointer(true),
            MacLimit:                                 models.ToPointer(0),
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum("access")),
            Mtu:                                      models.ToPointer(0),
            Networks:                                 []string{
                "string",
            },
            PersistMac:                               models.ToPointer(false),
            PoeDisabled:                              models.ToPointer(false),
            PortAuth:                                 models.NewOptional(models.ToPointer(models.SwitchPortUsageDot1xEnum("dot1x"))),
            PortNetwork:                              models.ToPointer("string"),
            Speed:                                    models.ToPointer(models.SwitchPortUsageSpeedEnum("auto")),
            StormControl:                             models.ToPointer(models.SwitchPortUsageStormControl{
                NoBroadcast:           models.ToPointer(false),
                NoMulticast:           models.ToPointer(false),
                NoRegisteredMulticast: models.ToPointer(false),
                NoUnknownUnicast:      models.ToPointer(false),
                Percentage:            models.ToPointer(80),
            }),
            StpEdge:                                  models.ToPointer(true),
            VoipNetwork:                              models.ToPointer("string"),
        },
        "property2": models.SwitchPortUsage{
            AllNetworks:                              models.ToPointer(false),
            AllowDhcpd:                               models.ToPointer(true),
            BypassAuthWhenServerDown:                 models.ToPointer(true),
            Description:                              models.ToPointer("string"),
            DisableAutoneg:                           models.ToPointer(false),
            Disabled:                                 models.ToPointer(false),
            Duplex:                                   models.ToPointer(models.SwitchPortUsageDuplexEnum("auto")),
            EnableMacAuth:                            models.ToPointer(true),
            EnableQos:                                models.ToPointer(true),
            GuestNetwork:                             models.NewOptional(models.ToPointer("string")),
            MacAuthOnly:                              models.ToPointer(true),
            MacLimit:                                 models.ToPointer(0),
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum("access")),
            Mtu:                                      models.ToPointer(0),
            Networks:                                 []string{
                "string",
            },
            PersistMac:                               models.ToPointer(false),
            PoeDisabled:                              models.ToPointer(false),
            PortNetwork:                              models.ToPointer("string"),
            Speed:                                    models.ToPointer(models.SwitchPortUsageSpeedEnum("auto")),
            StormControl:                             models.ToPointer(models.SwitchPortUsageStormControl{
                NoBroadcast:           models.ToPointer(false),
                NoMulticast:           models.ToPointer(false),
                NoRegisteredMulticast: models.ToPointer(false),
                NoUnknownUnicast:      models.ToPointer(false),
                Percentage:            models.ToPointer(80),
            }),
            StpEdge:                                  models.ToPointer(true),
            VoipNetwork:                              models.ToPointer("string"),
        },
    },
    Proxy:                           models.ToPointer(models.Proxy{
        Url:                  models.ToPointer("http://proxy.internal:8080/*"),
    }),
    Rogue:                           models.ToPointer(models.SiteRogue{
        Enabled:              models.ToPointer(false),
        HoneypotEnabled:      models.ToPointer(false),
        MinDuration:          models.ToPointer(10),
        MinRssi:              models.ToPointer(-80),
        WhitelistedBssids:    []string{
            "NeighborSSID",
        },
        WhitelistedSsids:     []string{
            "cc:8e:6f:d4:bf:16",
            "cc-8e-6f-d4-bf-16",
            "cc-73-*",
            "cc:82:*",
        },
    }),
    SimpleAlert:                     models.ToPointer(models.SimpleAlert{
        ArpFailure:           models.ToPointer(models.SimpleAlertArpFailure{
            ClientCount:          models.ToPointer(10),
            Duration:             models.ToPointer(20),
            IncidentCount:        models.ToPointer(10),
        }),
        DhcpFailure:          models.ToPointer(models.SimpleAlertDhcpFailure{
            ClientCount:          models.ToPointer(10),
            Duration:             models.ToPointer(10),
            IncidentCount:        models.ToPointer(20),
        }),
        DnsFailure:           models.ToPointer(models.SimpleAlertDnsFailure{
            ClientCount:          models.ToPointer(20),
            Duration:             models.ToPointer(10),
            IncidentCount:        models.ToPointer(30),
        }),
    }),
    Skyatp:                          models.ToPointer(models.SiteSettingSkyatp{
        Enabled:              models.ToPointer(true),
        SendIpMacMapping:     models.ToPointer(true),
    }),
    SrxApp:                          models.ToPointer(models.SiteSettingSrxApp{
        Enabled:              models.ToPointer(false),
    }),
    SshKeys:                         []string{
        "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host",
    },
    Ssr:                             models.ToPointer(models.SiteSettingSsr{
        ConductorHosts:       []string{
            "\"1.1.1.1\", \"2.2.2.2\"",
        },
        DisableStats:         models.ToPointer(true),
    }),
    StatusPortal:                    models.ToPointer(models.SiteSettingStatusPortal{
        Enabled:              models.ToPointer(false),
        Hostnames:            []string{
            "my.misty.com",
        },
    }),
    Vars:                            map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    Vna:                             models.ToPointer(models.SiteSettingVna{
        Enabled:              models.ToPointer(false),
    }),
    WanVna:                          models.ToPointer(models.SiteSettingWanVna{
        Enabled:              models.ToPointer(false),
    }),
    Wids:                            models.ToPointer(models.SiteWids{
        RepeatedAuthFailures: models.ToPointer(models.SiteWidsRepeatedAuthFailures{
            Duration:             models.ToPointer(60),
            Threshold:            models.ToPointer(0),
        }),
    }),
    Wifi:                            models.ToPointer(models.SiteWifi{
        CiscoEnabled:                      models.ToPointer(true),
        Disable11k:                        models.ToPointer(false),
        DisableRadiosWhenPowerConstrained: models.ToPointer(false),
        EnableArpSpoofCheck:               models.ToPointer(false),
        EnableSharedRadioScanning:         models.ToPointer(true),
        Enabled:                           models.ToPointer(true),
        LocateConnected:                   models.ToPointer(true),
        LocateUnconnected:                 models.ToPointer(false),
        MeshAllowDfs:                      models.ToPointer(false),
        MeshEnableCrm:                     models.ToPointer(false),
        MeshEnabled:                       models.ToPointer(false),
        MeshPsk:                           models.NewOptional(models.ToPointer("string")),
        MeshSsid:                          models.NewOptional(models.ToPointer("string")),
        ProxyArp:                          models.NewOptional(models.ToPointer(models.SiteWifiProxyArpEnum("default"))),
    }),
    WiredVna:                        models.ToPointer(models.SiteSettingWiredVna{
        Enabled:              models.ToPointer(false),
    }),
    ZoneOccupancyAlert:              models.ToPointer(models.SiteZoneOccupancyAlert{
        EmailNotifiers:       []string{
            "foo@juniper.net",
            "bar@juniper.net",
        },
        Enabled:              models.ToPointer(false),
        Threshold:            models.ToPointer(5),
    }),
}

apiResponse, err := sitesSetting.UpdateSiteSettings(ctx, siteId, &body)
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
  "additional_config_cmds": [
    "set snmp community public"
  ],
  "analytic": {
    "enabled": false
  },
  "ap_matching": {
    "enabled": true,
    "rules": [
      {
        "eth1,eth2": {
          "port_vlan_id": 1,
          "vlan_ids": [
            1,
            10,
            50
          ]
        }
      }
    ]
  },
  "ap_port_config": {
    "model_specific": {
      "AP32": {
        "eth1,eth2": {
          "port_vlan_id": 1,
          "vlan_ids": [
            1,
            10,
            50
          ]
        }
      }
    }
  },
  "auto_placement": {
    "orientation": 45,
    "x": 30,
    "y": 60
  },
  "auto_upgrade": {
    "custom_versions": {
      "AP21": "stable",
      "AP41": "0.1.5135",
      "AP61": "0.1.7215"
    },
    "day_of_week": "sun",
    "enabled": false,
    "time_of_day": "12:00",
    "version": "beta"
  },
  "blacklist_url": "https://papi.s3.amazonaws.com/blacklist/xxx...",
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
    "power": 6,
    "power_mode": "custom"
  },
  "config_auto_revert": false,
  "created_time": 0,
  "device_updown_threshold": 0,
  "dns_servers": [
    "string"
  ],
  "dns_suffix": [
    "string"
  ],
  "engagement": {
    "dwell_tag_names": {
      "bounce": "Bounce",
      "engaged": "Engaged",
      "passerby": "Passer By",
      "stationed": "Stationed"
    },
    "dwell_tags": {
      "bounce": null,
      "engaged": "300-14400",
      "passerby": null,
      "stationed": "14400-43200"
    },
    "hours": {
      "fri": "09:00-17:00",
      "mon": "09:00-17:00",
      "sat": "09:00-12:00",
      "sun": "09:00-12:00",
      "thu": "09:00-17:00",
      "tue": "09:00-17:00",
      "wed": "09:00-17:00"
    },
    "max_dwell": 43200,
    "min_dwell": 0
  },
  "evpn_options": {
    "auto_loopback_subnet": "100.101.0.0/16",
    "auto_router_id_subnet": "100.100.0.0/24",
    "core_as_border": false,
    "overlay": {
      "as": 65000
    },
    "per_vlan_vga_v4_mac": false,
    "routed_at": "edge",
    "underlay": {
      "as_base": 65001,
      "routed_id_prefix": "/24",
      "subnet": "10.255.240.0/20"
    }
  },
  "flags": {
    "property1": "string",
    "property2": "string"
  },
  "for_site": true,
  "gateway_additional_config_cmds": [
    "set snmp community public"
  ],
  "gateway_mgmt": {
    "admin_sshkeys": [
      "string"
    ],
    "app_probing": {
      "apps": [
        "string"
      ],
      "custom_apps": [
        {
          "app_type": "string",
          "hostname": [
            "string"
          ],
          "name": "string",
          "protocol": "http"
        }
      ],
      "enabled": true
    },
    "app_usage": true,
    "auto_signature_update": {
      "day_of_week": "mon",
      "enable": true,
      "time_of_day": "string"
    },
    "config_revert_timer": 10,
    "probe_hosts": [
      "string"
    ],
    "root_password": "string",
    "security_log_source_address": "192.168.1.1",
    "security_log_source_interface": "string"
  },
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f09",
  "led": {
    "brightness": 255,
    "enabled": true
  },
  "modified_time": 0,
  "mxedge": {
    "mist_das": {
      "coa_servers": [
        {
          "disable_event_timestamp_check": false,
          "enabled": true,
          "host": "string",
          "port": 3799,
          "secret": "string"
        }
      ],
      "enabled": false
    },
    "radsec": {
      "acct_servers": [
        {
          "host": "string",
          "port": 1813,
          "secret": "string",
          "ssids": [
            "string"
          ]
        }
      ],
      "auth_servers": [
        {
          "host": "string",
          "keywrap_enabled": true,
          "keywrap_format": "hex",
          "keywrap_kek": "string",
          "keywrap_mack": "string",
          "port": 1812,
          "secret": "string",
          "ssids": [
            "string"
          ]
        }
      ],
      "enabled": true,
      "match_ssid": true,
      "proxy_hosts": [
        "string"
      ],
      "server_selection": "ordered",
      "source": "any"
    }
  },
  "mxedge_mgmt": {
    "mist_password": "MIST_PASSWORD",
    "root_password": "ROOT_PASSWORD"
  },
  "ntp_servers": [
    "pool.ntp.org"
  ],
  "occupancy": {
    "assets_enabled": false,
    "clients_enabled": true,
    "min_duration": 3000,
    "sdkclients_enabled": false,
    "unconnected_clients_enabled": false
  },
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "ospf_areas": {
    "property1": {
      "include_loopback": false,
      "networks": {
        "corp": {
          "auth_keys": {
            "1": "auth-key-1"
          },
          "auth_type": "md5",
          "bfd_minimum_interval": 500,
          "dead_interval": 40,
          "hello_interval": 10,
          "interface_type": "nbma",
          "metric": 10000
        },
        "guest": {
          "passive": true
        }
      },
      "type": "default"
    },
    "property2": {
      "include_loopback": false,
      "networks": {
        "corp": {
          "auth_keys": {
            "1": "auth-key-1"
          },
          "auth_type": "md5",
          "bfd_minimum_interval": 500,
          "dead_interval": 40,
          "hello_interval": 10,
          "interface_type": "nbma",
          "metric": 10000
        },
        "guest": {
          "passive": true
        }
      },
      "type": "default"
    }
  },
  "persist_config_on_device": false,
  "port_mirroring": {
    "property1": {
      "input_networks_ingress": [
        "corp"
      ],
      "input_port_ids_egress": [
        "ge-0/0/3"
      ],
      "input_port_ids_ingress": [
        "ge-0/0/3"
      ],
      "output_network": "analyze",
      "output_port_id": "ge-0/0/5"
    },
    "property2": {
      "input_networks_ingress": [
        "corp"
      ],
      "input_port_ids_egress": [
        "ge-0/0/3"
      ],
      "input_port_ids_ingress": [
        "ge-0/0/3"
      ],
      "output_network": "analyze",
      "output_port_id": "ge-0/0/5"
    }
  },
  "port_usages": {
    "dynamic": {
      "mode": "dynamic",
      "reset_default_when": "link_down",
      "rules": [
        {
          "equals": "string",
          "equals_any": [
            "string"
          ],
          "expression": "string",
          "src": "lldp_chassis_id",
          "usage": "string"
        }
      ]
    },
    "property1": {
      "all_networks": false,
      "allow_dhcpd": true,
      "authentication_protocol": "pap",
      "bypass_auth_when_server_down": true,
      "description": "string",
      "disable_autoneg": false,
      "disabled": false,
      "duplex": "auto",
      "enable_mac_auth": true,
      "enable_qos": true,
      "guest_network": "string",
      "mac_auth_only": true,
      "mac_limit": 0,
      "mode": "access",
      "mtu": 0,
      "networks": [
        "string"
      ],
      "persist_mac": false,
      "poe_disabled": false,
      "port_auth": "dot1x",
      "port_network": "string",
      "rejected_network": null,
      "speed": "auto",
      "storm_control": {
        "no_broadcast": false,
        "no_multicast": false,
        "no_registered_multicast": false,
        "no_unknown_unicast": false,
        "percentage": 80
      },
      "stp_edge": true,
      "voip_network": "string"
    },
    "property2": {
      "all_networks": false,
      "allow_dhcpd": true,
      "authentication_protocol": "pap",
      "bypass_auth_when_server_down": true,
      "description": "string",
      "disable_autoneg": false,
      "disabled": false,
      "duplex": "auto",
      "enable_mac_auth": true,
      "enable_qos": true,
      "guest_network": "string",
      "mac_auth_only": true,
      "mac_limit": 0,
      "mode": "access",
      "mtu": 0,
      "networks": [
        "string"
      ],
      "persist_mac": false,
      "poe_disabled": false,
      "port_network": "string",
      "rejected_network": null,
      "speed": "auto",
      "storm_control": {
        "no_broadcast": false,
        "no_multicast": false,
        "no_registered_multicast": false,
        "no_unknown_unicast": false,
        "percentage": 80
      },
      "stp_edge": true,
      "voip_network": "string"
    }
  },
  "proxy": {
    "url": "http://proxy.internal:8080/"
  },
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
  "remote_syslog": {
    "archive": {
      "files": 20,
      "size": "5m"
    },
    "console": {
      "contents": [
        {
          "facility": "config",
          "severity": "warning"
        }
      ]
    },
    "enabled": false,
    "files": [
      {
        "archive": {
          "files": 10,
          "size": "5m"
        },
        "contents": [
          {
            "facility": "config",
            "severity": "warning"
          }
        ],
        "explicit_priority": true,
        "file": "file-name",
        "match": "!alarm|ntp|errors.crc_error[chan]",
        "structured_data": true
      }
    ],
    "network": "default",
    "send_to_all_servers": false,
    "servers": [
      {
        "facility": "config",
        "host": "syslogd.internal",
        "port": 514,
        "protocol": "udp",
        "severity": "info",
        "tag": ""
      }
    ],
    "time_format": "millisecond",
    "users": [
      {
        "contents": [
          {
            "facility": "config",
            "severity": "warning"
          }
        ],
        "match": "\"!alarm|ntp|errors.crc_error[chan]\"",
        "user": "*"
      }
    ]
  },
  "report_gatt": false,
  "rogue": {
    "enabled": false,
    "honeypot_enabled": false,
    "min_duration": 10,
    "min_rssi": -80,
    "whitelisted_bssids": [
      "NeighborSSID"
    ],
    "whitelisted_ssids": [
      "cc:8e:6f:d4:bf:16",
      "cc-8e-6f-d4-bf-16",
      "cc-73-*",
      "cc:82:*"
    ]
  },
  "rtsa": {
    "app_waking": false,
    "disable_dead_reckoning": true,
    "disable_pressure_sensor": false,
    "enabled": true,
    "track_asset": false
  },
  "simple_alert": {
    "arp_failure": {
      "client_count": 10,
      "duration": 20,
      "incident_count": 10
    },
    "dhcp_failure": {
      "client_count": 10,
      "duration": 10,
      "incident_count": 20
    },
    "dns_failure": {
      "client_count": 20,
      "duration": 10,
      "incident_count": 30
    }
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "skyatp": {
    "enabled": true,
    "send_ip_mac_mapping": true
  },
  "srx_app": {
    "enabled": false
  },
  "ssh_keys": [
    "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host"
  ],
  "ssr": {
    "conductor_hosts": [
      "\"1.1.1.1\", \"2.2.2.2\""
    ],
    "disable_stats": true
  },
  "status_portal": {
    "enabled": false,
    "hostnames": [
      "my.misty.com"
    ]
  },
  "switch_mgmt": {
    "ap_affinity_threshold": 10,
    "config_revert_timer": 10,
    "dhcp_option_fqdn": false,
    "mxedge_proxy_host": "string",
    "mxedge_proxy_port": 2222,
    "root_password": "string",
    "tacacs": {
      "acct_servers": [
        {
          "host": "198.51.100.1",
          "port": "49",
          "secret": "string",
          "timeout": 10
        }
      ],
      "enabled": true,
      "network": "string",
      "tacplus_servers": [
        {
          "host": "198.51.100.1",
          "port": "49",
          "secret": "string",
          "timeout": 10
        }
      ]
    },
    "use_mxedge_proxy": true
  },
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "vna": {
    "enabled": false
  },
  "vrf_instances": {
    "guest": {
      "extra_routes": {
        "0.0.0.0/0": {
          "via": "192.168.31.1"
        }
      },
      "networks": [
        "guest"
      ]
    }
  },
  "vrrp_groups": {
    "property1": {
      "auth_key": "auth-key-1",
      "auth_password": "string",
      "auth_type": "md5",
      "networks": {
        "data": {
          "ip": "10.182.96.1"
        },
        "mgmt": {
          "ip": "10.182.104.1"
        },
        "v10": {
          "ip": "10.182.104.129"
        },
        "wap": {
          "ip": "10.182.102.1"
        }
      }
    },
    "property2": {
      "auth_key": "auth-key-1",
      "auth_password": "string",
      "auth_type": "md5",
      "networks": {
        "data": {
          "ip": "10.182.96.1"
        },
        "mgmt": {
          "ip": "10.182.104.1"
        },
        "v10": {
          "ip": "10.182.104.129"
        },
        "wap": {
          "ip": "10.182.102.1"
        }
      }
    }
  },
  "wan_vna": {
    "enabled": false
  },
  "watched_station_url": "https://papi.s3.amazonaws.com/watched_station/xxx...",
  "whitelist_url": "https://papi.s3.amazonaws.com/whitelist/xxx...",
  "wids": {
    "repeated_auth_failures": {
      "duration": 60,
      "threshold": 0
    }
  },
  "wifi": {
    "cisco_enabled": true,
    "disable_11k": false,
    "disable_radios_when_power_constrained": false,
    "enable_arp_spoof_check": false,
    "enable_shared_radio_scanning": true,
    "enabled": true,
    "locate_connected": true,
    "locate_unconnected": false,
    "mesh_allow_dfs": false,
    "mesh_enable_crm": false,
    "mesh_enabled": false,
    "mesh_psk": "string",
    "mesh_ssid": "string",
    "proxy_arp": "default"
  },
  "wired_vna": {
    "enabled": false
  },
  "zone_occupancy_alert": {
    "email_notifiers": [
      "foo@juniper.net",
      "bar@juniper.net"
    ],
    "enabled": false,
    "threshold": 5
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

