
# Stats Ap Usb Stat

USB peripheral status reported by an AP

## Structure

`StatsApUsbStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `models.Optional[int]` | Optional, Read-only | Radio channel used by the USB peripheral |
| `Connected` | `models.Optional[bool]` | Optional, Read-only | Whether the USB peripheral is connected |
| `LastActivity` | `models.Optional[int]` | Optional, Read-only | Time of the last USB peripheral activity, in epoch seconds |
| `Type` | `models.Optional[string]` | Optional, Read-only | USB peripheral type reported by the AP |
| `Up` | `models.Optional[bool]` | Optional, Read-only | Whether the USB peripheral is operational |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApUsbStat := models.StatsApUsbStat{
        Channel:              models.NewOptional(models.ToPointer(3)),
        Connected:            models.NewOptional(models.ToPointer(true)),
        LastActivity:         models.NewOptional(models.ToPointer(1586873254)),
        Type:                 models.NewOptional(models.ToPointer("imagotag")),
        Up:                   models.NewOptional(models.ToPointer(true)),
    }

}
```

