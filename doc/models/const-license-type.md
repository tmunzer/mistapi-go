
# Const License Type

License type definition returned by the constants API

## Structure

`ConstLicenseType`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | Human-readable description of the license type |
| `EnforcementLevel` | [`*models.EnforcementLevelEnum`](../../doc/models/enforcement-level-enum.md) | Optional | Level at which the license is enforced |
| `EntitledLicenses` | `[]string` | Optional | License type keys this license type entitles |
| `Group` | `*string` | Optional | License group this license type belongs to |
| `Includes` | `[]string` | Optional | License SKU components included by a license type |
| `Key` | `*string` | Optional | Machine-readable license type key |
| `Name` | `*string` | Optional | Display name of the license type |
| `Type` | `*string` | Optional | License type identifier (SKU) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constLicenseType := models.ConstLicenseType{
        Description:          models.ToPointer("Wired Assurance 12"),
        EnforcementLevel:     models.ToPointer(models.EnforcementLevelEnum_ORG),
        EntitledLicenses:     []string{
            "sub_ex12",
            "sub_sadv1",
        },
        Group:                models.ToPointer("SUB-WIRED"),
        Includes:             []string{
            "sub_ex12a",
            "sub_ex12p",
        },
        Key:                  models.ToPointer("sub_ex12"),
        Name:                 models.ToPointer("SUB-EX12"),
        Type:                 models.ToPointer("SUB-EX12"),
    }

}
```

