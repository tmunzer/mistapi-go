
# Org Stats

Org statistics

## Structure

`OrgStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmtemplateId` | `uuid.UUID` | Required | - |
| `AllowMist` | `bool` | Required | - |
| `CreatedTime` | `float64` | Required | - |
| `Id` | `uuid.UUID` | Required | - |
| `ModifiedTime` | `float64` | Required | - |
| `MspId` | `uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `NumDevices` | `int` | Required | - |
| `NumDevicesConnected` | `int` | Required | - |
| `NumDevicesDisconnected` | `int` | Required | - |
| `NumInventory` | `int` | Required | - |
| `NumSites` | `int` | Required | - |
| `OrggroupIds` | `[]uuid.UUID` | Required | - |
| `SessionExpiry` | `int64` | Required | - |
| `Sle` | [`[]models.OrgSleStat`](../../doc/models/org-sle-stat.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "alarmtemplate_id": "00000fca-0000-0000-0000-000000000000",
  "allow_mist": false,
  "created_time": 116.02,
  "id": "00000d2c-0000-0000-0000-000000000000",
  "modified_time": 218.94,
  "msp_id": "00000b28-0000-0000-0000-000000000000",
  "name": "name2",
  "num_devices": 90,
  "num_devices_connected": 54,
  "num_devices_disconnected": 194,
  "num_inventory": 216,
  "num_sites": 180,
  "orggroup_ids": [
    "00000f3e-0000-0000-0000-000000000000",
    "00000f3f-0000-0000-0000-000000000000"
  ],
  "session_expiry": 76,
  "sle": [
    {
      "path": "path2",
      "user_minutes": {
        "ok": 13.84,
        "total": 12.38
      }
    }
  ]
}
```

