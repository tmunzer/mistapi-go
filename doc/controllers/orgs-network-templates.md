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
* [Update Org Network Template](../../doc/controllers/orgs-network-templates.md#update-org-network-template)


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.NetworkTemplate](../../doc/models/network-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NetworkTemplate{
    AdditionalConfigCmds:  []string{
        "set snmp community public",
    },
    DhcpSnooping:          models.ToPointer(models.DhcpSnooping{
        AllNetworks:          models.ToPointer(true),
        EnableArpSpoofCheck:  models.ToPointer(true),
        EnableIpSourceGuard:  models.ToPointer(true),
        Enabled:              models.ToPointer(true),
        Networks:             []string{
            "string",
        },
    }),
    DnsServers:            []string{
        "string",
    },
    DnsSuffix:             []string{
        "string",
    },
    ImportOrgNetworks:     []string{
        "ap",
    },
    MistNac:               models.ToPointer(models.SwitchMistNac{
        Enabled:              models.ToPointer(true),
        Network:              models.ToPointer("string"),
    }),
    Name:                  models.ToPointer("string"),
    Networks:              map[string]models.SwitchNetwork{
        "property1": models.SwitchNetwork{
            Subnet:               models.ToPointer("192.168.1.0/24"),
            VlanId:               models.VlanIdWithVariableContainer.FromNumber(10),
        },
        "property2": models.SwitchNetwork{
            Subnet:               models.ToPointer("192.168.1.0/24"),
            VlanId:               models.VlanIdWithVariableContainer.FromNumber(10),
        },
    },
    NtpServers:            []string{
        "string",
    },
    PortUsages:            map[string]models.SwitchPortUsage{
        "dynamic": models.SwitchPortUsage{
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum_DYNAMIC),
            ResetDefaultWhen:                         models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum_LINKDOWN),
            Rules:                                    []models.SwitchPortUsageDynamicRule{
                models.SwitchPortUsageDynamicRule{
                    Equals:               models.ToPointer("string"),
                    EqualsAny:            []string{
                        "string",
                    },
                    Expression:           models.ToPointer("string"),
                    Src:                  models.SwitchPortUsageDynamicRuleSrcEnum_LLDPCHASSISID,
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
            Duplex:                                   models.ToPointer(models.SwitchPortUsageDuplexEnum_AUTO),
            EnableMacAuth:                            models.ToPointer(true),
            EnableQos:                                models.ToPointer(true),
            GuestNetwork:                             models.NewOptional(models.ToPointer("string")),
            MacAuthOnly:                              models.ToPointer(true),
            MacLimit:                                 models.ToPointer(models.SwitchPortUsageMacLimitContainer.FromNumber(0)),
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum_ACCESS),
            Networks:                                 []string{
                "string",
            },
            PersistMac:                               models.ToPointer(false),
            PoeDisabled:                              models.ToPointer(false),
            PortAuth:                                 models.NewOptional(models.ToPointer(models.SwitchPortUsageDot1xEnum_DOT1X)),
            PortNetwork:                              models.ToPointer("string"),
            Speed:                                    models.ToPointer(models.SwitchPortUsageSpeedEnum_AUTO),
            StormControl:                             models.ToPointer(models.SwitchPortUsageStormControl{
                NoBroadcast:           models.ToPointer(false),
                NoMulticast:           models.ToPointer(false),
                NoRegisteredMulticast: models.ToPointer(false),
                NoUnknownUnicast:      models.ToPointer(false),
                Percentage:            models.ToPointer(80),
            }),
            StpEdge:                                  models.ToPointer(true),
            VoipNetwork:                              models.NewOptional(models.ToPointer("string")),
        },
        "property2": models.SwitchPortUsage{
            AllNetworks:                              models.ToPointer(false),
            AllowDhcpd:                               models.ToPointer(true),
            BypassAuthWhenServerDown:                 models.ToPointer(true),
            Description:                              models.ToPointer("string"),
            DisableAutoneg:                           models.ToPointer(false),
            Disabled:                                 models.ToPointer(false),
            Duplex:                                   models.ToPointer(models.SwitchPortUsageDuplexEnum_AUTO),
            EnableMacAuth:                            models.ToPointer(true),
            EnableQos:                                models.ToPointer(true),
            GuestNetwork:                             models.NewOptional(models.ToPointer("string")),
            MacAuthOnly:                              models.ToPointer(true),
            MacLimit:                                 models.ToPointer(models.SwitchPortUsageMacLimitContainer.FromNumber(0)),
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum_ACCESS),
            Networks:                                 []string{
                "string",
            },
            PersistMac:                               models.ToPointer(false),
            PoeDisabled:                              models.ToPointer(false),
            PortNetwork:                              models.ToPointer("string"),
            Speed:                                    models.ToPointer(models.SwitchPortUsageSpeedEnum_AUTO),
            StormControl:                             models.ToPointer(models.SwitchPortUsageStormControl{
                NoBroadcast:           models.ToPointer(false),
                NoMulticast:           models.ToPointer(false),
                NoRegisteredMulticast: models.ToPointer(false),
                NoUnknownUnicast:      models.ToPointer(false),
                Percentage:            models.ToPointer(80),
            }),
            StpEdge:                                  models.ToPointer(true),
            VoipNetwork:                              models.NewOptional(models.ToPointer("string")),
        },
    },
    SwitchMgmt:            models.ToPointer(models.SwitchMgmt{
        ProtectRe:             models.ToPointer(models.ProtectRe{
            Enabled:              models.ToPointer(false),
        }),
        RootPassword:          models.ToPointer("string"),
        Tacacs:                models.ToPointer(models.Tacacs{
            AcctServers:          []models.TacacsAcctServer{
                models.TacacsAcctServer{
                    Host:                 models.ToPointer("198.51.100.1"),
                    Port:                 models.ToPointer("49"),
                    Secret:               models.ToPointer("string"),
                    Timeout:              models.ToPointer(10),
                },
            },
            Enabled:              models.ToPointer(true),
            Network:              models.ToPointer("string"),
            TacplusServers:       []models.TacacsAuthServer{
                models.TacacsAuthServer{
                    Host:                 models.ToPointer("198.51.100.1"),
                    Port:                 models.ToPointer("49"),
                    Secret:               models.ToPointer("string"),
                    Timeout:              models.ToPointer(10),
                },
            },
        }),
        AdditionalProperties:  map[string]interface{}{
            "config_revert": interface{}("10"),
        },
    }),
    VrfConfig:             models.ToPointer(models.VrfConfig{
        Enabled:              models.ToPointer(false),
    }),
    VrfInstances:          map[string]models.SwitchVrfInstance{
        "property1": models.SwitchVrfInstance{
            ExtraRoutes:             map[string]models.VrfExtraRoute{
                "property1": models.VrfExtraRoute{
                    Via:                  models.ToPointer("192.0.2.10"),
                },
                "property2": models.VrfExtraRoute{
                    Via:                  models.ToPointer("198.51.100.1"),
                },
            },
            Networks:                []string{
                "string",
            },
        },
        "property2": models.SwitchVrfInstance{
            ExtraRoutes:             map[string]models.VrfExtraRoute{
                "property1": models.VrfExtraRoute{
                    Via:                  models.ToPointer("198.51.100.1"),
                },
                "property2": models.VrfExtraRoute{
                    Via:                  models.ToPointer("198.51.100.10"),
                },
            },
            Networks:                []string{
                "string",
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

## Example Response *(as JSON)*

```json
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
    "string"
  ],
  "dns_suffix": [
    "string"
  ],
  "extra_routes": {
    "property1": {
      "via": "string"
    },
    "property2": {
      "via": "string"
    }
  },
  "group_tags": {},
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6708",
  "import_org_networks": [
    "ap"
  ],
  "mist_nac": {
    "enabled": true,
    "network": "string"
  },
  "modified_time": 0,
  "name": "string",
  "networks": {
    "property1": {
      "subnet": "192.168.1.0/24",
      "vlan_id": 10
    },
    "property2": {
      "subnet": "192.168.1.0/24",
      "vlan_id": 10
    }
  },
  "ntp_servers": [
    "string"
  ],
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
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
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.NetworkTemplate](../../doc/models/network-template.md).

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Network Templates

Get List of Org Network Templates

```go
ListOrgNetworkTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.NetworkTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.NetworkTemplate](../../doc/models/network-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsNetworkTemplates.ListOrgNetworkTemplates(ctx, orgId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Network Template

Update Org Network Template

```go
UpdateOrgNetworkTemplate(
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.NetworkTemplate](../../doc/models/network-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networktemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NetworkTemplate{
    AdditionalConfigCmds:  []string{
        "set snmp community public",
    },
    DhcpSnooping:          models.ToPointer(models.DhcpSnooping{
        AllNetworks:          models.ToPointer(true),
        EnableArpSpoofCheck:  models.ToPointer(true),
        EnableIpSourceGuard:  models.ToPointer(true),
        Enabled:              models.ToPointer(true),
        Networks:             []string{
            "string",
        },
    }),
    DnsServers:            []string{
        "string",
    },
    DnsSuffix:             []string{
        "string",
    },
    ExtraRoutes:           map[string]models.ExtraRoute{
        "property1": models.ExtraRoute{
            Via:                  models.ToPointer("string"),
        },
        "property2": models.ExtraRoute{
            Via:                  models.ToPointer("string"),
        },
    },
    ImportOrgNetworks:     []string{
        "ap",
    },
    MistNac:               models.ToPointer(models.SwitchMistNac{
        Enabled:              models.ToPointer(true),
        Network:              models.ToPointer("string"),
    }),
    Name:                  models.ToPointer("string"),
    Networks:              map[string]models.SwitchNetwork{
        "property1": models.SwitchNetwork{
            Subnet:               models.ToPointer("192.168.1.0/24"),
            VlanId:               models.VlanIdWithVariableContainer.FromNumber(10),
        },
        "property2": models.SwitchNetwork{
            Subnet:               models.ToPointer("192.168.1.0/24"),
            VlanId:               models.VlanIdWithVariableContainer.FromNumber(10),
        },
    },
    NtpServers:            []string{
        "string",
    },
    PortUsages:            map[string]models.SwitchPortUsage{
        "dynamic": models.SwitchPortUsage{
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum_DYNAMIC),
            ResetDefaultWhen:                         models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum_LINKDOWN),
            Rules:                                    []models.SwitchPortUsageDynamicRule{
                models.SwitchPortUsageDynamicRule{
                    Equals:               models.ToPointer("string"),
                    EqualsAny:            []string{
                        "string",
                    },
                    Expression:           models.ToPointer("string"),
                    Src:                  models.SwitchPortUsageDynamicRuleSrcEnum_LLDPCHASSISID,
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
            Duplex:                                   models.ToPointer(models.SwitchPortUsageDuplexEnum_AUTO),
            EnableMacAuth:                            models.ToPointer(true),
            EnableQos:                                models.ToPointer(true),
            GuestNetwork:                             models.NewOptional(models.ToPointer("string")),
            MacAuthOnly:                              models.ToPointer(true),
            MacLimit:                                 models.ToPointer(models.SwitchPortUsageMacLimitContainer.FromNumber(0)),
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum_ACCESS),
            Networks:                                 []string{
                "string",
            },
            PersistMac:                               models.ToPointer(false),
            PoeDisabled:                              models.ToPointer(false),
            PortAuth:                                 models.NewOptional(models.ToPointer(models.SwitchPortUsageDot1xEnum_DOT1X)),
            PortNetwork:                              models.ToPointer("string"),
            Speed:                                    models.ToPointer(models.SwitchPortUsageSpeedEnum_AUTO),
            StormControl:                             models.ToPointer(models.SwitchPortUsageStormControl{
                NoBroadcast:           models.ToPointer(false),
                NoMulticast:           models.ToPointer(false),
                NoRegisteredMulticast: models.ToPointer(false),
                NoUnknownUnicast:      models.ToPointer(false),
                Percentage:            models.ToPointer(80),
            }),
            StpEdge:                                  models.ToPointer(true),
            VoipNetwork:                              models.NewOptional(models.ToPointer("string")),
        },
        "property2": models.SwitchPortUsage{
            AllNetworks:                              models.ToPointer(false),
            AllowDhcpd:                               models.ToPointer(true),
            BypassAuthWhenServerDown:                 models.ToPointer(true),
            Description:                              models.ToPointer("string"),
            DisableAutoneg:                           models.ToPointer(false),
            Disabled:                                 models.ToPointer(false),
            Duplex:                                   models.ToPointer(models.SwitchPortUsageDuplexEnum_AUTO),
            EnableMacAuth:                            models.ToPointer(true),
            EnableQos:                                models.ToPointer(true),
            GuestNetwork:                             models.NewOptional(models.ToPointer("string")),
            MacAuthOnly:                              models.ToPointer(true),
            MacLimit:                                 models.ToPointer(models.SwitchPortUsageMacLimitContainer.FromNumber(0)),
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum_ACCESS),
            Networks:                                 []string{
                "string",
            },
            PersistMac:                               models.ToPointer(false),
            PoeDisabled:                              models.ToPointer(false),
            PortNetwork:                              models.ToPointer("string"),
            Speed:                                    models.ToPointer(models.SwitchPortUsageSpeedEnum_AUTO),
            StormControl:                             models.ToPointer(models.SwitchPortUsageStormControl{
                NoBroadcast:           models.ToPointer(false),
                NoMulticast:           models.ToPointer(false),
                NoRegisteredMulticast: models.ToPointer(false),
                NoUnknownUnicast:      models.ToPointer(false),
                Percentage:            models.ToPointer(80),
            }),
            StpEdge:                                  models.ToPointer(true),
            VoipNetwork:                              models.NewOptional(models.ToPointer("string")),
        },
    },
    SwitchMgmt:            models.ToPointer(models.SwitchMgmt{
        ProtectRe:             models.ToPointer(models.ProtectRe{
            Enabled:              models.ToPointer(false),
        }),
        RootPassword:          models.ToPointer("string"),
        Tacacs:                models.ToPointer(models.Tacacs{
            AcctServers:          []models.TacacsAcctServer{
                models.TacacsAcctServer{
                    Host:                 models.ToPointer("198.51.100.1"),
                    Port:                 models.ToPointer("49"),
                    Secret:               models.ToPointer("string"),
                    Timeout:              models.ToPointer(10),
                },
            },
            Enabled:              models.ToPointer(true),
            Network:              models.ToPointer("string"),
            TacplusServers:       []models.TacacsAuthServer{
                models.TacacsAuthServer{
                    Host:                 models.ToPointer("198.51.100.1"),
                    Port:                 models.ToPointer("49"),
                    Secret:               models.ToPointer("string"),
                    Timeout:              models.ToPointer(10),
                },
            },
        }),
        AdditionalProperties:  map[string]interface{}{
            "config_revert": interface{}("10"),
        },
    }),
    VrfConfig:             models.ToPointer(models.VrfConfig{
        Enabled:              models.ToPointer(false),
    }),
    VrfInstances:          map[string]models.SwitchVrfInstance{
        "property1": models.SwitchVrfInstance{
            ExtraRoutes:             map[string]models.VrfExtraRoute{
                "property1": models.VrfExtraRoute{
                    Via:                  models.ToPointer("198.51.100.1"),
                },
                "property2": models.VrfExtraRoute{
                    Via:                  models.ToPointer("198.51.100.2"),
                },
            },
            Networks:                []string{
                "string",
            },
        },
        "property2": models.SwitchVrfInstance{
            ExtraRoutes:             map[string]models.VrfExtraRoute{
                "property1": models.VrfExtraRoute{
                    Via:                  models.ToPointer("198.51.100.1"),
                },
                "property2": models.VrfExtraRoute{
                    Via:                  models.ToPointer("198.51.100.2"),
                },
            },
            Networks:                []string{
                "string",
            },
        },
    },
    AdditionalProperties:  map[string]interface{}{
        "group_tags": interface{}(""),
    },
}

apiResponse, err := orgsNetworkTemplates.UpdateOrgNetworkTemplate(ctx, orgId, networktemplateId, &body)
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
    "string"
  ],
  "dns_suffix": [
    "string"
  ],
  "extra_routes": {
    "property1": {
      "via": "string"
    },
    "property2": {
      "via": "string"
    }
  },
  "import_org_networks": [
    "ap"
  ],
  "mist_nac": {
    "enabled": true,
    "network": "string"
  },
  "name": "string",
  "networks": {
    "property1": {
      "subnet": "192.168.1.0/24",
      "vlan_id": 10
    },
    "property2": {
      "subnet": "192.168.1.0/24",
      "vlan_id": 10
    }
  },
  "ntp_servers": [
    "string"
  ],
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
          "via": "198.51.100.2"
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
          "via": "198.51.100.2"
        }
      },
      "networks": [
        "string"
      ]
    }
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

