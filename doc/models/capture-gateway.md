
# Capture Gateway

Initiate a Gateway (SSR/SRX) Packet Capture

## Structure

`CaptureGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureGatewayFormatEnum`](../../doc/models/capture-gateway-format-enum.md) | Optional | enum: `stream`<br><br>**Default**: `"stream"` |
| `Gateways` | [`map[string]models.CaptureGatewayGateways`](../../doc/models/capture-gateway-gateways.md) | Required | List of SSRs. Property key is the SSR MAC |
| `MaxPktLen` | `models.Optional[int]` | Optional | minimum is 64 (SSR) / 68 (SRX) maximum is 10240 (SSR) / 1520 (SRX)<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 1520` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ports` | [`map[string]models.CaptureGatewayGatewaysPort`](../../doc/models/capture-gateway-gateways-port.md) | Optional | Property key is the port ID |
| `Type` | `string` | Required, Constant | enum: `gateway`<br><br>**Value**: `"gateway"` |

## Example (as JSON)

```json
{
  "duration": 300,
  "format": "stream",
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
  },
  "max_pkt_len": 128,
  "num_packets": 1000,
  "type": "gateway",
  "ports": {
    "key0": {
      "tcpdump_expression": "tcpdump_expression0"
    },
    "key1": {
      "tcpdump_expression": "tcpdump_expression0"
    },
    "key2": {
      "tcpdump_expression": "tcpdump_expression0"
    }
  }
}
```

