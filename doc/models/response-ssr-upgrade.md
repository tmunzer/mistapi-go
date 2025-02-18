
# Response Ssr Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSsrUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Counts` | [`models.ResponseSsrUpgradeCounts`](../../doc/models/response-ssr-upgrade-counts.md) | Required | - |
| `DeviceType` | `string` | Required | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `Status` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Strategy` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Versions` | `map[string]string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": "channel0",
  "counts": {
    "failed": 166,
    "queued": 234,
    "success": 90,
    "upgrading": 112,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "device_type": "device_type6",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "status": "status6",
  "strategy": "strategy4",
  "versions": {
    "key0": "versions9"
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

