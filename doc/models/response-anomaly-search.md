
# Response Anomaly Search

Paginated anomaly search response

## Structure

`ResponseAnomalySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Search window end timestamp for anomaly results, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of anomaly results requested per page |
| `Page` | `int` | Required | Current page number in the anomaly search results |
| `Results` | [`[]models.Anomaly`](../../doc/models/anomaly.md) | Required | Anomaly records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Search window start timestamp for anomaly results, in epoch seconds |
| `Total` | `*int` | Optional | Number of anomaly records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAnomalySearch := models.ResponseAnomalySearch{
        End:                  1711035686,
        Limit:                10,
        Page:                 1,
        Results:              []models.Anomaly{
            models.Anomaly{
                Events:               []string{
                    "events4",
                },
                Since:                models.ToPointer(float64(133.06)),
                SleBaseline:          float64(108.54),
                SleDeviation:         float64(23),
                Timestamp:            float64(2.64),
            },
        },
        Start:                1710949286,
        Total:                models.ToPointer(232),
    }

}
```

