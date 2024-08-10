
# Sdkstats Wireless Client

SDK Client Details statistics

## Structure

`SdkstatsWirelessClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required | - |
| `LastSeen` | `float64` | Required | last seen timestamp |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | map_id of the sdk client (if known), or null |
| `Name` | `*string` | Optional | name of the sdk client (if provided) |
| `NetworkConnection` | [`*models.StatsSdkclientNetworkConnection`](../../doc/models/stats-sdkclient-network-connection.md) | Optional | various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as |
| `Uuid` | `uuid.UUID` | Required | uuid of the sdk client |
| `Vbeacons` | [`[]models.SdkstatsWirelessClientVbeacon`](../../doc/models/sdkstats-wireless-client-vbeacon.md) | Optional | list of beacon_id’s of the sdk client is in and since when (if known)<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `X` | `*float64` | Optional | x (in pixels) of user location on the map (if known) |
| `Y` | `*float64` | Optional | y (in pixels) of user location on the map (if known) |
| `Zones` | [`[]models.SdkstatsWirelessClientZone`](../../doc/models/sdkstats-wireless-client-zone.md) | Optional | list of zone_id’s of the sdk client is in and since when (if known)<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "id": "00000d02-0000-0000-0000-000000000000",
  "last_seen": 207.0,
  "map_id": "0000112a-0000-0000-0000-000000000000",
  "name": "name0",
  "network_connection": {
    "mac": "mac2",
    "rssi": 235.8,
    "signal_level": 47.82,
    "type": "type2"
  },
  "uuid": "00000510-0000-0000-0000-000000000000",
  "vbeacons": [
    {
      "id": "00001cc0-0000-0000-0000-000000000000",
      "since": 115.2
    }
  ],
  "x": 48.84
}
```

