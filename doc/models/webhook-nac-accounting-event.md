
# Webhook Nac Accounting Event

## Structure

`WebhookNacAccountingEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | MAC address of the AP the client roamed or disconnected from |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap` |
| `Bssid` | `*string` | Optional | MAC physical address of the access point |
| `ClientIp` | `*string` | Optional | IP Address of client |
| `ClientType` | `*string` | Optional | Client type E.g. "wired", "wireless", "vty" |
| `Mac` | `*string` | Optional | Client's MAC Address |
| `NasVendor` | `*string` | Optional | NAS Device vendor name E.g. "Juniper", "Cisco" |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `*string` | Optional | ESSID |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `Type` | `*string` | Optional | Type of event. E.g. "ACCOUNTING_START", "ACCOUNTING_UPDATE", "ACCOUNTING_STOP" |
| `Username` | `*string` | Optional | Username authenticated with |

## Example (as JSON)

```json
{
  "ap": "5c5b355005be",
  "auth_type": "eap-tls",
  "bssid": "5c5b35546bb4",
  "client_ip": "172.16.87.4",
  "client_type": "wireless",
  "mac": "6e795836d5f9",
  "nas_vendor": "juniper-mist",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rx_pkts": 57770567,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "Test-CMR SSID",
  "tx_pkts": 812204062,
  "type": "NAC_ACCOUNTING_STOP",
  "username": "hi"
}
```

