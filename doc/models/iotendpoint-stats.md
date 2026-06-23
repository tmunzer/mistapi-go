
# Iotendpoint Stats

IoT endpoint statistics returned by a search response

## Structure

`IotendpointStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | MAC address of the AP the endpoint was seen on |
| `Id` | `*string` | Optional | Unique identifier for the IoT endpoint |
| `Lqi` | `*int` | Optional | Link Quality Indicator (0–255)<br><br>**Constraints**: `>= 0`, `<= 255` |
| `Mac` | `*string` | Optional | Endpoint MAC address reported in IoT statistics |
| `Mfg` | `*string` | Optional | Manufacturer name reported for the IoT endpoint |
| `Model` | `*string` | Optional | Device model reported for the IoT endpoint |
| `Timestamp` | `*float64` | Optional | Epoch timestamp of the last observation, in seconds |
| `Type` | `*string` | Optional | IoT endpoint type. enum: `zigbee` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    iotendpointStats := models.IotendpointStats{
        ApMac:                models.ToPointer("5c5b350e0001"),
        Id:                   models.ToPointer("63f9e299182b63f9"),
        Lqi:                  models.ToPointer(102),
        Mac:                  models.ToPointer("63f9e299182b63f9"),
        Mfg:                  models.ToPointer("Assa Abloy"),
        Model:                models.ToPointer("Assa Abloy"),
        Type:                 models.ToPointer("zigbee"),
    }

}
```

