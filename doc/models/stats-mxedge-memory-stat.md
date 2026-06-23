
# Stats Mxedge Memory Stat

Memory usage counters reported by a Mist Edge

## Structure

`StatsMxedgeMemoryStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Active` | `*int` | Optional | The amount of memory, in kilobytes, that has been used more recently and is usually not reclaimed unless absolutely necessary. |
| `Available` | `*int64` | Optional | An estimate of how much memory is available for starting new applications, without swapping. |
| `Buffers` | `*int` | Optional | The amount, in kilobytes, of temporary storage for raw disk blocks. |
| `Cached` | `*int` | Optional | The amount of physical RAM, in kilobytes, used as cache memory. |
| `Free` | `*int64` | Optional | The amount of physical RAM, in kilobytes, left unused by the system |
| `Inactive` | `*int` | Optional | The amount of memory, in kilobytes, that has been used less recently and is more eligible to be reclaimed for other purposes. |
| `SwapCached` | `*int` | Optional | The amount of memory, in kilobytes, that has once been moved into swap, then back into the main memory, but still also remains in the swapfile. |
| `SwapFree` | `*int` | Optional | The total amount of swap free, in kilobytes. |
| `SwapTotal` | `*int` | Optional | The total amount of swap available, in kilobytes. |
| `Total` | `*int64` | Optional | Usable RAM total, in kilobytes, which is physical RAM minus a number of reserved bits and the kernel binary code |
| `Usage` | `*int` | Optional | Memory utilization percentage reported by the Mist Edge |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeMemoryStat := models.StatsMxedgeMemoryStat{
        Active:               models.ToPointer(394936320),
        Available:            models.ToPointer(int64(4699291648)),
        Buffers:              models.ToPointer(107646976),
        Cached:               models.ToPointer(478060544),
        Free:                 models.ToPointer(int64(4330659840)),
        Inactive:             models.ToPointer(211980288),
        SwapCached:           models.ToPointer(0),
        SwapFree:             models.ToPointer(1022357504),
        SwapTotal:            models.ToPointer(1022357504),
        Total:                models.ToPointer(int64(8365957120)),
        Usage:                models.ToPointer(48),
    }

}
```

