
# Image Import

Multipart image upload payload

## Structure

`ImageImport`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `[]byte` | Required | Image file content uploaded as multipart form data |
| `Json` | `*string` | Optional | Optional JSON metadata submitted with the image upload |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    imageImport := models.ImageImport{
        File:                 nil,
        Json:                 models.ToPointer("json2"),
    }

}
```

