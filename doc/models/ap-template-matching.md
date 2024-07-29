
# Ap Template Matching

## Structure

`ApTemplateMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `Rules` | [`[]models.ApTemplateMatchingRule`](../../doc/models/ap-template-matching-rule.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

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
            }
          },
          "enable_mac_auth": false,
          "forwarding": "site_mxedge",
          "mac_auth_protocol": "eap-md5"
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
          "forwarding": "site_mxedge",
          "mac_auth_protocol": "eap-md5"
        }
      }
    },
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
            }
          },
          "enable_mac_auth": false,
          "forwarding": "site_mxedge",
          "mac_auth_protocol": "eap-md5"
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
          "forwarding": "site_mxedge",
          "mac_auth_protocol": "eap-md5"
        }
      }
    }
  ]
}
```

