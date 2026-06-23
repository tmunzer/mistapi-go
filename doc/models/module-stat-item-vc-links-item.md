
# Module Stat Item Vc Links Item

Virtual chassis link endpoint for a device module

## Structure

`ModuleStatItemVcLinksItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NeighborModuleIdx` | `*int` | Optional | Index of the neighboring module connected through this virtual chassis link |
| `NeighborPortId` | `*string` | Optional | Port identifier on the neighboring module for this virtual chassis link |
| `PortId` | `*string` | Optional | Local port identifier for this virtual chassis link |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemVcLinksItem := models.ModuleStatItemVcLinksItem{
        NeighborModuleIdx:    models.ToPointer(1),
        NeighborPortId:       models.ToPointer("vcp-255/1/0"),
        PortId:               models.ToPointer("vcp-255/1/0"),
    }

}
```

