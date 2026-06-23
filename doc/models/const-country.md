
# Const Country

Country definition returned by the constants API

## Structure

`ConstCountry`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alpha2` | `string` | Required | Country code, in two-character |
| `Certified` | `bool` | Required | Whether this country is certified for use in Mist APs |
| `Name` | `string` | Required | Human-readable country name for display |
| `Numeric` | `float64` | Required | Country code, ISO 3166-1 numeric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constCountry := models.ConstCountry{
        Alpha2:               "FR",
        Certified:            true,
        Name:                 "France",
        Numeric:              float64(250),
    }

}
```

