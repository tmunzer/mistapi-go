
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
            Channel:                models.NewOptional(models.ToPointer(80)),
            DynamicChainingEnabled: models.NewOptional(models.ToPointer(false)),
            Mac:                    models.NewOptional(models.ToPointer("mac4")),
            NoiseFloor:             models.NewOptional(models.ToPointer(180)),
        }),
        Band5:                models.ToPointer(models.ApRadioStat{
            Bandwidth:              models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
            Channel:                models.NewOptional(models.ToPointer(132)),
            DynamicChainingEnabled: models.NewOptional(models.ToPointer(false)),
            Mac:                    models.NewOptional(models.ToPointer("mac6")),
            NoiseFloor:             models.NewOptional(models.ToPointer(128)),
        }),
        Band6:                models.ToPointer(models.ApRadioStat{
            Bandwidth:              models.ToPointer(models.Dot11BandwidthEnum_ENUM80),
            Channel:                models.NewOptional(models.ToPointer(200)),
            DynamicChainingEnabled: models.NewOptional(models.ToPointer(false)),
            Mac:                    models.NewOptional(models.ToPointer("mac8")),
            NoiseFloor:             models.NewOptional(models.ToPointer(60)),
        }),
    }

}
```

