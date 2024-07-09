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
    AdditionalConfigCmds: []string{
        "set snmp community public",
    },
    DhcpSnooping:         models.ToPointer(models.DhcpSnooping{
        AllNetworks:         models.ToPointer(true),
        EnableArpSpoofCheck: models.ToPointer(true),
        EnableIpSourceGuard: models.ToPointer(true),
        Enabled:             models.ToPointer(true),
        Networks:            []string{
            "string",
        },
    }),
    DnsServers:           []string{
        "string",
    },
    DnsSuffix:            []string{
        "string",
    },
    ImportOrgNetworks:    []string{
        "ap",
    },
    MistNac:              models.ToPointer(models.SwitchMistNac{
        Enabled: models.ToPointer(true),
        Network: models.ToPointer("string"),
    }),
    Name:                 models.ToPointer("string"),
    Networks:             map[string]models.SwitchNetwork{
        "property1": models.SwitchNetwork{
            Subnet:          models.ToPointer("192.168.1.0/24"),
            VlanId:          10,
        },
        "property2": models.SwitchNetwork{
            Subnet:          models.ToPointer("192.168.1.0/24"),
            VlanId:          10,
        },
    },
    NtpServers:           []string{
        "string",
    },
    PortUsages:           map[string]models.SwitchPortUsage{
        "dynamic": models.SwitchPortUsage{
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum("dynamic")),
            ResetDefaultWhen:                         models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum("link_down")),
            Rules:                                    []models.SwitchPortUsageDynamicRule{
                models.SwitchPortUsageDynamicRule{
                    Equals:     models.ToPointer("string"),
                    EqualsAny:  []string{
                        "string",
                    },
                    Expression: models.ToPointer("string"),
                    Src:        models.SwitchPortUsageDynamicRuleSrcEnum("lldp_chassis_id"),
                    Usage:      models.ToPointer("string"),
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
            PortAuth:                                 models.ToPointer("string"),
            PortNetwork:                              models.ToPointer("string"),
            RejectedNetwork:                          models.NewOptional(models.ToPointer("rejected_network4")),
            Speed:                                    models.ToPointer("string"),
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
            PortAuth:                                 models.ToPointer("string"),
            PortNetwork:                              models.ToPointer("string"),
            RejectedNetwork:                          models.NewOptional(models.ToPointer("rejected_network4")),
            Speed:                                    models.ToPointer("string"),
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
    SwitchMgmt:           models.ToPointer(models.SwitchMgmt{
        ConfigRevert: models.ToPointer(10),
        ProtectRe:    models.ToPointer(models.ProtectRe{
            Enabled:         models.ToPointer(false),
        }),
        RootPassword: models.ToPointer("string"),
        Tacacs:       models.ToPointer(models.Tacacs{
            AcctServers:    []models.TacacsAcctServer{
                models.TacacsAcctServer{
                    Host:    models.ToPointer("198.51.100.1"),
                    Port:    models.ToPointer("49"),
                    Secret:  models.ToPointer("string"),
                    Timeout: models.ToPointer(10),
                },
            },
            Enabled:        models.ToPointer(true),
            Network:        models.ToPointer("string"),
            TacplusServers: []models.TacacsAuthServer{
                models.TacacsAuthServer{
                    Host:    models.ToPointer("198.51.100.1"),
                    Port:    models.ToPointer("49"),
                    Secret:  models.ToPointer("string"),
                    Timeout: models.ToPointer(10),
                },
            },
        }),
    }),
    VrfConfig:            models.ToPointer(models.VrfConfig{
        Enabled: models.ToPointer(false),
    }),
    VrfInstances:         map[string]models.VrfInstance{
        "property1": models.VrfInstance{
            ExtraRoutes: map[string]models.VrfExtraRoute{
                "property1": models.VrfExtraRoute{
                    Via: models.ToPointer("192.0.2.10"),
                },
                "property2": models.VrfExtraRoute{
                    Via: models.ToPointer("198.51.100.1"),
                },
            },
            Networks:    []string{
                "string",
            },
        },
        "property2": models.VrfInstance{
            ExtraRoutes: map[string]models.VrfExtraRoute{
                "property1": models.VrfExtraRoute{
                    Via: models.ToPointer("198.51.100.1"),
                },
                "property2": models.VrfExtraRoute{
                    Via: models.ToPointer("198.51.100.10"),
                },
            },
            Networks:    []string{
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
      "mode": "access",
      "mtu": 0,
      "networks": [
        "string"
      ],
      "persist_mac": false,
      "poe_disabled": false,
      "port_auth": "string",
      "port_network": "string",
      "rejected_network": null,
      "speed": "string",
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
      "port_auth": "string",
      "port_network": "string",
      "rejected_network": null,
      "speed": "string",
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
    AdditionalConfigCmds: []string{
        "set snmp community public",
    },
    DhcpSnooping:         models.ToPointer(models.DhcpSnooping{
        AllNetworks:         models.ToPointer(true),
        EnableArpSpoofCheck: models.ToPointer(true),
        EnableIpSourceGuard: models.ToPointer(true),
        Enabled:             models.ToPointer(true),
        Networks:            []string{
            "string",
        },
    }),
    DnsServers:           []string{
        "string",
    },
    DnsSuffix:            []string{
        "string",
    },
    ExtraRoutes:          map[string]models.ExtraRouteProperties{
        "property1": models.ExtraRouteProperties{
            Via:           models.ToPointer("string"),
        },
        "property2": models.ExtraRouteProperties{
            Via:           models.ToPointer("string"),
        },
    },
    ImportOrgNetworks:    []string{
        "ap",
    },
    MistNac:              models.ToPointer(models.SwitchMistNac{
        Enabled: models.ToPointer(true),
        Network: models.ToPointer("string"),
    }),
    Name:                 models.ToPointer("string"),
    Networks:             map[string]models.SwitchNetwork{
        "property1": models.SwitchNetwork{
            Subnet:          models.ToPointer("192.168.1.0/24"),
            VlanId:          10,
        },
        "property2": models.SwitchNetwork{
            Subnet:          models.ToPointer("192.168.1.0/24"),
            VlanId:          10,
        },
    },
    NtpServers:           []string{
        "string",
    },
    PortUsages:           map[string]models.SwitchPortUsage{
        "dynamic": models.SwitchPortUsage{
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum("dynamic")),
            ResetDefaultWhen:                         models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum("link_down")),
            Rules:                                    []models.SwitchPortUsageDynamicRule{
                models.SwitchPortUsageDynamicRule{
                    Equals:     models.ToPointer("string"),
                    EqualsAny:  []string{
                        "string",
                    },
                    Expression: models.ToPointer("string"),
                    Src:        models.SwitchPortUsageDynamicRuleSrcEnum("lldp_chassis_id"),
                    Usage:      models.ToPointer("string"),
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
            PortAuth:                                 models.ToPointer("string"),
            PortNetwork:                              models.ToPointer("string"),
            RejectedNetwork:                          models.NewOptional(models.ToPointer("rejected_network4")),
            Speed:                                    models.ToPointer("string"),
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
            PortAuth:                                 models.ToPointer("string"),
            PortNetwork:                              models.ToPointer("string"),
            RejectedNetwork:                          models.NewOptional(models.ToPointer("rejected_network4")),
            Speed:                                    models.ToPointer("string"),
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
    SwitchMgmt:           models.ToPointer(models.SwitchMgmt{
        ConfigRevert: models.ToPointer(10),
        ProtectRe:    models.ToPointer(models.ProtectRe{
            Enabled:         models.ToPointer(false),
        }),
        RootPassword: models.ToPointer("string"),
        Tacacs:       models.ToPointer(models.Tacacs{
            AcctServers:    []models.TacacsAcctServer{
                models.TacacsAcctServer{
                    Host:    models.ToPointer("198.51.100.1"),
                    Port:    models.ToPointer("49"),
                    Secret:  models.ToPointer("string"),
                    Timeout: models.ToPointer(10),
                },
            },
            Enabled:        models.ToPointer(true),
            Network:        models.ToPointer("string"),
            TacplusServers: []models.TacacsAuthServer{
                models.TacacsAuthServer{
                    Host:    models.ToPointer("198.51.100.1"),
                    Port:    models.ToPointer("49"),
                    Secret:  models.ToPointer("string"),
                    Timeout: models.ToPointer(10),
                },
            },
        }),
    }),
    VrfConfig:            models.ToPointer(models.VrfConfig{
        Enabled: models.ToPointer(false),
    }),
    VrfInstances:         map[string]models.VrfInstance{
        "property1": models.VrfInstance{
            ExtraRoutes: map[string]models.VrfExtraRoute{
                "property1": models.VrfExtraRoute{
                    Via: models.ToPointer("198.51.100.1"),
                },
                "property2": models.VrfExtraRoute{
                    Via: models.ToPointer("198.51.100.2"),
                },
            },
            Networks:    []string{
                "string",
            },
        },
        "property2": models.VrfInstance{
            ExtraRoutes: map[string]models.VrfExtraRoute{
                "property1": models.VrfExtraRoute{
                    Via: models.ToPointer("198.51.100.1"),
                },
                "property2": models.VrfExtraRoute{
                    Via: models.ToPointer("198.51.100.2"),
                },
            },
            Networks:    []string{
                "string",
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
      "mtu": 0,
      "networks": [
        "string"
      ],
      "persist_mac": false,
      "poe_disabled": false,
      "port_auth": "string",
      "port_network": "string",
      "rejected_network": null,
      "speed": "string",
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
      "port_auth": "string",
      "port_network": "string",
      "rejected_network": null,
      "speed": "string",
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

