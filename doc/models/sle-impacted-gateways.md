
# Sle Impacted Gateways

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedGateways`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `*string` | Optional | - |
| `End` | `*int` | Optional | - |
| `Failure` | `*string` | Optional | - |
| `Gateways` | [`[]models.SleImpactedGatewaysGateway`](../../doc/models/sle-impacted-gateways-gateway.md) | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Metric` | `*string` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Start` | `*int` | Optional | - |
| `TotalCount` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "classifier": "classifier8",
  "end": 122,
  "failure": "failure6",
  "gateways": [
    {
      "degraded": 114.84,
      "duration": 70,
      "gateway_mac": "gateway_mac2",
      "gateway_model": "gateway_model6",
      "gateway_version": "gateway_version8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "limit": 48,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

