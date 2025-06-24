
# Tunterm Dhcpd Config Property

*This model accepts additional fields of type interface{}.*

## Structure

`TuntermDhcpdConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Servers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Type` | [`*models.TuntermDhcpdTypeEnum`](../../doc/models/tunterm-dhcpd-type-enum.md) | Optional | enum: `relay`<br><br>**Default**: `"relay"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "relay",
  "servers": [
    "servers7",
    "servers8",
    "servers9"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

