
# Response Dswitches Metrics

## Structure

`ResponseDswitchesMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InactiveWiredVlans` | [`models.DswitchesMetricsInactiveWiredVlans`](../../doc/models/dswitches-metrics-inactive-wired-vlans.md) | Required | - |
| `PoeCompliance` | [`models.DswitchesMetricsPoeCompliance`](../../doc/models/dswitches-metrics-poe-compliance.md) | Required | - |
| `SwitchApAffinity` | [`models.DswitchesMetricsSwitchApAffinity`](../../doc/models/dswitches-metrics-switch-ap-affinity.md) | Required | - |
| `VersionCompliance` | [`models.DswitchesMetricsVersionCompliance`](../../doc/models/dswitches-metrics-version-compliance.md) | Required | - |

## Example (as JSON)

```json
{
  "inactive_wired_vlans": {
    "details": {
      "key1": "val1",
      "key2": "val2"
    },
    "score": 46.48
  },
  "poe_compliance": {
    "details": {
      "total_aps": 132,
      "total_power": 137.54
    },
    "score": 170.96
  },
  "switch_ap_affinity": {
    "details": {
      "system_name": [
        "system_name0",
        "system_name1"
      ],
      "threshold": 56.78
    },
    "score": 16.64
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
          ]
        }
      ],
      "total_switch_count": 14
    },
    "score": 149.42
  }
}
```

