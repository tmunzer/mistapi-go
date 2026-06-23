
# Client Nac

NAC client authentication and access state

## Structure

`ClientNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `[]string` | Optional, Read-only | AP MAC addresses observed for the NAC client during the selected time range |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap` |
| `CertCn` | `[]string` | Optional, Read-only | When certificate based authentication is used, the CN from the certificates used for the specified duration |
| `CertIssuer` | `[]string` | Optional, Read-only | When certificate based authentication is used, the Issuer from the certificates used for the specified duration |
| `CertSerial` | `[]string` | Optional, Read-only | When certificate based authentication is used, the Serial from the certificates used for the specified duration |
| `CertSubject` | `[]string` | Optional, Read-only | When certificate based authentication is used, the Subject from the certificates used for the specified duration |
| `ClientIp` | `[]string` | Optional, Read-only | The known IP addresses used by the client for the specified duration |
| `DeviceMac` | `*string` | Optional, Read-only | MAC address of the AP or switch the client is connected to |
| `EdrManaged` | `*bool` | Optional | Whether the NAC client is managed by an EDR provider |
| `EdrProvider` | [`*models.EdrProviderEnum`](../../doc/models/edr-provider-enum.md) | Optional | EDR provider associated with the NAC client. enum: `crowdstrike`, `sentinelone` |
| `EdrStatus` | [`*models.EdrStatusEnum`](../../doc/models/edr-status-enum.md) | Optional | EDR Status of the NAC client. enum: `sentinelone_healthy`, `sentinelone_infected`, `crowdstrike_low`, `crowdstrike_medium`, `crowdstrike_high`, `crowdstrike_critical`, `crowdstrike_informational` |
| `Group` | `*string` | Optional | User group associated with the NAC client |
| `IdpId` | `*string` | Optional | Identity Provider identifier used during NAC authentication |
| `IdpRole` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `LastAp` | `*string` | Optional | Latest AP where the client is/was connected to |
| `LastCertCn` | `*string` | Optional | When certificate based authentication is used, the CN from the latest certificate used |
| `LastCertExpiry` | `*float64` | Optional | When certificate based authentication is used, the expiration date from the latest certificate used |
| `LastCertIssuer` | `*string` | Optional | When certificate based authentication is used, the Issuer from the latest certificate used |
| `LastCertSerial` | `*string` | Optional | When certificate based authentication is used, the Serial from the latest certificate used |
| `LastCertSubject` | `*string` | Optional | When certificate based authentication is used, the Subject from the latest certificate used |
| `LastClientIp` | `*string` | Optional | The last known IP address for the client |
| `LastNacruleId` | `*string` | Optional | ID of the latest NAC Rule used to authenticate the client |
| `LastNacruleName` | `*string` | Optional | Name of the latest NAC Rule used to authenticate the client |
| `LastNasVendor` | `*string` | Optional | Vendor name of the NAS for the latest authentication |
| `LastPortId` | `*string` | Optional | If Wired authentication, the latest Port-id the client was connected to |
| `LastSsid` | `*string` | Optional | If Wireless authentication, the latest SSID the client was connected to |
| `LastStatus` | [`*models.NacClientLastStatusEnum`](../../doc/models/nac-client-last-status-enum.md) | Optional | Latest Authentication status of the client. enum: `denied`, `permitted`, `session_started`, `session_stopped` |
| `LastUsername` | `*string` | Optional | If dot1x authentication, the username used during the latest authentication. Otherwise, the MAC address of the client |
| `LastVlan` | `*int` | Optional | Latest VLAN ID assigned to the client |
| `Mac` | `*string` | Optional, Read-only | Client MAC address observed in the NAC event |
| `NacruleId` | `[]string` | Optional, Read-only | IDs of the NAC Rules used to authenticate the client for the specified duration |
| `NacruleMatched` | `*bool` | Optional | Whether a NAC policy rule matched the client |
| `NacruleName` | `[]string` | Optional, Read-only | Name of the NAC Rules used to authenticate the client for the specified duration |
| `NasIp` | `*string` | Optional | IP address of the NAS device used for authentication |
| `NasVendor` | `[]string` | Optional, Read-only | Vendor name of the NAS for the specified duration |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PortId` | `[]string` | Optional, Read-only | Port-ids the client was connected to for the specified duration |
| `RandomMac` | [`*models.RandomMacEnum`](../../doc/models/random-mac-enum.md) | Optional | Whether the client is using randomized MAC address or not. enum: `true`, `false` |
| `RespAttrs` | `[]string` | Optional | List of RADIUS AVP returned by the Authentication Server<br><br>**Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Ssid` | `[]string` | Optional | SSIDs the client was connected to for the specified duration |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | [`*models.NacAccessTypeEnum`](../../doc/models/nac-access-type-enum.md) | Optional | Type of network access. enum: `wireless`, `wired`, `vty` |
| `UsermacLabel` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Username` | `[]string` | Optional, Read-only | List of usernames that have been assigned to the client |
| `Vlan` | `[]string` | Optional, Read-only | List of vlans that have been assigned to the client |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    clientNac := models.ClientNac{
        Ap:                   []string{
            "5c5b35bf16bb",
            "d4dc090041b4",
        },
        AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTLS),
        CertCn:               []string{
            "john@mycorp.net",
        },
        CertIssuer:           []string{
            "/C=US/ST=CA/CN=MyCorp",
        },
        CertSerial:           []string{
            "2c63510123456789",
        },
        CertSubject:          []string{
            "/C=US/O=MyCorp/CN=john@mycorp.net/emailAddress=john@mycorp.net",
        },
        ClientIp:             []string{
            "10.100.0.157",
        },
        DeviceMac:            models.ToPointer("60c78d8c7f6f"),
        LastAp:               models.ToPointer("a83a79a947ee"),
        LastCertCn:           models.ToPointer("john@mycorp.net"),
        LastCertExpiry:       models.ToPointer(float64(1746711240)),
        LastCertIssuer:       models.ToPointer("/C=US/ST=CA/CN=MyCorp"),
        LastCertSerial:       models.ToPointer("2c63510123456789"),
        LastCertSubject:      models.ToPointer("/C=US/O=MyCorp/CN=john@mycorp.net/emailAddress=john@mycorp.net"),
        LastClientIp:         models.ToPointer("10.100.0.157"),
        LastNacruleId:        models.ToPointer("603b62db-d839-4152-9f7f-f2578443de8d"),
        LastNacruleName:      models.ToPointer("Wireless Cert Auth"),
        LastNasVendor:        models.ToPointer("juniper-mist"),
        LastPortId:           models.ToPointer("ge-0/0/17.0"),
        LastSsid:             models.ToPointer("MyCorp-NAC"),
        LastStatus:           models.ToPointer(models.NacClientLastStatusEnum_PERMITTED),
        LastUsername:         models.ToPointer("john@mycorp.net"),
        LastVlan:             models.ToPointer(10),
        Mac:                  models.ToPointer("ac3eb179e535"),
        NacruleId:            []string{
            "603b62db-d839-4152-9f7f-f2578443de8d",
        },
        NacruleName:          []string{
            "Wireless Cert Auth",
        },
        NasVendor:            []string{
            "juniper-mist",
        },
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PortId:               []string{
            "ge-0/0/17.0",
        },
        RespAttrs:            []string{
            "Tunnel-Type=VLAN",
            "Tunnel-Medium-Type=IEEE-802",
            "Tunnel-Private-Group-Id=750",
            "User-Name=anonymous",
        },
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Ssid:                 []string{
            "MyCorp-NAC",
        },
        Type:                 models.ToPointer(models.NacAccessTypeEnum_WIRELESS),
    }

}
```

