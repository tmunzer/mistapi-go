
# Gateway Compliance Version

version compliance score, major version for gateway, type

## Structure

`GatewayComplianceVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorVersion` | [`map[string]models.GatewayComplianceMajorVersionProperties`](../../doc/models/gateway-compliance-major-version-properties.md) | Optional | - |
| `Score` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "score": 99.9,
  "type": "gateway",
  "major_version": {
    "key0": {
      "major_count": 80,
      "major_version": "major_version0"
    }
  }
}
```

