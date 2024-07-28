
# Evpn Topology Switch

## Structure

`EvpnTopologySwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Config` | [`*models.DeviceSwitch`](../../doc/models/device-switch.md) | Optional | Switch Configuration.<br>You can configure `port_usages` and `networks` settings at the device level, but most of the time it's better use the Site Setting to achieve better consistency and be able to re-use the same settings across switches entries defined here will "replace" those defined in Site Setting/Network Template |
| `DeviceprofileId` | `*uuid.UUID` | Optional | - |
| `Downlinks` | `[]string` | Optional | - |
| `Esilaglinks` | `[]string` | Optional | - |
| `EvpnId` | `*int` | Optional | **Constraints**: `>= 1` |
| `Mac` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Model` | `*string` | Optional | - |
| `Pod` | `*int` | Optional | optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.<br><br>* for CLOS, to group dist / access switches into pods<br>* for ERB/CRB, to group dist / esilag-access into pods<br>**Default**: `1`<br>**Constraints**: `>= 1`, `<= 255` |
| `Pods` | `[]int` | Optional | by default, core switches are assumed to be connecting all pods.<br>if you want to limit the pods, you can specify pods. |
| `Role` | [`*models.EvpnTopologySwitchRoleEnum`](../../doc/models/evpn-topology-switch-role-enum.md) | Optional | use `role`==`none` to remove a switch from the topology<br>**Constraints**: *Minimum Length*: `1` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SuggestedDownlinks` | `[]string` | Optional | - |
| `SuggestedEsilaglinks` | `[]string` | Optional | - |
| `SuggestedUplinks` | `[]string` | Optional | - |
| `Uplinks` | `[]string` | Optional | if not specified in the request, suggested ones will be used |

## Example (as JSON)

```json
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
        "type": "any"
      }
    },
    "additional_config_cmds": [
      "additional_config_cmds0",
      "additional_config_cmds9"
    ],
    "created_time": 40.26,
    "deviceprofile_id": "00001f46-0000-0000-0000-000000000000"
  },
  "downlinks": [
    "downlinks6",
    "downlinks5"
  ],
  "esilaglinks": [
    "esilaglinks4",
    "esilaglinks5"
  ],
  "evpn_id": 94
}
```

