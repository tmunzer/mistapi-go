
# Ap Template Matching

*This model accepts additional fields of type interface{}.*

## Structure

`ApTemplateMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `Rules` | [`[]models.ApTemplateMatchingRule`](../../doc/models/ap-template-matching-rule.md) | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "rules": [
    {
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
          "forwarding": "site_mxedge",
          "mac_auth_preferred": false,
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
          "forwarding": "site_mxedge",
          "mac_auth_preferred": false,
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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

