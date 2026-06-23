
# Switch Virtual Chassis Member

Virtual Chassis member identified by MAC address and role

## Structure

`SwitchVirtualChassisMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Virtual Chassis member MAC address; for FPC0 this matches the device ID MAC |
| `MemberId` | `int` | Required | Virtual Chassis member identifier |
| `VcRole` | [`models.SwitchVirtualChassisMemberVcRoleEnum`](../../doc/models/switch-virtual-chassis-member-vc-role-enum.md) | Required | Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config. enum: `backup`, `linecard`, `master` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchVirtualChassisMember := models.SwitchVirtualChassisMember{
        Mac:                  "aff827549235",
        MemberId:             234,
        VcRole:               models.SwitchVirtualChassisMemberVcRoleEnum_BACKUP,
    }

}
```

