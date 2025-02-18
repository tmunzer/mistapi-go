
# User Mac Import

*This model accepts additional fields of type interface{}.*

## Structure

`UserMacImport`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Added` | `[]string` | Optional | - |
| `Errors` | `[]string` | Optional | - |
| `Updated` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "added": [
    "921b638445cd"
  ],
  "errors": [
    "921b638445ce - mac invalid",
    "921b638445cf - mac already provided"
  ],
  "updated": [
    "721b638445ef",
    "721b638445ee"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

