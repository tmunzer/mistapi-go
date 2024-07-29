
# Dhcpd Config Option

## Structure

`DhcpdConfigOption`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`*models.DhcpdConfigOptionTypeEnum`](../../doc/models/dhcpd-config-option-type-enum.md) | Optional | enum: `boolean`, `hex`, `int16`, `int32`, `ip`, `string`, `uint16`, `uint32` |
| `Value` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "type": "int16",
  "value": "value6"
}
```

