
# Admin Privilege View Enum

UI view name allowed by a custom admin role. enum: `lobby_admin`, `location`, `marketing`, `mxedge_admin`, `reporting`, `security`, `super_observer`, `switch_admin`

## Enumeration

`AdminPrivilegeViewEnum`

## Fields

| Name |
|  --- |
| `lobby_admin` |
| `location` |
| `marketing` |
| `mxedge_admin` |
| `reporting` |
| `security` |
| `super_observer` |
| `switch_admin` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    adminPrivilegeView := models.AdminPrivilegeViewEnum_MARKETING

}
```

