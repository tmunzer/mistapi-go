
# Utils Send Ble Beacon

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsSendBleBeacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconFrame` | `*string` | Optional | - |
| `BeaconFreq` | `*int` | Optional | - |
| `Duration` | `*int` | Optional | **Constraints**: `>= 1`, `<= 60` |
| `Macs` | `[]string` | Optional | - |
| `MapIds` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "beacon_frame": "68b329da9893e34099c7d8ad5cb9c940",
  "beacon_freq": 100,
  "duration": 10,
  "macs": [
    "5c5b35584a6f",
    "5c5b350ea3b3"
  ],
  "map_ids": [
    "845a23bf-bed9-e43c-4c86-6fa474be7ae5"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

