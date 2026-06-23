
# Const License Type

License type definition returned by the constants API

## Structure

`ConstLicenseType`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | Human-readable description of the license type |
| `Includes` | `[]string` | Optional | License SKU components included by a license type |
| `Key` | `*string` | Optional | Machine-readable license type key |
| `Name` | `*string` | Optional | Display name of the license type |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constLicenseType := models.ConstLicenseType{
        Description:          models.ToPointer("Wired Assurance 12"),
        Includes:             []string{
            "sub_ex12a",
            "sub_ex12p",
        },
        Key:                  models.ToPointer("sub_ex12"),
        Name:                 models.ToPointer("SUB-EX12"),
    }

}
```

