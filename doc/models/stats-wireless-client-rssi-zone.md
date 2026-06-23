
# Stats Wireless Client Rssi Zone

RSSI zone currently containing a wireless client

## Structure

`StatsWirelessClientRssiZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Since` | `*float64` | Optional | Time when the wireless client entered the RSSI zone, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWirelessClientRssiZone := models.StatsWirelessClientRssiZone{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Since:                models.ToPointer(float64(108.64)),
    }

}
```

