
# Webhook Minis Reachability Event

Marvis Minis reachability synthetic test result

## Structure

`WebhookMinisReachabilityEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvgLatency` | `*float64` | Optional | Average latency in milliseconds |
| `DeviceMac` | `*string` | Optional | MAC address of the device performing the test |
| `LossPercentage` | `*float64` | Optional | Percentage of packets lost during the reachability test |
| `MaxLatency` | `*float64` | Optional | Maximum latency in milliseconds |
| `MinLatency` | `*float64` | Optional | Minimum latency in milliseconds |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `ProbeName` | `*string` | Optional | Name of the probe |
| `ProbeTarget` | `*string` | Optional | Target host or IP for the probe |
| `ProbeType` | `*string` | Optional | Probe category used for the Minis reachability test |
| `Protocol` | `*string` | Optional | Network protocol used for the reachability test |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TestType` | `*string` | Optional | Type of test performed |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Vlan` | `*int` | Optional | Network VLAN ID used for the reachability test |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookMinisReachabilityEvent := models.WebhookMinisReachabilityEvent{
        AvgLatency:           models.ToPointer(float64(224.12)),
        DeviceMac:            models.ToPointer("7cb68d8f0440"),
        LossPercentage:       models.ToPointer(float64(129.96)),
        MaxLatency:           models.ToPointer(float64(99.4)),
        MinLatency:           models.ToPointer(float64(10.5)),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        ProbeName:            models.ToPointer("google ping"),
        ProbeTarget:          models.ToPointer("google.com"),
        ProbeType:            models.ToPointer("reachability"),
        Protocol:             models.ToPointer("icmp"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        TestType:             models.ToPointer("ping"),
        Vlan:                 models.ToPointer(12),
    }

}
```

