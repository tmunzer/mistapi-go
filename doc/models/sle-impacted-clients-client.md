
# Sle Impacted Clients Client

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedClientsClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*int` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Switches` | [`[]models.SleImpactedClientsClientSwitch`](../../doc/models/sle-impacted-clients-client-switch.md) | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 248,
  "duration": 98,
  "mac": "mac2",
  "name": "name8",
  "switches": [
    {
      "interfaces": [
        "interfaces9"
      ],
      "switch_mac": "switch_mac6",
      "switch_name": "switch_name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

