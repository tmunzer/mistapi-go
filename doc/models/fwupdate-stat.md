
# Fwupdate Stat

Firmware update status for a device

## Structure

`FwupdateStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Progress` | `models.Optional[int]` | Optional, Read-only | Firmware update progress percentage, or null when unavailable<br><br>**Constraints**: `>= 0`, `<= 100` |
| `Status` | [`models.Optional[models.FwupdateStatStatusEnum]`](../../doc/models/fwupdate-stat-status-enum.md) | Optional, Read-only | enum: `inprogress`, `failed`, `upgraded`, `success`, `scheduled`, `error` |
| `StatusId` | `models.Optional[int]` | Optional, Read-only | Numeric firmware update status identifier |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `WillRetry` | `models.Optional[bool]` | Optional, Read-only | Whether the firmware update process will retry after the current status |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    fwupdateStat := models.FwupdateStat{
        Progress:             models.NewOptional(models.ToPointer(10)),
        Status:               models.NewOptional(models.ToPointer(models.FwupdateStatStatusEnum_INPROGRESS)),
        StatusId:             models.NewOptional(models.ToPointer(5)),
        Timestamp:            models.ToPointer(float64(115.22)),
        WillRetry:            models.NewOptional(models.ToPointer(false)),
    }

}
```

