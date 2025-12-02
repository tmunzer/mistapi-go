
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
      "match_model": "string",
      "name": "string",
      "port_config": {
        "eth1,eth2": {
          "disabled": true,
          "dynamic_vlan": {
            "default_vlan_id": 999,
            "enabled": true,
            "type": "type6",
            "vlans": {
              "key0": "vlans1"
            }
          },
          "port_vlan_id": 1,
          "vlan_id": 9,
          "vlan_ids": [
            1,
            10,
            50
          ],
          "enable_mac_auth": false,
          "forwarding": "site_mxedge",
          "mac_auth_preferred": false
        }
      }
    }
  ],
  "enabled": false
}
```

