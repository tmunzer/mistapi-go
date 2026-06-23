
# Admin Privilege Scope Enum

enum: `msp`, `org`, `orggroup`, `site`, `sitegroup`

## Enumeration

`AdminPrivilegeScopeEnum`

## Fields

| Name |
|  --- |
| `msp` |
| `org` |
| `orggroup` |
| `site` |
| `sitegroup` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    adminPrivilegeScope := models.AdminPrivilegeScopeEnum_ORG

}
```

