
# Client Nac

*This model accepts additional fields of type interface{}.*

## Structure

`ClientNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `[]string` | Optional | - |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `peap-tls`, `psk` |
| `CertCn` | `[]string` | Optional | When certificate based authentication is used, the CN from the certificates used for the specified duration |
| `CertIssuer` | `[]string` | Optional | When certificate based authentication is used, the Issuer from the certificates used for the specified duration |
| `CertSerial` | `[]string` | Optional | When certificate based authentication is used, the Serial from the certificates used for the specified duration |
| `CertSubject` | `[]string` | Optional | When certificate based authentication is used, the Subject from the certificates used for the specified duration |
| `ClientIp` | `[]string` | Optional | The known IP Addresses used by the client for the specified duration |
| `DeviceMac` | `*string` | Optional | MAC Address of the device (AP, Switch) the client is connected to |
| `Group` | `*string` | Optional | - |
| `IdpId` | `*string` | Optional | - |
| `IdpRole` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `LastAp` | `*string` | Optional | Latest AP where the client is/was connected to |
| `LastCertCn` | `*string` | Optional | When certificate based authentication is used, the CN from the latest certificate used |
| `LastCertExpiry` | `*float64` | Optional | When certificate based authentication is used, the expiration date from the latest certificate used |
| `LastCertIssuer` | `*string` | Optional | When certificate based authentication is used, the Issuer from the latest certificate used |
| `LastCertSerial` | `*string` | Optional | When certificate based authentication is used, the Serial from the latest certificate used |
| `LastCertSubject` | `*string` | Optional | When certificate based authentication is used, the Subject from the latest certificate used |
| `LastClientIp` | `*string` | Optional | The last known IP Address for the client |
| `LastNacruleId` | `*string` | Optional | ID of the latest NAC Rule used to authenticate the client |
| `LastNacruleName` | `*string` | Optional | Name of the latest NAC Rule used to authenticate the client |
| `LastNasVendor` | `*string` | Optional | Vendor name of the NAS for the latest authentication |
| `LastPortId` | `*string` | Optional | If Wired authentication, the latest Port-id the client was connected to |
| `LastSsid` | `*string` | Optional | If Wireless authentication, the latest SSID the client was connected to |
| `LastStatus` | [`*models.LastStatusEnum`](../../doc/models/last-status-enum.md) | Optional | Latest Authentication status of the client. enum: `denied`, `permitted`, `session_started`, `session_stopped` |
| `LastUsername` | `*string` | Optional | If dot1x authentication, the username used during the latest authentication. Otherwise, the MAC address of the client |
| `LastVlan` | `*int` | Optional | Latest VLAN ID assigned to the client |
| `Mac` | `*string` | Optional | Client MAC address |
| `NacruleId` | `[]string` | Optional | IDs of the NAC Rules used to authenticate the client for the specified duration |
| `NacruleMatched` | `*bool` | Optional | - |
| `NacruleName` | `[]string` | Optional | Name of the NAC Rules used to authenticate the client for the specified duration |
| `NasVendor` | `[]string` | Optional | Vendor name of the NAS for the specified duration |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortId` | `[]string` | Optional | Port-ids the client was connected to  for the specified duration |
| `RandomMac` | `*bool` | Optional | Whether the client is using randomized MAC Address or not |
| `RespAttrs` | `[]string` | Optional | List of Radius AVP returned by the Authentication Server<br>**Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `[]string` | Optional | SSIDs the client was connected to  for the specified duration |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Type` | [`*models.NacAccessTypeEnum`](../../doc/models/nac-access-type-enum.md) | Optional | Type of network access. enum: `wireless`, `wired` |
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
  "cert_cn": [
    "john@mycorp.net"
  ],
  "cert_issuer": [
    "/C=US/ST=CA/CN=MyCorp"
  ],
  "cert_serial": [
    "2c63510123456789"
  ],
  "cert_subject": [
    "/C=US/O=MyCorp/CN=john@mycorp.net/emailAddress=john@mycorp.net"
  ],
  "client_ip": [
    "10.100.0.157"
  ],
  "device_mac": "60c78d8c7f6f",
  "last_ap": "a83a79a947ee",
  "last_cert_cn": "john@mycorp.net",
  "last_cert_expiry": 1746711240,
  "last_cert_issuer": "/C=US/ST=CA/CN=MyCorp",
  "last_cert_serial": "2c63510123456789",
  "last_cert_subject": "/C=US/O=MyCorp/CN=john@mycorp.net/emailAddress=john@mycorp.net",
  "last_client_ip": "10.100.0.157",
  "last_nacrule_id": "603b62db-d839-4152-9f7f-f2578443de8d",
  "last_nacrule_name": "Wireless Cert Auth",
  "last_nas_vendor": "juniper-mist",
  "last_port_id": "ge-0/0/17.0",
  "last_ssid": "MyCorp-NAC",
  "last_status": "permitted",
  "last_username": "john@mycorp.net",
  "last_vlan": 10,
  "mac": "ac3eb179e535",
  "nacrule_id": [
    "603b62db-d839-4152-9f7f-f2578443de8d"
  ],
  "nacrule_name": [
    "Wireless Cert Auth"
  ],
  "nas_vendor": [
    "juniper-mist"
  ],
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "port_id": [
    "ge-0/0/17.0"
  ],
  "resp_attrs": [
    "Tunnel-Type=VLAN",
    "Tunnel-Medium-Type=IEEE-802",
    "Tunnel-Private-Group-Id=750",
    "User-Name=anonymous"
  ],
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": [
    "MyCorp-NAC"
  ],
  "type": "wireless",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

