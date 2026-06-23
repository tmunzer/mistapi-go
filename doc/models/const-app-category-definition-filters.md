
# Const App Category Definition Filters

Platform-specific application category filters

## Structure

`ConstAppCategoryDefinitionFilters`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Srx` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Ssr` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constAppCategoryDefinitionFilters := models.ConstAppCategoryDefinitionFilters{
        Srx:                  []string{
            "Enhanced_Images_Media",
            "Enhanced_Web_Images",
            "Enhanced_Image_Servers",
        },
        Ssr:                  []string{
            "ssr6",
        },
    }

}
```

