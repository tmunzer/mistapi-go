
# Count Pbn Distinct Enum

Fields that can be used to group PBN advisory count results. enum: `versions`, `models`, `customer_risk`, `bug_type`

## Enumeration

`CountPbnDistinctEnum`

## Fields

| Name |
|  --- |
| `versions` |
| `models` |
| `customer_risk` |
| `bug_type` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    countPbnDistinct := models.CountPbnDistinctEnum_CUSTOMERRISK

}
```

