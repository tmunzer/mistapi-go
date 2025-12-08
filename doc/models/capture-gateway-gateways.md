
# Capture Gateway Gateways

## Structure

`CaptureGatewayGateways`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | [`map[string]models.CaptureGatewayGatewaysPort`](../../doc/models/capture-gateway-gateways-port.md) | Optional | Property key is the port ID |

## Example (as JSON)

```json
{
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

