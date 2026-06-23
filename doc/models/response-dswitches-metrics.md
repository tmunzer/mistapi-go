
# Response Dswitches Metrics

Metrics summary returned for discovered switch compliance checks

## Structure

`ResponseDswitchesMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InactiveWiredVlans` | [`models.DswitchesMetricsInactiveWiredVlans`](../../doc/models/dswitches-metrics-inactive-wired-vlans.md) | Required | Inactive wired VLAN metric for APs connected to discovered switches |
| `PoeCompliance` | [`models.DswitchesMetricsPoeCompliance`](../../doc/models/dswitches-metrics-poe-compliance.md) | Required | PoE compliance metric for APs connected to discovered switches |
| `SwitchApAffinity` | [`models.DswitchesMetricsSwitchApAffinity`](../../doc/models/dswitches-metrics-switch-ap-affinity.md) | Required | Switch/AP affinity metric for discovered switches |
| `VersionCompliance` | [`models.DswitchesMetricsVersionCompliance`](../../doc/models/dswitches-metrics-version-compliance.md) | Required | Version compliance metric for discovered switches |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseDswitchesMetrics := models.ResponseDswitchesMetrics{
        InactiveWiredVlans:   models.DswitchesMetricsInactiveWiredVlans{
            Details:              interface{}("[key1, val1][key2, val2]"),
            Score:                float64(46.48),
        },
        PoeCompliance:        models.DswitchesMetricsPoeCompliance{
            Details:              models.DswitchesMetricsPoeComplianceDetails{
                TotalAps:             132,
                TotalPower:           float64(137.54),
            },
            Score:                float64(170.96),
        },
        SwitchApAffinity:     models.DswitchesMetricsSwitchApAffinity{
            Details:              models.DswitchesMetricsSwitchApAffinityDetails{
                SystemName:           []string{
                    "system_name0",
                    "system_name1",
                },
                Threshold:            float64(56.78),
            },
            Score:                float64(16.64),
        },
        VersionCompliance:    models.DswitchesMetricsVersionCompliance{
            Details:              models.DswitchesMetricsVersionComplianceDetails{
                MajorVersions:        []models.DswitchesComplianceMajorVersion{
                    models.DswitchesComplianceMajorVersion{
                        MajorCount:           float64(47.78),
                        Model:                "model6",
                        SystemNames:          []string{
                            "system_names6",
                            "system_names7",
                        },
                    },
                },
                TotalSwitchCount:     14,
            },
            Score:                float64(149.42),
        },
    }

}
```

