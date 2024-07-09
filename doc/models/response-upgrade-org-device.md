
# Response Upgrade Org Device

## Structure

`ResponseUpgradeOrgDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Upgrade` | [`*models.UpgradeOrgDeviceUpgrade`](../../doc/models/upgrade-org-device-upgrade.md) | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "upgrade": {
    "id": "000016ac-0000-0000-0000-000000000000",
    "start_time": 2.28,
    "status": "downloaded",
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
}
```

