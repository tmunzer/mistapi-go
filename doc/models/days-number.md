
# Days Number

Request body specifying CRL truncation retention in days

*This model accepts additional fields of type interface{}.*

## Structure

`DaysNumber`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Days` | `*int` | Optional | Number of most recent days of CRL entries to retain after truncation<br><br>**Default**: `30` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    daysNumber := models.DaysNumber{
        Days:                 models.ToPointer(30),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

