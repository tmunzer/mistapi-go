
# Const App Category Definition

*This model accepts additional fields of type interface{}.*

## Structure

`ConstAppCategoryDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | Description of the app category |
| `Filters` | [`*models.ConstAppCategoryDefinitionFilters`](../../doc/models/const-app-category-definition-filters.md) | Optional | - |
| `Includes` | `[]string` | Optional | List of other App Categories contained by this one |
| `Key` | `string` | Required | Key name of the app category |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "display": "Images",
  "filters": {
    "srx": [
      "Enhanced_Images_Media",
      "Enhanced_Web_Images",
      "Enhanced_Image_Servers"
    ],
    "ssr": [
      "ssr6",
      "ssr5",
      "ssr4"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "key": "Images",
  "includes": [
    "includes5"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

