
# Guest

Guest authorization record at site scope

*This model accepts additional fields of type interface{}.*

## Structure

`Guest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessCodeEmail` | `*string` | Optional, Read-only | If `auth_method`==`email`, the email address where the authorization code has been sent to |
| `ApMac` | `*string` | Optional, Read-only | MAC address of the AP used during guest registration |
| `AuthMethod` | `*string` | Optional, Read-only | Guest authentication method used for the authorization |
| `Authorized` | `*bool` | Optional | Whether the guest is currently authorized<br><br>**Default**: `true` |
| `AuthorizedExpiringTime` | `*float64` | Optional, Read-only | Unix timestamp when the guest authorization expires |
| `AuthorizedTime` | `*float64` | Optional, Read-only | Unix timestamp when the guest was authorized |
| `Company` | `*string` | Optional | Optional company name provided by the guest during registration |
| `Email` | `*string` | Optional | Optional email address provided by the guest during registration |
| `Field1` | `*string` | Optional | Optional custom field 1 value provided by the guest during registration |
| `Field2` | `*string` | Optional | Optional custom field 2 value provided by the guest during registration |
| `Field3` | `*string` | Optional | Optional custom field 3 value provided by the guest during registration |
| `Field4` | `*string` | Optional | Optional custom field 4 value provided by the guest during registration |
| `Mac` | `*string` | Optional, Read-only | Device MAC address captured during guest registration |
| `Minutes` | `*int` | Optional | Authorization duration, in minutes. Default is 1440 minutes (1 day), maximum is 259200 (180 days)<br><br>**Default**: `1440`<br><br>**Constraints**: `>= 0`, `<= 259200` |
| `Name` | `*string` | Optional, Read-only | Optional name provided by the guest during registration |
| `RandomMac` | `*bool` | Optional, Read-only | Whether the guest device used a randomized MAC address to connect to the SSID |
| `Ssid` | `*string` | Optional, Read-only | Name of the SSID |
| `WlanId` | `*uuid.UUID` | Optional, Read-only | Identifier of the WLAN used for the guest authorization |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    guest := models.Guest{
        AccessCodeEmail:        models.ToPointer("access_code_email4"),
        ApMac:                  models.ToPointer("ap_mac4"),
        AuthMethod:             models.ToPointer("auth_method6"),
        Authorized:             models.ToPointer(true),
        AuthorizedExpiringTime: models.ToPointer(float64(1480704955)),
        AuthorizedTime:         models.ToPointer(float64(1480704355)),
        Company:                models.ToPointer("abc"),
        Email:                  models.ToPointer("john@abc.com"),
        Minutes:                models.ToPointer(1440),
        Name:                   models.ToPointer("John Smith"),
        Ssid:                   models.ToPointer("Guest-SSID"),
        WlanId:                 models.ToPointer(uuid.MustParse("6748cfa6-4e12-11e6-9188-0242ac110007")),
        AdditionalProperties:   map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

