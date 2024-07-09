
# Sle Impacted Gateways

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

## Example (as JSON)

```json
{
  "classifier": "classifier2",
  "end": 130,
  "failure": "failure4",
  "gateways": [
    {
      "degraded": 114.84,
      "duration": 70,
      "gateway_mac": "gateway_mac2",
      "gateway_model": "gateway_model6",
      "gateway_version": "gateway_version8"
    }
  ],
  "limit": 40
}
```
