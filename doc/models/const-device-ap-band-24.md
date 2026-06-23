
# Const Device Ap Band 24

2.4 GHz radio capability limits for an AP model

## Structure

`ConstDeviceApBand24`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band5ChannelsOp` | `*string` | Optional | 5 GHz channel set used when this radio operates on 5 GHz |
| `MaxClients` | `*int` | Optional | Maximum client count supported on the 2.4 GHz radio |
| `MaxPower` | `*int` | Optional | Maximum transmit power for the 2.4 GHz radio, in dBm |
| `MinPower` | `*int` | Optional | Minimum transmit power for the 2.4 GHz radio, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceApBand24 := models.ConstDeviceApBand24{
        Band5ChannelsOp:      models.ToPointer("low"),
        MaxClients:           models.ToPointer(128),
        MaxPower:             models.ToPointer(19),
        MinPower:             models.ToPointer(8),
    }

}
```

