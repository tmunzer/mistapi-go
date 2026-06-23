
# Service Policy Skyatp Http Inspection

Sky ATP HTTP inspection settings

## Structure

`ServicePolicySkyatpHttpInspection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether Sky ATP HTTP inspection is enabled |
| `Profile` | [`*models.ServicePolicySkyatpHttpInspectionProfileEnum`](../../doc/models/service-policy-skyatp-http-inspection-profile-enum.md) | Optional | Sky ATP HTTP inspection profile to apply. enum: `standard`, `strict` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicySkyatpHttpInspection := models.ServicePolicySkyatpHttpInspection{
        Enabled:              models.ToPointer(false),
        Profile:              models.ToPointer(models.ServicePolicySkyatpHttpInspectionProfileEnum_STANDARD),
    }

}
```

