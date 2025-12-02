
# Stats Ap Lldp Stat

LLDP neighbor information and power negotiations. For backward compatibility, when multiple neighbors exist, only information from the first neighbor is displayed.

## Structure

`StatsApLldpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisId` | `models.Optional[string]` | Optional | - |
| `LldpMedSupported` | `models.Optional[bool]` | Optional | Whether it support LLDP-MED |
| `MgmtAddr` | `models.Optional[string]` | Optional | Management IP address of the switch |
| `MgmtAddrs` | `[]string` | Optional | List of management IP addresses (IPv4 and IPv6) |
| `PortDesc` | `models.Optional[string]` | Optional | Port description, e.g. “2/20”, “Port 2 on Switch0” |
| `PortId` | `models.Optional[string]` | Optional | Port identifier |
| `PowerAllocated` | `models.Optional[float64]` | Optional | In mW, power allocated by PSE |
| `PowerAvail` | `*int` | Optional | In mW, total Power Avail at AP from pwr source |
| `PowerBudget` | `*int` | Optional | In mW, surplus if positive or deficit if negative |
| `PowerConstrained` | `*bool` | Optional | Whether power is insufficient |
| `PowerDraw` | `models.Optional[float64]` | Optional | In mW, total power needed by PD |
| `PowerNeeded` | `*int` | Optional | In mW, total Power needed incl Peripherals |
| `PowerOpmode` | `*string` | Optional | Constrained mode |
| `PowerRequestCount` | `models.Optional[int]` | Optional | Number of negotiations, if it keeps increasing, we don’ t have a stable power |
| `PowerRequested` | `models.Optional[float64]` | Optional | In mW, power requested by PD |
| `PowerSrc` | `*string` | Optional | Single power source (DC Input / PoE 802.3at / PoE 802.3af / PoE 802.3bt / MULTI-PD / LLDP / ? (unknown)). |
| `PowerSrcs` | `[]string` | Optional | List of management IP addresses (IPv4 and IPv6) |
| `SystemDesc` | `models.Optional[string]` | Optional | Description provided by switch |
| `SystemName` | `models.Optional[string]` | Optional | Name of the switch |

## Example (as JSON)

```json
{
  "port_desc": "2/20",
  "port_id": "ge-0/0/4",
  "system_desc": "uniper Networks, Inc. ex4300-48t internet router, kernel JUNOS 20.4R3-S7.2, Build date: 2023-04-21 19:47:18 UTC Copyright (c) 1996-2023 Juniper Networks, Inc.",
  "system_name": "Core-AE23",
  "chassis_id": "chassis_id6",
  "lldp_med_supported": false,
  "mgmt_addr": "mgmt_addr4",
  "mgmt_addrs": [
    "mgmt_addrs9",
    "mgmt_addrs0"
  ]
}
```

