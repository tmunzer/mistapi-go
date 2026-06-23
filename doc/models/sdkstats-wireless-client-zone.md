
# Sdkstats Wireless Client Zone

Zone currently containing an SDK client

## Structure

`SdkstatsWirelessClientZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Since` | `float64` | Required | Time when the SDK client entered the zone, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    sdkstatsWirelessClientZone := models.SdkstatsWirelessClientZone{
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Since:                float64(197.62),
    }

}
```

