
# Upgrade Org Devices Upgrade Info

## Structure

`UpgradeOrgDevicesUpgradeInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `StartTime` | `*int` | Optional | - |
| `Status` | [`*models.UpgradeDeviceStatusEnum`](../../doc/models/upgrade-device-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`, `queued` |
| `Targets` | [`*models.UpgradeDevicesTargets`](../../doc/models/upgrade-devices-targets.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "start_time": 1717658765,
  "status": "queued",
  "targets": {
    "download_requested": [
      "download_requested6"
    ],
    "downloaded": [
      "downloaded0",
      "downloaded1",
      "downloaded2"
    ],
    "downloading": [
      "downloading6"
    ],
    "failed": [
      "failed6"
    ],
    "reboot_in_progress": [
      "reboot_in_progress3",
      "reboot_in_progress4"
    ]
  }
}
```

