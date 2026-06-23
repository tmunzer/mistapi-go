
# Binary Stream

Binary file upload payload

## Structure

`BinaryStream`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `[]byte` | Required | Binary file payload to upload with this request |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    binaryStream := models.BinaryStream{
        File:                 nil,
    }

}
```

