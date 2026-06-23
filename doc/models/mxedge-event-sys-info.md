
# Mxedge Event Sys Info

System resource details for a Mist Edge event

## Structure

`MxedgeEventSysInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Resource` | `*string` | Optional | System resource referenced by the event |
| `Severity` | [`*models.EventSeverityEnum`](../../doc/models/event-severity-enum.md) | Optional | Severity level for an event. enum: `normal`, `critical`, `high`, `warning` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeEventSysInfo := models.MxedgeEventSysInfo{
        Resource:             models.ToPointer("resource6"),
        Severity:             models.ToPointer(models.EventSeverityEnum_HIGH),
    }

}
```

