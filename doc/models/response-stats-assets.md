
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
  "end": 170,
  "limit": 0,
  "results": [
    {
      "battery_voltage": 2970.0,
      "beam": 6,
      "by": "asset",
      "device_id": "00000000-0000-0000-1000-5c5b35000001",
      "device_name": "a",
      "duration": 120,
      "eddystone_uid_instance": "5c5b35000001",
      "eddystone_uid_namespace": "2818e3868dec25629ede",
      "eddystone_url_url": "https://www.abc.com",
      "ibeacon_major": 1234,
      "ibeacon_minor": 1234,
      "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "last_seen": 1470417522,
      "mac": "6fa474be7ae5",
      "map_id": "c45be59f-854d-4ef7-b782-dcd6309c84a9",
      "mfg_company_id": 935,
      "mfg_data": "648520a1020000",
      "name": "6fa474be7ae5",
      "rssi": -60,
      "temperature": 23,
      "x": 280.19918140310193,
      "y": 420.2987721046529,
      "_ttl": 162
    }
  ],
  "start": 128,
  "total": 162,
  "next": "next8"
}
```

