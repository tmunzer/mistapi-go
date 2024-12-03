
# Response Map Import Summary

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseMapImportSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumApAssigned` | `int` | Required | - |
| `NumInvAssigned` | `int` | Required | - |
| `NumMapAssigned` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "num_ap_assigned": 120,
  "num_inv_assigned": 228,
  "num_map_assigned": 218,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

