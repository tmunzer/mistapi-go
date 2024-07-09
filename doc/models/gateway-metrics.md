
# Gateway Metrics

## Structure

`GatewayMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigSuccess` | `*float64` | Optional | config success score |
| `VersionCompliance` | [`*models.GatewayComplianceVersion`](../../doc/models/gateway-compliance-version.md) | Optional | version compliance score, major version for gateway, type |

## Example (as JSON)

```json
{
  "config_success": 99.9,
  "version_compliance": {
    "major_version": {
      "key0": {
        "major_count": 80,
        "major_version": "major_version0"
      },
      "key1": {
        "major_count": 80,
        "major_version": "major_version0"
      },
      "key2": {
        "major_count": 80,
        "major_version": "major_version0"
      }
    },
    "score": 149.42,
    "type": "type2"
  }
}
```

