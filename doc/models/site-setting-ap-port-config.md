
# Site Setting Ap Port Config

## Structure

`SiteSettingApPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ModelSpecific` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | Property key is the AP model (e.g "AP32") |

## Example (as JSON)

```json
{
  "model_specific": {
    "AP32": {
      "eth1,eth2": {
        "port_vlan_id": 1,
        "vlan_ids": [
          1,
          10,
          50
        ]
      },
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
      "forwarding": "wxtunnel",
      "mac_auth_preferred": false
    }
  }
}
```

