
# Stats Wan Client

WAN client record returned by WAN client search

## Structure

`StatsWanClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DhcpExpireTime` | `*float64` | Optional | DHCP lease expiration time for the WAN client, in epoch seconds |
| `DhcpStartTime` | `*float64` | Optional | DHCP lease start time for the WAN client, in epoch seconds |
| `Hostname` | `[]string` | Optional | Hostnames observed for a WAN client |
| `Ip` | `[]string` | Optional | IP addresses observed for a WAN client |
| `IpSrc` | `*string` | Optional | Source used to learn the WAN client IP address, such as dhcp |
| `LastHostname` | `*string` | Optional | Most recent hostname observed for the WAN client |
| `LastIp` | `*string` | Optional | Most recent IP address observed for the WAN client |
| `Mfg` | `*string` | Optional | Manufacturer inferred for the WAN client |
| `Network` | `*string` | Optional | Mist network name associated with the WAN client |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Wcid` | `*string` | Optional | WAN client identifier associated with the record |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWanClient := models.StatsWanClient{
        DhcpExpireTime:       models.ToPointer(float64(74.98)),
        DhcpStartTime:        models.ToPointer(float64(189.3)),
        Hostname:             []string{
            "hostname6",
        },
        Ip:                   []string{
            "ip1",
        },
        IpSrc:                models.ToPointer("dhcp"),
        LastHostname:         models.ToPointer("sonoszp"),
        LastIp:               models.ToPointer("192.168.1.139"),
        Mfg:                  models.ToPointer("Sonos"),
        Network:              models.ToPointer("lan"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Wcid:                 models.ToPointer("8bbe7389-212b-c65d-2208-00fab2017936"),
    }

}
```

