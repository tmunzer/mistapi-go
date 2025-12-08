
# Site Occupancy Analytics

Occupancy Analytics settings

## Structure

`SiteOccupancyAnalytics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetsEnabled` | `*bool` | Optional | Indicate whether named BLE assets should be included in the zone occupancy calculation<br><br>**Default**: `false` |
| `ClientsEnabled` | `*bool` | Optional | Indicate whether connected Wi-Fi clients should be included in the zone occupancy calculation<br><br>**Default**: `true` |
| `MinDuration` | `*int` | Optional | Minimum duration<br><br>**Default**: `3000` |
| `SdkclientsEnabled` | `*bool` | Optional | Indicate whether SDK clients should be included in the zone occupancy calculation<br><br>**Default**: `false` |
| `UnconnectedClientsEnabled` | `*bool` | Optional | Indicate whether unconnected Wi-Fi clients should be included in the zone occupancy calculation<br><br>**Default**: `false` |

## Example (as JSON)

```json
{
  "assets_enabled": false,
  "clients_enabled": true,
  "min_duration": 3000,
  "sdkclients_enabled": false,
  "unconnected_clients_enabled": false
}
```

