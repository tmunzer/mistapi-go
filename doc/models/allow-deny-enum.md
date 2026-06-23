
# Allow Deny Enum

Policy action value that either allows or denies matching traffic. enum: `allow`, `deny`

## Enumeration

`AllowDenyEnum`

## Fields

| Name |
|  --- |
| `allow` |
| `deny` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    allowDeny := models.AllowDenyEnum_ALLOW

}
```

