
# Sle Impacted Gateways Gateway

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedGatewaysGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*float64` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `GatewayMac` | `*string` | Optional | - |
| `GatewayModel` | `*string` | Optional | - |
| `GatewayVersion` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 231.66,
  "duration": 232,
  "gateway_mac": "gateway_mac4",
  "gateway_model": "gateway_model8",
  "gateway_version": "gateway_version0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

