
# Response Autoplacement

Auto-placement start response returned after scheduling the run

## Structure

`ResponseAutoplacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Devices` | [`map[string]models.ResponseAutoplacementDevice`](../../doc/models/response-autoplacement-device.md) | Optional, Read-only | Property key is the AP MAC address. Contains the validation status of each device. |
| `EstimatedRuntime` | `*int` | Optional, Read-only | Estimated runtime for the process in seconds. |
| `Reason` | `*string` | Optional, Read-only | Provides the reason for the status. |
| `Started` | `*bool` | Optional, Read-only | Indicates whether the autoplacement process has started. |
| `Valid` | `*bool` | Optional, Read-only | Indicates whether the autoplacement request is valid. |
| `WifiInterrupting` | `*bool` | Optional | Indicates whether the auto placement process will interrupt WiFi traffic. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoplacement := models.ResponseAutoplacement{
    }

}
```

