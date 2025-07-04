
# Capture Mxedge

Initiate a Wireless Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `<= 86400` |
| `Format` | [`*models.CaptureMxedgeFormatEnum`](../../doc/models/capture-mxedge-format-enum.md) | Optional | PCAP format. enum:<br><br>* `stream`: to Mist cloud<br>* `tzsp`: stream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)<br><br>**Default**: `"stream"` |
| `MaxPktLen` | `*int` | Optional | Max_len of each packet to capture<br><br>**Default**: `128`<br><br>**Constraints**: `<= 2048` |
| `Mxedges` | [`map[string]models.CaptureMxedgeMxedges`](../../doc/models/capture-mxedge-mxedges.md) | Optional | - |
| `NumPackets` | `*int` | Optional | Number of packets to capture, 0 for unlimited<br><br>**Default**: `1024`<br><br>**Constraints**: `<= 10000` |
| `Type` | `string` | Required, Constant | enum: `mxedge`<br><br>**Value**: `"mxedge"` |
| `TzspHost` | `*string` | Optional | Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets. |
| `TzspPort` | `*int` | Optional | If `format`==`tzsp`. Port on remote host for receiving the captured packets<br><br>**Default**: `37008`<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 68,
  "num_packets": 100,
  "type": "mxedge",
  "tzsp_host": "192.168.1.2",
  "tzsp_port": 37008,
  "mxedges": {
    "key0": {
      "interfaces": {
        "key0": {
          "tcpdump_expression": "tcpdump_expression4",
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
    },
    "key1": {
      "interfaces": {
        "key0": {
          "tcpdump_expression": "tcpdump_expression4",
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
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

