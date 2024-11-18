
# Evpn Topology

## Structure

`EvpnTopology`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EvpnOptions` | [`*models.EvpnOptions`](../../doc/models/evpn-options.md) | Optional | EVPN Options |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Name` | `*string` | Optional | - |
| `Overwrite` | `*bool` | Optional | - |
| `PodNames` | `map[string]string` | Optional | Property key is the pod number |
| `Switches` | [`[]models.EvpnTopologySwitch`](../../doc/models/evpn-topology-switch.md) | Required | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "CC",
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
        "port_config": {
          "key0": {
            "ae_disable_lacp": false,
            "ae_idx": 230,
            "ae_lacp_slow": false,
            "aggregated": false,
            "critical": false,
            "usage": "usage6"
          },
          "key1": {
            "ae_disable_lacp": false,
            "ae_idx": 230,
            "ae_lacp_slow": false,
            "aggregated": false,
            "critical": false,
            "usage": "usage6"
          }
        }
      },
      "downlink_ips": [
        "downlink_ips6"
      ]
    }
  ],
  "evpn_options": {
    "auto_loopback_subnet": "auto_loopback_subnet4",
    "auto_loopback_subnet6": "auto_loopback_subnet60",
    "auto_router_id_subnet": "auto_router_id_subnet8",
    "auto_router_id_subnet6": "auto_router_id_subnet60",
    "core_as_border": false
  },
  "overwrite": false,
  "pod_names": {
    "key0": "pod_names2"
  }
}
```

