
# License Usage Site

Site license usage and entitlement response

## Structure

`LicenseUsageSite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgEntitled` | `map[string]int` | Required | License entitlement for the entire org |
| `SvnaEnabled` | `bool` | Required | Eligibility for the Switch SLE |
| `TrialEnabled` | `bool` | Required | Whether trial licensing is enabled for the site |
| `Usages` | `map[string]int` | Required | Subscriptions and their quantities |
| `VnaEligible` | `bool` | Required | Eligibility for the AP/Client SLE |
| `VnaUi` | `bool` | Required | If True, Conversational Assistant and Marvis Action available |
| `WvnaEligible` | `bool` | Required | Eligibility for the WAN SLE |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    licenseUsageSite := models.LicenseUsageSite{
        OrgEntitled:          map[string]int{
            "SUB-LOC": 30,
            "SUB-MAN": 60,
        },
        SvnaEnabled:          false,
        TrialEnabled:         false,
        Usages:               map[string]int{
            "SUB-LOC": 30,
            "SUB-MAN": 60,
        },
        VnaEligible:          false,
        VnaUi:                false,
        WvnaEligible:         false,
    }

}
```

