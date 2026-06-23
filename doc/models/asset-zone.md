
# Asset Zone

Map zone membership for an individual asset statistic

## Structure

`AssetZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Since` | `*float64` | Optional | Timestamp when the asset entered this zone |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    assetZone := models.AssetZone{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Since:                models.ToPointer(float64(7.36)),
    }

}
```

