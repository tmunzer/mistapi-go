
# Capture Gateway

Initiate a Gateway (SSR) Packet Capture

*This model accepts additional fields of type interface{}.*

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
| `Ports` | [`map[string]models.CaptureGatewayGatewaysPort`](../../doc/models/capture-gateway-gateways-port.md) | Optional | Property key is the port ID |
| `Type` | `string` | Required, Constant | enum: `gateway`<br>**Default**: `"gateway"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

