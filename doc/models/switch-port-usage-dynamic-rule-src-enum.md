
# Switch Port Usage Dynamic Rule Src Enum

enum: `link_peermac`, `lldp_chassis_id`, `lldp_hardware_revision`, `lldp_manufacturer_name`, `lldp_oui`, `lldp_serial_number`, `lldp_system_description`, `lldp_system_name`, `radius_dynamicfilter`, `radius_usermac`, `radius_username`

## Enumeration

`SwitchPortUsageDynamicRuleSrcEnum`

## Fields

| Name |
|  --- |
| `link_peermac` |
| `lldp_chassis_id` |
| `lldp_hardware_revision` |
| `lldp_manufacturer_name` |
| `lldp_oui` |
| `lldp_serial_number` |
| `lldp_system_description` |
| `lldp_system_name` |
| `radius_dynamicfilter` |
| `radius_usermac` |
| `radius_username` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortUsageDynamicRuleSrc := models.SwitchPortUsageDynamicRuleSrcEnum_LLDPMANUFACTURERNAME

}
```

