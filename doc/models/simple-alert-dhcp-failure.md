
# Simple Alert Dhcp Failure

Thresholds for DHCP failure heuristic alerts

## Structure

`SimpleAlertDhcpFailure`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientCount` | `*int` | Optional | Number of distinct clients that must encounter DHCP failures before alerting<br><br>**Default**: `10` |
| `Duration` | `*int` | Optional | Time window in minutes for evaluating DHCP failures<br><br>**Default**: `10`<br><br>**Constraints**: `>= 5`, `<= 60` |
| `IncidentCount` | `*int` | Optional | Number of DHCP failure incidents required within the duration window<br><br>**Default**: `20` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    simpleAlertDhcpFailure := models.SimpleAlertDhcpFailure{
        ClientCount:          models.ToPointer(10),
        Duration:             models.ToPointer(10),
        IncidentCount:        models.ToPointer(20),
    }

}
```

