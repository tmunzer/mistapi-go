
# Capture Switch

Initiate a Switch (Junos) Packet Capture

## Structure

`CaptureSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `>= 0`, `<= 86400` |
| `Format` | [`*models.CaptureSwitchFormatEnum`](../../doc/models/capture-switch-format-enum.md) | Optional | enum: `stream`<br>**Default**: `"stream"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `512`<br>**Constraints**: `>= 128`, `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024` |
| `Ports` | `[]string` | Optional | dict of port which uses port id as the key<br>**Constraints**: *Maximum Items*: `6` |
| `Switches` | [`map[string]models.CaptureSwitchSwitches`](../../doc/models/capture-switch-switches.md) | Optional | Property key is the switch mac |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `Type` | `string` | Required, Constant | enum: `switch`<br>**Default**: `"switch"` |

## Example (as JSON)

```json
{
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 1500,
  "num_packets": 100,
  "tcpdump_expression": "port 443",
  "type": "switch",
  "ports": [
    "ports3",
    "ports4"
  ]
}
```

