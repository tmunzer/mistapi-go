
# Service Failover Policy Enum

enum: `non_revertible`, `none`, `revertible`

## Enumeration

`ServiceFailoverPolicyEnum`

## Fields

| Name |
|  --- |
| `non_revertible` |
| `none` |
| `revertible` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    serviceFailoverPolicy := models.ServiceFailoverPolicyEnum_NONREVERTIBLE

}
```

