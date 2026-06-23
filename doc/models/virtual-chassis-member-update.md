
# Virtual Chassis Member Update

Member update for a Virtual Chassis add, remove, or preprovision operation

## Structure

`VirtualChassisMemberUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Required if `op`==`add` or `op`==`preprovision`; MAC address of the member to add or preprovision |
| `Member` | `*int` | Optional | Required if `op`==`remove`; member ID to remove from the Virtual Chassis |
| `MemberId` | `*int` | Optional | Required if `op`==`preprovision`. Optional if `op`==`add`; target member ID for the Virtual Chassis member |
| `VcPorts` | `[]string` | Optional | Required if `op`==`add` or `op`==`preprovision` |
| `VcRole` | [`*models.VirtualChassisMemberUpdateVcRoleEnum`](../../doc/models/virtual-chassis-member-update-vc-role-enum.md) | Optional | Required if `op`==`add` or `op`==`preprovision`. enum: `backup`, `linecard`, `master` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    virtualChassisMemberUpdate := models.VirtualChassisMemberUpdate{
        Mac:                  models.ToPointer("mac2"),
        Member:               models.ToPointer(8),
        MemberId:             models.ToPointer(30),
        VcPorts:              []string{
            "vc_ports2",
            "vc_ports3",
            "vc_ports4",
        },
        VcRole:               models.ToPointer(models.VirtualChassisMemberUpdateVcRoleEnum_MASTER),
    }

}
```

