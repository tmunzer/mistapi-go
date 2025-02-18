
# Vrrp Config

Junos VRRP config

*This model accepts additional fields of type interface{}.*

## Structure

`VrrpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `Groups` | [`map[string]models.VrrpConfigGroup`](../../doc/models/vrrp-config-group.md) | Optional | Property key is the VRRP name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "groups": {
    "key0": {
      "priority": 102,
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

