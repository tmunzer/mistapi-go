
# Client Nac

*This model accepts additional fields of type interface{}.*

## Structure

`ClientNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `[]string` | Optional | - |
| `AuthType` | `*string` | Optional | Type of authentication used by the client |
| `CertCn` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `CertIssuer` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `CertSerial` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `CertSubject` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `ClientIp` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `DeviceMac` | `*string` | Optional | - |
| `Group` | `*string` | Optional | - |
| `IdpId` | `*string` | Optional | - |
| `IdpRole` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `LastAp` | `*string` | Optional | - |
| `LastCertCn` | `*string` | Optional | - |
| `LastCertExpiry` | `*int` | Optional | - |
| `LastCertIssuer` | `*string` | Optional | - |
| `LastCertSerial` | `*string` | Optional | - |
| `LastCertSubject` | `*string` | Optional | - |
| `LastClientIp` | `*string` | Optional | - |
| `LastNacruleId` | `*string` | Optional | - |
| `LastNacruleName` | `*string` | Optional | - |
| `LastNasVendor` | `*string` | Optional | - |
| `LastPortId` | `*string` | Optional | - |
| `LastSsid` | `*string` | Optional | - |
| `LastStatus` | `*string` | Optional | - |
| `LastUsername` | `*string` | Optional | - |
| `LastVlan` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `NacruleId` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `NacruleMatched` | `*bool` | Optional | - |
| `NacruleName` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `NasVendor` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortId` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `RandomMac` | `*bool` | Optional | Whether the client is using randomized MAC Address or not |
| `RespAttr` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | Type of client (wired, wireless) |
| `Username` | `[]string` | Optional | List of usernames that have been assigned to the client |
| `Vlan` | `[]string` | Optional | List of vlans that have been assigned to the client |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "cert_serial": [
    "cert_serial8",
    "cert_serial9",
    "cert_serial0"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

