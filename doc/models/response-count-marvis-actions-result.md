
# Response Count Marvis Actions Result

Marvis action count result for one distinct value

*This model accepts additional fields of type string.*

## Structure

`ResponseCountMarvisActionsResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | Number of Marvis actions matching this distinct value |
| `AdditionalProperties` | `map[string]string` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseCountMarvisActionsResult := models.ResponseCountMarvisActionsResult{
        Count:                models.ToPointer(24),
        AdditionalProperties: map[string]string{
            "exampleAdditionalProperty": "response_count_marvis_actions_result_additionalProperties8",
        },
    }

}
```

