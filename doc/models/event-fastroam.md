
# Event Fastroam

Fast-roaming event observed for a wireless client

## Structure

`EventFastroam`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | Destination AP MAC address for the roam |
| `ClientMac` | `string` | Required | Roaming client MAC address for the event |
| `Fromap` | `string` | Required | Source AP MAC address reported for the roam |
| `Latency` | `float64` | Required | Roaming latency measured for the client event, in seconds |
| `Ssid` | `string` | Required | Wireless network SSID involved in the roam |
| `Subtype` | `*string` | Optional | Detailed roaming event subtype |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | [`*models.EventFastroamTypeEnum`](../../doc/models/event-fastroam-type-enum.md) | Optional | enum: `fail`, `none`, `pingpong`, `poor`, `slow`, `success` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    eventFastroam := models.EventFastroam{
        ApMac:                "ap_mac8",
        ClientMac:            "client_mac4",
        Fromap:               "fromap4",
        Latency:              float64(95.14),
        Ssid:                 "ssid4",
        Subtype:              models.ToPointer("subtype8"),
        Timestamp:            float64(157.64),
        Type:                 models.ToPointer(models.EventFastroamTypeEnum_SLOW),
    }

}
```

