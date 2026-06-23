
# Service Policy Skyatp Dns Tunnel Detection

Sky ATP DNS tunneling detection settings

## Structure

`ServicePolicySkyatpDnsTunnelDetection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether Sky ATP DNS tunneling detection is enabled |
| `Profile` | [`*models.ServicePolicySkyatpDnsTunnelDetectionProfileEnum`](../../doc/models/service-policy-skyatp-dns-tunnel-detection-profile-enum.md) | Optional | enum: `default`, `standard`, `strict` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicySkyatpDnsTunnelDetection := models.ServicePolicySkyatpDnsTunnelDetection{
        Enabled:              models.ToPointer(false),
        Profile:              models.ToPointer(models.ServicePolicySkyatpDnsTunnelDetectionProfileEnum_STRICT),
    }

}
```

