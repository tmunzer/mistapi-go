
# Stats Wireless Client Vbeacon

Virtual beacon currently associated with a wireless client

## Structure

`StatsWirelessClientVbeacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Since` | `*float64` | Optional | Time when the wireless client began matching the virtual beacon, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWirelessClientVbeacon := models.StatsWirelessClientVbeacon{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Since:                models.ToPointer(float64(242.84)),
    }

}
```

