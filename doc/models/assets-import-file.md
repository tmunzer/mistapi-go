
# Assets Import File

CSV file upload payload for importing BLE assets

## Structure

`AssetsImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*[]byte` | Optional | CSV file containing asset records to import |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    assetsImportFile := models.AssetsImportFile{
        File:                 models.ToPointer(nil),
    }

}
```

