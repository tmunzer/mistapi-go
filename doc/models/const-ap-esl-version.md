
# Const Ap Esl Version

Supported Electronic Shelf Label (ESL) version for an AP model

## Structure

`ConstApEslVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EslVersion` | `*string` | Optional, Read-only | Electronic Shelf Label (ESL) package version supported by the AP model |
| `Model` | `*string` | Optional, Read-only | AP model that supports the listed ESL version |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constApEslVersion := models.ConstApEslVersion{
        EslVersion:           models.ToPointer("2.5.1"),
        Model:                models.ToPointer("AP34"),
    }

}
```

