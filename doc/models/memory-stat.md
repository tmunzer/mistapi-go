
# Memory Stat

Memory utilization statistics for a device; in a virtual chassis, this reports the master Routing Engine

## Structure

`MemoryStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Usage` | `float64` | Required | Current memory utilization percentage for the device or master Routing Engine |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    memoryStat := models.MemoryStat{
        Usage:                float64(93.62),
    }

}
```

