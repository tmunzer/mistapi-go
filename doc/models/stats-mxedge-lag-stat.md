
# Stats Mxedge Lag Stat

Link aggregation group status for a Mist Edge

## Structure

`StatsMxedgeLagStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ActivePorts` | `[]string` | Optional | List of ports active on the LAG defined by the LACP |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeLagStat := models.StatsMxedgeLagStat{
        ActivePorts:          []string{
            "active_ports9",
        },
    }

}
```

