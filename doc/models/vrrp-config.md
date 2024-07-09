
# Vrrp Config

Junos VRRP config

## Structure

`VrrpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `Groups` | [`map[string]models.VrrpConfigGroup`](../../doc/models/vrrp-config-group.md) | Optional | Property key is the VRRP name |

## Example (as JSON)

```json
{
  "enabled": false,
  "groups": {
    "key0": {
      "priority": 102
    },
    "key1": {
      "priority": 102
    },
    "key2": {
      "priority": 102
    }
  }
}
```

