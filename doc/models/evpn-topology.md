
# Evpn Topology

*This model accepts additional fields of type interface{}.*

## Structure

`EvpnTopology`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `EvpnOptions` | [`*models.EvpnOptions`](../../doc/models/evpn-options.md) | Optional | EVPN Options |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Overwrite` | `*bool` | Optional | - |
| `PodNames` | `map[string]string` | Optional | Property key is the pod number |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SwitchConfigs` | [`map[string]models.EvpnTopologySwitchConfig`](../../doc/models/evpn-topology-switch-config.md) | Optional | Property key is the switch mac |
| `Switches` | [`[]models.EvpnTopologySwitch`](../../doc/models/evpn-topology-switch.md) | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "CC",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "switches": [
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
      "role": "esilag-access",
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
            "bypass_auth_when_server_down_for_unknown_client": false,
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
            "bypass_auth_when_server_down_for_unknown_client": false,
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
            "bypass_auth_when_server_down_for_unknown_client": false,
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
        "downlink_ips6"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "created_time": 252.34,
  "evpn_options": {
    "auto_loopback_subnet": "auto_loopback_subnet4",
    "auto_loopback_subnet6": "auto_loopback_subnet60",
    "auto_router_id_subnet": "auto_router_id_subnet8",
    "auto_router_id_subnet6": "auto_router_id_subnet60",
    "core_as_border": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "modified_time": 82.62,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

