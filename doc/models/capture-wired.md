
# Capture Wired

Initiate a Wired Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureWired`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | AP MAC address used to target the wired packet capture |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureWiredFormatEnum`](../../doc/models/capture-wired-format-enum.md) | Optional | pcap format. enum: `pcap`, `stream`<br><br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | Maximum bytes captured from each packet, or null to use the default<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump filter expression applied to packet capture traffic |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for wired captures. enum: `wired`<br><br>**Value**: `"wired"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureWired := models.CaptureWired{
        ApMac:                models.NewOptional(models.ToPointer("ap_mac2")),
        Duration:             models.NewOptional(models.ToPointer(300)),
        Format:               models.ToPointer(models.CaptureWiredFormatEnum_PCAP),
        MaxPktLen:            models.NewOptional(models.ToPointer(128)),
        NumPackets:           models.NewOptional(models.ToPointer(1000)),
        TcpdumpExpression:    models.NewOptional(models.ToPointer("tcp port 80")),
        Type:                 "wired",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

