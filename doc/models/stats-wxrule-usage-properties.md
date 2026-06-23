
# Stats Wxrule Usage Properties

Usage counters for one WxLAN rule destination tag

## Structure

`StatsWxruleUsageProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumFlows` | `*int` | Optional | Number of flows counted for this WxLAN rule and destination tag |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsWxruleUsageProperties := models.StatsWxruleUsageProperties{
        NumFlows:             models.ToPointer(24),
    }

}
```

