
# Capture Gateway

Initiate a Gateway (SSR) Packet Capture

## Structure

`CaptureGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `>= 0`, `<= 86400` |
| `Format` | [`*models.CaptureGatewayFormatEnum`](../../doc/models/capture-gateway-format-enum.md) | Optional | enum: `stream`<br>**Default**: `"stream"` |
| `Gateways` | [`map[string]models.CaptureGatewayGateways`](../../doc/models/capture-gateway-gateways.md) | Optional | List of SSRs. Property key is the SSR MAC |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `128`<br>**Constraints**: `>= 0`, `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024`<br>**Constraints**: `>= 0` |
| `Ports` | `[]string` | Optional | dict of port which uses port id as the key |
| `Type` | `string` | Required, Constant | enum: `gateway`<br>**Default**: `"gateway"` |

## Example (as JSON)

```json
{
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 1500,
  "num_packets": 100,
  "type": "gateway",
  "gateways": {
    "key0": {
      "ports": {
        "key0": {
          "tcpdump_expression": "tcpdump_expression0"
        },
        "key1": {
          "tcpdump_expression": "tcpdump_expression0"
        }
      }
    }
  }
}
```

