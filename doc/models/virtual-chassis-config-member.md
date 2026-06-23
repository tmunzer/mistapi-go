
# Virtual Chassis Config Member

Virtual Chassis member configuration

## Structure

`VirtualChassisConfigMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Locating` | `*bool` | Optional, Read-only | Whether this Virtual Chassis member is currently in locate mode |
| `Mac` | `string` | Required | Member MAC address; for FPC0 this matches the device MAC address |
| `MemberId` | `*int` | Optional | Member ID used for a pre-provisioned Virtual Chassis |
| `VcPorts` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `VcRole` | [`models.VirtualChassisConfigMemberVcRoleEnum`](../../doc/models/virtual-chassis-config-member-vc-role-enum.md) | Required | enum: `backup`, `linecard`, `master` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    virtualChassisConfigMember := models.VirtualChassisConfigMember{
        Locating:             models.ToPointer(false),
        Mac:                  "mac8",
        MemberId:             models.ToPointer(180),
        VcPorts:              []string{
            "vc_ports8",
            "vc_ports9",
            "vc_ports0",
        },
        VcRole:               models.VirtualChassisConfigMemberVcRoleEnum_LINECARD,
    }

}
```

