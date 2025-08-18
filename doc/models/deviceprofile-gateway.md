
# Deviceprofile Gateway

Gateway Template is applied to a site for gateway(s) in a site.

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceprofileGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `BgpConfig` | [`map[string]models.BgpConfig`](../../doc/models/bgp-config.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DhcpdConfig` | [`*models.DhcpdConfig`](../../doc/models/dhcpd-config.md) | Optional | - |
| `DnsOverride` | `*bool` | Optional | **Default**: `false` |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `ExtraRoutes` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8"), the destination Network name or a variable (e.g. "{{myvar}}") |
| `ExtraRoutes6` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64"), the destination Network name or a variable (e.g. "{{myvar}}") |
| `GatewayMatching` | [`*models.GatewayMatching`](../../doc/models/gateway-matching.md) | Optional | Gateway matching |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IdpProfiles` | [`map[string]models.IdpProfile`](../../doc/models/idp-profile.md) | Optional | Property key is the profile name |
| `IpConfigs` | [`map[string]models.GatewayIpConfigProperty`](../../doc/models/gateway-ip-config-property.md) | Optional | Property key is the network name |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `Networks` | [`[]models.Network`](../../doc/models/network.md) | Optional | - |
| `NtpOverride` | `*bool` | Optional | **Default**: `false` |
| `NtpServers` | `[]string` | Optional | List of NTP servers specific to this device. By default, those in Site Settings will be used |
| `OobIpConfig` | [`*models.GatewayOobIpConfig`](../../doc/models/gateway-oob-ip-config.md) | Optional | Out-of-band (vme/em0/fxp0) IP config |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathPreferences` | [`map[string]models.GatewayPathPreferences`](../../doc/models/gateway-path-preferences.md) | Optional | Property key is the path name |
| `PortConfig` | [`map[string]models.GatewayPortConfig`](../../doc/models/gateway-port-config.md) | Optional | Property key is the port(s) name or range (e.g. "ge-0/0/0-10") |
| `RouterId` | `*string` | Optional | Auto assigned if not set |
| `RoutingPolicies` | [`map[string]models.RoutingPolicy`](../../doc/models/routing-policy.md) | Optional | Property key is the routing policy name |
| `ServicePolicies` | [`[]models.ServicePolicy`](../../doc/models/service-policy.md) | Optional | - |
| `TunnelConfigs` | [`map[string]models.TunnelConfig`](../../doc/models/tunnel-config.md) | Optional | Property key is the tunnel name |
| `TunnelProviderOptions` | [`*models.TunnelProviderOptions`](../../doc/models/tunnel-provider-options.md) | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `gateway`<br><br>**Value**: `"gateway"` |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`map[string]models.GatewayVrfInstance`](../../doc/models/gateway-vrf-instance.md) | Optional | Property key is the network name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dnsOverride": false,
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "gw_template",
  "ntpOverride": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "router_id": "10.2.1.10",
  "type": "gateway",
  "vrf_instances": {
    "CORP_VRF": {
      "networks": [
        "CORP_NET",
        "MGMT_NET"
      ]
    }
  },
  "additional_config_cmds": [
    "additional_config_cmds2",
    "additional_config_cmds1"
  ],
  "bgp_config": {
    "key0": {
      "auth_key": "auth_key8",
      "bfd_minimum_interval": 212,
      "bfd_multiplier": 90,
      "disable_bfd": false,
      "export": "export6",
      "via": "vpn",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "auth_key": "auth_key8",
      "bfd_minimum_interval": 212,
      "bfd_multiplier": 90,
      "disable_bfd": false,
      "export": "export6",
      "via": "vpn",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "auth_key": "auth_key8",
      "bfd_minimum_interval": 212,
      "bfd_multiplier": 90,
      "disable_bfd": false,
      "export": "export6",
      "via": "vpn",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "created_time": 177.78,
  "dhcpd_config": {
    "enabled": false,
    "exampleAdditionalProperty": {
      "dns_servers": [
        "dns_servers2",
        "dns_servers3",
        "dns_servers4"
      ],
      "dns_suffix": [
        "dns_suffix1",
        "dns_suffix0",
        "dns_suffix9"
      ],
      "fixed_bindings": {
        "key0": null
      },
      "gateway": "gateway8",
      "ip6_end": "ip6_end8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

