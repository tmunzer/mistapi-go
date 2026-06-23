
# Const Device Ap Band 6

6 GHz radio capability limits for an AP model

## Structure

`ConstDeviceApBand6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaxClients` | `*int` | Optional | Maximum client count supported on the 6 GHz radio |
| `MaxPower` | `*int` | Optional | Maximum transmit power for the 6 GHz radio, in dBm |
| `MinPower` | `*int` | Optional | Minimum transmit power for the 6 GHz radio, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceApBand6 := models.ConstDeviceApBand6{
        MaxClients:           models.ToPointer(128),
        MaxPower:             models.ToPointer(17),
        MinPower:             models.ToPointer(8),
    }

}
```

