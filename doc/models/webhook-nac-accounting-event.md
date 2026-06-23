
# Webhook Nac Accounting Event

NAC accounting event for a client session with traffic counters

## Structure

`WebhookNacAccountingEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | MAC address of the AP the client roamed or disconnected from |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap` |
| `Bssid` | `*string` | Optional | Wireless BSSID used for the NAC accounting session |
| `ClientIp` | `*string` | Optional | Client IP address observed for the NAC accounting session |
| `ClientType` | `*string` | Optional | Client type E.g. "wired", "wireless", "vty" |
| `Mac` | `*string` | Optional | Client MAC address for the NAC accounting session |
| `NasVendor` | `*string` | Optional | NAS Device vendor name E.g. "Juniper", "Cisco" |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Ssid` | `*string` | Optional | Wireless SSID used for the NAC accounting session |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Type` | `*string` | Optional | NAC accounting event type, such as `ACCOUNTING_START`, `ACCOUNTING_UPDATE`, or `ACCOUNTING_STOP` |
| `Username` | `*string` | Optional | Client-presented username for NAC authentication |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookNacAccountingEvent := models.WebhookNacAccountingEvent{
        Ap:                   models.ToPointer("5c5b355005be"),
        AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTLS),
        Bssid:                models.ToPointer("5c5b35546bb4"),
        ClientIp:             models.ToPointer("172.16.87.4"),
        ClientType:           models.ToPointer("wireless"),
        Mac:                  models.ToPointer("6e795836d5f9"),
        NasVendor:            models.ToPointer("juniper-mist"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Ssid:                 models.ToPointer("Test-CMR SSID"),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        Type:                 models.ToPointer("NAC_ACCOUNTING_STOP"),
        Username:             models.ToPointer("hi"),
    }

}
```

