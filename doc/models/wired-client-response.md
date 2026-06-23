
# Wired Client Response

Wired client record returned by a wired client search

## Structure

`WiredClientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthMethod` | `*string` | Optional | Method used to authenticate the wired client |
| `AuthState` | `*string` | Optional | State reported for wired client authentication |
| `DeviceMac` | `[]string` | Optional, Read-only | Switch or gateway MAC addresses where the wired client was observed |
| `DeviceMacPort` | [`[]models.WiredClientResponseDeviceMacPortItem`](../../doc/models/wired-client-response-device-mac-port-item.md) | Optional, Read-only | Per-port switch or gateway observations for a wired client<br><br>**Constraints**: *Unique Items Required* |
| `DhcpClientIdentifier` | `*string` | Optional | Identifier value reported by the wired client in DHCP |
| `DhcpClientOptions` | [`[]models.DhcpClientOption`](../../doc/models/dhcp-client-option.md) | Optional | DHCP options observed for a wired client |
| `DhcpFqdn` | `*string` | Optional | Fully qualified domain name reported by the wired client through DHCP |
| `DhcpHostname` | `*string` | Optional | Hostname reported by the wired client through DHCP |
| `DhcpRequestParams` | `*string` | Optional | Parameter request list advertised by the wired client in DHCP |
| `DhcpVendorClassIdentifier` | `*string` | Optional | Vendor class identifier reported by the wired client in DHCP |
| `Ip` | `[]string` | Optional, Read-only | Client IP addresses observed for a wired client |
| `Mac` | `*string` | Optional, Read-only | Client MAC address for the wired client record |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PortId` | `[]string` | Optional, Read-only | Switch or gateway port identifiers where a wired client was observed |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Vlan` | `[]int` | Optional, Read-only | Client VLAN IDs observed for a wired client |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    wiredClientResponse := models.WiredClientResponse{
        AuthMethod:                models.ToPointer("mac_auth"),
        AuthState:                 models.ToPointer("authenticated"),
        DeviceMac:                 []string{
            "device_mac5",
            "device_mac6",
        },
        DeviceMacPort:             []models.WiredClientResponseDeviceMacPortItem{
            models.WiredClientResponseDeviceMacPortItem{
                DeviceMac:            models.ToPointer("device_mac8"),
                Ip:                   models.ToPointer("ip8"),
                PortId:               models.ToPointer("port_id4"),
                PortParent:           models.ToPointer("port_parent6"),
                Start:                models.ToPointer("start8"),
            },
            models.WiredClientResponseDeviceMacPortItem{
                DeviceMac:            models.ToPointer("device_mac8"),
                Ip:                   models.ToPointer("ip8"),
                PortId:               models.ToPointer("port_id4"),
                PortParent:           models.ToPointer("port_parent6"),
                Start:                models.ToPointer("start8"),
            },
            models.WiredClientResponseDeviceMacPortItem{
                DeviceMac:            models.ToPointer("device_mac8"),
                Ip:                   models.ToPointer("ip8"),
                PortId:               models.ToPointer("port_id4"),
                PortParent:           models.ToPointer("port_parent6"),
                Start:                models.ToPointer("start8"),
            },
        },
        DhcpClientIdentifier:      models.ToPointer("MAC address 00155df6d500"),
        DhcpFqdn:                  models.ToPointer("ITS-VMMT0-D1N02.mgthub.local"),
        DhcpHostname:              models.ToPointer("ITS-VMMT0-D1N02"),
        DhcpRequestParams:         models.ToPointer("1 3 6 15 31 33 43 44 46 47 119 121 249 252"),
        DhcpVendorClassIdentifier: models.ToPointer("MSFT 5.0"),
        OrgId:                     models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:                    models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

