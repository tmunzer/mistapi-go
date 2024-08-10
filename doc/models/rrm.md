
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
| `Timestamp` | `float64` | Required | time where the status was updated |

## Example (as JSON)

```json
{
  "band_24": {
    "key0": {
      "bandwidth": 20,
      "channel": 80,
      "curr_bandwidht": 80,
      "curr_channel": 200,
      "curr_power": 116
    },
    "key1": {
      "bandwidth": 20,
      "channel": 80,
      "curr_bandwidht": 80,
      "curr_channel": 200,
      "curr_power": 116
    },
    "key2": {
      "bandwidth": 20,
      "channel": 80,
      "curr_bandwidht": 80,
      "curr_channel": 200,
      "curr_power": 116
    }
  },
  "band_24_metric": {
    "cochannel_neighbors": 161.38,
    "density": 107.0,
    "interferences": {
      "149": {
        "radar": 0.3
      },
      "153": {
        "radar": 0.2
      }
    },
    "neighbors": 12.0,
    "noise": 252.7
  },
  "band_5": {
    "key0": {
      "bandwidth": 20,
      "channel": 132,
      "curr_bandwidht": 80,
      "curr_channel": 252,
      "curr_power": 64
    },
    "key1": {
      "bandwidth": 20,
      "channel": 132,
      "curr_bandwidht": 80,
      "curr_channel": 252,
      "curr_power": 64
    }
  },
  "band_5_metric": {
    "cochannel_neighbors": 78.16,
    "density": 23.78,
    "interferences": {
      "149": {
        "radar": 0.3
      },
      "153": {
        "radar": 0.2
      }
    },
    "neighbors": 184.78,
    "noise": 86.52
  },
  "rftemplate": {
    "name": "name6",
    "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
    "ant_gain_24": 220,
    "ant_gain_5": 132,
    "ant_gain_6": 248,
    "band_24": {
      "allow_rrm_disable": false,
      "ant_gain": 0,
      "antenna_mode": "1x1",
      "bandwidth": 20,
      "channels": [
        221
      ]
    },
    "band_24_usage": "24"
  },
  "rftemplate_id": "00000cf4-0000-0000-0000-000000000000",
  "rftemplate_name": "rftemplate_name8",
  "status": "ready",
  "timestamp": 36.02,
  "band_6": {
    "key0": {
      "bandwidth": 20,
      "channel": 200,
      "curr_bandwidht": 80,
      "curr_channel": 64,
      "curr_power": 252
    }
  },
  "band_6_metric": {
    "cochannel_neighbors": 154.0,
    "density": 99.62,
    "interferences": {
      "key0": {
        "radar": 1.76
      }
    },
    "neighbors": 4.62,
    "noise": 10.68
  }
}
```

