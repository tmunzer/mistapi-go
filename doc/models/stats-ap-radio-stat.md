
# Stats Ap Radio Stat

Per-band radio statistics reported by an AP

## Structure

`StatsApRadioStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Runtime radio statistics for an access point radio |
| `Band5` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Runtime radio statistics for an access point radio |
| `Band6` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Runtime radio statistics for an access point radio |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApRadioStat := models.StatsApRadioStat{
        Band24:               models.ToPointer(models.ApRadioStat{
            Bandwidth:              models.ToPointer(models.Dot11BandwidthEnum_ENUM160),
        }),
        Band5:                models.ToPointer(models.ApRadioStat{
            Bandwidth:              models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
        }),
        Band6:                models.ToPointer(models.ApRadioStat{
            Bandwidth:              models.ToPointer(models.Dot11BandwidthEnum_ENUM80),
        }),
    }

}
```

