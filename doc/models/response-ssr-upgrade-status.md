
# Response Ssr Upgrade Status

## Structure

`ResponseSsrUpgradeStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `DeviceType` | `*string` | Optional | - |
| `Id` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Status` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Targets` | [`models.ResponseSsrUpgradeStatusTargets`](../../doc/models/response-ssr-upgrade-status-targets.md) | Required | - |
| `Versions` | `interface{}` | Required | - |

## Example (as JSON)

```json
{
  "channel": "channel8",
  "device_type": "device_type4",
  "id": "id6",
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
    ]
  },
  "versions": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

