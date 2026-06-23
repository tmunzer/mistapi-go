
# Jsi Warranty Type Enum

Warranty type label for Juniper Support Insight (JSI) devices. enum: `Standard Hardware Warranty`, `Enhanced Hardware Warranty`, `Dead On Arrival Warranty`, `Limited Lifetime Warranty`, `Software Warranty`, `Limited Lifetime Warranty for WLA`, `Warranty-JCPO EOL (DOA Not Included)`, `MIST Enhanced Hardware Warranty`, `MIST Standard Warranty`, `Determine Lifetime warranty`

## Enumeration

`JsiWarrantyTypeEnum`

## Fields

| Name |
|  --- |
| `Standard Hardware Warranty` |
| `Enhanced Hardware Warranty` |
| `Dead On Arrival Warranty` |
| `Limited Lifetime Warranty` |
| `Software Warranty` |
| `Limited Lifetime Warranty for WLA` |
| `Warranty-JCPO EOL (DOA Not Included)` |
| `MIST Enhanced Hardware Warranty` |
| `MIST Standard Warranty` |
| `Determine Lifetime warranty` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    jsiWarrantyType := models.JsiWarrantyTypeEnum_ENUMSOFTWAREWARRANTY

}
```

