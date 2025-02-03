
# Protect Re Custom

Custom acls

*This model accepts additional fields of type interface{}.*

## Structure

`ProtectReCustom`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | Matched dst port, "0" means any<br>**Default**: `"0"` |
| `Protocol` | [`*models.ProtectReCustomProtocolEnum`](../../doc/models/protect-re-custom-protocol-enum.md) | Optional | enum: `any`, `icmp`, `tcp`, `udp`<br>**Default**: `"any"` |
| `Subnets` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port_range": "80,1035-1040",
  "protocol": "any",
  "subnets": [
    "subnets5",
    "subnets4"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

