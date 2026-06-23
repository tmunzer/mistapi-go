
# Const App Category Definition

Application category definition returned by the constants API

## Structure

`ConstAppCategoryDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | Description of the app category |
| `Filters` | [`*models.ConstAppCategoryDefinitionFilters`](../../doc/models/const-app-category-definition-filters.md) | Optional | Platform-specific application category filters |
| `Includes` | `[]string` | Optional | List of other App Categories contained by this one |
| `Key` | `string` | Required | Machine-readable application category key |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constAppCategoryDefinition := models.ConstAppCategoryDefinition{
        Display:              "Images",
        Filters:              models.ToPointer(models.ConstAppCategoryDefinitionFilters{
            Srx:                  []string{
                "Enhanced_Images_Media",
                "Enhanced_Web_Images",
                "Enhanced_Image_Servers",
            },
            Ssr:                  []string{
                "ssr6",
                "ssr5",
                "ssr4",
            },
        }),
        Includes:             []string{
            "includes9",
            "includes0",
        },
        Key:                  "Images",
    }

}
```

