
# Response Nac Crl Files

Response containing uploaded NAC CRL file metadata

## Structure

`ResponseNacCrlFiles`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.NacCrlFile`](../../doc/models/nac-crl-file.md) | Optional | List of uploaded NAC CRL file metadata |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseNacCrlFiles := models.ResponseNacCrlFiles{
        Results:              []models.NacCrlFile{
            models.NacCrlFile{
                Name:                 models.ToPointer("name6"),
                Url:                  models.ToPointer("url0"),
            },
            models.NacCrlFile{
                Name:                 models.ToPointer("name6"),
                Url:                  models.ToPointer("url0"),
            },
        },
    }

}
```

