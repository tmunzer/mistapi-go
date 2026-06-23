
# Stats Gateway Port

Gateway port statistics record returned by stats APIs

## Structure

`StatsGatewayPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Active` | `*bool` | Optional, Read-only | Indicates if interface is active/inactive |
| `AuthState` | [`*models.PortAuthStateEnum`](../../doc/models/port-auth-state-enum.md) | Optional | Port authentication state. enum: `""`, `authenticated`, `authenticating`, `held`, `init` |
| `Disabled` | `*bool` | Optional, Read-only | Indicates if interface is disabled |
| `ForSite` | `*bool` | Optional, Read-only | Whether the reporting device is scoped to a specific site |
| `FullDuplex` | `*bool` | Optional | Indicates full or half duplex |
| `Jitter` | `*float64` | Optional, Read-only | Last sampled jitter of the interface |
| `Latency` | `*float64` | Optional, Read-only | Last sampled latency of the interface |
| `Loss` | `*float64` | Optional, Read-only | Last sampled loss of the interface |
| `LteIccid` | `models.Optional[string]` | Optional | LTE ICCID value, Check for null/empty |
| `LteImei` | `models.Optional[string]` | Optional | LTE IMEI value, Check for null/empty |
| `LteImsi` | `models.Optional[string]` | Optional | LTE IMSI value, Check for null/empty |
| `MacCount` | `*int` | Optional, Read-only | Number of MAC addresses in the forwarding table |
| `MacLimit` | `*int` | Optional, Read-only | Limit on number of dynamically learned macs<br><br>**Constraints**: `>= 0` |
| `NeighborMac` | `string` | Required, Read-only | chassis identifier of the chassis type listed |
| `NeighborPortDesc` | `*string` | Optional, Read-only | Description supplied by the system on the interface E.g. "GigabitEthernet2/0/39" |
| `NeighborSystemName` | `*string` | Optional, Read-only | Name supplied by the system on the interface E.g. neighbor system name E.g. "Kumar-Acc-SW.mist.local" |
| `PoeDisabled` | `*bool` | Optional, Read-only | Is the POE configured not be disabled. |
| `PoeMode` | [`*models.StatsSwitchPortPoeModeEnum`](../../doc/models/stats-switch-port-poe-mode-enum.md) | Optional | enum: `802.3af`, `802.3at`, `802.3bt` |
| `PoeOn` | `*bool` | Optional, Read-only | Is the device attached to POE |
| `PortId` | `string` | Required, Read-only | Identifier of the port reporting these statistics |
| `PortMac` | `string` | Required, Read-only | MAC address assigned to the interface |
| `PortUsage` | `*string` | Optional | Logical usage assigned to the port |
| `PowerDraw` | `*float64` | Optional, Read-only | Amount of power being used by the interface at the time the command is executed. Unit in watts. |
| `RxBcastPkts` | `*int` | Optional, Read-only | Number of broadcast packets received on the interface |
| `RxBps` | `models.Optional[int64]` | Optional, Read-only | Rate of receiving traffic, bits/seconds, last known |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxErrors` | `*int` | Optional, Read-only | Number of receive errors observed on the interface |
| `RxMcastPkts` | `*int` | Optional, Read-only | Number of multicast packets received on the interface |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `Speed` | `*int` | Optional, Read-only | Current link speed of the port, in Mbps |
| `StpRole` | [`*models.PortStpRoleEnum`](../../doc/models/port-stp-role-enum.md) | Optional | Spanning Tree Protocol role for the port. enum: `""`, `alternate`, `backup`, `designated`, `disabled`, `root`, `root-prevented` |
| `StpState` | [`*models.PortStpStateEnum`](../../doc/models/port-stp-state-enum.md) | Optional | Spanning Tree Protocol state for the port. enum: `""`, `blocking`, `disabled`, `forwarding`, `learning`, `listening` |
| `TxBcastPkts` | `*int` | Optional, Read-only | Number of broadcast packets transmitted on the interface |
| `TxBps` | `models.Optional[int64]` | Optional, Read-only | Rate of transmitting traffic, bits/seconds, last known |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxErrors` | `*int` | Optional, Read-only | Number of transmit errors observed on the interface |
| `TxMcastPkts` | `*int` | Optional, Read-only | Number of multicast packets transmitted on the interface |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Type` | [`*models.StatsSwitchPortTypeEnum`](../../doc/models/stats-switch-port-type-enum.md) | Optional | device type. enum: `ap`, `ble`, `gateway`, `mxedge`, `nac`, `switch` |
| `Unconfigured` | `*bool` | Optional, Read-only | Indicates if interface is unconfigured |
| `Up` | `*bool` | Optional, Read-only | Indicates if interface is up |
| `XcvrModel` | `*string` | Optional, Read-only | Optic Slot ModelName, Check for null/empty |
| `XcvrPartNumber` | `*string` | Optional, Read-only | Optic Slot Partnumber, Check for null/empty |
| `XcvrSerial` | `*string` | Optional, Read-only | Optic Slot SerialNumber, Check for null/empty |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsGatewayPort := models.StatsGatewayPort{
        Active:               models.ToPointer(false),
        AuthState:            models.ToPointer(models.PortAuthStateEnum_HELD),
        Disabled:             models.ToPointer(false),
        ForSite:              models.ToPointer(false),
        FullDuplex:           models.ToPointer(true),
        NeighborMac:          "64d814353400",
        NeighborPortDesc:     models.ToPointer("GigabitEthernet1/0/21"),
        NeighborSystemName:   models.ToPointer("CORP-D-SW-2"),
        PortId:               "ge-0/0/0",
        PortMac:              "5c4527a96580",
        PortUsage:            models.ToPointer("lan"),
        RxBps:                models.NewOptional(models.ToPointer(int64(60003))),
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        Speed:                models.ToPointer(1000),
        TxBps:                models.NewOptional(models.ToPointer(int64(634301))),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        Type:                 models.ToPointer(models.StatsSwitchPortTypeEnum_GATEWAY),
        XcvrModel:            models.ToPointer("SFP+-10G-SR"),
        XcvrPartNumber:       models.ToPointer("740-021487"),
        XcvrSerial:           models.ToPointer("N6AA9HT"),
    }

}
```

