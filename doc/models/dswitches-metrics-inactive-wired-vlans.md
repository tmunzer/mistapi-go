
# Dswitches Metrics Inactive Wired Vlans

Inactive wired VLAN metric for APs connected to discovered switches

## Structure

`DswitchesMetricsInactiveWiredVlans`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | `interface{}` | Required | Metric-specific details for inactive wired VLAN findings |
| `Score` | `float64` | Required | Compliance score for the inactive wired VLAN metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dswitchesMetricsInactiveWiredVlans := models.DswitchesMetricsInactiveWiredVlans{
        Details:              interface{}("[key1, val1][key2, val2]"),
        Score:                float64(244.56),
    }

}
```

