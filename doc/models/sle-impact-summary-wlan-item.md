
# Sle Impact Summary Wlan Item

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactSummaryWlanItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `float64` | Required | - |
| `Duration` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | - |
| `WlanId` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 245.48,
  "duration": 118.54,
  "name": "name8",
  "total": 238.52,
  "wlan_id": "wlan_id0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

