
# Tunnel Provider Options Zscaler

For zscaler-ipsec and zscaler-gre

## Structure

`TunnelProviderOptionsZscaler`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AupBlockInternetUntilAccepted` | `*bool` | Optional | Whether Zscaler blocks internet access until the Acceptable Use Policy is accepted<br><br>**Default**: `false` |
| `AupEnabled` | `*bool` | Optional | Can only be `true` when `auth_required`==`false`, display Acceptable Use Policy (AUP)<br><br>**Default**: `false` |
| `AupForceSslInspection` | `*bool` | Optional | Proxy HTTPs traffic, requiring Zscaler cert to be installed in browser<br><br>**Default**: `false` |
| `AupTimeoutInDays` | `*int` | Optional | Required if `aup_enabled`==`true`. Days before AUP is requested again<br><br>**Constraints**: `>= 1`, `<= 180` |
| `AuthRequired` | `*bool` | Optional | Enable this option to enforce user authentication<br><br>**Default**: `false` |
| `CautionEnabled` | `*bool` | Optional | Can only be `true` when `auth_required`==`false`, display caution notification for non-authenticated users<br><br>**Default**: `false` |
| `DnBandwidth` | `models.Optional[float64]` | Optional | Download bandwidth cap of the link, in Mbps. Disabled if not set<br><br>**Constraints**: `>= 0.1`, `<= 99999` |
| `IdleTimeInMinutes` | `*int` | Optional | Required if `surrogate_IP`==`true`, idle Time to Disassociation<br><br>**Constraints**: `>= 0`, `<= 43200` |
| `OfwEnabled` | `*bool` | Optional | If `true`, enable the firewall control option<br><br>**Default**: `false` |
| `SubLocations` | [`[]models.TunnelProviderOptionsZscalerSubLocation`](../../doc/models/tunnel-provider-options-zscaler-sub-location.md) | Optional | `sub-locations` can be used for specific uses cases to define different configuration based on the user network |
| `SurrogateIP` | `*bool` | Optional | Can only be `true` when `auth_required`==`true`. Map a user to a private IP address so it applies the user's policies, instead of the location's policies<br><br>**Default**: `false` |
| `SurrogateIPEnforcedForKnownBrowsers` | `*bool` | Optional | Can only be `true` when `surrogate_IP`==`true`, enforce surrogate IP for known browsers |
| `SurrogateRefreshTimeInMinutes` | `*int` | Optional | Required if `surrogate_IP_enforced_for_known_browsers`==`true`, must be lower or equal than `idle_time_in_minutes`, refresh Time for re-validation of Surrogacy<br><br>**Constraints**: `>= 1`, `<= 43200` |
| `UpBandwidth` | `models.Optional[float64]` | Optional | Download bandwidth cap of the link, in Mbps. Disabled if not set<br><br>**Constraints**: `>= 0.1`, `<= 99999` |
| `XffForwardEnabled` | `*bool` | Optional | Location uses proxy chaining to forward traffic<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelProviderOptionsZscaler := models.TunnelProviderOptionsZscaler{
        AupBlockInternetUntilAccepted:       models.ToPointer(false),
        AupEnabled:                          models.ToPointer(false),
        AupForceSslInspection:               models.ToPointer(false),
        AupTimeoutInDays:                    models.ToPointer(50),
        AuthRequired:                        models.ToPointer(false),
        CautionEnabled:                      models.ToPointer(false),
        DnBandwidth:                         models.NewOptional(models.ToPointer(float64(200))),
        OfwEnabled:                          models.ToPointer(false),
        SurrogateIP:                         models.ToPointer(false),
        UpBandwidth:                         models.NewOptional(models.ToPointer(float64(200))),
        XffForwardEnabled:                   models.ToPointer(false),
    }

}
```

