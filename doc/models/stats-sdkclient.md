
# Stats Sdkclient

SDK Client statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsSdkclient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `LastSeen` | `float64` | Required | Last seen timestamp |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | Map_id of the sdk client (if known), or null |
| `Name` | `*string` | Optional | Name of the sdk client (if provided) |
| `NetworkConnection` | [`models.StatsSdkclientNetworkConnection`](../../doc/models/stats-sdkclient-network-connection.md) | Required | Various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as |
| `Uuid` | `uuid.UUID` | Required | UUID of the sdk client |
| `X` | `*float64` | Optional | X (in pixels) of user location on the map (if known) |
| `Y` | `*float64` | Optional | Y (in pixels) of user location on the map (if known) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "last_seen": 1428939600.0,
  "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
  "name": "John's iPhone",
  "network_connection": {
    "mac": "mac2",
    "rssi": 235.8,
    "signal_level": 47.82,
    "type": "type2",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "uuid": "ada72f8f-1643-e5c6-94db-f2a5636f1a64",
  "x": 60.0,
  "y": 80.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

