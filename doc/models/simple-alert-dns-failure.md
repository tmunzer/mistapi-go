
# Simple Alert Dns Failure

Thresholds for DNS failure heuristic alerts

## Structure

`SimpleAlertDnsFailure`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientCount` | `*int` | Optional | Number of distinct clients that must encounter DNS failures before alerting<br><br>**Default**: `20` |
| `Duration` | `*int` | Optional | Time window in minutes for evaluating DNS failures<br><br>**Default**: `10`<br><br>**Constraints**: `>= 5`, `<= 60` |
| `IncidentCount` | `*int` | Optional | Number of DNS failure incidents required within the duration window<br><br>**Default**: `30` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    simpleAlertDnsFailure := models.SimpleAlertDnsFailure{
        ClientCount:          models.ToPointer(20),
        Duration:             models.ToPointer(10),
        IncidentCount:        models.ToPointer(30),
    }

}
```

