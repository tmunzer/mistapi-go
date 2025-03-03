
# Stats Unconnected Client

Unconnected clients statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsUnconnectedClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | MAC address of the AP that heard the client |
| `LastSeen` | `float64` | Required | Last seen timestamp |
| `Mac` | `string` | Required | MAC address of the (unconnected) client |
| `Manufacture` | `string` | Required | Device manufacture, through fingerprinting or OUI |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | Map_id of the client (if known), or null |
| `Rssi` | `int` | Required | Client RSSI observed by the AP that heard the client (in dBm) |
| `X` | `*float64` | Optional | X (in pixels) of user location on the map (if known) |
| `Y` | `float64` | Required | Y (in pixels) of user location on the map (if known) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_mac": "ap_mac4",
  "last_seen": 109.82,
  "mac": "mac6",
  "manufacture": "manufacture6",
  "map_id": "00000b60-0000-0000-0000-000000000000",
  "rssi": 166,
  "x": 109.98,
  "y": 241.26,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

