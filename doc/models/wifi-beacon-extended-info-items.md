
# Wifi Beacon Extended Info Items

Extended Wi-Fi beacon packet metadata with frame and sequence control fields

## Structure

`WifiBeaconExtendedInfoItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FrameCtrl` | `*int` | Optional | Frame control field of 802.11 header |
| `Payload` | `*string` | Optional | Extended Info Payload associated with frame |
| `SeqCtrl` | `*int` | Optional | Sequence control field of 802.11 header |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wifiBeaconExtendedInfoItems := models.WifiBeaconExtendedInfoItems{
        FrameCtrl:            models.ToPointer(234),
        Payload:              models.ToPointer("payload8"),
        SeqCtrl:              models.ToPointer(146),
    }

}
```

