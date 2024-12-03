
# Gateway Metrics

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigSuccess` | `*float64` | Optional | config success score |
| `VersionCompliance` | [`*models.GatewayComplianceVersion`](../../doc/models/gateway-compliance-version.md) | Optional | version compliance score, major version for gateway, type |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "config_success": 99.9,
  "version_compliance": {
    "major_version": {
      "key0": {
        "major_count": 80,
        "major_version": "major_version0",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key1": {
        "major_count": 80,
        "major_version": "major_version0",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key2": {
        "major_count": 80,
        "major_version": "major_version0",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    },
    "score": 149.42,
    "type": "type2",
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

