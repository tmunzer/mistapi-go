
# Switch Port Usage Dynamic Rule

## Structure

`SwitchPortUsageDynamicRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Equals` | `*string` | Optional | - |
| `EqualsAny` | `[]string` | Optional | use `equals_any` to match any item in a list |
| `Expression` | `*string` | Optional | "[0:3]":"abcdef" -> "abc"<br>"split(.)[1]": "a.b.c" -> "b"<br>"split(-)[1][0:3]: "a1234-b5678-c90" -> "b56" |
| `Src` | [`models.SwitchPortUsageDynamicRuleSrcEnum`](../../doc/models/switch-port-usage-dynamic-rule-src-enum.md) | Required | enum: `link_peermac`, `lldp_chassis_id`, `lldp_hardware_revision`, `lldp_manufacturer_name`, `lldp_oui`, `lldp_serial_number`, `lldp_system_name`, `radius_dynamicfilter`, `radius_usermac`, `radius_username` |
| `Usage` | `*string` | Optional | `port_usage` name |

## Example (as JSON)

```json
{
  "equals": "equals4",
  "equals_any": [
    "equals_any3",
    "equals_any2",
    "equals_any1"
  ],
  "expression": "expression2",
  "src": "link_peermac",
  "usage": "usage2"
}
```

