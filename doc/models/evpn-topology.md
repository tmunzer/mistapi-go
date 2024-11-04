
# Evpn Topology

## Structure

`EvpnTopology`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Config` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch settings |
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
      "mac": "5c5b35000003",
      "pod": 1,
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
      "config": {
        "acl_policies": [
          {
            "actions": [
              {
                "action": "allow",
                "dst_tag": "dst_tag0"
              }
            ],
            "name": "name2",
            "src_tags": [
              "src_tags1",
              "src_tags0"
            ]
          },
          {
            "actions": [
              {
                "action": "allow",
                "dst_tag": "dst_tag0"
              }
            ],
            "name": "name2",
            "src_tags": [
              "src_tags1",
              "src_tags0"
            ]
          }
        ],
        "acl_tags": {
          "key0": {
            "gbp_tag": 14,
            "macs": [
              "macs1"
            ],
            "network": "network2",
            "radius_group": "radius_group8",
            "specs": [
              {
                "port_range": "port_range8",
                "protocol": "protocol6"
              }
            ],
            "type": "dynamic_gbp"
          }
        },
        "additional_config_cmds": [
          "additional_config_cmds0",
          "additional_config_cmds9"
        ],
        "created_time": 40.26,
        "deviceprofile_id": "00001f46-0000-0000-0000-000000000000",
        "type": "type4"
      },
      "downlinks": [
        "downlinks6",
        "downlinks7"
      ],
      "esilaglinks": [
        "esilaglinks4",
        "esilaglinks3"
      ],
      "evpn_id": 192
    }
  ],
  "config": {
    "ap_affinity_threshold": 164,
    "cli_banner": "cli_banner4",
    "cli_idle_timeout": 76,
    "config_revert_timer": 122,
    "dhcp_option_fqdn": false
  },
  "overwrite": false,
  "pod_names": {
    "key0": "pod_names2"
  }
}
```

