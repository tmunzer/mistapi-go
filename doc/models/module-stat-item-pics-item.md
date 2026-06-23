
# Module Stat Item Pics Item

Physical Interface Card summary for a device module

## Structure

`ModuleStatItemPicsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Index` | `*int` | Optional | PIC index within the device module |
| `ModelNumber` | `*string` | Optional | Model number reported for the PIC |
| `PortGroups` | [`[]models.ModuleStatItemPicsItemPortGroupsItem`](../../doc/models/module-stat-item-pics-item-port-groups-item.md) | Optional | Port group summaries for a PIC |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemPicsItem := models.ModuleStatItemPicsItem{
        Index:                models.ToPointer(6),
        ModelNumber:          models.ToPointer("model_number0"),
        PortGroups:           []models.ModuleStatItemPicsItemPortGroupsItem{
            models.ModuleStatItemPicsItemPortGroupsItem{
                Count:                models.ToPointer(32),
                Type:                 models.ToPointer("type6"),
            },
            models.ModuleStatItemPicsItemPortGroupsItem{
                Count:                models.ToPointer(32),
                Type:                 models.ToPointer("type6"),
            },
            models.ModuleStatItemPicsItemPortGroupsItem{
                Count:                models.ToPointer(32),
                Type:                 models.ToPointer("type6"),
            },
        },
    }

}
```

