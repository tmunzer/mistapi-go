
# License

Organization license entitlement, subscription, and usage summary

## Structure

`License`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amendments` | [`[]models.LicenseAmendment`](../../doc/models/license-amendment.md) | Optional, Read-only | Read-only list of license amendments<br><br>**Constraints**: *Unique Items Required* |
| `Entitled` | `map[string]int` | Optional, Read-only | Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses entitled. |
| `FullyLoaded` | `map[string]int` | Optional, Read-only | Maximum number of licenses that may be required if the service is enabled on all the Organization Devices. Property key is the service name (e.g. "SUB-MAN"). |
| `Licenses` | [`[]models.LicenseSub`](../../doc/models/license-sub.md) | Optional | License subscriptions for the organization |
| `Summary` | `map[string]int` | Optional, Read-only | Number of licenses currently consumed. Property key is license type (e.g. SUB-MAN). |
| `Usages` | `map[string]int` | Optional, Read-only | Number of available licenses. Property key is the service name (e.g. "SUB-MAN"). |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    license := models.License{
        Licenses:             []models.LicenseSub{
            models.LicenseSub{
            },
            models.LicenseSub{
            },
        },
    }

}
```

