
# Stats Ap Zigbee Stat

ZigBee statistics reported by an AP, present only when ZigBee is enabled on the AP

## Structure

`StatsApZigbeeStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IotproxyStatus` | `models.Optional[string]` | Optional, Read-only | Connection status of the IoT proxy |
| `NumIotendpoints` | `models.Optional[int]` | Optional, Read-only | Number of IoT endpoints connected through the AP |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApZigbeeStat := models.StatsApZigbeeStat{
        IotproxyStatus:       models.NewOptional(models.ToPointer("connected")),
        NumIotendpoints:      models.NewOptional(models.ToPointer(3)),
    }

}
```

