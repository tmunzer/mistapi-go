
# Response Auto Map Assignment

Auto map assignment start response

## Structure

`ResponseAutoMapAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Devices` | [`map[string]models.ResponseAutoMapAssignmentDevice`](../../doc/models/response-auto-map-assignment-device.md) | Optional | Contains the validation status of each device. The property key is the device MAC address. |
| `EstimatedRuntime` | `*int` | Optional | Estimated runtime for the process in seconds |
| `Reason` | `*string` | Optional | Provides the reason for the status |
| `Started` | `*bool` | Optional | Indicates whether the auto map assignment process has started |
| `Valid` | `*bool` | Optional | Indicates whether the auto map assignment request is valid |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoMapAssignment := models.ResponseAutoMapAssignment{
        Devices:              map[string]models.ResponseAutoMapAssignmentDevice{
            "key0": models.ResponseAutoMapAssignmentDevice{
                Reason:               models.ToPointer("reason0"),
                Valid:                models.ToPointer(false),
            },
            "key1": models.ResponseAutoMapAssignmentDevice{
                Reason:               models.ToPointer("reason0"),
                Valid:                models.ToPointer(false),
            },
            "key2": models.ResponseAutoMapAssignmentDevice{
                Reason:               models.ToPointer("reason0"),
                Valid:                models.ToPointer(false),
            },
        },
        EstimatedRuntime:     models.ToPointer(44),
        Reason:               models.ToPointer("reason4"),
        Started:              models.ToPointer(false),
        Valid:                models.ToPointer(false),
    }

}
```

