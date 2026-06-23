
# Switch Mist Nac

Mist NAC RadSec settings for a switch

## Structure

`SwitchMistNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether Mist NAC RadSec is enabled for the switch |
| `Network` | `*string` | Optional | Switch network used for Mist NAC RadSec connectivity |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMistNac := models.SwitchMistNac{
        Enabled:              models.ToPointer(false),
        Network:              models.ToPointer("network8"),
    }

}
```

