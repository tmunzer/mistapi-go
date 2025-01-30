
# Upgrade Org Devices Upgrade Info

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeOrgDevicesUpgradeInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `StartTime` | `*int` | Optional | - |
| `Status` | [`*models.UpgradeDeviceStatusEnum`](../../doc/models/upgrade-device-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading` |
| `Targets` | [`*models.UpgradeDevicesTargetIds`](../../doc/models/upgrade-devices-target-ids.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "start_time": 1717658765,
  "status": "downloading",
  "targets": {
    "download_requested": [
      "download_requested6"
    ],
    "downloaded": [
      "downloaded0",
      "downloaded1",
      "downloaded2"
    ],
    "failed": [
      "failed6"
    ],
    "reboot_in_progress": [
      "reboot_in_progress3",
      "reboot_in_progress4"
    ],
    "rebooted": [
      "rebooted5"
    ],
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

