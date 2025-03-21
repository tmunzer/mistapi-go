
# Sle Impacted Interfaces

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedInterfaces`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `*string` | Optional | - |
| `End` | `*int` | Optional | - |
| `Failure` | `*string` | Optional | - |
| `Interfaces` | [`[]models.SleImpactedInterfacesInterface`](../../doc/models/sle-impacted-interfaces-interface.md) | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Metric` | `*string` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Start` | `*int` | Optional | - |
| `TotalCount` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "classifier": "classifier0",
  "end": 186,
  "failure": "failure8",
  "interfaces": [
    {
      "degraded": 98.24,
      "duration": 227.3,
      "interface_name": "interface_name4",
      "switch_mac": "switch_mac2",
      "switch_name": "switch_name6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "degraded": 98.24,
      "duration": 227.3,
      "interface_name": "interface_name4",
      "switch_mac": "switch_mac2",
      "switch_name": "switch_name6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "degraded": 98.24,
      "duration": 227.3,
      "interface_name": "interface_name4",
      "switch_mac": "switch_mac2",
      "switch_name": "switch_name6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "limit": 16,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

