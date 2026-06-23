
# Port Stp Role Enum

Spanning Tree Protocol role for the port. enum: `""`, `alternate`, `backup`, `designated`, `disabled`, `root`, `root-prevented`

## Enumeration

`PortStpRoleEnum`

## Fields

| Name |
|  --- |
| `alternate` |
| `backup` |
| `designated` |
| `disabled` |
| `root` |
| `root-prevented` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    portStpRole := models.PortStpRoleEnum_ROOT

}
```

