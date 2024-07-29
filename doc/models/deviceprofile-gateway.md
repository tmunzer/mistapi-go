
# Deviceprofile Gateway

Gateway Template is applied to a site for gateway(s) in a site.

## Structure

`DeviceprofileGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config<br><br>**Note**: no check is done |
| `BgpConfig` | [`map[string]models.BgpConfig`](../../doc/models/bgp-config.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `DhcpdConfig` | [`*models.DhcpdConfig`](../../doc/models/dhcpd-config.md) | Optional | - |
| `DnsOverride` | `*bool` | Optional | **Default**: `false` |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `ExtraRoutes` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `GatewayMatching` | [`*models.GatewayMatching`](../../doc/models/gateway-matching.md) | Optional | Gateway matching |
| `Id` | `*uuid.UUID` | Optional | - |
| `IdpProfiles` | [`map[string]models.IdpProfile`](../../doc/models/idp-profile.md) | Optional | Property key is the profile name |
| `IpConfigs` | [`map[string]models.GatewayIpConfigProperty`](../../doc/models/gateway-ip-config-property.md) | Optional | Property key is the network name |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | - |
| `Networks` | [`[]models.Network`](../../doc/models/network.md) | Optional | - |
| `NtpOverride` | `*bool` | Optional | **Default**: `false` |
| `NtpServers` | `[]string` | Optional | list of NTP servers specific to this device. By default, those in Site Settings will be used |
| `OobIpConfig` | [`*models.GatewayOobIpConfig`](../../doc/models/gateway-oob-ip-config.md) | Optional | out-of-band (vme/em0/fxp0) IP config |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathPreferences` | [`map[string]models.GatewayPathPreferences`](../../doc/models/gateway-path-preferences.md) | Optional | Property key is the path name |
| `PortConfig` | [`map[string]models.GatewayPortConfig`](../../doc/models/gateway-port-config.md) | Optional | Property key is the port(s) name or range (e.g. "ge-0/0/0-10") |
| `RouterId` | `*string` | Optional | auto assigned if not set |
| `RoutingPolicies` | [`map[string]models.RoutingPolicy`](../../doc/models/routing-policy.md) | Optional | Property key is the routing policy name |
| `ServicePolicies` | [`[]models.ServicePolicy`](../../doc/models/service-policy.md) | Optional | - |
| `TunnelConfigs` | [`map[string]models.TunnelConfigs`](../../doc/models/tunnel-configs.md) | Optional | Property key is the tunnel name |
| `TunnelProviderOptions` | [`*models.TunnelProviderOptions`](../../doc/models/tunnel-provider-options.md) | Optional | - |
| `Type` | [`*models.DeviceTypeGatewayEnum`](../../doc/models/device-type-gateway-enum.md) | Optional | Device Type. enum: `gateway` |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`map[string]models.GatewayVrfInstance`](../../doc/models/gateway-vrf-instance.md) | Optional | Property key is the network name |

## Example (as JSON)

```json
{
  "dnsOverride": false,
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "name": "gw_template",
  "ntpOverride": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "router_id": "10.2.1.10",
  "vrf_instances": {
    "CORP_VRF": {
      "networks": [
        "CORP_NET",
        "MGMT_NET"
      ]
    }
  },
  "additional_config_cmds": [
    "additional_config_cmds6"
  ],
  "bgp_config": {
    "key0": {
      "auth_key": "auth_key8",
      "bfd_minimum_interval": 212,
      "bfd_multiplier": 90,
      "communities": [
        {
          "id": "id8",
          "local_preference": 56,
          "vpn_name": "vpn_name0"
        }
      ],
      "disable_bfd": false
    }
  },
  "created_time": 141.8,
  "dhcpd_config": {
    "enabled": false
  }
}
```

