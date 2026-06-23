
# Evpn Options Overlay

EVPN overlay BGP settings

## Structure

`EvpnOptionsOverlay`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `As` | `*int` | Optional | Overlay BGP Local AS Number<br><br>**Default**: `65000`<br><br>**Constraints**: `>= 1`, `<= 65535` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnOptionsOverlay := models.EvpnOptionsOverlay{
        As:                   models.ToPointer(65000),
    }

}
```

