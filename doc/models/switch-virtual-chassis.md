
# Switch Virtual Chassis

Required for preprovisioned Virtual Chassis

## Structure

`SwitchVirtualChassis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | [`[]models.SwitchVirtualChassisMember`](../../doc/models/switch-virtual-chassis-member.md) | Optional | List of Virtual Chassis members |
| `Preprovisioned` | `*bool` | Optional | To configure whether the VC is preprovisioned or nonprovisioned<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchVirtualChassis := models.SwitchVirtualChassis{
        Members:              []models.SwitchVirtualChassisMember{
            models.SwitchVirtualChassisMember{
                Mac:                  "mac2",
                MemberId:             58,
                VcRole:               models.SwitchVirtualChassisMemberVcRoleEnum_MASTER,
            },
        },
        Preprovisioned:       models.ToPointer(false),
    }

}
```

