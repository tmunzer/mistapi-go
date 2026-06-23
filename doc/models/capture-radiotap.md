
# Capture Radiotap

Initiate a Radiotap Packet Capture

## Structure

`CaptureRadiotap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | AP MAC address used to target the radiotap packet capture |
| `Band` | [`*models.CaptureRadiotapBandEnum`](../../doc/models/capture-radiotap-band-enum.md) | Optional | enum: `24`, `24,5,6`, `5`, `6`<br><br>**Default**: `"24"` |
| `ClientMac` | `*string` | Optional | Client MAC address used to filter the radiotap packet capture |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureRadiotapFormatEnum`](../../doc/models/capture-radiotap-format-enum.md) | Optional | Output format for the radiotap packet capture. enum: `pcap`, `stream`<br><br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | Maximum bytes captured from each packet, or null to use the default<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ssid` | `*string` | Optional | Wireless network SSID used to filter the radiotap packet capture |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump filter expression applied to packet capture traffic |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for radiotap captures. enum: `radiotap`<br><br>**Value**: `"radiotap"` |
| `WlanId` | `*uuid.UUID` | Optional | WLAN id associated with the respective ssid. |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    captureRadiotap := models.CaptureRadiotap{
        ApMac:                models.ToPointer("a83a79a947ee"),
        Band:                 models.ToPointer(models.CaptureRadiotapBandEnum_ENUM24),
        ClientMac:            models.ToPointer("38f9d3972ff1"),
        Duration:             models.NewOptional(models.ToPointer(300)),
        Format:               models.ToPointer(models.CaptureRadiotapFormatEnum_STREAM),
        MaxPktLen:            models.NewOptional(models.ToPointer(128)),
        NumPackets:           models.NewOptional(models.ToPointer(1000)),
        Ssid:                 models.ToPointer("test"),
        TcpdumpExpression:    models.NewOptional(models.ToPointer("tcp port 80")),
        Type:                 "radiotap",
        WlanId:               models.ToPointer(uuid.MustParse("fac8e973-feb9-421a-b381-aabbc4b61f5a")),
    }

}
```

