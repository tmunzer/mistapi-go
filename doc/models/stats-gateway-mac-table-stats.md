
# Stats Gateway Mac Table Stats

Gateway MAC table utilization counters

## Structure

`StatsGatewayMacTableStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MacTableCount` | `*int` | Optional | Number of MAC table entries currently learned by the gateway |
| `MaxMacEntriesSupported` | `*int` | Optional | Maximum MAC table entries supported by the gateway |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsGatewayMacTableStats := models.StatsGatewayMacTableStats{
        MacTableCount:          models.ToPointer(34),
        MaxMacEntriesSupported: models.ToPointer(178),
    }

}
```

