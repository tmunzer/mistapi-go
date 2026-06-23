
# Response Count

Distinct count response for time-bounded search results

## Structure

`ResponseCount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Distinct` | `string` | Required | Field used to group the count results |
| `End` | `int` | Required | Search window end timestamp for the count request, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of distinct count results requested |
| `Results` | [`[]models.CountResult`](../../doc/models/count-result.md) | Required | List of count result rows<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Search window start timestamp for the count request, in epoch seconds |
| `Total` | `int` | Required | Number of distinct result buckets returned |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseCount := models.ResponseCount{
        Distinct:             "distinct0",
        End:                  128,
        Limit:                42,
        Results:              []models.CountResult{
            models.CountResult{
                Count:                226,
                AdditionalProperties: map[string]string{
                    "exampleAdditionalProperty": "count_result_additionalProperties2",
                },
            },
        },
        Start:                86,
        Total:                136,
    }

}
```

