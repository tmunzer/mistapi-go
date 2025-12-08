
# Site Setting Ap Matching Rule

## Structure

`SiteSettingApMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MatchModel` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | Property key is the interface(s) (e.g. "eth1,eth2") |

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
        }
      },
      "enable_mac_auth": false,
      "forwarding": "site_mxedge",
      "mac_auth_preferred": false
    }
  }
}
```

