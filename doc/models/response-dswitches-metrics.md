
# Response Dswitches Metrics

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDswitchesMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InactiveWiredVlans` | [`models.DswitchesMetricsInactiveWiredVlans`](../../doc/models/dswitches-metrics-inactive-wired-vlans.md) | Required | - |
| `PoeCompliance` | [`models.DswitchesMetricsPoeCompliance`](../../doc/models/dswitches-metrics-poe-compliance.md) | Required | - |
| `SwitchApAffinity` | [`models.DswitchesMetricsSwitchApAffinity`](../../doc/models/dswitches-metrics-switch-ap-affinity.md) | Required | - |
| `VersionCompliance` | [`models.DswitchesMetricsVersionCompliance`](../../doc/models/dswitches-metrics-version-compliance.md) | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "inactive_wired_vlans": {
    "details": {
      "key1": "val1",
      "key2": "val2"
    },
    "score": 46.48,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "poe_compliance": {
    "details": {
      "total_aps": 132,
      "total_power": 137.54,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "score": 170.96,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "switch_ap_affinity": {
    "details": {
      "system_name": [
        "system_name0",
        "system_name1"
      ],
      "threshold": 56.78,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "score": 16.64,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "version_compliance": {
    "details": {
      "major_versions": [
        {
          "major_count": 47.78,
          "model": "model6",
          "system_names": [
            "system_names6",
            "system_names7"
          ],
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "total_switch_count": 14,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "score": 149.42,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

