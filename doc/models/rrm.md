
# Rrm

RRM

## Structure

`Rrm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`map[string]models.RrmBand`](../../doc/models/rrm-band.md) | Required | proposal on band 2.4G, key is ap_id, value is the proposal |
| `Band24Metric` | [`models.RrmBandMetric`](../../doc/models/rrm-band-metric.md) | Required | - |
| `Band5` | [`map[string]models.RrmBand`](../../doc/models/rrm-band.md) | Required | proposal on band 5G, key is ap_id, value is the proposal |
| `Band5Metric` | [`models.RrmBandMetric`](../../doc/models/rrm-band-metric.md) | Required | - |
| `Band6` | [`map[string]models.RrmBand`](../../doc/models/rrm-band.md) | Optional | proposal on band 6G, key is ap_id, value is the proposal |
| `Band6Metric` | [`*models.RrmBandMetric`](../../doc/models/rrm-band-metric.md) | Optional | - |
| `Rftemplate` | [`models.RfTemplate`](../../doc/models/rf-template.md) | Required | RF Template |
| `RftemplateId` | `uuid.UUID` | Required | - |
| `RftemplateName` | `string` | Required | - |
| `Status` | [`models.RrmStatusEnum`](../../doc/models/rrm-status-enum.md) | Required | enum: `ready`, `unknown`, `updating` |
| `Timestamp` | `float64` | Required | Epoch (seconds) |

## Example (as JSON)

```json
{
  "band_24": {
    "key0": {
      "bandwidth": 160,
      "channel": 80,
      "curr_bandwidth": 80,
      "curr_channel": 200,
      "curr_power": 116
    },
    "key1": {
      "bandwidth": 160,
      "channel": 80,
      "curr_bandwidth": 80,
      "curr_channel": 200,
      "curr_power": 116
    },
    "key2": {
      "bandwidth": 160,
      "channel": 80,
      "curr_bandwidth": 80,
      "curr_channel": 200,
      "curr_power": 116
    }
  },
  "band_24_metric": {
    "cochannel_neighbors": 161.38,
    "density": 1.0,
    "interferences": {
      "149": {
        "radar": 0.3
      },
      "153": {
        "radar": 0.2
      }
    },
    "neighbors": 12.0,
    "noise": 252.7,
    "avg_aps_per_channel": 75.02,
    "channel_distribution_uniformity": 91.38,
    "naps_by_channel": {
      "key0": 11.57
    },
    "naps_by_power": {
      "key0": 128.78,
      "key1": 128.77
    }
  },
  "band_5": {
    "key0": {
      "bandwidth": 20,
      "channel": 132,
      "curr_bandwidth": 0,
      "curr_channel": 252,
      "curr_power": 64
    },
    "key1": {
      "bandwidth": 20,
      "channel": 132,
      "curr_bandwidth": 0,
      "curr_channel": 252,
      "curr_power": 64
    }
  },
  "band_5_metric": {
    "cochannel_neighbors": 78.16,
    "density": 1.0,
    "interferences": {
      "149": {
        "radar": 0.3
      },
      "153": {
        "radar": 0.2
      }
    },
    "neighbors": 184.78,
    "noise": 86.52,
    "avg_aps_per_channel": 158.24,
    "channel_distribution_uniformity": 174.6,
    "naps_by_channel": {
      "key0": 57.41,
      "key1": 57.42
    },
    "naps_by_power": {
      "key0": 96.48
    }
  },
  "rftemplate": {
    "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
    "name": "name6",
    "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
    "ant_gain_24": 220,
    "ant_gain_5": 132,
    "ant_gain_6": 248,
    "band_24": {
      "allow_rrm_disable": false,
      "ant_gain": 0,
      "antenna_mode": "1x1",
      "bandwidth": 0,
      "channels": [
        221
      ]
    },
    "band_24_usage": "24",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "rftemplate_id": "00000cf4-0000-0000-0000-000000000000",
  "rftemplate_name": "rftemplate_name8",
  "status": "ready",
  "timestamp": 36.02,
  "band_6": {
    "key0": {
      "bandwidth": 80,
      "channel": 200,
      "curr_bandwidth": 80,
      "curr_channel": 64,
      "curr_power": 252
    }
  },
  "band_6_metric": {
    "avg_aps_per_channel": 82.4,
    "channel_distribution_uniformity": 98.76,
    "cochannel_neighbors": 154.0,
    "density": 1.0,
    "interferences": {
      "key0": null
    },
    "naps_by_channel": {
      "key0": 18.43,
      "key1": 18.42
    },
    "naps_by_power": {
      "key0": 20.64
    },
    "neighbors": 4.62,
    "noise": 10.68
  }
}
```

