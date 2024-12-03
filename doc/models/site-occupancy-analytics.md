
# Site Occupancy Analytics

Occupancy Analytics settings

*This model accepts additional fields of type interface{}.*

## Structure

`SiteOccupancyAnalytics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetsEnabled` | `*bool` | Optional | indicate whether named BLE assets should be included in the zone occupancy calculation<br>**Default**: `false` |
| `ClientsEnabled` | `*bool` | Optional | indicate whether connected WiFi clients should be included in the zone occupancy calculation<br>**Default**: `true` |
| `MinDuration` | `*int` | Optional | minimum duration<br>**Default**: `3000` |
| `SdkclientsEnabled` | `*bool` | Optional | indicate whether SDK clients should be included in the zone occupancy calculation<br>**Default**: `false` |
| `UnconnectedClientsEnabled` | `*bool` | Optional | indicate whether unconnected WiFi clients should be included in the zone occupancy calculation<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "assets_enabled": false,
  "clients_enabled": true,
  "min_duration": 3000,
  "sdkclients_enabled": false,
  "unconnected_clients_enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

