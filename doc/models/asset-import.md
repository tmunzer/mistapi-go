
# Asset Import

Asset record supplied in a JSON import payload

## Structure

`AssetImport`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Bluetooth MAC address used to identify the imported BLE asset<br><br>**Constraints**: *Minimum Length*: `1` |
| `Name` | `string` | Required | Display name for the imported BLE asset<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    assetImport := models.AssetImport{
        Mac:                  "mac2",
        Name:                 "name8",
    }

}
```

