
# Site Sle Metrics

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSleMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Supported` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": [
    "enabled3",
    "enabled4"
  ],
  "supported": [
    "supported0",
    "supported9",
    "supported8"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

