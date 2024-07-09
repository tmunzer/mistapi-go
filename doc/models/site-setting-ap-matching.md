
# Site Setting Ap Matching

## Structure

`SiteSettingApMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `Rules` | [`[]models.SiteSettingApMatchingRule`](../../doc/models/site-setting-ap-matching-rule.md) | Optional | - |

## Example (as JSON)

```json
{
  "rules": [
    {
      "eth1,eth2": {
        "port_vlan_id": 1,
        "vlan_ids": [
          1,
          10,
          50
        ]
      },
      "match_model": "match_model0",
      "name": "name8",
      "port_config": {
        "key0": {
          "disabled": false,
          "dynamic_vlan": {
            "default_vlan_id": 34,
            "enabled": false,
            "type": "type6",
            "vlans": {
              "key0": "vlans1"
            }
          },
          "enable_mac_auth": false,
          "forwarding": "mxtunnel",
          "mac_auth_protocol": "pap"
        },
        "key1": {
          "disabled": false,
          "dynamic_vlan": {
            "default_vlan_id": 34,
            "enabled": false,
            "type": "type6",
            "vlans": {
              "key0": "vlans1"
            }
          },
          "enable_mac_auth": false,
          "forwarding": "mxtunnel",
          "mac_auth_protocol": "pap"
        }
      }
    }
  ],
  "enabled": false
}
```

