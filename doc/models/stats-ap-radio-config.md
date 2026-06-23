
# Stats Ap Radio Config

Radio configuration reported by an AP

## Structure

`StatsApRadioConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.StatsApRadioConfigBand`](../../doc/models/stats-ap-radio-config-band.md) | Optional | Per-band radio configuration reported by an AP |
| `Band24Usage` | `models.Optional[string]` | Optional, Read-only | Current operating use for the 2.4 GHz radio |
| `Band5` | [`*models.StatsApRadioConfigBand`](../../doc/models/stats-ap-radio-config-band.md) | Optional | Per-band radio configuration reported by an AP |
| `Band6` | [`*models.StatsApRadioConfigBand`](../../doc/models/stats-ap-radio-config-band.md) | Optional | Per-band radio configuration reported by an AP |
| `ScanningEnabled` | `*bool` | Optional | Whether radio scanning is enabled on the AP |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApRadioConfig := models.StatsApRadioConfig{
        Band24:               models.ToPointer(models.StatsApRadioConfigBand{
            AllowRrmDisable:        models.NewOptional(models.ToPointer(false)),
            Bandwidth:              models.NewOptional(models.ToPointer(float64(4.04))),
            Channel:                models.ToPointer(80),
            Disabled:               models.NewOptional(models.ToPointer(false)),
            DynamicChainingEnabled: models.NewOptional(models.ToPointer(false)),
        }),
        Band24Usage:          models.NewOptional(models.ToPointer("5")),
        Band5:                models.ToPointer(models.StatsApRadioConfigBand{
            AllowRrmDisable:        models.NewOptional(models.ToPointer(false)),
            Bandwidth:              models.NewOptional(models.ToPointer(float64(218.56))),
            Channel:                models.ToPointer(132),
            Disabled:               models.NewOptional(models.ToPointer(false)),
            DynamicChainingEnabled: models.NewOptional(models.ToPointer(false)),
        }),
        Band6:                models.ToPointer(models.StatsApRadioConfigBand{
            AllowRrmDisable:        models.NewOptional(models.ToPointer(false)),
            Bandwidth:              models.NewOptional(models.ToPointer(float64(77.08))),
            Channel:                models.ToPointer(200),
            Disabled:               models.NewOptional(models.ToPointer(false)),
            DynamicChainingEnabled: models.NewOptional(models.ToPointer(false)),
        }),
        ScanningEnabled:      models.ToPointer(false),
    }

}
```

