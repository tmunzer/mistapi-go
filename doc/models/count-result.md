
# Count Result

Count result row with the matching distinct field values

*This model accepts additional fields of type string.*

## Structure

`CountResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `int` | Required | Number of matching items for the distinct value or values in this result |
| `AdditionalProperties` | `map[string]string` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    countResult := models.CountResult{
        Count:                78,
        AdditionalProperties: map[string]string{
            "exampleAdditionalProperty": "count_result_additionalProperties2",
        },
    }

}
```

