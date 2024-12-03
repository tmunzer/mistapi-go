
# Capture Switch

Initiate a Switch (Junos) Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `>= 0`, `<= 86400` |
| `Format` | [`*models.CaptureSwitchFormatEnum`](../../doc/models/capture-switch-format-enum.md) | Optional | enum: `stream`<br>**Default**: `"stream"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `512`<br>**Constraints**: `>= 128`, `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024` |
| `Ports` | [`map[string]models.CaptureSwitchPortsTcpdumpExpression`](../../doc/models/capture-switch-ports-tcpdump-expression.md) | Optional | Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request |
| `Switches` | [`map[string]models.CaptureSwitchSwitches`](../../doc/models/capture-switch-switches.md) | Optional | Property key is the switch mac |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `Type` | `string` | Required, Constant | enum: `switch`<br>**Default**: `"switch"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 1500,
  "num_packets": 100,
  "tcpdump_expression": "port 443",
  "type": "switch",
  "ports": {
    "key0": {
      "tcpdump_expression": "tcpdump_expression0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "tcpdump_expression": "tcpdump_expression0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "tcpdump_expression": "tcpdump_expression0",
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

