
# Simple Alert Arp Failure

Thresholds for ARP failure heuristic alerts

## Structure

`SimpleAlertArpFailure`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientCount` | `*int` | Optional | Number of distinct clients that must encounter ARP failures before alerting<br><br>**Default**: `10` |
| `Duration` | `*int` | Optional | Time window in minutes for evaluating ARP failures<br><br>**Default**: `20`<br><br>**Constraints**: `>= 5`, `<= 60` |
| `IncidentCount` | `*int` | Optional | Number of ARP failure incidents required within the duration window<br><br>**Default**: `10` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    simpleAlertArpFailure := models.SimpleAlertArpFailure{
        ClientCount:          models.ToPointer(10),
        Duration:             models.ToPointer(20),
        IncidentCount:        models.ToPointer(10),
    }

}
```

