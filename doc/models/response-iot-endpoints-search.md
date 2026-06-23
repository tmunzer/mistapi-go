
# Response Iot Endpoints Search

Time-bounded response for IoT endpoint search results

## Structure

`ResponseIotEndpointsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Epoch timestamp, in seconds, for the end of the IoT endpoint search window |
| `Results` | [`[]models.IotendpointStats`](../../doc/models/iotendpoint-stats.md) | Required | IoT endpoint statistics returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | Epoch timestamp, in seconds, for the start of the IoT endpoint search window |
| `Total` | `int` | Required | Number of IoT endpoint records matching the search filters |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseIotEndpointsSearch := models.ResponseIotEndpointsSearch{
        End:                  float64(48.6),
        Results:              []models.IotendpointStats{
            models.IotendpointStats{
                ApMac:                models.ToPointer("5c5b350e0001"),
                Id:                   models.ToPointer("63f9e299182b63f9"),
                Lqi:                  models.ToPointer(168),
                Mac:                  models.ToPointer("63f9e299182b63f9"),
                Mfg:                  models.ToPointer("Assa Abloy"),
                Model:                models.ToPointer("Assa Abloy"),
                Type:                 models.ToPointer("zigbee"),
            },
        },
        Start:                float64(4.66),
        Total:                244,
    }

}
```

