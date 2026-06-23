
# Extra Route 6 Next Qualified Properties

Qualified next-hop attributes for an IPv6 static route

## Structure

`ExtraRoute6NextQualifiedProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metric` | `models.Optional[int]` | Optional | Route metric for this qualified IPv6 next hop |
| `Preference` | `models.Optional[int]` | Optional | Route preference for this qualified IPv6 next hop |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    extraRoute6NextQualifiedProperties := models.ExtraRoute6NextQualifiedProperties{
        Metric:               models.NewOptional(models.ToPointer(114)),
        Preference:           models.NewOptional(models.ToPointer(34)),
    }

}
```

