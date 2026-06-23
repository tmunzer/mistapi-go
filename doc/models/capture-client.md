
# Capture Client

Initiate a Client Packet Capture

## Structure

`CaptureClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | AP MAC address used to target the client packet capture |
| `ClientMac` | `models.Optional[string]` | Optional | Client MAC address, required if `type`==`client`; optional otherwise |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `IncludesMcast` | `*bool` | Optional | Whether to include multicast traffic in the packet capture<br><br>**Default**: `false` |
| `MaxPktLen` | `models.Optional[int]` | Optional | Maximum bytes captured from each packet, or null to use the default<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ssid` | `models.Optional[string]` | Optional | Optional SSID filter for the packet capture |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for client captures. enum: `client`<br><br>**Value**: `"client"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureClient := models.CaptureClient{
        ApMac:                models.NewOptional(models.ToPointer("ap_mac8")),
        ClientMac:            models.NewOptional(models.ToPointer("60a10a773412")),
        Duration:             models.NewOptional(models.ToPointer(300)),
        IncludesMcast:        models.ToPointer(false),
        MaxPktLen:            models.NewOptional(models.ToPointer(128)),
        NumPackets:           models.NewOptional(models.ToPointer(1000)),
        Type:                 "client",
    }

}
```

