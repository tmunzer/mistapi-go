
# Webhook Client Latency Event

Site-level client latency metrics for authentication, DHCP, and DNS operations

## Structure

`WebhookClientLatencyEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvgAuth` | `*float64` | Optional | Average authentication latency observed during the reporting interval |
| `AvgDhcp` | `*float64` | Optional | Average DHCP latency observed during the reporting interval |
| `AvgDns` | `*float64` | Optional | Average DNS latency observed during the reporting interval |
| `MaxAuth` | `*float64` | Optional | Maximum authentication latency observed during the reporting interval |
| `MaxDhcp` | `*float64` | Optional | Maximum DHCP latency observed during the reporting interval |
| `MaxDns` | `*float64` | Optional | Maximum DNS latency observed during the reporting interval |
| `MinAuth` | `*float64` | Optional | Minimum authentication latency observed during the reporting interval |
| `MinDhcp` | `*float64` | Optional | Minimum DHCP latency observed during the reporting interval |
| `MinDns` | `*float64` | Optional | Minimum DNS latency observed during the reporting interval |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookClientLatencyEvent := models.WebhookClientLatencyEvent{
        AvgAuth:              models.ToPointer(float64(0.17170219)),
        AvgDhcp:              models.ToPointer(float64(0.017828934)),
        AvgDns:               models.ToPointer(float64(0.024532124)),
        MaxAuth:              models.ToPointer(float64(0.18170219)),
        MaxDhcp:              models.ToPointer(float64(0.027828934)),
        MaxDns:               models.ToPointer(float64(0.022532124)),
        MinAuth:              models.ToPointer(float64(0.16050219)),
        MinDhcp:              models.ToPointer(float64(0.015828934)),
        MinDns:               models.ToPointer(float64(0.029532124)),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

