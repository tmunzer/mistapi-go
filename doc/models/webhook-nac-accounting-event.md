
# Webhook Nac Accounting Event

## Structure

`WebhookNacAccountingEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | mac address of the AP the client roamed or disconnected from |
| `AuthType` | `*string` | Optional | radius authentication type |
| `Bssid` | `*string` | Optional | it’s the MAC physical address of the access point |
| `ClientIp` | `*string` | Optional | IP Address of client |
| `ClientType` | `*string` | Optional | client type E.g. “wired”, “wireless”, “vty” |
| `Mac` | `*string` | Optional | the client’s mac |
| `NasVendor` | `*string` | Optional | NAS Device vendor name E.g. “Juniper”, “Cisco” |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RxPkts` | `*int` | Optional | number of packets received |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `*string` | Optional | ESSID |
| `Timestamp` | `*float64` | Optional | sampling time (in epoch) |
| `TxPkts` | `*int` | Optional | number of packets sent |
| `Type` | `*string` | Optional | type of event. E.g. “ACCOUNTING_START”, “ACCOUNTING_UPDATE”, “ACCOUNTING_STOP” |
| `Username` | `*string` | Optional | username authenticated with |

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
  "timestamp": 1547235620.89,
  "tx_pkts": 5233,
  "type": "NAC_ACCOUNTING_STOP",
  "username": "hi"
}
```

