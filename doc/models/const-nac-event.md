
# Const Nac Event

NAC event example payload returned by the constants API

## Structure

`ConstNacEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | Access point MAC address associated with the NAC event |
| `Bssid` | `*string` | Optional | Wireless BSSID associated with the NAC event |
| `CertCn` | `*string` | Optional | Certificate common name presented during NAC authentication |
| `CertExpiry` | `*int` | Optional | Certificate expiration timestamp presented during NAC authentication |
| `CertIssuer` | `*string` | Optional | Certificate issuer presented during NAC authentication |
| `CertSanUpn` | `[]string` | Optional | Certificate Subject Alternative Name UPN values |
| `CertSerial` | `*string` | Optional | Certificate serial number presented during NAC authentication |
| `CertSubject` | `*string` | Optional | Certificate subject presented during NAC authentication |
| `EapType` | `*string` | Optional | EAP method used for the NAC authentication |
| `NasVendor` | `*string` | Optional | Network access server vendor reported for the NAC event |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RandomMac` | `*bool` | Optional | Whether the client used a randomized MAC address in the NAC event |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Ssid` | `*string` | Optional | Wireless SSID associated with the NAC event |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | NAC event type key |
| `Username` | `*string` | Optional | User identity associated with the NAC event |
| `Wcid` | `*uuid.UUID` | Optional | Wireless client identifier associated with the NAC event |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    constNacEvent := models.ConstNacEvent{
        Ap:                   models.ToPointer("5c5b355008c0"),
        Bssid:                models.ToPointer("5c5b35548892"),
        CertCn:               models.ToPointer("suriyas"),
        CertExpiry:           models.ToPointer(1711557441),
        CertIssuer:           models.ToPointer("/DC=net/DC=jnpr/CN=Juniper Networks Issuing AWS1 CA"),
        CertSanUpn:           []string{
            "suriyas@juniper.net",
        },
        CertSerial:           models.ToPointer("1300103d29e56ef083797bedc2000100103d29"),
        CertSubject:          models.ToPointer("/CN=suriyas/emailAddress=suriyas@juniper.net"),
        EapType:              models.ToPointer("EAP-TLS"),
        NasVendor:            models.ToPointer("Mist"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RandomMac:            models.ToPointer(true),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Ssid:                 models.ToPointer("Test_Suriya-SSID"),
        Type:                 models.ToPointer("NAC_CLIENT_CERT_CHECK_SUCCESS"),
        Username:             models.ToPointer("suriyas@juniper.net"),
        Wcid:                 models.ToPointer(uuid.MustParse("b43637b0-f0d9-0a1d-1ec2-73c394a9f679")),
    }

}
```

