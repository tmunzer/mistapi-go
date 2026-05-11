
# Capture Mxedge

Initiate a Mist Edge Packet Capture

## Structure

`CaptureMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Duration of the capture, in seconds. Default is 600, minimum is 60 and maximum is 10800 (3h)<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 10800` |
| `Format` | [`*models.CaptureMxedgeFormatEnum`](../../doc/models/capture-mxedge-format-enum.md) | Optional | PCAP format. enum:<br><br>* `stream`: to Mist cloud<br>* `tzsp`: stream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)<br><br>**Default**: `"stream"` |
| `MaxPktLen` | `*int` | Optional | Max_len of each packet to capture. Default is 512, minimum is 64 and maximum is 2048<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `Mxedges` | [`map[string]models.CaptureMxedgeMxedges`](../../doc/models/capture-mxedge-mxedges.md) | Optional | Dict of Mist Edges to capture on, property key is the Mist Edge ID. Property value is a dict of interfaces to capture for the given mxedge (e.g. port1, kni0, lacp0, ipsec, drop, oobm) |
| `NumPackets` | `*int` | Optional | Number of packets to capture. Default is 1024, maximum is 10000, minimum 1, or 0 for unlimited (local/remote streaming only)<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, applicable across all interfaces if specified at top level. An interface-specific value (under the `interfaces` dict) overrides this top-level value. |
| `Type` | `string` | Required, Constant | enum: `mxedge`<br><br>**Value**: `"mxedge"` |
| `TzspHost` | `*string` | Optional | Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets |
| `TzspPort` | `*int` | Optional | Optional port on remote host for receiving the captured packets. Default is 37008 (TZSP)<br><br>**Default**: `37008`<br><br>**Constraints**: `>= 1`, `<= 65535` |

## Example (as JSON)

```json
{
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 512,
  "num_packets": 100,
  "type": "mxedge",
  "tzsp_host": "192.168.1.2",
  "tzsp_port": 37008,
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

