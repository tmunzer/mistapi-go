
# Module Stat Item Pics Item Port Groups Item

Port group summary for a PIC

## Structure

`ModuleStatItemPicsItemPortGroupsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | Number of ports in this PIC port group |
| `Type` | `*string` | Optional | Port media or interface type for this PIC port group |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemPicsItemPortGroupsItem := models.ModuleStatItemPicsItemPortGroupsItem{
        Count:                models.ToPointer(38),
        Type:                 models.ToPointer("type8"),
    }

}
```

