
# Virtual Chassis Update

Virtual Chassis member update request

*This model accepts additional fields of type interface{}.*

## Structure

`VirtualChassisUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `*int` | Optional | Only if `op`==`renumber`; existing member ID to renumber |
| `Members` | [`[]models.VirtualChassisMemberUpdate`](../../doc/models/virtual-chassis-member-update.md) | Optional | Member updates for a Virtual Chassis operation |
| `NewMember` | `*int` | Optional | Only if `op`==`renumber`; new member ID to assign |
| `Op` | [`*models.VirtualChassisUpdateOpEnum`](../../doc/models/virtual-chassis-update-op-enum.md) | Optional | enum: `add`, `preprovision`, `remove`, `renumber` |
| `RemoveInventory` | `*bool` | Optional | Only if `op`==`preprovision`. When removing members from a pre-provisioned VC, set to `true` to delete the inventory records for removed members (e.g. for RMA). Members being removed must be in "not-present" state.<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    virtualChassisUpdate := models.VirtualChassisUpdate{
        Member:               models.ToPointer(184),
        Members:              []models.VirtualChassisMemberUpdate{
            models.VirtualChassisMemberUpdate{
                Mac:                  models.ToPointer("mac2"),
                Member:               models.ToPointer(176),
                MemberId:             models.ToPointer(58),
                VcPorts:              []string{
                    "vc_ports2",
                    "vc_ports3",
                    "vc_ports4",
                },
                VcRole:               models.ToPointer(models.VirtualChassisMemberUpdateVcRoleEnum_MASTER),
            },
        },
        NewMember:            models.ToPointer(118),
        Op:                   models.ToPointer(models.VirtualChassisUpdateOpEnum_REMOVE),
        RemoveInventory:      models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

