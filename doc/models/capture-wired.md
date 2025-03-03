
# Capture Wired

Initiate a Wired Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureWired`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureWiredFormatEnum`](../../doc/models/capture-wired-format-enum.md) | Optional | pcap format. enum: `pcap`, `stream`<br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `512`<br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br>**Default**: `1024`<br>**Constraints**: `>= 0`, `<= 10000` |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump expression |
| `Type` | `string` | Required, Constant | enum: `wired`<br>**Value**: `"wired"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 300,
  "format": "pcap",
  "max_pkt_len": 128,
  "num_packets": 1000,
  "tcpdump_expression": "tcp port 80",
  "type": "wired",
  "ap_mac": "ap_mac8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

