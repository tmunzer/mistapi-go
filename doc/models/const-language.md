
# Const Language

*This model accepts additional fields of type interface{}.*

## Structure

`ConstLanguage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | - |
| `DisplayNative` | `string` | Required | - |
| `Key` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "display": "English (US)",
  "display_native": "English (US)",
  "key": "en-US",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

