
# Note String

Request body containing a note value

*This model accepts additional fields of type interface{}.*

## Structure

`NoteString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Note` | `*string` | Optional | Some text note describing the intent |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    noteString := models.NoteString{
        Note:                 models.ToPointer("maintenance window"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

