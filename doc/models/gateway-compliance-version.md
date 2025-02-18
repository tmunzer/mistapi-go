
# Gateway Compliance Version

Version compliance score, major version for gateway, type

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayComplianceVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorVersion` | [`map[string]models.GatewayComplianceMajorVersionProperties`](../../doc/models/gateway-compliance-major-version-properties.md) | Optional | - |
| `Score` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "score": 99.9,
  "type": "gateway",
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

