
# Nac Client Event

NAC authentication event for a wired or wireless client

## Structure

`NacClientEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | Access point MAC address for the client session |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap` |
| `Bssid` | `*string` | Optional | Wireless BSSID used for the client session |
| `ClientType` | [`*models.NacAccessTypeEnum`](../../doc/models/nac-access-type-enum.md) | Optional | Type of network access. enum: `wireless`, `wired`, `vty` |
| `DeviceMac` | `*string` | Optional, Read-only | MAC address of the AP or switch the client is connected to |
| `DryrunNacruleId` | `*uuid.UUID` | Optional, Read-only | NAC Policy Dry Run Rule ID, if present and matched |
| `DryrunNacruleMatched` | `*bool` | Optional, Read-only | `true` if dryrun rule present and matched with priority, `false` if not matched or not present |
| `IdpId` | `*uuid.UUID` | Optional, Read-only | If IDP is used, the id of the IDP configuration used |
| `IdpRole` | `[]string` | Optional | Identity provider roles or groups returned for a NAC client event |
| `IdpUsername` | `*string` | Optional, Read-only | If IDP is used, the username presented to the Identity Provider |
| `Mac` | `*string` | Optional, Read-only | Client MAC address observed in the NAC event |
| `MxedgeId` | `*string` | Optional | Mist Edge ID used to connect to cloud |
| `NacruleId` | `*uuid.UUID` | Optional, Read-only | NAC Policy Rule ID, if matched |
| `NacruleMatched` | `*bool` | Optional, Read-only | NAC Policy Rule Matched |
| `NasVendor` | `*string` | Optional, Read-only | Vendor name of the NAS |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PortId` | `*string` | Optional, Read-only | Port ID where the NAC client event occurred |
| `PortType` | [`*models.NacAccessTypeEnum`](../../doc/models/nac-access-type-enum.md) | Optional | Type of network access. enum: `wireless`, `wired`, `vty` |
| `RandomMac` | [`*models.RandomMacEnum`](../../doc/models/random-mac-enum.md) | Optional | Whether the client is using randomized MAC address or not. enum: `true`, `false` |
| `RespAttrs` | `[]string` | Optional | List of RADIUS AVP returned by the Authentication Server<br><br>**Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Ssid` | `*string` | Optional, Read-only | SSIDs the client was connecting to |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional, Read-only | Event type, e.g. NAC_CLIENT_PERMIT. Use the [List NAC Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions) endpoint to get the full list of available values. |
| `UsermacLabel` | `[]string` | Optional | Labels derived from usermac entry |
| `Username` | `*string` | Optional, Read-only | username assigned to the client |
| `Vlan` | `*string` | Optional, Read-only | vlan that assigned to the client |
| `VlanSource` | `*string` | Optional | Source of the assigned VLAN, for example `nactag` or `usermac` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    nacClientEvent := models.NacClientEvent{
        Ap:                   models.ToPointer("5c5b35513227"),
        AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTLS),
        Bssid:                models.ToPointer("5c5b355fafcc"),
        ClientType:           models.ToPointer(models.NacAccessTypeEnum_WIRELESS),
        DeviceMac:            models.ToPointer("60c78d8c7f6f"),
        DryrunNacruleId:      models.ToPointer(uuid.MustParse("32f27e7d-ff26-4a9b-b3d1-ff9bcb264012")),
        IdpId:                models.ToPointer(uuid.MustParse("912ef72e-2239-4996-b81e-469e87a27cd6")),
        IdpRole:              []string{
            "itsuperusers",
            "vip",
        },
        IdpUsername:          models.ToPointer("user@deaflyz.net"),
        Mac:                  models.ToPointer("ac3eb179e535"),
        NacruleId:            models.ToPointer(uuid.MustParse("32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62")),
        NasVendor:            models.ToPointer("juniper-mist"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PortId:               models.ToPointer("ge-0/0/17.0"),
        PortType:             models.ToPointer(models.NacAccessTypeEnum_WIRELESS),
        RespAttrs:            []string{
            "Tunnel-Type=VLAN",
            "Tunnel-Medium-Type=IEEE-802",
            "Tunnel-Private-Group-Id=750",
            "User-Name=anonymous",
        },
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Ssid:                 models.ToPointer("MyCorp-NAC"),
        Type:                 models.ToPointer("NAC_CLIENT_PERMIT"),
        UsermacLabel:         []string{
            "bldg5",
            "printer",
        },
        VlanSource:           models.ToPointer("nactag"),
    }

}
```

