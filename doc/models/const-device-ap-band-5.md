
# Const Device Ap Band 5

5 GHz radio capability limits for an AP model

## Structure

`ConstDeviceApBand5`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaxClients` | `*int` | Optional | Maximum client count supported on the 5 GHz radio |
| `MaxPower` | `*int` | Optional | Maximum transmit power for the 5 GHz radio, in dBm |
| `MinPower` | `*int` | Optional | Minimum transmit power for the 5 GHz radio, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceApBand5 := models.ConstDeviceApBand5{
        MaxClients:           models.ToPointer(128),
        MaxPower:             models.ToPointer(17),
        MinPower:             models.ToPointer(8),
    }

}
```

