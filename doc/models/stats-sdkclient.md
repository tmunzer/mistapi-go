
# Stats Sdkclient

Detailed statistics for an individual SDK client

## Structure

`StatsSdkclient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | Map identifier for the SDK client's location, if known |
| `Name` | `*string` | Optional | Display name provided for the SDK client |
| `NetworkConnection` | [`models.StatsSdkclientNetworkConnection`](../../doc/models/stats-sdkclient-network-connection.md) | Required | Current network connection details reported for an SDK client |
| `Uuid` | `uuid.UUID` | Required | Application UUID for the SDK client |
| `X` | `*float64` | Optional | Horizontal map coordinate of the SDK client location, in pixels, if known |
| `Y` | `*float64` | Optional | Vertical map coordinate of the SDK client location, in pixels, if known |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsSdkclient := models.StatsSdkclient{
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        MapId:                models.NewOptional(models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5"))),
        Name:                 models.ToPointer("John's iPhone"),
        NetworkConnection:    models.StatsSdkclientNetworkConnection{
            Mac:                  "mac2",
            Rssi:                 float64(235.8),
            SignalLevel:          float64(47.82),
            Type:                 "type2",
        },
        Uuid:                 uuid.MustParse("ada72f8f-1643-e5c6-94db-f2a5636f1a64"),
        X:                    models.ToPointer(float64(60)),
        Y:                    models.ToPointer(float64(80)),
    }

}
```

