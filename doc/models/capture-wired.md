
# Capture Wired

Initiate a Wired Packet Capture

## Structure

`CaptureWired`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `Format` | [`*models.CaptureWiredFormatEnum`](../../doc/models/capture-wired-format-enum.md) | Optional | pcap format. enum: `pcap`, `stream`<br>**Default**: `"pcap"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `128`<br>**Constraints**: `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024` |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump expression |
| `Type` | `string` | Required, Constant | enum: `wired`<br>**Default**: `"wired"` |

## Example (as JSON)

```json
{
  "duration": 600,
  "format": "pcap",
  "max_pkt_len": 68,
  "num_packets": 100,
  "tcpdump_expression": "tcp port 80",
  "type": "wired",
  "ap_mac": "ap_mac8"
}
```

