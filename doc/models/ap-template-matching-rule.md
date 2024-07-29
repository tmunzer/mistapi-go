
# Ap Template Matching Rule

## Structure

`ApTemplateMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MatchModel` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Name` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | Property key is the interface(s) name (e.g. "eth1,eth2") |

## Example (as JSON)

```json
{
  "match_model": "match_model2",
  "name": "name0",
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
    },
    "key2": {
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
```

