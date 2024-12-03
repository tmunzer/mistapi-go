
# Mxedge Upgrade Info Items

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeUpgradeInfoItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Default` | `*bool` | Optional | - |
| `Package` | `string` | Required | - |
| `Version` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "default": false,
  "package": "package0",
  "version": "version4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

