
# Device Gateway

device gateway

## Structure

`DeviceGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config<br><br>**Note**: no check is done |
| `BgpConfig` | [`map[string]models.BgpConfig`](../../doc/models/bgp-config.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `DeviceprofileId` | `*uuid.UUID` | Optional | - |
| `DhcpdConfig` | [`*models.DhcpdConfig`](../../doc/models/dhcpd-config.md) | Optional | - |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `ExtraRoutes` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `IdpProfiles` | [`map[string]models.IdpProfile`](../../doc/models/idp-profile.md) | Optional | Property key is the profile name |
| `Image1Url` | `models.Optional[string]` | Optional | - |
| `Image2Url` | `models.Optional[string]` | Optional | - |
| `Image3Url` | `models.Optional[string]` | Optional | - |
| `IpConfigs` | [`map[string]models.GatewayIpConfigProperty`](../../doc/models/gateway-ip-config-property.md) | Optional | Property key is the network name |
| `Mac` | `*string` | Optional | device MAC address |
| `Managed` | `*bool` | Optional | - |
| `MapId` | `*uuid.UUID` | Optional | map where the device belongs to |
| `Model` | `*string` | Optional | device Model |
| `ModifiedTime` | `*float64` | Optional | - |
| `MspId` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Networks` | [`[]models.Network`](../../doc/models/network.md) | Optional | - |
| `Notes` | `*string` | Optional | - |
| `NtpServers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `OobIpConfig` | [`*models.GatewayOobIpConfig`](../../doc/models/gateway-oob-ip-config.md) | Optional | out-of-band (vme/em0/fxp0) IP config |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathPreferences` | [`map[string]models.GatewayPathPreferences`](../../doc/models/gateway-path-preferences.md) | Optional | Property key is the path name |
| `PortConfig` | [`map[string]models.GatewayPortConfig`](../../doc/models/gateway-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortMirroring` | [`*models.GatewayPortMirroring`](../../doc/models/gateway-port-mirroring.md) | Optional | - |
| `RouterId` | `*string` | Optional | auto assigned if not set |
| `RoutingPolicies` | [`map[string]models.RoutingPolicy`](../../doc/models/routing-policy.md) | Optional | Property key is the routing policy name |
| `Serial` | `*string` | Optional | device Serial |
| `ServicePolicies` | [`[]models.ServicePolicy`](../../doc/models/service-policy.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TunnelConfigs` | [`map[string]models.TunnelConfigs`](../../doc/models/tunnel-configs.md) | Optional | Property key is the tunnel name |
| `TunnelProviderOptions` | [`*models.TunnelProviderOptions`](../../doc/models/tunnel-provider-options.md) | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `gateway`<br>**Default**: `"gateway"` |
| `Vars` | `map[string]string` | Optional | a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`map[string]models.GatewayVrfInstance`](../../doc/models/gateway-vrf-instance.md) | Optional | Property key is the network name |
| `X` | `*float64` | Optional | x in pixel |
| `Y` | `*float64` | Optional | y in pixel |

## Example (as JSON)

```json
{
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "router_id": "10.2.1.10",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "gateway",
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "vrf_instances": {
    "CORP_VRF": {
      "networks": [
        "CORP_NET",
        "MGMT_NET"
      ]
    }
  },
  "x": 53.5,
  "y": 173.1,
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
  "created_time": 176.02,
  "deviceprofile_id": "0000062e-0000-0000-0000-000000000000",
  "dhcpd_config": {
    "enabled": false
  }
}
```

