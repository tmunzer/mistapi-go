
# Webhook Guest Authorizations Event

Guest portal authorization event with registration and WLAN context

## Structure

`WebhookGuestAuthorizationsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | MAC address of the AP handling the guest authorization |
| `AuthMethod` | `*string` | Optional | Guest portal authentication method used for this authorization event |
| `AuthorizedExpiringTime` | `*int` | Optional | Unix timestamp when the guest authorization expires |
| `AuthorizedTime` | `*int` | Optional | Unix timestamp when the guest authorization was granted |
| `Carrier` | `*string` | Optional | Mobile carrier used when authentication relies on a cellular provider |
| `Client` | `*string` | Optional | MAC address of the guest client device |
| `Company` | `*string` | Optional | Guest company name provided during registration |
| `Email` | `*string` | Optional | Guest email address provided during registration |
| `Field1` | `*string` | Optional | Value submitted for custom guest field 1 |
| `Field2` | `*string` | Optional | Value submitted for custom guest field 2 |
| `Field3` | `*string` | Optional | Value submitted for custom guest field 3 |
| `Field4` | `*string` | Optional | Value submitted for custom guest field 4 |
| `Mobile` | `*string` | Optional | Guest mobile phone number provided during registration |
| `Name` | `*string` | Optional | Full name provided by the guest during registration |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SmsGateway` | `*string` | Optional | SMS provider used for text-message authentication |
| `SponsorEmail` | `*string` | Optional | Sponsor email address associated with the guest authorization |
| `Ssid` | `*string` | Optional | WLAN SSID on which the guest was authorized |
| `WlanId` | `*string` | Optional | WLAN identifier on which the guest was authorized |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookGuestAuthorizationsEvent := models.WebhookGuestAuthorizationsEvent{
        Ap:                     models.ToPointer("5c5b350e55c8"),
        AuthMethod:             models.ToPointer("passphrase"),
        AuthorizedExpiringTime: models.ToPointer(1677076639),
        AuthorizedTime:         models.ToPointer(1677076519),
        Carrier:                models.ToPointer("docomo"),
        Client:                 models.ToPointer("ac2316eca70a"),
        Company:                models.ToPointer("MIST"),
        Email:                  models.ToPointer("abcd@abcd.com"),
        Field1:                 models.ToPointer("field1 value"),
        Field2:                 models.ToPointer("field2 value"),
        Field3:                 models.ToPointer("field3 value"),
        Field4:                 models.ToPointer("field4 value"),
        Mobile:                 models.ToPointer("+0123456789"),
        Name:                   models.ToPointer("Dr Strange"),
        OrgId:                  models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:                 models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        SmsGateway:             models.ToPointer("Telstra"),
        SponsorEmail:           models.ToPointer("sponsor@gmail.com"),
        Ssid:                   models.ToPointer("Portal Auth"),
        WlanId:                 models.ToPointer("7681be9a-044a-4622-90cf-3accde5ad853"),
    }

}
```

