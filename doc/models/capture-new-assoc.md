
# Capture New Assoc

Initiate a packet Capture for New Wireless Client Associations

## Structure

`CaptureNewAssoc`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | AP MAC address used to target new association capture |
| `ClientMac` | `*string` | Optional | Client MAC address, required if `type`==`client`; optional otherwise |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `IncludesMcast` | `*bool` | Optional | Whether to include multicast traffic in the packet capture<br><br>**Default**: `false` |
| `MaxPktLen` | `models.Optional[int]` | Optional | Maximum bytes captured from each packet, or null to use the default<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ssid` | `*string` | Optional | Optional SSID filter for new association capture |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for new association captures. enum: `new_assoc`<br><br>**Value**: `"new_assoc"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureNewAssoc := models.CaptureNewAssoc{
        ApMac:                models.ToPointer("a83a79a947ee"),
        ClientMac:            models.ToPointer("60a10a773412"),
        Duration:             models.NewOptional(models.ToPointer(300)),
        IncludesMcast:        models.ToPointer(false),
        MaxPktLen:            models.NewOptional(models.ToPointer(128)),
        NumPackets:           models.NewOptional(models.ToPointer(1000)),
        Type:                 "new_assoc",
    }

}
```

