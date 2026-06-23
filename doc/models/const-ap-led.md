
# Const Ap Led

AP LED status and `last_trouble` code definition returned by the constants API

## Structure

`ConstApLed`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `string` | Required | Numeric LED status and `last_trouble` code reported by the AP |
| `Description` | `string` | Required | Human-readable explanation of the AP status |
| `Key` | `string` | Required | Machine-readable AP status key |
| `Name` | `string` | Required | Display name for the AP status |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constApLed := models.ConstApLed{
        Code:                 "01",
        Description:          "LED not working",
        Key:                  "LED_FAILURE",
        Name:                 "LED Failure",
    }

}
```

