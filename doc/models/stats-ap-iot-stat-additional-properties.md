
# Stats Ap Iot Stat Additional Properties

IoT input statistic value

## Structure

`StatsApIotStatAdditionalProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Value` | `models.Optional[int]` | Optional, Read-only | Reported value for this IoT input |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApIotStatAdditionalProperties := models.StatsApIotStatAdditionalProperties{
        Value:                models.NewOptional(models.ToPointer(116)),
    }

}
```

