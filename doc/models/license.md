
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
    "github.com/google/uuid"
)

func main() {
    license := models.License{
        Amendments:           []models.LicenseAmendment{
            models.LicenseAmendment{
                CreatedTime:          models.ToPointer(float64(132.88)),
                EndTime:              models.ToPointer(210),
                Id:                   models.ToPointer(uuid.MustParse("00000292-0000-0000-0000-000000000000")),
                ModifiedTime:         models.ToPointer(float64(202.08)),
                Quantity:             models.ToPointer(182),
            },
            models.LicenseAmendment{
                CreatedTime:          models.ToPointer(float64(132.88)),
                EndTime:              models.ToPointer(210),
                Id:                   models.ToPointer(uuid.MustParse("00000292-0000-0000-0000-000000000000")),
                ModifiedTime:         models.ToPointer(float64(202.08)),
                Quantity:             models.ToPointer(182),
            },
        },
        Entitled:             map[string]int{
            "key0": 120,
            "key1": 119,
        },
        FullyLoaded:          map[string]int{
            "key0": 214,
            "key1": 215,
            "key2": 216,
        },
        Licenses:             []models.LicenseSub{
            models.LicenseSub{
                CreatedTime:          models.ToPointer(float64(186.08)),
                EndTime:              models.ToPointer(154),
                Id:                   models.ToPointer(uuid.MustParse("0000017a-0000-0000-0000-000000000000")),
                ModifiedTime:         models.ToPointer(float64(148.88)),
                OrderId:              models.ToPointer("order_id2"),
            },
            models.LicenseSub{
                CreatedTime:          models.ToPointer(float64(186.08)),
                EndTime:              models.ToPointer(154),
                Id:                   models.ToPointer(uuid.MustParse("0000017a-0000-0000-0000-000000000000")),
                ModifiedTime:         models.ToPointer(float64(148.88)),
                OrderId:              models.ToPointer("order_id2"),
            },
        },
        Summary:              map[string]int{
            "key0": 22,
        },
    }

}
```

