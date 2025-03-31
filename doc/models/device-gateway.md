
# Device Gateway

Device gateway

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `BgpConfig` | [`map[string]models.BgpConfig`](../../doc/models/bgp-config.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceprofileId` | `*uuid.UUID` | Optional | - |
| `DhcpdConfig` | [`*models.DhcpdConfig`](../../doc/models/dhcpd-config.md) | Optional | - |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `ExtraRoutes` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8"), the destination Network name or a variable (e.g. "{{myvar}}") |
| `ExtraRoutes6` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64"), the destination Network name or a variable (e.g. "{{myvar}}") |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IdpProfiles` | [`map[string]models.IdpProfile`](../../doc/models/idp-profile.md) | Optional | Property key is the profile name |
| `Image1Url` | `models.Optional[string]` | Optional | - |
| `Image2Url` | `models.Optional[string]` | Optional | - |
| `Image3Url` | `models.Optional[string]` | Optional | - |
| `IpConfigs` | [`map[string]models.GatewayIpConfigProperty`](../../doc/models/gateway-ip-config-property.md) | Optional | Property key is the network name |
| `Mac` | `*string` | Optional | Device MAC address |
| `Managed` | `*bool` | Optional | - |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `Model` | `*string` | Optional | Device Model |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Networks` | [`[]models.Network`](../../doc/models/network.md) | Optional | - |
| `Notes` | `*string` | Optional | - |
| `NtpServers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `OobIpConfig` | [`*models.GatewayOobIpConfig`](../../doc/models/gateway-oob-ip-config.md) | Optional | Out-of-band (vme/em0/fxp0) IP config |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathPreferences` | [`map[string]models.GatewayPathPreferences`](../../doc/models/gateway-path-preferences.md) | Optional | Property key is the path name |
| `PortConfig` | [`map[string]models.GatewayPortConfig`](../../doc/models/gateway-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortMirroring` | [`*models.GatewayPortMirroring`](../../doc/models/gateway-port-mirroring.md) | Optional | - |
| `RouterId` | `*string` | Optional | Auto assigned if not set |
| `RoutingPolicies` | [`map[string]models.RoutingPolicy`](../../doc/models/routing-policy.md) | Optional | Property key is the routing policy name |
| `Serial` | `*string` | Optional | Device Serial |
| `ServicePolicies` | [`[]models.ServicePolicy`](../../doc/models/service-policy.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TunnelConfigs` | [`map[string]models.TunnelConfig`](../../doc/models/tunnel-config.md) | Optional | Property key is the tunnel name |
| `TunnelProviderOptions` | [`*models.TunnelProviderOptions`](../../doc/models/tunnel-provider-options.md) | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `gateway`<br>**Value**: `"gateway"` |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`map[string]models.GatewayVrfInstance`](../../doc/models/gateway-vrf-instance.md) | Optional | Property key is the network name |
| `X` | `*float64` | Optional | X in pixel |
| `Y` | `*float64` | Optional | Y in pixel |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
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
      "disable_bfd": false,
      "export": "export6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "created_time": 176.02,
  "deviceprofile_id": "0000062e-0000-0000-0000-000000000000",
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
        "key0": {
          "ip": "ip0",
          "name": "name6",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      },
      "gateway": "gateway8",
      "ip_end": "ip_end4",
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

