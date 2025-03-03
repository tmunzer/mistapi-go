
# Capture Gateway

Initiate a Gateway (SSR/SRX) Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureGatewayFormatEnum`](../../doc/models/capture-gateway-format-enum.md) | Optional | enum: `stream`<br>**Default**: `"stream"` |
| `Gateways` | [`map[string]models.CaptureGatewayGateways`](../../doc/models/capture-gateway-gateways.md) | Required | List of SSRs. Property key is the SSR MAC |
| `MaxPktLen` | `models.Optional[int]` | Optional | minimum is 64 (SSR) / 68 (SRX) maximum is 10240 (SSR) / 1520 (SRX)<br>**Default**: `512`<br>**Constraints**: `>= 64`, `<= 1520` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br>**Default**: `1024`<br>**Constraints**: `>= 0`, `<= 10000` |
| `Ports` | [`map[string]models.CaptureGatewayGatewaysPort`](../../doc/models/capture-gateway-gateways-port.md) | Optional | Property key is the port ID |
| `Type` | `string` | Required, Constant | enum: `gateway`<br>**Value**: `"gateway"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 300,
  "format": "stream",
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
  "max_pkt_len": 128,
  "num_packets": 1000,
  "type": "gateway",
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

