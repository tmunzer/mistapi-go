
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
  "alarmtemplate_id": "000009f4-0000-0000-0000-000000000000",
  "allow_mist": false,
  "created_time": 45.08,
  "id": "00000756-0000-0000-0000-000000000000",
  "modified_time": 33.88,
  "msp_id": "00000552-0000-0000-0000-000000000000",
  "name": "name8",
  "num_devices": 164,
  "num_devices_connected": 236,
  "num_devices_disconnected": 12,
  "num_inventory": 142,
  "num_sites": 254,
  "orggroup_ids": [
    "00000968-0000-0000-0000-000000000000"
  ],
  "session_expiry": 2,
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

