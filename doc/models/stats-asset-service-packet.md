
# Stats Asset Service Packet

Service data advertisement from a BLE asset

## Structure

`StatsAssetServicePacket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | `*string` | Optional | Service data payload (hex encoded) |
| `LastRxTime` | `*int` | Optional | Unix timestamp when this service data was last received |
| `RxCnt` | `*int` | Optional | Total number of times this service data was received |
| `Uuid` | `*string` | Optional | BLE service UUID advertised by the asset service packet |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsAssetServicePacket := models.StatsAssetServicePacket{
        Data:                 models.ToPointer("640"),
        LastRxTime:           models.ToPointer(1645855923),
        RxCnt:                models.ToPointer(213065),
        Uuid:                 models.ToPointer("00003e10-0000-1000-8000-00805f9b34fb"),
    }

}
```

