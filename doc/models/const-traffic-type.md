
# Const Traffic Type

Traffic type definition returned by the constants API

## Structure

`ConstTrafficType`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `*string` | Optional | Human-readable traffic type label |
| `Dscp` | `*int` | Optional | Differentiated Services Code Point value used for this traffic type |
| `FailoverPolicy` | `*string` | Optional | Failover behavior associated with this traffic type |
| `MaxJitter` | `*int` | Optional | Maximum jitter threshold for this traffic type, in milliseconds |
| `MaxLatency` | `*int` | Optional | Maximum latency threshold for this traffic type, in milliseconds |
| `MaxLoss` | `*int` | Optional | Maximum packet loss threshold for this traffic type, in percent |
| `Name` | `*string` | Optional | Machine-readable traffic type name |
| `TrafficClass` | `*string` | Optional | Traffic class associated with this traffic type |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constTrafficType := models.ConstTrafficType{
        Display:              models.ToPointer("VoIP Video"),
        Dscp:                 models.ToPointer(32),
        FailoverPolicy:       models.ToPointer("non_revertible"),
        MaxJitter:            models.ToPointer(250),
        MaxLatency:           models.ToPointer(1500),
        MaxLoss:              models.ToPointer(35),
        Name:                 models.ToPointer("voip_video"),
        TrafficClass:         models.ToPointer("medium"),
    }

}
```

