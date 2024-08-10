
# Stats Sdkclient

SDK Client statistics

## Structure

`StatsSdkclient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required | - |
| `LastSeen` | `float64` | Required | last seen timestamp |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | map_id of the sdk client (if known), or null |
| `Name` | `*string` | Optional | name of the sdk client (if provided) |
| `NetworkConnection` | [`models.StatsSdkclientNetworkConnection`](../../doc/models/stats-sdkclient-network-connection.md) | Required | various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as |
| `Uuid` | `uuid.UUID` | Required | uuid of the sdk client |
| `X` | `*float64` | Optional | x (in pixels) of user location on the map (if known) |
| `Y` | `*float64` | Optional | y (in pixels) of user location on the map (if known) |

## Example (as JSON)

```json
{
  "id": "de87bf9d-183f-e383-cc68-6ba43947d403",
  "last_seen": 1428939600.0,
  "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
  "name": "John's iPhone",
  "network_connection": {
    "mac": "mac2",
    "rssi": 235.8,
    "signal_level": 47.82,
    "type": "type2"
  },
  "uuid": "ada72f8f-1643-e5c6-94db-f2a5636f1a64",
  "x": 60.0,
  "y": 80.0
}
```

