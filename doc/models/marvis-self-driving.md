
# Marvis Self Driving

Self-driving network automation settings per domain

## Structure

`MarvisSelfDriving`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Wan` | [`*models.MarvisSelfDrivingDomain`](../../doc/models/marvis-self-driving-domain.md) | Optional | Self-driving automation settings for one Marvis domain |
| `Wired` | [`*models.MarvisSelfDrivingDomain`](../../doc/models/marvis-self-driving-domain.md) | Optional | Self-driving automation settings for one Marvis domain |
| `Wireless` | [`*models.MarvisSelfDrivingDomain`](../../doc/models/marvis-self-driving-domain.md) | Optional | Self-driving automation settings for one Marvis domain |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisSelfDriving := models.MarvisSelfDriving{
        Wan:                  models.ToPointer(models.MarvisSelfDrivingDomain{
            Enabled:              models.ToPointer(false),
        }),
        Wired:                models.ToPointer(models.MarvisSelfDrivingDomain{
            Enabled:              models.ToPointer(false),
        }),
        Wireless:             models.ToPointer(models.MarvisSelfDrivingDomain{
            Enabled:              models.ToPointer(false),
        }),
    }

}
```

