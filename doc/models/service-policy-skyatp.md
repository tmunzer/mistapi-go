
# Service Policy Skyatp

SRX Sky ATP threat inspection settings for a service policy

## Structure

`ServicePolicySkyatp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsDgaDetection` | [`*models.ServicePolicySkyatpDnsDgaDetection`](../../doc/models/service-policy-skyatp-dns-dga-detection.md) | Optional | Sky ATP DNS DGA detection settings |
| `DnsTunnelDetection` | [`*models.ServicePolicySkyatpDnsTunnelDetection`](../../doc/models/service-policy-skyatp-dns-tunnel-detection.md) | Optional | Sky ATP DNS tunneling detection settings |
| `HttpInspection` | [`*models.ServicePolicySkyatpHttpInspection`](../../doc/models/service-policy-skyatp-http-inspection.md) | Optional | Sky ATP HTTP inspection settings |
| `IotDevicePolicy` | [`*models.ServicePolicySkyatpIotDevicePolicy`](../../doc/models/service-policy-skyatp-iot-device-policy.md) | Optional | Sky ATP IoT device policy settings |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicySkyatp := models.ServicePolicySkyatp{
        DnsDgaDetection:      models.ToPointer(models.ServicePolicySkyatpDnsDgaDetection{
            Enabled:              models.ToPointer(false),
            Profile:              models.ToPointer(models.ServicePolicySkyatpDnsDgaDetectionProfileEnum_STANDARD),
        }),
        DnsTunnelDetection:   models.ToPointer(models.ServicePolicySkyatpDnsTunnelDetection{
            Enabled:              models.ToPointer(false),
            Profile:              models.ToPointer(models.ServicePolicySkyatpDnsTunnelDetectionProfileEnum_ENUMDEFAULT),
        }),
        HttpInspection:       models.ToPointer(models.ServicePolicySkyatpHttpInspection{
            Enabled:              models.ToPointer(false),
            Profile:              models.ToPointer(models.ServicePolicySkyatpHttpInspectionProfileEnum_STANDARD),
        }),
        IotDevicePolicy:      models.ToPointer(models.ServicePolicySkyatpIotDevicePolicy{
            Enabled:              models.ToPointer(false),
        }),
    }

}
```

