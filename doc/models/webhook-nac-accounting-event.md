
# Webhook Nac Accounting Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookNacAccountingEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | MAC address of the AP the client roamed or disconnected from |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `peap-tls`, `psk` |
| `Bssid` | `*string` | Optional | MAC physical address of the access point |
| `ClientIp` | `*string` | Optional | IP Address of client |
| `ClientType` | `*string` | Optional | Client type E.g. "wired", "wireless", "vty" |
| `Mac` | `*string` | Optional | Client's MAC Address |
| `NasVendor` | `*string` | Optional | NAS Device vendor name E.g. "Juniper", "Cisco" |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RxPkts` | `*int` | Optional | Number of packets received |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `*string` | Optional | ESSID |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `TxPkts` | `*int` | Optional | Number of packets sent |
| `Type` | `*string` | Optional | Type of event. E.g. "ACCOUNTING_START", "ACCOUNTING_UPDATE", "ACCOUNTING_STOP" |
| `Username` | `*string` | Optional | Username authenticated with |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "rx_pkts": 9523,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "Test-CMR SSID",
  "tx_pkts": 5233,
  "type": "NAC_ACCOUNTING_STOP",
  "username": "hi",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

