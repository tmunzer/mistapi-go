
# Sle Impacted Interfaces

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

## Example (as JSON)

```json
{
  "classifier": "classifier4",
  "end": 216,
  "failure": "failure2",
  "interfaces": [
    {
      "degraded": 98.24,
      "duration": 227.3,
      "interface_name": "interface_name4",
      "switch_mac": "switch_mac2",
      "switch_name": "switch_name6"
    },
    {
      "degraded": 98.24,
      "duration": 227.3,
      "interface_name": "interface_name4",
      "switch_mac": "switch_mac2",
      "switch_name": "switch_name6"
    },
    {
      "degraded": 98.24,
      "duration": 227.3,
      "interface_name": "interface_name4",
      "switch_mac": "switch_mac2",
      "switch_name": "switch_name6"
    }
  ],
  "limit": 46
}
```

