
# Capture Switch

Initiate a Switch (Junos) Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureSwitchFormatEnum`](../../doc/models/capture-switch-format-enum.md) | Optional | enum: `stream`<br><br>**Default**: `"stream"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ports` | [`map[string]models.CaptureSwitchPortsTcpdumpExpression`](../../doc/models/capture-switch-ports-tcpdump-expression.md) | Optional | Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request |
| `Switches` | [`map[string]models.CaptureSwitchSwitches`](../../doc/models/capture-switch-switches.md) | Required | Property key is the switch mac |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `Type` | `string` | Required, Constant | enum: `switch`<br><br>**Value**: `"switch"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 300,
  "format": "stream",
  "max_pkt_len": 128,
  "num_packets": 1000,
  "switches": {
    "key0": {
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
        }
      },
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
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
        }
      },
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
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
        }
      },
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
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

