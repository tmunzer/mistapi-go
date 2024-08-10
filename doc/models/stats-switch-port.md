
# Stats Switch Port

Switch port statistics

## Structure

`StatsSwitchPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Active` | `*bool` | Optional | Indicates if interface is active/inactive |
| `AuthState` | [`*models.StatsSwitchPortAuthStateEnum`](../../doc/models/stats-switch-port-auth-state-enum.md) | Optional | if `up`==`true` and has Authenticator role. enum: `authenticated`, `authenticating`, `held`, `init` |
| `ForSite` | `*bool` | Optional | - |
| `FullDuplex` | `*bool` | Optional | indicates full or half duplex |
| `Jitter` | `*float64` | Optional | Last sampled jitter of the interface |
| `Latency` | `*float64` | Optional | Last sampled latency of the interface |
| `Loss` | `*float64` | Optional | Last sampled loss of the interface |
| `LteIccid` | `models.Optional[string]` | Optional | LTE ICCID value, Check for null/empty |
| `LteImei` | `models.Optional[string]` | Optional | LTE IMEI value, Check for null/empty |
| `LteImsi` | `models.Optional[string]` | Optional | LTE IMSI value, Check for null/empty |
| `Mac` | `string` | Required | - |
| `MacCount` | `*int` | Optional | Number of mac addresses in the forwarding table |
| `MacLimit` | `*int` | Optional | Limit on number of dynamically learned macs<br>**Constraints**: `>= 0` |
| `NeighborMac` | `string` | Required | chassis identifier of the chassis type listed |
| `NeighborPortDesc` | `*string` | Optional | description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” |
| `NeighborSystemName` | `*string` | Optional | name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” |
| `OrgId` | `uuid.UUID` | Required | - |
| `PoeDisabled` | `*bool` | Optional | is the POE configured not be disabled. |
| `PoeMode` | [`*models.StatsSwitchPortPoeModeEnum`](../../doc/models/stats-switch-port-poe-mode-enum.md) | Optional | enum: `802.3af`, `802.3at`, `802.3bt` |
| `PoeOn` | `*bool` | Optional | is the device attached to POE |
| `PortId` | `string` | Required | - |
| `PortMac` | `string` | Required | interface mac address |
| `PortUsage` | [`*models.StatsSwitchPortPortUsageEnum`](../../doc/models/stats-switch-port-port-usage-enum.md) | Optional | gateway port usage. enum: `lan` |
| `PowerDraw` | `*float64` | Optional | Amount of power being used by the interface at the time the command is executed. Unit in watts. |
| `RxBcastPkts` | `*int` | Optional | Broadcast input packets |
| `RxBps` | `*int` | Optional | Input rate |
| `RxBytes` | `int64` | Required | rx bytes |
| `RxErrors` | `*int` | Optional | Input errors |
| `RxMcastPkts` | `*int` | Optional | Multicast input packets |
| `RxPkts` | `int` | Required | rx packets |
| `SiteId` | `uuid.UUID` | Required | - |
| `Speed` | `*int` | Optional | port speed |
| `StpRole` | [`*models.StatsSwitchPortStpRoleEnum`](../../doc/models/stats-switch-port-stp-role-enum.md) | Optional | if `up`==`true`. enum: `alternate`, `backup`, `designated`, `root`, `root-prevented` |
| `StpState` | [`*models.StatsSwitchPortStpStateEnum`](../../doc/models/stats-switch-port-stp-state-enum.md) | Optional | if `up`==`true`. enum: `blocking`, `disabled`, `forwarding`, `learning`, `listening` |
| `TxBcastPkts` | `*int` | Optional | Broadcast output packets |
| `TxBps` | `*int` | Optional | Output rate |
| `TxBytes` | `int64` | Required | tx bytes |
| `TxErrors` | `*int` | Optional | Output errors |
| `TxMcastPkts` | `*int` | Optional | Multicast output packets |
| `TxPkts` | `int` | Required | tx packets |
| `Type` | [`*models.StatsSwitchPortTypeEnum`](../../doc/models/stats-switch-port-type-enum.md) | Optional | device type. enum: `ap`, `ble`, `gateway`, `mxedge`, `nac`, `switch` |
| `Unconfigured` | `*bool` | Optional | indicates if interface is unconfigured |
| `Up` | `*bool` | Optional | indicates if interface is up |
| `XcvrModel` | `*string` | Optional | Optic Slot ModelName, Check for null/empty |
| `XcvrPartNumber` | `*string` | Optional | Optic Slot Partnumber, Check for null/empty |
| `XcvrSerial` | `*string` | Optional | Optic Slot SerialNumber, Check for null/empty |

## Example (as JSON)

```json
{
  "full_duplex": true,
  "mac": "5c4527a96580",
  "neighbor_mac": "64d814353400",
  "neighbor_port_desc": "GigabitEthernet1/0/21",
  "neighbor_system_name": "CORP-D-SW-2",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "port_id": "ge-0/0/0",
  "port_mac": "5c4527a96580",
  "rx_bytes": 4563443626,
  "rx_pkts": 0,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "speed": 1000,
  "tx_bytes": 11299516780,
  "tx_pkts": 492176,
  "type": "gateway",
  "xcvr_model": "SFP+-10G-SR",
  "xcvr_part_number": "740-021487",
  "xcvr_serial": "N6AA9HT",
  "active": false,
  "auth_state": "held",
  "for_site": false,
  "jitter": 198.02
}
```

