
# Response Stats Assets

## Structure

`ResponseStatsAssets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.StatsAsset`](../../doc/models/stats-asset.md) | Required | - |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 236,
  "limit": 190,
  "results": [
    {
      "battery_voltage": 2970.0,
      "beam": 6,
      "device_name": "a",
      "duration": 120,
      "eddystone_uid_instance": "5c5b35000001",
      "eddystone_uid_namespace": "2818e3868dec25629ede",
      "eddystone_url_url": "ttps://www.abc.com",
      "ibeacon_major": 12,
      "ibeacon_minor": 138,
      "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
      "last_seen": 1428939600,
      "mac": "6fa474be7ae5",
      "map_id": "c45be59f-854d-4ef7-b782-dcd6309c84a9",
      "name": "6fa474be7ae5",
      "rssi": -60,
      "temperatur": 23,
      "x": 280.19918140310193,
      "y": 420.2987721046529
    }
  ],
  "start": 194,
  "total": 28,
  "next": "next6"
}
```
