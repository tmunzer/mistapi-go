
# Upgrade Org Devices Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeOrgDevicesUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Upgrade` | [`*models.UpgradeOrgDevicesUpgradeInfo`](../../doc/models/upgrade-org-devices-upgrade-info.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "upgrade": {
    "id": "000016ac-0000-0000-0000-000000000000",
    "start_time": 228,
    "status": "cancelled",
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
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

