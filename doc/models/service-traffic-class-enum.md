
# Service Traffic Class Enum

when `traffic_type`==`custom`. enum: `best_effort`, `high`, `low`, `medium`

## Enumeration

`ServiceTrafficClassEnum`

## Fields

| Name |
|  --- |
| `best_effort` |
| `high` |
| `low` |
| `medium` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    serviceTrafficClass := models.ServiceTrafficClassEnum_LOW

}
```

