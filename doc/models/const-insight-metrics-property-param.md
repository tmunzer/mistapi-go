
# Const Insight Metrics Property Param

Query parameter definition for an insight metric

## Structure

`ConstInsightMetricsPropertyParam`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Required` | `*bool` | Optional | Whether this query parameter is required |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constInsightMetricsPropertyParam := models.ConstInsightMetricsPropertyParam{
        Required:             models.ToPointer(false),
    }

}
```

