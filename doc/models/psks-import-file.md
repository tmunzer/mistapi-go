
# Psks Import File

Multipart upload payload for importing PSKs

## Structure

`PsksImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*[]byte` | Optional | Uploaded CSV or binary file containing PSKs to import |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    psksImportFile := models.PsksImportFile{
        File:                 models.ToPointer(nil),
    }

}
```

