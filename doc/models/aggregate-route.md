
# Aggregate Route

Aggregate route configuration for a network or routing instance

## Structure

`AggregateRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | Whether to install the aggregate route as a discard route<br><br>**Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | Routing metric assigned to the aggregate route<br><br>**Constraints**: `>= 0`, `<= 4294967295` |
| `Preference` | `models.Optional[int]` | Optional | Route preference assigned to the aggregate route<br><br>**Constraints**: `>= 0`, `<= 4294967295` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    aggregateRoute := models.AggregateRoute{
        Discard:              models.ToPointer(false),
        Metric:               models.NewOptional(models.ToPointer(240)),
        Preference:           models.NewOptional(models.ToPointer(192)),
    }

}
```

