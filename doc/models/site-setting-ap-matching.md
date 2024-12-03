
# Site Setting Ap Matching

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingApMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `Rules` | [`[]models.SiteSettingApMatchingRule`](../../doc/models/site-setting-ap-matching-rule.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
            },
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          },
          "enable_mac_auth": false,
          "flow_control": false,
          "forwarding": "site_mxedge",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        "key1": {
          "disabled": false,
          "dynamic_vlan": {
            "default_vlan_id": 34,
            "enabled": false,
            "type": "type6",
            "vlans": {
              "key0": "vlans1"
            },
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          },
          "enable_mac_auth": false,
          "flow_control": false,
          "forwarding": "site_mxedge",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      }
    }
  ],
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

