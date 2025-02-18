
# Evpn Topology Switch

*This model accepts additional fields of type interface{}.*

## Structure

`EvpnTopologySwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Config` | [`*models.EvpnTopologySwitchConfig`](../../doc/models/evpn-topology-switch-config.md) | Optional | - |
| `DeviceprofileId` | `*uuid.UUID` | Optional | - |
| `DownlinkIps` | `[]string` | Optional | - |
| `Downlinks` | `[]string` | Optional | - |
| `Esilaglinks` | `[]string` | Optional | - |
| `EvpnId` | `*int` | Optional | **Constraints**: `>= 1` |
| `Mac` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Model` | `*string` | Optional | - |
| `Pod` | `*int` | Optional | Optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.<br><br>* for CLOS, to group dist / access switches into pods<br>* for ERB/CRB, to group dist / esilag-access into pods<br>**Default**: `1`<br>**Constraints**: `>= 1`, `<= 255` |
| `Pods` | `[]int` | Optional | By default, core switches are assumed to be connecting all pods.<br>if you want to limit the pods, you can specify pods. |
| `Role` | [`models.EvpnTopologySwitchRoleEnum`](../../doc/models/evpn-topology-switch-role-enum.md) | Required | use `role`==`none` to remove a switch from the topology. enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none`<br>**Constraints**: *Minimum Length*: `1` |
| `RouterId` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SuggestedDownlinks` | `[]string` | Optional | - |
| `SuggestedEsilaglinks` | `[]string` | Optional | - |
| `SuggestedUplinks` | `[]string` | Optional | - |
| `Uplinks` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "deviceprofile_id": "6a1deab1-96df-4fa2-8455-d5253f943d06",
  "downlinks": [
    "5c5b35000005",
    "5c5b35000006"
  ],
  "esilaglinks": [
    "5c5b35000005",
    "5c5b35000006"
  ],
  "mac": "5c5b35000003",
  "model": "QFX10002-36Q",
  "pod": 1,
  "role": "access",
  "router_id": "172.16.254.4",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "suggested_downlinks": [
    "5c5b35000005",
    "5c5b35000006"
  ],
  "suggested_esilaglinks": [
    "5c5b35000005",
    "5c5b35000006"
  ],
  "suggested_uplinks": [
    "5c5b35000005",
    "5c5b35000006"
  ],
  "uplinks": [
    "5c5b35000005",
    "5c5b35000006"
  ],
  "config": {
    "dhcpd_config": {
      "enabled": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "networks": {
      "key0": {
        "gateway": "gateway8",
        "gateway6": "gateway64",
        "isolation": false,
        "isolation_vlan_id": "isolation_vlan_id8",
        "subnet": "subnet6",
        "vlan_id": "String7",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key1": {
        "gateway": "gateway8",
        "gateway6": "gateway64",
        "isolation": false,
        "isolation_vlan_id": "isolation_vlan_id8",
        "subnet": "subnet6",
        "vlan_id": "String7",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    },
    "other_ip_configs": {
      "key0": {
        "evpn_anycast": false,
        "ip": "ip4",
        "ip6": "ip60",
        "netmask": "netmask0",
        "netmask6": "netmask60",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key1": {
        "evpn_anycast": false,
        "ip": "ip4",
        "ip6": "ip60",
        "netmask": "netmask0",
        "netmask6": "netmask60",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key2": {
        "evpn_anycast": false,
        "ip": "ip4",
        "ip6": "ip60",
        "netmask": "netmask0",
        "netmask6": "netmask60",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    },
    "port_config": {
      "key0": {
        "ae_disable_lacp": false,
        "ae_idx": 230,
        "ae_lacp_slow": false,
        "aggregated": false,
        "critical": false,
        "usage": "usage6",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key1": {
        "ae_disable_lacp": false,
        "ae_idx": 230,
        "ae_lacp_slow": false,
        "aggregated": false,
        "critical": false,
        "usage": "usage6",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    },
    "port_usages": {
      "key0": {
        "all_networks": false,
        "allow_dhcpd": false,
        "allow_multiple_supplicants": false,
        "bypass_auth_when_server_down": false,
        "bypass_auth_when_server_down_for_unkown_client": false,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key1": {
        "all_networks": false,
        "allow_dhcpd": false,
        "allow_multiple_supplicants": false,
        "bypass_auth_when_server_down": false,
        "bypass_auth_when_server_down_for_unkown_client": false,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key2": {
        "all_networks": false,
        "allow_dhcpd": false,
        "allow_multiple_supplicants": false,
        "bypass_auth_when_server_down": false,
        "bypass_auth_when_server_down_for_unkown_client": false,
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
  },
  "downlink_ips": [
    "downlink_ips0",
    "downlink_ips1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

