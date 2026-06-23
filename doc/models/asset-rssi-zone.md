
# Asset Rssi Zone

RSSI zone membership for an individual asset statistic

## Structure

`AssetRssiZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Since` | `*float64` | Optional | Timestamp when the asset entered this RSSI zone |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    assetRssiZone := models.AssetRssiZone{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Since:                models.ToPointer(float64(44.42)),
    }

}
```

