
# Jse Device

JSE device reported for an organization

## Structure

`JseDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExtIp` | `*string` | Optional | External IP address for the JSE device, when available |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `*string` | Optional | Device MAC address for the JSE device |
| `Model` | `*string` | Optional | Device model for the JSE device |
| `Serial` | `*string` | Optional | Device serial number for the JSE device |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    jseDevice := models.JseDevice{
        ExtIp:                models.ToPointer("ext_ip0"),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  models.ToPointer("mac6"),
        Model:                models.ToPointer("model0"),
        Serial:               models.ToPointer("serial2"),
    }

}
```

