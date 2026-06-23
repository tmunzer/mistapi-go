
# Capture Radiotapwired

Initiate a Radiotap Packet Capture and Wired Packet Capture

## Structure

`CaptureRadiotapwired`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | AP MAC address used to target the combined radiotap and wired packet capture |
| `Band` | [`*models.CaptureRadiotapwiredBandEnum`](../../doc/models/capture-radiotapwired-band-enum.md) | Optional | only used for radiotap. enum: `24`, `24,5,6`, `5`, `6`<br><br>**Default**: `"24"` |
| `ClientMac` | `models.Optional[string]` | Optional | Client MAC address used to filter the combined packet capture |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureRadiotapwiredFormatEnum`](../../doc/models/capture-radiotapwired-format-enum.md) | Optional | Output format for the combined packet capture. enum: `pcap`, `stream`<br><br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | Maximum bytes captured from each packet, or null to use the default<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `RadiotapTcpdumpExpression` | `*string` | Optional | tcpdump expression for radiotap interface (802.11 + radio headers) |
| `Ssid` | `models.Optional[string]` | Optional | Wireless network SSID used to filter the radiotap portion |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump filter expression applied to packet capture traffic |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for combined radiotap and wired captures. enum: `radiotap,wired`<br><br>**Value**: `"radiotap,wired"` |
| `WiredTcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump filter expression applied to packet capture traffic |
| `WirelessTcpdumpExpression` | `*string` | Optional | tcpdump expression for radiotap interface (802.11) |
| `WlanId` | `models.Optional[string]` | Optional | WLAN id associated with the respective ssid. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureRadiotapwired := models.CaptureRadiotapwired{
        ApMac:                     models.NewOptional(models.ToPointer("ap_mac6")),
        Band:                      models.ToPointer(models.CaptureRadiotapwiredBandEnum_ENUM24),
        ClientMac:                 models.NewOptional(models.ToPointer("38f9d3972ff1")),
        Duration:                  models.NewOptional(models.ToPointer(300)),
        Format:                    models.ToPointer(models.CaptureRadiotapwiredFormatEnum_STREAM),
        MaxPktLen:                 models.NewOptional(models.ToPointer(128)),
        NumPackets:                models.NewOptional(models.ToPointer(1000)),
        RadiotapTcpdumpExpression: models.ToPointer("type"),
        Ssid:                      models.NewOptional(models.ToPointer("test")),
        TcpdumpExpression:         models.NewOptional(models.ToPointer("tcp port 80")),
        Type:                      "radiotap,wired",
        WiredTcpdumpExpression:    models.NewOptional(models.ToPointer("tcp port 80")),
        WlanId:                    models.NewOptional(models.ToPointer("fac8e973-feb9-421a-b381-aabbc4b61f5a")),
    }

}
```

