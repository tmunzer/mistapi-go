
# Module Stat Item Network Resource

Network resource usage counter reported by a device module

## Structure

`ModuleStatItemNetworkResource`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | Current utilization count for the network resource<br><br>**Constraints**: `>= 0` |
| `Limit` | `*int` | Optional | Maximum supported count for the network resource<br><br>**Constraints**: `>= 0` |
| `Type` | `*string` | Optional | Network resource category, such as FIB or FLOW |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemNetworkResource := models.ModuleStatItemNetworkResource{
        Count:                models.ToPointer(17),
        Limit:                models.ToPointer(768000),
        Type:                 models.ToPointer("FIB"),
    }

}
```

