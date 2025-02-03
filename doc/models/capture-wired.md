
# Capture Wired

Initiate a Wired Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureWired`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Duration` | `*int` | Optional | Duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `Format` | [`*models.CaptureWiredFormatEnum`](../../doc/models/capture-wired-format-enum.md) | Optional | pcap format. enum: `pcap`, `stream`<br>**Default**: `"pcap"` |
| `MaxPktLen` | `*int` | Optional | Max_len of each packet to capture<br>**Default**: `128`<br>**Constraints**: `<= 2048` |
| `NumPackets` | `*int` | Optional | Number of packets to capture, 0 for unlimited<br>**Default**: `1024` |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump expression |
| `Type` | `string` | Required, Constant | enum: `wired`<br>**Value**: `"wired"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 600,
  "format": "pcap",
  "max_pkt_len": 68,
  "num_packets": 100,
  "tcpdump_expression": "tcp port 80",
  "type": "wired",
  "ap_mac": "ap_mac8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

