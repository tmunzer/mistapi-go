
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
  "battery_voltage": 43.64,
  "eddystone_instance": "eddystone_instance0",
  "eddystone_namespace": "eddystone_namespace4",
  "last_seen": 119.08,
  "mac": "mac2",
  "map_id": "000007c2-0000-0000-0000-000000000000",
  "name": "name8",
  "power": 30,
  "type": "type8",
  "x": 119.24,
  "y": 250.52
}
```

