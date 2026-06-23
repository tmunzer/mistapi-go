
# Utils Send Ble Beacon

Request to transmit an arbitrary BLE beacon frame from selected APs

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsSendBleBeacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconFrame` | `*string` | Optional | Raw BLE beacon frame payload to transmit |
| `BeaconFreq` | `*int` | Optional | Transmission interval for the arbitrary BLE beacon frame |
| `Duration` | `*int` | Optional | Number of seconds to continue sending the BLE beacon frame<br><br>**Constraints**: `>= 1`, `<= 60` |
| `Macs` | `[]string` | Optional | AP MAC addresses of devices that should transmit the BLE beacon |
| `MapIds` | `[]string` | Optional | Map identifiers used to select APs that should transmit the BLE beacon |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsSendBleBeacon := models.UtilsSendBleBeacon{
        BeaconFrame:          models.ToPointer("68b329da9893e34099c7d8ad5cb9c940"),
        BeaconFreq:           models.ToPointer(100),
        Duration:             models.ToPointer(10),
        Macs:                 []string{
            "5c5b35584a6f",
            "5c5b350ea3b3",
        },
        MapIds:               []string{
            "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

