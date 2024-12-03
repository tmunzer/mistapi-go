
# Stats Beacon

*This model accepts additional fields of type interface{}.*

## Structure

`StatsBeacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatteryVoltage` | `*float64` | Optional | battery voltage, in mV |
| `EddystoneInstance` | `*string` | Optional | - |
| `EddystoneNamespace` | `*string` | Optional | - |
| `LastSeen` | `float64` | Required | - |
| `Mac` | `string` | Required | - |
| `MapId` | `uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `Power` | `int` | Required | - |
| `Type` | `string` | Required | - |
| `X` | `float64` | Required | - |
| `Y` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "battery_voltage": 33.46,
  "eddystone_instance": "eddystone_instance8",
  "eddystone_namespace": "eddystone_namespace4",
  "last_seen": 108.9,
  "mac": "mac4",
  "map_id": "0000106c-0000-0000-0000-000000000000",
  "name": "name0",
  "power": 220,
  "type": "type0",
  "x": 146.94,
  "y": 240.34,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

