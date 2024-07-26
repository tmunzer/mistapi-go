# Orgs Network Templates

```go
orgsNetworkTemplates := client.OrgsNetworkTemplates()
```

## Class Name

`OrgsNetworkTemplates`

## Methods

* [Create Org Network Template](../../doc/controllers/orgs-network-templates.md#create-org-network-template)
* [Delete Org Network Template](../../doc/controllers/orgs-network-templates.md#delete-org-network-template)
* [Get Org Network Template](../../doc/controllers/orgs-network-templates.md#get-org-network-template)
* [List Org Network Templates](../../doc/controllers/orgs-network-templates.md#list-org-network-templates)
* [Update Org Network Templates](../../doc/controllers/orgs-network-templates.md#update-org-network-templates)


# Create Org Network Template

Update Org Network Templates

```go
CreateOrgNetworkTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NetworkTemplate) (
    models.ApiResponse[models.NetworkTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NetworkTemplate`](../../doc/models/network-template.md) | Body, Optional | Request Body |

## Response Type

[`models.NetworkTemplate`](../../doc/models/network-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NetworkTemplate{
    ExtraRoutes6:         map[string]models.ExtraRoute6{
        "2a02:1234:420a:10c9::/64": models.ExtraRoute6{
            Via:           models.ToPointer("2a02:1234:200a::100"),
        },
    },
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    VrfInstances:         map[string]models.VrfInstance{
        "guest": models.VrfInstance{
            ExtraRoutes: map[string]models.VrfExtraRoute{
                "0.0.0.0/0": models.VrfExtraRoute{
                    Via: models.ToPointer("192.168.31.1"),
                },
            },
            Networks:    []string{
                "guest",
            },
        },
    },
}

apiResponse, err := orgsNetworkTemplates.CreateOrgNetworkTemplate(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Delete Org Network Template

Delete Org Network Template

```go
DeleteOrgNetworkTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    networktemplateId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networktemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networktemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsNetworkTemplates.DeleteOrgNetworkTemplate(ctx, orgId, networktemplateId)
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


# Get Org Network Template

Get Org Network Templates Details

```go
GetOrgNetworkTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    networktemplateId uuid.UUID) (
    models.ApiResponse[models.NetworkTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networktemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.NetworkTemplate`](../../doc/models/network-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networktemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsNetworkTemplates.GetOrgNetworkTemplate(ctx, orgId, networktemplateId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Network Templates

Get List of Org Network Templates

```go
ListOrgNetworkTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.NetworkTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.NetworkTemplate`](../../doc/models/network-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsNetworkTemplates.ListOrgNetworkTemplates(ctx, orgId, &page, &limit)
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
    "additional_config_cmds": [
      "set snmp community public"
    ],
    "created_time": 0,
    "dhcp_snooping": {
      "all_networks": true,
      "enable_arp_spoof_check": true,
      "enable_ip_source_guard": true,
      "enabled": true,
      "networks": [
        "string"
      ]
    },
    "dns_servers": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "dns_suffix": [
      ".mist.local",
      ".mist.com"
    ],
    "extra_routes": {
      "0.0.0.0/0": {
        "via": "1.2.3.4"
      }
    },
    "group_tags": {},
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6808",
    "import_org_networks": [
      "ap"
    ],
    "mist_nac": {
      "enabled": true,
      "network": "default"
    },
    "modified_time": 0,
    "name": "template_name",
    "networks": {
      "corp": {
        "vlan_id": 600
      },
      "default": {
        "subnet": "192.168.1.0/24",
        "vlan_id": 1
      },
      "guest": {
        "vlan_id": 700
      },
      "mgmt": {
        "vlan_id": 500
      }
    },
    "ntp_servers": [
      "192.168.1.10"
    ],
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "port_usages": {
      "ap": {
        "all_networks": false,
        "allow_dhcpd": true,
        "authentication_protocol": "pap",
        "bypass_auth_when_server_down": true,
        "description": "WAP",
        "disable_autoneg": false,
        "disabled": false,
        "duplex": "auto",
        "enable_mac_auth": false,
        "enable_qos": true,
        "mac_auth_only": false,
        "mac_limit": 0,
        "mode": "trunk",
        "mtu": 9192,
        "networks": [
          "guest",
          "corp"
        ],
        "persist_mac": false,
        "poe_disabled": false,
        "port_network": "default",
        "rejected_network": null,
        "storm_control": {
          "no_broadcast": false,
          "no_multicast": false,
          "no_registered_multicast": false,
          "no_unknown_unicast": false,
          "percentage": 80
        },
        "stp_edge": true
      },
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
      "iot": {
        "allow_dhcpd": true,
        "mode": "access",
        "port_network": "default",
        "stp_edge": true
      },
      "uplink": {
        "all_networks": true,
        "enable_qos": false,
        "mode": "trunk",
        "port_network": "default",
        "stp_edge": false
      }
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
      "network": "default"
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
    "switch_matching": {
      "enable": true,
      "rules": [
        {
          "additional_config_cmds": [
            "set snmp community public"
          ],
          "match_model": "EX4300",
          "match_name[0:3]": "abc",
          "name": "match by name",
          "port_config": {
            "ge-0/0/0": {
              "usage": "uplink"
            },
            "ge-0/0/8-16ge-1/0/0-47": {
              "usage": "ap"
            }
          }
        },
        {
          "additional_config_cmds": [
            "set snmp community public2"
          ],
          "match_role": "access",
          "name": "match by role",
          "port_config": {
            "ge-0/0/0": {
              "usage": "uplink"
            }
          }
        }
      ]
    },
    "switch_mgmt": {
      "config_revert": 10,
      "protect_re": {
        "enabled": false
      },
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
      }
    },
    "vrf_config": {
      "enabled": false
    },
    "vrf_instances": {
      "property1": {
        "extra_routes": {
          "property1": {
            "via": "198.51.100.1"
          },
          "property2": {
            "via": "198.51.100.10"
          }
        },
        "networks": [
          "string"
        ]
      },
      "property2": {
        "extra_routes": {
          "property1": {
            "via": "198.51.100.1"
          },
          "property2": {
            "via": "198.51.100.10"
          }
        },
        "networks": [
          "string"
        ]
      }
    }
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


# Update Org Network Templates

Update Org Network Template

```go
UpdateOrgNetworkTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    networktemplateId uuid.UUID,
    body *models.NetworkTemplate) (
    models.ApiResponse[models.NetworkTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networktemplateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NetworkTemplate`](../../doc/models/network-template.md) | Body, Optional | Request Body |

## Response Type

[`models.NetworkTemplate`](../../doc/models/network-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networktemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NetworkTemplate{
    ExtraRoutes6:         map[string]models.ExtraRoute6{
        "2a02:1234:420a:10c9::/64": models.ExtraRoute6{
            Via:           models.ToPointer("2a02:1234:200a::100"),
        },
    },
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    VrfInstances:         map[string]models.VrfInstance{
        "guest": models.VrfInstance{
            ExtraRoutes: map[string]models.VrfExtraRoute{
                "0.0.0.0/0": models.VrfExtraRoute{
                    Via: models.ToPointer("192.168.31.1"),
                },
            },
            Networks:    []string{
                "guest",
            },
        },
    },
}

apiResponse, err := orgsNetworkTemplates.UpdateOrgNetworkTemplates(ctx, orgId, networktemplateId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

