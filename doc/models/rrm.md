
# Rrm

Current RRM channel-planning state for a site

## Structure

`Rrm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`map[string]models.RrmBand`](../../doc/models/rrm-band.md) | Required | RRM proposed channel, power, and usage settings for 2.4 GHz radios; property key is AP ID |
| `Band24Metric` | [`models.RrmBandMetric`](../../doc/models/rrm-band-metric.md) | Required | Aggregate RRM metrics for a radio band |
| `Band5` | [`map[string]models.RrmBand`](../../doc/models/rrm-band.md) | Required | RRM proposed channel, power, and usage settings for 5 GHz radios; property key is AP ID |
| `Band5Metric` | [`models.RrmBandMetric`](../../doc/models/rrm-band-metric.md) | Required | Aggregate RRM metrics for a radio band |
| `Band6` | [`map[string]models.RrmBand`](../../doc/models/rrm-band.md) | Optional | RRM proposed channel, power, and usage settings for 6 GHz radios; property key is AP ID |
| `Band6Metric` | [`*models.RrmBandMetric`](../../doc/models/rrm-band-metric.md) | Optional | Aggregate RRM metrics for a radio band |
| `Rftemplate` | [`models.RfTemplate`](../../doc/models/rf-template.md) | Required | RF template used by the current RRM calculation |
| `RftemplateId` | `uuid.UUID` | Required | RF template identifier used by the current RRM calculation |
| `RftemplateName` | `string` | Required | RF template name used by the current RRM calculation |
| `Status` | [`models.RrmStatusEnum`](../../doc/models/rrm-status-enum.md) | Required | enum: `ready`, `unknown`, `updating` |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    rrm := models.Rrm{
        Band24:               map[string]models.RrmBand{
            "key0": models.RrmBand{
                Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM160),
                Channel:              models.ToPointer(80),
                CurrBandwidth:        models.ToPointer(models.Dot11BandwidthEnum_ENUM80),
                CurrChannel:          models.ToPointer(200),
                CurrPower:            models.ToPointer(116),
            },
            "key1": models.RrmBand{
                Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM160),
                Channel:              models.ToPointer(80),
                CurrBandwidth:        models.ToPointer(models.Dot11BandwidthEnum_ENUM80),
                CurrChannel:          models.ToPointer(200),
                CurrPower:            models.ToPointer(116),
            },
            "key2": models.RrmBand{
                Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM160),
                Channel:              models.ToPointer(80),
                CurrBandwidth:        models.ToPointer(models.Dot11BandwidthEnum_ENUM80),
                CurrChannel:          models.ToPointer(200),
                CurrPower:            models.ToPointer(116),
            },
        },
        Band24Metric:         models.RrmBandMetric{
            AvgApsPerChannel:              models.ToPointer(float64(75.02)),
            ChannelDistributionUniformity: models.ToPointer(float64(91.38)),
            CochannelNeighbors:            float64(161.38),
            Density:                       float64(1),
            Interferences:                 map[string]models.RrmBandMetricInterference{
                "149": models.RrmBandMetricInterference{
                    Radar:                models.ToPointer(float64(0.3)),
                },
                "153": models.RrmBandMetricInterference{
                    Radar:                models.ToPointer(float64(0.2)),
                },
            },
            NapsByChannel:                 map[string]float64{
                "key0": float64(11.57),
            },
            NapsByPower:                   map[string]float64{
                "key0": float64(128.78),
                "key1": float64(128.77),
            },
            Neighbors:                     float64(12),
            Noise:                         float64(252.7),
        },
        Band5:                map[string]models.RrmBand{
            "key0": models.RrmBand{
                Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
                Channel:              models.ToPointer(132),
                CurrBandwidth:        models.ToPointer(models.Dot11BandwidthEnum_ENUM0),
                CurrChannel:          models.ToPointer(252),
                CurrPower:            models.ToPointer(64),
            },
            "key1": models.RrmBand{
                Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
                Channel:              models.ToPointer(132),
                CurrBandwidth:        models.ToPointer(models.Dot11BandwidthEnum_ENUM0),
                CurrChannel:          models.ToPointer(252),
                CurrPower:            models.ToPointer(64),
            },
        },
        Band5Metric:          models.RrmBandMetric{
            AvgApsPerChannel:              models.ToPointer(float64(158.24)),
            ChannelDistributionUniformity: models.ToPointer(float64(174.6)),
            CochannelNeighbors:            float64(78.16),
            Density:                       float64(1),
            Interferences:                 map[string]models.RrmBandMetricInterference{
                "149": models.RrmBandMetricInterference{
                    Radar:                models.ToPointer(float64(0.3)),
                },
                "153": models.RrmBandMetricInterference{
                    Radar:                models.ToPointer(float64(0.2)),
                },
            },
            NapsByChannel:                 map[string]float64{
                "key0": float64(57.41),
                "key1": float64(57.42),
            },
            NapsByPower:                   map[string]float64{
                "key0": float64(96.48),
            },
            Neighbors:                     float64(184.78),
            Noise:                         float64(86.52),
        },
        Band6:                map[string]models.RrmBand{
            "key0": models.RrmBand{
                Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM80),
                Channel:              models.ToPointer(200),
                CurrBandwidth:        models.ToPointer(models.Dot11BandwidthEnum_ENUM80),
                CurrChannel:          models.ToPointer(64),
                CurrPower:            models.ToPointer(252),
            },
        },
        Band6Metric:          models.ToPointer(models.RrmBandMetric{
            AvgApsPerChannel:              models.ToPointer(float64(82.4)),
            ChannelDistributionUniformity: models.ToPointer(float64(98.76)),
            CochannelNeighbors:            float64(154),
            Density:                       float64(1),
            Interferences:                 map[string]models.RrmBandMetricInterference{
                "key0": nil,
            },
            NapsByChannel:                 map[string]float64{
                "key0": float64(18.43),
                "key1": float64(18.42),
            },
            NapsByPower:                   map[string]float64{
                "key0": float64(20.64),
            },
            Neighbors:                     float64(4.62),
            Noise:                         float64(10.68),
        }),
        Rftemplate:           models.RfTemplate{
            AntGain24:            models.ToPointer(220),
            AntGain5:             models.ToPointer(132),
            AntGain6:             models.ToPointer(248),
            Band24:               models.ToPointer(models.RftemplateRadioBand24{
                AllowRrmDisable:      models.ToPointer(false),
                AntGain:              models.NewOptional(models.ToPointer(0)),
                AntennaMode:          models.ToPointer(models.RadioBandAntennaModeEnum_ENUM1X1),
                Bandwidth:            models.ToPointer(models.Dot11Bandwidth24Enum_ENUM0),
                Channels:             models.NewOptional(models.ToPointer([]int{
                    221,
                })),
            }),
            Band24Usage:          models.ToPointer(models.RadioBand24UsageEnum_ENUM24),
            Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
            Name:                 "name6",
            OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
            AdditionalProperties: map[string]interface{}{
                "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
            },
        },
        RftemplateId:         uuid.MustParse("00000cf4-0000-0000-0000-000000000000"),
        RftemplateName:       "rftemplate_name8",
        Status:               models.RrmStatusEnum_READY,
        Timestamp:            float64(36.02),
    }

}
```

