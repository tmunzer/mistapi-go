
# Ap Stats Ble Config

## Structure

`ApStatsBleConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconRate` | `*int` | Optional | - |
| `BeaconRateModel` | `*string` | Optional | - |
| `BeamDisabled` | `[]int` | Optional | - |
| `Power` | `*int` | Optional | - |
| `PowerMode` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "beacon_rate_model": "custom",
  "beam_disabled": [
    29,
    30,
    31
  ],
  "power_mode": "custom",
  "beacon_rate": 26,
  "power": 128
}
```

