
# Switch Matching Rule Ip Config

In-Band Management interface configuration

## Structure

`SwitchMatchingRuleIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Network` | `*string` | Optional | VLAN Name for the management interface |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |

## Example (as JSON)

```json
{
  "type": "static",
  "network": "network8"
}
```

