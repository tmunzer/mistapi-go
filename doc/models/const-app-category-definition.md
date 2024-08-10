
# Const App Category Definition

## Structure

`ConstAppCategoryDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | Description of the app category |
| `Filters` | [`*models.ConstAppCategoryDefinitionFilters`](../../doc/models/const-app-category-definition-filters.md) | Optional | - |
| `Includes` | `[]string` | Optional | List of other App Categories contained by this one |
| `Key` | `string` | Required | Key name of the app category |

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
    ]
  },
  "key": "Images",
  "includes": [
    "includes5"
  ]
}
```

