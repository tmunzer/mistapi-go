
# Capture Wireless

Initiate a Wireless Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureWireless`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | AP MAC address used to target the wireless packet capture |
| `Band` | [`*models.CaptureWirelessBandEnum`](../../doc/models/capture-wireless-band-enum.md) | Optional | enum: `24`, `5`, `6`<br><br>**Default**: `"24"` |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureWirelessFormatEnum`](../../doc/models/capture-wireless-format-enum.md) | Optional | pcap format. enum: `pcap`, `stream`<br><br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | Maximum bytes captured from each packet, or null to use the default<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ssid` | `*string` | Optional | Wireless network SSID used to filter the packet capture |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for wireless captures. enum: `wireless`<br><br>**Value**: `"wireless"` |
| `WlanId` | `*uuid.UUID` | Optional | WLAN identifier used to filter the wireless packet capture |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureWireless := models.CaptureWireless{
        ApMac:                models.NewOptional(models.ToPointer("ap_mac0")),
        Band:                 models.ToPointer(models.CaptureWirelessBandEnum_ENUM24),
        Duration:             models.NewOptional(models.ToPointer(300)),
        Format:               models.ToPointer(models.CaptureWirelessFormatEnum_PCAP),
        MaxPktLen:            models.NewOptional(models.ToPointer(128)),
        NumPackets:           models.NewOptional(models.ToPointer(1000)),
        Type:                 "wireless",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

