
# Response Ssr Upgrade Status

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSsrUpgradeStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `DeviceType` | `*string` | Optional | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Status` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Targets` | [`models.ResponseSsrUpgradeStatusTargets`](../../doc/models/response-ssr-upgrade-status-targets.md) | Required | - |
| `Versions` | `interface{}` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": "channel8",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "status": "status8",
  "targets": {
    "failed": [
      "failed6"
    ],
    "queued": [
      "queued8"
    ],
    "success": [
      "success2",
      "success3",
      "success4"
    ],
    "upgrading": [
      "upgrading9"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "versions": {
    "key1": "val1",
    "key2": "val2"
  },
  "device_type": "device_type4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

