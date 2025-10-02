
# Capture Scan

Initiate a Scan Radio Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureScan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | Filter by ap_mac |
| `Aps` | [`map[string]models.CaptureScanAps`](../../doc/models/capture-scan-aps.md) | Optional | Dictionary key is AP mac and value is a dictionary which contains key "band", "bandwidth", "channel" and "tcpdump_expression". In case keys are missed we will take parent value if parent values are not set we will use default value |
| `Band` | [`models.Optional[models.CaptureScanBandEnum]`](../../doc/models/capture-scan-band-enum.md) | Optional | Only Single value allowed, default value gets applied when user provides wrong values. enum: `24`, `5`, `6`<br><br>**Default**: `"5"` |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `0`(disabled, response only), `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `*int` | Optional | Specify the channel value where scan PCAP has to be started, default value gets applied when user provides wrong values<br><br>**Default**: `1` |
| `ClientMac` | `models.Optional[string]` | Optional | Filter by client mac |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureScanFormatEnum`](../../doc/models/capture-scan-format-enum.md) | Optional | enum: `pcap`, `stream`<br><br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `Type` | `string` | Required, Constant | enum: `scan`<br><br>**Value**: `"scan"` |
| `Width` | `*string` | Optional | Specify the bandwidth value with respect to the channel. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "24",
  "bandwidth": 20,
  "channel": 1,
  "client_mac": "38f9d3972ff1",
  "duration": 300,
  "format": "stream",
  "max_pkt_len": 128,
  "num_packets": 1000,
  "tcpdump_expression": "tcp port 80",
  "type": "scan",
  "ap_mac": "ap_mac4",
  "aps": {
    "key0": {
      "band": "24",
      "channel": "channel0",
      "tcpdump_expression": "tcpdump_expression4",
      "width": "width2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "band": "24",
      "channel": "channel0",
      "tcpdump_expression": "tcpdump_expression4",
      "width": "width2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "band": "24",
      "channel": "channel0",
      "tcpdump_expression": "tcpdump_expression4",
      "width": "width2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

