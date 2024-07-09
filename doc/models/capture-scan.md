
# Capture Scan

Initiate a Scan Radio Packet Capture

## Structure

`CaptureScan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | filter by ap_mac |
| `Aps` | [`map[string]models.CaptureScanAps`](../../doc/models/capture-scan-aps.md) | Optional | dictionary key is AP mac and value is a dictionary which contains key “band”, “bandwidth”, “channel” and “tcpdump_expression”. In case keys are missed we will take parent value if parent values are not set we will use default value |
| `Band` | [`models.Optional[models.CaptureScanBandEnum]`](../../doc/models/capture-scan-band-enum.md) | Optional | Only Single value allowed, default value gets applied when user provides wrong values<br>**Default**: `"5"` |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band * `80` is only applicable for band_5 and band_6 * `160` is only for band_6 |
| `Channel` | `*int` | Optional | specify the channel value where scan PCAP has to be started, default value gets applied when user provides wrong values<br>**Default**: `1` |
| `ClientMac` | `models.Optional[string]` | Optional | filter by client mac |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `>= 0`, `<= 86400` |
| `Format` | [`*models.CaptureScanFormatEnum`](../../doc/models/capture-scan-format-enum.md) | Optional | **Default**: `"pcap"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `512`<br>**Constraints**: `>= 0`, `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024` |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `Type` | `string` | Required, Constant | **Default**: `"scan"` |
| `Width` | `*string` | Optional | specify the bandwidth value with respect to the channel. |

## Example (as JSON)

```json
{
  "band": "24",
  "bandwidth": 20,
  "channel": 1,
  "client_mac": "38f9d3972ff1",
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 68,
  "num_packets": 100,
  "tcpdump_expression": "tcp port 80",
  "type": "scan",
  "ap_mac": "ap_mac2",
  "aps": {
    "key0": {
      "band": "24",
      "channel": "channel0",
      "tcpdump_expression": "tcpdump_expression4",
      "width": "width2"
    },
    "key1": {
      "band": "24",
      "channel": "channel0",
      "tcpdump_expression": "tcpdump_expression4",
      "width": "width2"
    },
    "key2": {
      "band": "24",
      "channel": "channel0",
      "tcpdump_expression": "tcpdump_expression4",
      "width": "width2"
    }
  }
}
```
