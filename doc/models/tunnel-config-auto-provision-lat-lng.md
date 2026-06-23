
# Tunnel Config Auto Provision Lat Lng

Geographic coordinate override for tunnel POP selection

## Structure

`TunnelConfigAutoProvisionLatLng`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Lat` | `float64` | Required | Geographic latitude used for POP selection override |
| `Lng` | `float64` | Required | Geographic longitude used for POP selection override |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigAutoProvisionLatLng := models.TunnelConfigAutoProvisionLatLng{
        Lat:                  float64(37.295833),
        Lng:                  float64(-122.032946),
    }

}
```

