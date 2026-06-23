
# Capture Switch

Initiate a Switch (Junos) Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureSwitchFormatEnum`](../../doc/models/capture-switch-format-enum.md) | Optional | Output format for the switch packet capture. enum: `stream`<br><br>**Default**: `"stream"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | Maximum bytes captured from each packet, or null to use the default<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ports` | [`map[string]models.CaptureSwitchPortsTcpdumpExpression`](../../doc/models/capture-switch-ports-tcpdump-expression.md) | Optional | Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request |
| `Switches` | [`map[string]models.CaptureSwitchSwitches`](../../doc/models/capture-switch-switches.md) | Required | Property key is the switch MAC address |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for switch captures. enum: `switch`<br><br>**Value**: `"switch"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureSwitch := models.CaptureSwitch{
        Duration:             models.NewOptional(models.ToPointer(300)),
        Format:               models.ToPointer(models.CaptureSwitchFormatEnum_STREAM),
        MaxPktLen:            models.NewOptional(models.ToPointer(128)),
        NumPackets:           models.NewOptional(models.ToPointer(1000)),
        Ports:                map[string]models.CaptureSwitchPortsTcpdumpExpression{
            "key0": models.CaptureSwitchPortsTcpdumpExpression{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
            },
            "key1": models.CaptureSwitchPortsTcpdumpExpression{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
            },
        },
        Switches:             map[string]models.CaptureSwitchSwitches{
            "key0": models.CaptureSwitchSwitches{
                Ports:                map[string]models.CaptureSwitchPortsTcpdumpExpression{
                    "key0": models.CaptureSwitchPortsTcpdumpExpression{
                        TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
                    },
                    "key1": models.CaptureSwitchPortsTcpdumpExpression{
                        TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
                    },
                },
            },
            "key1": models.CaptureSwitchSwitches{
                Ports:                map[string]models.CaptureSwitchPortsTcpdumpExpression{
                    "key0": models.CaptureSwitchPortsTcpdumpExpression{
                        TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
                    },
                    "key1": models.CaptureSwitchPortsTcpdumpExpression{
                        TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
                    },
                },
            },
            "key2": models.CaptureSwitchSwitches{
                Ports:                map[string]models.CaptureSwitchPortsTcpdumpExpression{
                    "key0": models.CaptureSwitchPortsTcpdumpExpression{
                        TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
                    },
                    "key1": models.CaptureSwitchPortsTcpdumpExpression{
                        TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
                    },
                },
            },
        },
        TcpdumpExpression:    models.ToPointer("port 443"),
        Type:                 "switch",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

