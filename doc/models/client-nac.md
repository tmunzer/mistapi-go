
# Client Nac

## Structure

`ClientNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `[]string` | Optional | - |
| `AuthType` | `*string` | Optional | - |
| `CertCn` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `CertIssuer` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `IdpId` | `*string` | Optional | - |
| `IdpRole` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `LastAp` | `*string` | Optional | - |
| `LastCertCn` | `*string` | Optional | - |
| `LastCertExpiry` | `*int` | Optional | - |
| `LastCertIssuer` | `*string` | Optional | - |
| `LastNacruleId` | `*string` | Optional | - |
| `LastNacruleName` | `*string` | Optional | - |
| `LastNasVendor` | `*string` | Optional | - |
| `LastSsid` | `*string` | Optional | - |
| `LastStatus` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `NacruleId` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `NacruleMatched` | `*bool` | Optional | - |
| `NacruleName` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `NasVendor` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RandomMac` | `*bool` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "ap": [
    "5c5b35bf16bb",
    "d4dc090041b4"
  ],
  "auth_type": "eap-tls",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 1694689718.612,
  "type": "wireless",
  "cert_cn": [
    "cert_cn9",
    "cert_cn0"
  ],
  "cert_issuer": [
    "cert_issuer4",
    "cert_issuer3",
    "cert_issuer2"
  ],
  "idp_id": "idp_id0"
}
```

