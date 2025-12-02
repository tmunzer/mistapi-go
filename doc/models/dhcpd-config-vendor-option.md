
# Dhcpd Config Vendor Option

## Structure

`DhcpdConfigVendorOption`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`*models.DhcpdConfigVendorOptionTypeEnum`](../../doc/models/dhcpd-config-vendor-option-type-enum.md) | Optional | enum: `boolean`, `hex`, `int16`, `int32`, `ip`, `string`, `uint16`, `uint32` |
| `Value` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "type": "uint16",
  "value": "value8"
}
```

