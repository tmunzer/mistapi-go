
# Routing Policy Term Matching Vpn Path Sla

*This model accepts additional fields of type interface{}.*

## Structure

`RoutingPolicyTermMatchingVpnPathSla`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaxJitter` | `models.Optional[int]` | Optional | - |
| `MaxLatency` | `models.Optional[int]` | Optional | - |
| `MaxLoss` | `models.Optional[int]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "max_latency": 1500,
  "max_loss": 30,
  "max_jitter": 142,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

