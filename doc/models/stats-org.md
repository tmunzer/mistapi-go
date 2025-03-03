
# Stats Org

Org statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmtemplateId` | `uuid.UUID` | Required | - |
| `AllowMist` | `bool` | Required | - |
| `CreatedTime` | `float64` | Required | When the object has been created, in epoch |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `float64` | Required | When the object has been modified for the last time, in epoch |
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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "alarmtemplate_id": "00001c2c-0000-0000-0000-000000000000",
  "allow_mist": false,
  "created_time": 91.72,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "modified_time": 243.24,
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
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
        "total": 12.38,
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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

