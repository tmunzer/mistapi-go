
# Upgrade Org Device Upgrade

## Structure

`UpgradeOrgDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | - |
| `StartTime` | `*float64` | Optional | - |
| `Status` | [`*models.DeviceUpgradeStatusEnum`](../../doc/models/device-upgrade-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading` |
| `Targets` | [`*models.UpgradeOrgDeviceTargets`](../../doc/models/upgrade-org-device-targets.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "473f6eca-6276-4993-bfeb-53cbbbba6f18",
  "start_time": 1717658765.0,
  "status": "created",
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
    ]
  }
}
```

