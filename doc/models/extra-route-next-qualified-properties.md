
# Extra Route Next Qualified Properties

Qualified next-hop attributes for an IPv4 static route

## Structure

`ExtraRouteNextQualifiedProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metric` | `models.Optional[int]` | Optional | Route metric for this qualified IPv4 next hop |
| `Preference` | `models.Optional[int]` | Optional | Route preference for this qualified IPv4 next hop |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    extraRouteNextQualifiedProperties := models.ExtraRouteNextQualifiedProperties{
        Metric:               models.NewOptional(models.ToPointer(246)),
        Preference:           models.NewOptional(models.ToPointer(186)),
    }

}
```

