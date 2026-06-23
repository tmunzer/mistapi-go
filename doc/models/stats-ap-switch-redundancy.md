
# Stats Ap Switch Redundancy

Switch redundancy status reported by an AP

## Structure

`StatsApSwitchRedundancy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumRedundantAps` | `models.Optional[int]` | Optional, Read-only | Number of redundant APs available for switch redundancy |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApSwitchRedundancy := models.StatsApSwitchRedundancy{
        NumRedundantAps:      models.NewOptional(models.ToPointer(1)),
    }

}
```

