
# Sle Impacted Aps

## Structure

`SleImpactedAps`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aps` | [`[]models.SleImpactedApsAp`](../../doc/models/sle-impacted-aps-ap.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Classifier` | `string` | Required | - |
| `End` | `float64` | Required | - |
| `Failure` | `string` | Required | - |
| `Limit` | `float64` | Required | - |
| `Metric` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Page` | `float64` | Required | - |
| `Start` | `float64` | Required | - |
| `TotalCount` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "aps": [
    {
      "ap_mac": "ap_mac6",
      "degraded": 30.24,
      "duration": 159.3,
      "name": "name4",
      "total": 197.76
    }
  ],
  "classifier": "classifier2",
  "end": 101.8,
  "failure": "failure4",
  "limit": 148.22,
  "metric": "metric6",
  "page": 21.36,
  "start": 57.86,
  "total_count": 171.98
}
```

