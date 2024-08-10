
# Unconnected Client Stat

Unconnected clients statistics

## Structure

`UnconnectedClientStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | mac address of the AP that heard the client |
| `LastSeen` | `float64` | Required | last seen timestamp |
| `Mac` | `string` | Required | mac address of the (unconnected) client |
| `Manufacture` | `string` | Required | device manufacture, through fingerprinting or OUI |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | map_id of the client (if known), or null |
| `Rssi` | `int` | Required | client RSSI observered by the AP that heard the client (in dBm) |
| `X` | `*float64` | Optional | x (in pixels) of user location on the map (if known) |
| `Y` | `float64` | Required | y (in pixels) of user location on the map (if known) |

## Example (as JSON)

```json
{
  "ap_mac": "ap_mac0",
  "last_seen": 154.38,
  "mac": "mac2",
  "manufacture": "manufacture2",
  "map_id": "000025b8-0000-0000-0000-000000000000",
  "rssi": 14,
  "x": 101.46,
  "y": 29.82
}
```

