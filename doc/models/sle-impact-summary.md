
# Sle Impact Summary

SLE impact summary grouped by client, device, AP, WLAN, and band dimensions

## Structure

`SleImpactSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | [`[]models.SleImpactSummaryApItem`](../../doc/models/sle-impact-summary-ap-item.md) | Required | Impact summary rows grouped by AP<br><br>**Constraints**: *Unique Items Required* |
| `Band` | [`[]models.SleImpactSummaryBandItem`](../../doc/models/sle-impact-summary-band-item.md) | Required | Impact summary rows grouped by radio band<br><br>**Constraints**: *Unique Items Required* |
| `Classifier` | `string` | Required | Requested SLE classifier filter applied to the impact summary |
| `DeviceOs` | [`[]models.SleImpactSummaryDeviceOsItem`](../../doc/models/sle-impact-summary-device-os-item.md) | Required | Impact summary rows grouped by client device OS<br><br>**Constraints**: *Unique Items Required* |
| `DeviceType` | [`[]models.SleImpactSummaryDeviceTypeItem`](../../doc/models/sle-impact-summary-device-type-item.md) | Required | Impact summary rows grouped by client device type<br><br>**Constraints**: *Unique Items Required* |
| `End` | `float64` | Required | Last timestamp in the impact summary window |
| `Failure` | `string` | Required | Requested SLE failure filter applied to the impact summary |
| `Metric` | `string` | Required | SLE metric name summarized by this response<br><br>**Constraints**: *Minimum Length*: `1` |
| `Start` | `float64` | Required | First timestamp in the impact summary window |
| `Wlan` | [`[]models.SleImpactSummaryWlanItem`](../../doc/models/sle-impact-summary-wlan-item.md) | Required | Impact summary rows grouped by WLAN<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactSummary := models.SleImpactSummary{
        Ap:                   []models.SleImpactSummaryApItem{
            models.SleImpactSummaryApItem{
                ApMac:                "ap_mac6",
                Degraded:             float64(149.14),
                Duration:             float64(22.2),
                Name:                 "name4",
                Total:                float64(78.86),
            },
        },
        Band:                 []models.SleImpactSummaryBandItem{
            models.SleImpactSummaryBandItem{
                Band:                 "band4",
                Degraded:             float64(250.12),
                Duration:             float64(123.18),
                Name:                 "name2",
                Total:                float64(22.12),
            },
        },
        Classifier:           "classifier0",
        DeviceOs:             []models.SleImpactSummaryDeviceOsItem{
            models.SleImpactSummaryDeviceOsItem{
                Degraded:             float64(94.1),
                DeviceOs:             "device_os0",
                Duration:             float64(223.16),
                Name:                 "name0",
                Total:                float64(122.1),
            },
        },
        DeviceType:           []models.SleImpactSummaryDeviceTypeItem{
            models.SleImpactSummaryDeviceTypeItem{
                Degraded:             float64(131.9),
                DeviceType:           "device_type0",
                Duration:             float64(4.96),
                Name:                 "name0",
                Total:                float64(159.9),
            },
        },
        End:                  float64(192.82),
        Failure:              "failure8",
        Metric:               "metric6",
        Start:                float64(148.88),
        Wlan:                 []models.SleImpactSummaryWlanItem{
            models.SleImpactSummaryWlanItem{
                Degraded:             float64(169.9),
                Duration:             float64(42.96),
                Name:                 "name0",
                Total:                float64(197.9),
                WlanId:               "wlan_id2",
            },
        },
    }

}
```

