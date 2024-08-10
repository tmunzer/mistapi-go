
# Stats Org

Org statistics

## Structure

`StatsOrg`

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
| `Sle` | [`[]models.StatsOrgSle`](../../doc/models/stats-org-sle.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "alarmtemplate_id": "00001c2c-0000-0000-0000-000000000000",
  "allow_mist": false,
  "created_time": 91.72,
  "id": "0000198e-0000-0000-0000-000000000000",
  "modified_time": 243.24,
  "msp_id": "0000178a-0000-0000-0000-000000000000",
  "name": "name2",
  "num_devices": 220,
  "num_devices_connected": 180,
  "num_devices_disconnected": 68,
  "num_inventory": 86,
  "num_sites": 54,
  "orggroup_ids": [
    "00001ba0-0000-0000-0000-000000000000",
    "00001ba1-0000-0000-0000-000000000000",
    "00001ba2-0000-0000-0000-000000000000"
  ],
  "session_expiry": 202,
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

