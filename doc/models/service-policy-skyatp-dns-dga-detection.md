
# Service Policy Skyatp Dns Dga Detection

Sky ATP DNS DGA detection settings

## Structure

`ServicePolicySkyatpDnsDgaDetection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether Sky ATP DNS DGA detection is enabled |
| `Profile` | [`*models.ServicePolicySkyatpDnsDgaDetectionProfileEnum`](../../doc/models/service-policy-skyatp-dns-dga-detection-profile-enum.md) | Optional | enum: `default`, `standard`, `strict` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicySkyatpDnsDgaDetection := models.ServicePolicySkyatpDnsDgaDetection{
        Enabled:              models.ToPointer(false),
        Profile:              models.ToPointer(models.ServicePolicySkyatpDnsDgaDetectionProfileEnum_STANDARD),
    }

}
```

