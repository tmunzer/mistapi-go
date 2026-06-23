
# Map Micello

Micello map import configuration

## Structure

`MapMicello`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountKey` | `string` | Required | Account key that has access to the map |
| `DefaultLevelId` | `int` | Required | Micello default floor or level identifier |
| `MapId` | `uuid.UUID` | Required | Micello map identifier to import |
| `VendorName` | `string` | Required, Constant | The vendor ‘micello’. enum: `micello`<br><br>**Value**: `"micello"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mapMicello := models.MapMicello{
        AccountKey:           "",
        DefaultLevelId:       5,
        MapId:                uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"),
        VendorName:           "micello",
    }

}
```

