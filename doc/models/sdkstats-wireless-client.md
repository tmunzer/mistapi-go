
# Sdkstats Wireless Client

SDK Client Details statistics

*This model accepts additional fields of type interface{}.*

## Structure

`SdkstatsWirelessClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `LastSeen` | `float64` | Required | Last seen timestamp |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | Map_id of the sdk client (if known), or null |
| `Name` | `*string` | Optional | Name of the sdk client (if provided) |
| `NetworkConnection` | [`*models.StatsSdkclientNetworkConnection`](../../doc/models/stats-sdkclient-network-connection.md) | Optional | Various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as |
| `Uuid` | `uuid.UUID` | Required | UUID of the sdk client |
| `Vbeacons` | [`[]models.SdkstatsWirelessClientVbeacon`](../../doc/models/sdkstats-wireless-client-vbeacon.md) | Optional | List of beacon_id’s of the sdk client is in and since when (if known)<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `X` | `*float64` | Optional | X (in pixels) of user location on the map (if known) |
| `Y` | `*float64` | Optional | Y (in pixels) of user location on the map (if known) |
| `Zones` | [`[]models.SdkstatsWirelessClientZone`](../../doc/models/sdkstats-wireless-client-zone.md) | Optional | List of zone_id’s of the sdk client is in and since when (if known)<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "last_seen": 207.0,
  "uuid": "00000510-0000-0000-0000-000000000000",
  "map_id": "0000112a-0000-0000-0000-000000000000",
  "name": "name0",
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
  "vbeacons": [
    {
      "id": "00001cc0-0000-0000-0000-000000000000",
      "since": 115.2,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "x": 48.84,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

