
# Sdkstats Wireless Client Vbeacon

Virtual beacon currently associated with an SDK client

## Structure

`SdkstatsWirelessClientVbeacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Since` | `float64` | Required | Time when the SDK client began matching the virtual beacon, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    sdkstatsWirelessClientVbeacon := models.SdkstatsWirelessClientVbeacon{
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Since:                float64(77.62),
    }

}
```

