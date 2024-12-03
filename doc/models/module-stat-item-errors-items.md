
# Module Stat Item Errors Items

*This model accepts additional fields of type interface{}.*

## Structure

`ModuleStatItemErrorsItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | `*string` | Optional | - |
| `MinimumVersion` | `*string` | Optional | - |
| `Reason` | `*string` | Optional | - |
| `Since` | `int` | Required | - |
| `Type` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "feature": "Mist-Management",
  "minimum_version": "128T-6.0.0-1",
  "since": 1657497600,
  "type": "FW_UPGRADE_REQUIRED_BY_FEATURE",
  "reason": "reason8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

