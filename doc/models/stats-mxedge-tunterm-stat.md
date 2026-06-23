
# Stats Mxedge Tunterm Stat

Tunnel termination monitoring status reported by a Mist Edge

## Structure

`StatsMxedgeTuntermStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MonitoringFailed` | `*bool` | Optional | Whether tunnel termination monitoring is currently failing |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeTuntermStat := models.StatsMxedgeTuntermStat{
        MonitoringFailed:     models.ToPointer(false),
    }

}
```

