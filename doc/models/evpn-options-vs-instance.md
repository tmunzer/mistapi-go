
# Evpn Options Vs Instance

EVPN virtual-switch instance network mapping

## Structure

`EvpnOptionsVsInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnOptionsVsInstance := models.EvpnOptionsVsInstance{
        Networks:             []string{
            "networks0",
        },
    }

}
```

