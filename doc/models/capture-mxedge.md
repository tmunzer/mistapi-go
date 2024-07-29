
# Capture Mxedge

Initiate a Wireless Packet Capture

## Structure

`CaptureMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `Format` | [`*models.CaptureMxedgeFormatEnum`](../../doc/models/capture-mxedge-format-enum.md) | Optional | pcap format. enum: `stream`<br>**Default**: `"stream"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `128`<br>**Constraints**: `<= 2048` |
| `Mxedges` | [`map[string]models.CaptureMxedgeMxedges`](../../doc/models/capture-mxedge-mxedges.md) | Optional | - |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024`<br>**Constraints**: `<= 10000` |
| `Type` | `string` | Required, Constant | enum: `mxedge`<br>**Default**: `"mxedge"` |

## Example (as JSON)

```json
{
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 68,
  "num_packets": 100,
  "type": "mxedge",
  "mxedges": {
    "key0": {
      "interfaces": {
        "key0": {
          "tcpdump_expression": "tcpdump_expression4"
        }
      }
    },
    "key1": {
      "interfaces": {
        "key0": {
          "tcpdump_expression": "tcpdump_expression4"
        }
      }
    }
  }
}
```

