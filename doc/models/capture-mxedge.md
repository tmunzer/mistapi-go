
# Capture Mxedge

Initiate a Wireless Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `Format` | [`*models.CaptureMxedgeFormatEnum`](../../doc/models/capture-mxedge-format-enum.md) | Optional | pcap format. enum:<br><br>* `stream`: to Mist cloud<br>* `tzsp`: tream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)<br>**Default**: `"stream"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `128`<br>**Constraints**: `<= 2048` |
| `Mxedges` | [`map[string]models.CaptureMxedgeMxedges`](../../doc/models/capture-mxedge-mxedges.md) | Optional | - |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024`<br>**Constraints**: `<= 10000` |
| `Type` | `string` | Required, Constant | enum: `mxedge`<br>**Value**: `"mxedge"` |
| `TzspHost` | `*string` | Optional | Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets. |
| `TzspPort` | `*int` | Optional | if `format`==`tzsp`. Port on remote host for receiving the captured packets<br>**Default**: `37008`<br>**Constraints**: `>= 1`, `<= 65535` |
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

