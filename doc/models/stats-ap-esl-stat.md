
# Stats Ap Esl Stat

Electronic shelf label dongle status reported by an AP

## Structure

`StatsApEslStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `models.Optional[int]` | Optional, Read-only | Radio channel used by the ESL dongle |
| `Connected` | `models.Optional[bool]` | Optional, Read-only | Whether the ESL dongle is connected |
| `Ip` | `models.Optional[string]` | Optional, Read-only | Network IP address of Hanshow and SoluM dongles |
| `Mac` | `models.Optional[string]` | Optional, Read-only | Dongle MAC address for Hanshow and SoluM dongles |
| `ProductId` | `models.Optional[string]` | Optional, Read-only | Product ID of Hanshow and SoluM dongles |
| `Type` | `models.Optional[string]` | Optional, Read-only | ESL dongle type reported by the AP |
| `Up` | `models.Optional[bool]` | Optional, Read-only | Whether the ESL dongle is operational |
| `VendorId` | `models.Optional[string]` | Optional, Read-only | Vendor ID of Hanshow and SoluM dongles |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApEslStat := models.StatsApEslStat{
        Channel:              models.NewOptional(models.ToPointer(118)),
        Connected:            models.NewOptional(models.ToPointer(false)),
        Ip:                   models.NewOptional(models.ToPointer("172.16.2.249")),
        Mac:                  models.NewOptional(models.ToPointer("98-6d-35-79-76-3b")),
        ProductId:            models.NewOptional(models.ToPointer("A4A2")),
        Type:                 models.NewOptional(models.ToPointer("imagotag")),
        VendorId:             models.NewOptional(models.ToPointer("0525")),
    }

}
```

