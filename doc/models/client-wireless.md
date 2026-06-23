
# Client Wireless

Wireless client record observed during a client search or stats query

## Structure

`ClientWireless`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `[]string` | Optional | List of AP MAC addresses the client was connected to |
| `AppVersion` | `[]string` | Optional | Only when client has the Marvis Client app running. List of the versions of the Marvis Client |
| `Band` | `*string` | Optional | Wi-Fi radio band used by the wireless client |
| `Device` | `[]string` | Optional | Only when client has the Marvis Client app running. List of the type of device type detected |
| `Ftc` | `*bool` | Optional | Whether fast transition information is reported for the wireless client |
| `Hardware` | `*string` | Optional | Only when client has the Marvis Client app running. Type of Wi-Fi adapter |
| `Hostname` | `[]string` | Optional | List of hostname detected for this client |
| `Ip` | `[]string` | Optional | List if the IP addresses detected for this client |
| `LastAp` | `*string` | Optional | Latest AP where the client is/was connected to |
| `LastDevice` | `*string` | Optional | Latest type of device we identified (e.g. iPhone, Mac, ...) |
| `LastFirmware` | `*string` | Optional | Only when client has the Marvis Client app running. Same as "firmware" |
| `LastHostname` | `*string` | Optional | Latest hostname we detected for the client |
| `LastIp` | `*string` | Optional | The last known IP address for the client |
| `LastModel` | `*string` | Optional | Only when client has the Marvis Client app running. latest client hardware model we detected for the client |
| `LastOs` | `*string` | Optional | Only when client has the Marvis Client app running. Latest version of OS Type we detected for the client |
| `LastOsVersion` | `*string` | Optional | Only when client has the Marvis Client app running. Latest version of OS Version we detected for the client |
| `LastPskId` | `*uuid.UUID` | Optional | Only for PPSK authentication. Latest PPSK ID used by the client |
| `LastPskName` | `*string` | Optional | Only for PPSK authentication. Latest PPSK Name used by the client |
| `LastSsid` | `*string` | Optional | If dot1x authentication, the username used during the latest authentication. Otherwise, the MAC address of the client |
| `LastUsername` | `*string` | Optional | Most recent username associated with the wireless client |
| `LastVlan` | `*int` | Optional | Latest VLAN ID assigned to the client |
| `LastWlanId` | `*uuid.UUID` | Optional | ID of the latest SSID (WLAN) the client is/was connected to |
| `Mac` | `*string` | Optional | Client MAC address for the wireless client |
| `Mfg` | `*string` | Optional | Manufacturer of the client hardware (MAC OUI based) |
| `Model` | `*string` | Optional | Only when client has the Marvis Client app running. Client hardware model |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Os` | `[]string` | Optional | Only when client is having the Marvis Client app running. List of OS detected for the client |
| `OsVersion` | `[]string` | Optional | Only when client is having the Marvis Client app running. List of OS version detected for the client |
| `Protocol` | `*string` | Optional | 802.11 protocol amendment used by the wireless client |
| `PskId` | `[]uuid.UUID` | Optional | List of IDs of the PPSK used by the client |
| `PskName` | `[]string` | Optional | List of names of the PPSK used by the client |
| `RandomMac` | `*bool` | Optional | Whether the wireless client uses a randomized MAC address |
| `SdkVersion` | `[]string` | Optional | Only when client has the Marvis Client app running. List of Marvis Client SDK version detected for the client |
| `SiteId` | `*uuid.UUID` | Optional | Mist Site ID where the client is connected |
| `SiteIds` | `[]uuid.UUID` | Optional | List of Mist Site IDs where the client was connected |
| `Ssid` | `[]string` | Optional | List of the WLAN names the client was connected to |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Username` | `[]string` | Optional | Only for 802.1X authentication. List of usernames used by the client |
| `Vlan` | `[]int` | Optional | List of vlans that have been assigned to the client |
| `WlanId` | `[]uuid.UUID` | Optional | List of IDs of WLANs the client was connected to |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    clientWireless := models.ClientWireless{
        Ap:                   []string{
            "a83a79a947ee",
            "003e73170b4c",
        },
        AppVersion:           []string{
            "0.100.3",
        },
        Band:                 models.ToPointer("5"),
        Device:               []string{
            "Mac",
        },
        Ftc:                  models.ToPointer(false),
        Hardware:             models.ToPointer("Apple Wi-Fi adapter"),
        Hostname:             []string{
            "hostname-a",
            "hostname-b",
        },
        Ip:                   []string{
            "10.5.23.43",
            "192.168.0.2",
        },
        LastAp:               models.ToPointer("a83a79a947ee"),
        LastDevice:           models.ToPointer("Zebra"),
        LastFirmware:         models.ToPointer("wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675"),
        LastHostname:         models.ToPointer("hostname-a"),
        LastIp:               models.ToPointer("10.100.0.157"),
        LastModel:            models.ToPointer("MBP 16\\\" M1 2021"),
        LastOs:               models.ToPointer("Sonoma"),
        LastOsVersion:        models.ToPointer("14.4.1 (Build 23E224)"),
        LastPskId:            models.ToPointer(uuid.MustParse("abf7dc5c-bb51-4bb7-93b6-5547400ffe11")),
        LastPskName:          models.ToPointer("iot"),
        LastSsid:             models.ToPointer("john@mycorp.net"),
        LastVlan:             models.ToPointer(10),
        LastWlanId:           models.ToPointer(uuid.MustParse("e5d67b07-aae8-494b-8584-cbc20c8110aa")),
        Mac:                  models.ToPointer("bcd074000000"),
        Mfg:                  models.ToPointer("Apple"),
        Model:                models.ToPointer("MBP 16\\\" M1 2021"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Os:                   []string{
            "Sonoma",
        },
        OsVersion:            []string{
            "14.4.1 (Build 23E224)",
        },
        Protocol:             models.ToPointer("ax"),
        PskId:                []uuid.UUID{
            uuid.MustParse("abf7dc5c-bb51-4bb7-93b6-5547400ffe11"),
        },
        PskName:              []string{
            "iot",
        },
        SdkVersion:           []string{
            "0.100.3",
        },
        SiteId:               models.ToPointer(uuid.MustParse("25ff5219-9be7-4db9-907d-0c9b60445147")),
        SiteIds:              []uuid.UUID{
            uuid.MustParse("25ff5219-9be7-4db9-907d-0c9b60445147"),
        },
        Ssid:                 []string{
            "IoT SSID",
        },
        Username:             []string{
            "user@corp.com",
        },
        Vlan:                 []int{
            10,
        },
        WlanId:               []uuid.UUID{
            uuid.MustParse("e5d67b07-aae8-494b-8584-cbc20c8110aa"),
        },
    }

}
```

