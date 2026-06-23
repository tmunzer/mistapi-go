
# Response Mxtunnels Preempt Aps

Result of preempting APs onto preferred MxTunnel peers

## Structure

`ResponseMxtunnelsPreemptAps`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreemptedAps` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseMxtunnelsPreemptAps := models.ResponseMxtunnelsPreemptAps{
        PreemptedAps:         []string{
            "preempted_aps0",
        },
    }

}
```

