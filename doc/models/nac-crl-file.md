
# Nac Crl File

Metadata for an uploaded NAC CRL file

## Structure

`NacCrlFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*string` | Optional, Read-only | Unique ID for the uploaded CRL file, used to reference the file |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Issuer name for the CRL file |
| `Url` | `*string` | Optional | Download URL for the uploaded NAC CRL file |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacCrlFile := models.NacCrlFile{
        Id:                   models.ToPointer("a1ca26f3-44dd-4833-9a7b-97bbb2ab5230"),
        Name:                 models.ToPointer("SampleCertificateSigner"),
        Url:                  models.ToPointer("http://url/to/crl_file"),
    }

}
```

