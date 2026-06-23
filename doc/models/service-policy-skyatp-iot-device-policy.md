
# Service Policy Skyatp Iot Device Policy

Sky ATP IoT device policy settings

## Structure

`ServicePolicySkyatpIotDevicePolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether Sky ATP IoT device policy inspection is enabled |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicySkyatpIotDevicePolicy := models.ServicePolicySkyatpIotDevicePolicy{
        Enabled:              models.ToPointer(false),
    }

}
```

