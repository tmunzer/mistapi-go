
# Sle Impacted Applications

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "apps": [
    {
      "app": "app4",
      "degraded": 58,
      "duration": 164,
      "name": "name6",
      "threshold": 20,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "app": "app4",
      "degraded": 58,
      "duration": 164,
      "name": "name6",
      "threshold": 20,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "app": "app4",
      "degraded": 58,
      "duration": 164,
      "name": "name6",
      "threshold": 20,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "classifier": "classifier2",
  "end": 220,
  "failure": "failure4",
  "limit": "limit2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

