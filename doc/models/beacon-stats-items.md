
# Beacon Stats Items

## Structure

`BeaconStatsItems`

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

## Example (as JSON)

```json
{
  "battery_voltage": 72.28,
  "eddystone_instance": "eddystone_instance4",
  "eddystone_namespace": "eddystone_namespace8",
  "last_seen": 147.72,
  "mac": "mac6",
  "map_id": "00001272-0000-0000-0000-000000000000",
  "name": "name2",
  "power": 78,
  "type": "type8",
  "x": 147.88,
  "y": 23.16
}
```

