
# Response Auto Orientation

Auto orientation start response

## Structure

`ResponseAutoOrientation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Devices` | [`map[string]models.ResponseAutoOrientationDevice`](../../doc/models/response-auto-orientation-device.md) | Optional | Contains the validation status of each device. The property key is the device MAC address. |
| `EstimatedRuntime` | `*int` | Optional | Estimated runtime for the process in seconds |
| `Reason` | `*string` | Optional | Provides the reason for the status. |
| `Started` | `*bool` | Optional | Indicates whether the auto orient process has started. |
| `Valid` | `*bool` | Optional | Indicates whether the auto orient request is valid. |
| `WifiInterrupting` | `*bool` | Optional | Indicates whether the auto orient process will interrupt WiFi traffic. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoOrientation := models.ResponseAutoOrientation{
        Devices:              map[string]models.ResponseAutoOrientationDevice{
            "key0": models.ResponseAutoOrientationDevice{
                Reason:               models.ToPointer("reason0"),
                Valid:                models.ToPointer(false),
            },
        },
        EstimatedRuntime:     models.ToPointer(20),
        Reason:               models.ToPointer("reason4"),
        Started:              models.ToPointer(false),
        Valid:                models.ToPointer(false),
    }

}
```

