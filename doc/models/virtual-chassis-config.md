
# Virtual Chassis Config

Virtual Chassis creation or configuration request

*This model accepts additional fields of type interface{}.*

## Structure

`VirtualChassisConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Locating` | `*bool` | Optional, Read-only | Whether the Virtual Chassis is currently in locate mode |
| `Members` | [`[]models.VirtualChassisConfigMember`](../../doc/models/virtual-chassis-config-member.md) | Optional | Virtual Chassis member configurations |
| `Preprovisioned` | `*bool` | Optional | Whether to create the Virtual Chassis in pre-provisioned mode<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    virtualChassisConfig := models.VirtualChassisConfig{
        Locating:             models.ToPointer(false),
        Members:              []models.VirtualChassisConfigMember{
            models.VirtualChassisConfigMember{
                Locating:             models.ToPointer(false),
                Mac:                  "mac2",
                MemberId:             models.ToPointer(58),
                VcPorts:              []string{
                    "vc_ports2",
                    "vc_ports3",
                    "vc_ports4",
                },
                VcRole:               models.VirtualChassisConfigMemberVcRoleEnum_MASTER,
            },
        },
        Preprovisioned:       models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

