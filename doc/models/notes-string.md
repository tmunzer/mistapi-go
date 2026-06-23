
# Notes String

Request body containing notes text

*This model accepts additional fields of type interface{}.*

## Structure

`NotesString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Notes` | `*string` | Optional | Text to attach to the target resource as notes |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    notesString := models.NotesString{
        Notes:                models.ToPointer("wired pcap test"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

