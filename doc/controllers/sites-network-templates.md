# Sites Network Templates

```go
sitesNetworkTemplates := client.SitesNetworkTemplates()
```

## Class Name

`SitesNetworkTemplates`


# List Site Network Template Derived

Get derived Network Templates for Site

```go
ListSiteNetworkTemplateDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[models.NetworkTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | Whether resolve the site variables |

## Response Type

[`models.NetworkTemplate`](../../doc/models/network-template.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := sitesNetworkTemplates.ListSiteNetworkTemplateDerived(ctx, siteId, nil)
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

