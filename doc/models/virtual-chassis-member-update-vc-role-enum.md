
# Virtual Chassis Member Update Vc Role Enum

Required if `op`==`add` or `op`==`preprovision`. enum: `backup`, `linecard`, `master`

## Enumeration

`VirtualChassisMemberUpdateVcRoleEnum`

## Fields

| Name |
|  --- |
| `backup` |
| `linecard` |
| `master` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    virtualChassisMemberUpdateVcRole := models.VirtualChassisMemberUpdateVcRoleEnum_MASTER

}
```

