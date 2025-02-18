
# Capture Gateway Gateways

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureGatewayGateways`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | [`map[string]models.CaptureGatewayGatewaysPort`](../../doc/models/capture-gateway-gateways-port.md) | Optional | Property key is the port ID |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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

