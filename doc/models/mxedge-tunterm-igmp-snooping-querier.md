
# Mxedge Tunterm Igmp Snooping Querier

IGMP querier settings for tunnel termination

## Structure

`MxedgeTuntermIgmpSnoopingQuerier`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaxResponseTime` | `*int` | Optional | Querier's query response interval, in tenths-of-seconds |
| `Mtu` | `*int` | Optional | The MTU we use (needed when forming large IGMPv3 Reports) |
| `QueryInterval` | `*int` | Optional | Querier's query interval, in seconds |
| `Robustness` | `*int` | Optional | IGMP querier robustness variable<br><br>**Constraints**: `>= 1`, `<= 7` |
| `Version` | `*int` | Optional | Querier's maximum protocol version |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermIgmpSnoopingQuerier := models.MxedgeTuntermIgmpSnoopingQuerier{
        MaxResponseTime:      models.ToPointer(10),
        Mtu:                  models.ToPointer(1500),
        QueryInterval:        models.ToPointer(125),
        Robustness:           models.ToPointer(7),
        Version:              models.ToPointer(3),
    }

}
```

