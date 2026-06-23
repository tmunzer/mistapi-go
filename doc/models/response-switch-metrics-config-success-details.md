
# Response Switch Metrics Config Success Details

Detail values for the switch configuration success metric

## Structure

`ResponseSwitchMetricsConfigSuccessDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigSuccessCount` | `*int` | Optional | Number of switches with successful configuration status |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSwitchMetricsConfigSuccessDetails := models.ResponseSwitchMetricsConfigSuccessDetails{
        ConfigSuccessCount:   models.ToPointer(28),
    }

}
```

