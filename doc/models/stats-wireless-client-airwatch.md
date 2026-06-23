
# Stats Wireless Client Airwatch

AirWatch authorization information reported for a wireless client

## Structure

`StatsWirelessClientAirwatch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Authorized` | `bool` | Required | Whether the wireless client is authorized by AirWatch |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsWirelessClientAirwatch := models.StatsWirelessClientAirwatch{
        Authorized:           false,
    }

}
```

