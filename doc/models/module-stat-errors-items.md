
# Module Stat Errors Items

## Structure

`ModuleStatErrorsItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | `*string` | Optional | - |
| `MinimumVersion` | `*string` | Optional | - |
| `Reason` | `*string` | Optional | - |
| `Since` | `int` | Required | - |
| `Type` | `string` | Required | - |

## Example (as JSON)

```json
{
  "feature": "Mist-Management",
  "minimum_version": "128T-6.0.0-1",
  "since": 1657497600,
  "type": "FW_UPGRADE_REQUIRED_BY_FEATURE",
  "reason": "reason4"
}
```

