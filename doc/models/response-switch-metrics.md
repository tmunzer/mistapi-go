
# Response Switch Metrics

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSwitchMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ActivePortsSummary` | [`*models.ResponseSwitchMetricsActivePortsSummary`](../../doc/models/response-switch-metrics-active-ports-summary.md) | Optional | - |
| `ConfigSuccess` | [`*models.ResponseSwitchMetricsConfigSuccess`](../../doc/models/response-switch-metrics-config-success.md) | Optional | - |
| `VersionCompliance` | [`*models.ResponseSwitchMetricsVersionCompliance`](../../doc/models/response-switch-metrics-version-compliance.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "active_ports_summary": {
    "details": {
      "active_port_count": 136,
      "total_port_count": 42,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "score": 238,
    "total_switch_count": 26,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "config_success": {
    "details": {
      "config_success_count": 166,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "score": 52,
    "total_switch_count": 160,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "version_compliance": {
    "details": {
      "major_versions": [
        {
          "major_count": 170,
          "major_version": "major_version0",
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
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "score": 94,
    "total_switch_count": 138,
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

