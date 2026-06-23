
# Stats Wireless Client Zone

Zone currently containing a wireless client

## Structure

`StatsWirelessClientZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Since` | `*float64` | Optional | Time when the wireless client entered the zone, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWirelessClientZone := models.StatsWirelessClientZone{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Since:                models.ToPointer(float64(103.24)),
    }

}
```

