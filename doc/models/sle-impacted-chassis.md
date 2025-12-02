
# Sle Impacted Chassis

## Structure

`SleImpactedChassis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Chassis` | [`[]models.SleImpactedChassisChassisItem`](../../doc/models/sle-impacted-chassis-chassis-item.md) | Optional | - |
| `Classifier` | `*string` | Optional | - |
| `End` | `*int` | Optional | - |
| `Failure` | `*string` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Metric` | `*string` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Start` | `*int` | Optional | - |
| `TotalCount` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "chassis": [
    {
      "chassis": "chassis0",
      "degraded": 171.46,
      "duration": 44.52,
      "role": "role0",
      "switch_mac": "switch_mac4"
    }
  ],
  "classifier": "classifier0",
  "end": 146,
  "failure": "failure8",
  "limit": 24
}
```

