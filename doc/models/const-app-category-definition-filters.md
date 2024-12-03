
# Const App Category Definition Filters

*This model accepts additional fields of type interface{}.*

## Structure

`ConstAppCategoryDefinitionFilters`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Srx` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Ssr` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "srx": [
    "Enhanced_Images_Media",
    "Enhanced_Web_Images",
    "Enhanced_Image_Servers"
  ],
  "ssr": [
    "ssr8",
    "ssr9"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

