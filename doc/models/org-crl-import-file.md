
# Org Crl Import File

Multipart upload payload for importing an organization CRL file

## Structure

`OrgCrlImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*[]byte` | Optional | a PEM or DER formatted CRL file |
| `Json` | `*string` | Optional | a JSON string with "name" field for CRL file issuer (optional) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgCrlImportFile := models.OrgCrlImportFile{
        File:                 models.ToPointer(nil),
        Json:                 models.ToPointer("json6"),
    }

}
```

