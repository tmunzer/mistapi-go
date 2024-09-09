
# Const Nac Event

## Structure

`ConstNacEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | - |
| `Bssid` | `*string` | Optional | - |
| `CertCn` | `*string` | Optional | - |
| `CertExpiry` | `*int` | Optional | - |
| `CertIssuer` | `*string` | Optional | - |
| `CertSanUpn` | `[]string` | Optional | - |
| `CertSerial` | `*string` | Optional | - |
| `CertSubject` | `*string` | Optional | - |
| `EapType` | `*string` | Optional | - |
| `NasVendor` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RandomMac` | `*bool` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Username` | `*string` | Optional | - |
| `Wcid` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "ap": "5c5b355008c0",
  "bssid": "5c5b35548892",
  "cert_cn": "suriyas",
  "cert_expiry": 1711557441,
  "cert_issuer": "/DC=net/DC=jnpr/CN=Juniper Networks Issuing AWS1 CA",
  "cert_san_upn": [
    "suriyas@juniper.net"
  ],
  "cert_serial": "1300103d29e56ef083797bedc2000100103d29",
  "cert_subject": "/CN=suriyas/emailAddress=suriyas@juniper.net",
  "eap_type": "EAP-TLS",
  "nas_vendor": "Mist",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "random_mac": true,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "Test_Suriya-SSID",
  "timestamp": 1685658478.43899,
  "type": "NAC_CLIENT_CERT_CHECK_SUCCESS",
  "username": "suriyas@juniper.net",
  "wcid": "b43637b0-f0d9-0a1d-1ec2-73c394a9f679"
}
```

