
# Ap Led

Indicator light settings for an access point

## Structure

`ApLed`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Brightness` | `*int` | Optional | Indicator LED brightness level from 0 to 255<br><br>**Default**: `255`<br><br>**Constraints**: `>= 0`, `<= 255` |
| `Enabled` | `*bool` | Optional | Whether the AP indicator LED is enabled<br><br>**Default**: `true` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apLed := models.ApLed{
        Brightness:           models.ToPointer(255),
        Enabled:              models.ToPointer(true),
    }

}
```

