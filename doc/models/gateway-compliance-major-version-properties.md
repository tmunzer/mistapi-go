
# Gateway Compliance Major Version Properties

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayComplianceMajorVersionProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorCount` | `*int` | Optional | - |
| `MajorVersion` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "major_version": "19.4R2-S1.2",
  "major_count": 56,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

