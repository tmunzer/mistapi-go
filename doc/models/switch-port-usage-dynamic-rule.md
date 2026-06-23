
# Switch Port Usage Dynamic Rule

Dynamic port usage rule evaluated against LLDP, RADIUS, or peer MAC attributes

## Structure

`SwitchPortUsageDynamicRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | Optional description of the rule |
| `Equals` | `*string` | Optional | Exact value that the selected source attribute must match |
| `EqualsAny` | `[]string` | Optional | Use `equals_any` to match any item in a list |
| `Expression` | `*string` | Optional | "[0:3]":"abcdef" -> "abc"<br>"split(.)[1]": "a.b.c" -> "b"<br>"split(-)[1][0:3]: "a1234-b5678-c90" -> "b56" |
| `Src` | [`models.SwitchPortUsageDynamicRuleSrcEnum`](../../doc/models/switch-port-usage-dynamic-rule-src-enum.md) | Required | enum: `link_peermac`, `lldp_chassis_id`, `lldp_hardware_revision`, `lldp_manufacturer_name`, `lldp_oui`, `lldp_serial_number`, `lldp_system_description`, `lldp_system_name`, `radius_dynamicfilter`, `radius_usermac`, `radius_username` |
| `Usage` | `*string` | Optional | Port usage name to apply when this dynamic rule matches |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortUsageDynamicRule := models.SwitchPortUsageDynamicRule{
        Description:          models.ToPointer("description4"),
        Equals:               models.ToPointer("equals4"),
        EqualsAny:            []string{
            "equals_any3",
            "equals_any2",
        },
        Expression:           models.ToPointer("expression8"),
        Src:                  models.SwitchPortUsageDynamicRuleSrcEnum_LLDPHARDWAREREVISION,
        Usage:                models.ToPointer("usage2"),
    }

}
```

