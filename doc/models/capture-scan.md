
# Capture Scan

Initiate a Scan Radio Packet Capture

## Structure

`CaptureScan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | AP MAC address used to filter the scan radio packet capture |
| `Aps` | [`map[string]models.CaptureScanAps`](../../doc/models/capture-scan-aps.md) | Optional | Dictionary key is AP mac and value is a dictionary which contains key "band", "bandwidth", "channel" and "tcpdump_expression". In case keys are missed we will take parent value if parent values are not set we will use default value |
| `Band` | [`models.Optional[models.CaptureScanBandEnum]`](../../doc/models/capture-scan-band-enum.md) | Optional | Only Single value allowed, default value gets applied when user provides wrong values. enum: `24`, `5`, `6`<br><br>**Default**: `"5"` |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `0`(disabled, response only), `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `*int` | Optional | Specify the channel value where scan PCAP has to be started, default value gets applied when user provides wrong values<br><br>**Default**: `1` |
| `ClientMac` | `models.Optional[string]` | Optional | Client MAC address used to filter the scan radio packet capture |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureScanFormatEnum`](../../doc/models/capture-scan-format-enum.md) | Optional | Output format for the scan radio packet capture. enum: `pcap`, `stream`<br><br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | Maximum bytes captured from each packet, or null to use the default<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for scan radio captures. enum: `scan`<br><br>**Value**: `"scan"` |
| `Width` | `*string` | Optional | Specify the bandwidth value with respect to the channel. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureScan := models.CaptureScan{
        ApMac:                models.NewOptional(models.ToPointer("ap_mac4")),
        Aps:                  map[string]models.CaptureScanAps{
            "key0": models.CaptureScanAps{
                Band:                 models.ToPointer(models.CaptureScanApsBandEnum_ENUM24),
                Channel:              models.ToPointer("channel0"),
                TcpdumpExpression:    models.ToPointer("tcpdump_expression4"),
                Width:                models.ToPointer("width2"),
            },
            "key1": models.CaptureScanAps{
                Band:                 models.ToPointer(models.CaptureScanApsBandEnum_ENUM24),
                Channel:              models.ToPointer("channel0"),
                TcpdumpExpression:    models.ToPointer("tcpdump_expression4"),
                Width:                models.ToPointer("width2"),
            },
            "key2": models.CaptureScanAps{
                Band:                 models.ToPointer(models.CaptureScanApsBandEnum_ENUM24),
                Channel:              models.ToPointer("channel0"),
                TcpdumpExpression:    models.ToPointer("tcpdump_expression4"),
                Width:                models.ToPointer("width2"),
            },
        },
        Band:                 models.NewOptional(models.ToPointer(models.CaptureScanBandEnum_ENUM24)),
        Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
        Channel:              models.ToPointer(1),
        ClientMac:            models.NewOptional(models.ToPointer("38f9d3972ff1")),
        Duration:             models.NewOptional(models.ToPointer(300)),
        Format:               models.ToPointer(models.CaptureScanFormatEnum_STREAM),
        MaxPktLen:            models.NewOptional(models.ToPointer(128)),
        NumPackets:           models.NewOptional(models.ToPointer(1000)),
        TcpdumpExpression:    models.ToPointer("tcp port 80"),
        Type:                 "scan",
    }

}
```

