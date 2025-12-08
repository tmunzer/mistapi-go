
# Sle Impacted Applications

## Structure

`SleImpactedApplications`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | [`[]models.SleImpactedApplicationsApp`](../../doc/models/sle-impacted-applications-app.md) | Optional | - |
| `Classifier` | `*string` | Optional | - |
| `End` | `*int` | Optional | - |
| `Failure` | `*string` | Optional | - |
| `Limit` | `*string` | Optional | - |
| `Metric` | `*string` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Start` | `*int` | Optional | - |
| `TotalCount` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "apps": [
    {
      "app": "app4",
      "degraded": 58,
      "duration": 164,
      "name": "name6",
      "threshold": 20
    },
    {
      "app": "app4",
      "degraded": 58,
      "duration": 164,
      "name": "name6",
      "threshold": 20
    },
    {
      "app": "app4",
      "degraded": 58,
      "duration": 164,
      "name": "name6",
      "threshold": 20
    }
  ],
  "classifier": "classifier2",
  "end": 220,
  "failure": "failure4",
  "limit": "limit2"
}
```

