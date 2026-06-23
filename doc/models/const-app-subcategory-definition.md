
# Const App Subcategory Definition

Application subcategory definition returned by the constants API

## Structure

`ConstAppSubcategoryDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | Description of the app subcategory |
| `Key` | `string` | Required | Machine-readable application subcategory key |
| `TrafficType` | `string` | Required | Type of traffic (QoS) of the app subcategory |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constAppSubcategoryDefinition := models.ConstAppSubcategoryDefinition{
        Display:              "Office Document",
        Key:                  "Office_Documents",
        TrafficType:          "Images",
    }

}
```

