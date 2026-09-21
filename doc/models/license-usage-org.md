
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
        NumDevices:           0,
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Usages:               0,
    }

}
```

