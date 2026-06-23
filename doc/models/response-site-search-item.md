
# Response Site Search Item

Site record returned by organization site search

## Structure

`ResponseSiteSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgradeEnabled` | `bool` | Required | Whether automatic firmware upgrade is enabled for the site |
| `AutoUpgradeVersion` | `string` | Required, Read-only | Desired automatic firmware upgrade version track for the site |
| `CountryCode` | `models.Optional[string]` | Optional, Read-only | Country code configured for the site, or null if unset |
| `HoneypotEnabled` | `bool` | Required | Whether honeypot detection is enabled for the site |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Name` | `string` | Required, Read-only | Display name of the site |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Timezone` | `string` | Required, Read-only | Time zone configured for the site |
| `VnaEnabled` | `bool` | Required | Whether Virtual Network Assistant is enabled for the site |
| `WifiEnabled` | `bool` | Required | Whether Wi-Fi service is enabled for the site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSiteSearchItem := models.ResponseSiteSearchItem{
        AutoUpgradeEnabled:   false,
        AutoUpgradeVersion:   "auto_upgrade_version2",
        CountryCode:          models.NewOptional(models.ToPointer("country_code4")),
        HoneypotEnabled:      false,
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Name:                 "name4",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Timestamp:            float64(105.12),
        Timezone:             "timezone6",
        VnaEnabled:           false,
        WifiEnabled:          false,
    }

}
```

