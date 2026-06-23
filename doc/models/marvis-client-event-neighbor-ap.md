
# Marvis Client Event Neighbor Ap

A neighboring AP observed in a Marvis Client event

## Structure

`MarvisClientEventNeighborAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Wi-Fi band the AP is operating on |
| `Bssid` | `*string` | Optional | BSSID of the neighboring AP |
| `Channel` | `*int` | Optional | Channel the neighboring AP is on |
| `Rssi` | `*int` | Optional | RSSI of the neighboring AP signal, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisClientEventNeighborAp := models.MarvisClientEventNeighborAp{
        Band:                 models.ToPointer("band4"),
        Bssid:                models.ToPointer("bssid6"),
        Channel:              models.ToPointer(66),
        Rssi:                 models.ToPointer(32),
    }

}
```

