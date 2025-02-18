
# Site Setting Ap Matching Rule

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingApMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MatchModel` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | Property key is the interface(s) (e.g. "eth1,eth2") |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "match_model": "AP12",
  "name": "AP12",
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
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

