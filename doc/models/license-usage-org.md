
# License Usage Org

License usage record for an organization usage scope

## Structure

`LicenseUsageOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForSite` | `*bool` | Optional, Read-only | Whether this license usage record is scoped to a site |
| `FullyLoaded` | `map[string]int` | Optional, Read-only | Maximum number of licenses that may be required if the service is enabled on all the Organization Devices. Property key is the service name (e.g. "SUB-MAN"). |
| `NumDevices` | `int` | Required, Read-only | Number of devices counted in this license usage scope |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Summary` | `map[string]int` | Optional, Read-only | Number of licenses currently consumed. Property key is license type (e.g. SUB-MAN). |
| `Usages` | `map[string]int` | Required, Read-only | Number of available licenses. Property key is the service name (e.g. "SUB-MAN"). |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    licenseUsageOrg := models.LicenseUsageOrg{
        ForSite:              models.ToPointer(false),
        FullyLoaded:          map[string]int{
            "key0": 162,
            "key1": 163,
        },
        NumDevices:           104,
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Summary:              map[string]int{
            "key0": 226,
            "key1": 227,
            "key2": 228,
        },
        Usages:               map[string]int{
            "key0": 55,
            "key1": 56,
        },
    }

}
```

