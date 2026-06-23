
# Name String

Request body containing a name value

*This model accepts additional fields of type interface{}.*

## Structure

`NameString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Value to create or update as the target resource name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nameString := models.NameString{
        Name:                 models.ToPointer("name2"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

