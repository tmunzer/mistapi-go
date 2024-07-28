
# Ap Stats Lldp Stat

LLDP Stat (neighbor information, power negotiations)

## Structure

`ApStatsLldpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisId` | `models.Optional[string]` | Optional | - |
| `LldpMedSupported` | `models.Optional[bool]` | Optional | whether it support LLDP-MED |
| `MgmtAddr` | `models.Optional[string]` | Optional | switch’s management address (if advertised), can be IPv4, IPv6, or MAC |
| `MgmtAddrs` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `PortDesc` | `models.Optional[string]` | Optional | ge-0/0/4 |
| `PortId` | `models.Optional[string]` | Optional | - |
| `PowerAllocated` | `models.Optional[float64]` | Optional | in mW, provided/allocated by PSE |
| `PowerDraw` | `models.Optional[float64]` | Optional | in mW, total power needed by PD |
| `PowerRequestCount` | `models.Optional[int]` | Optional | number of negotiations, if it keeps increasing, we don’ t have a stable power |
| `PowerRequested` | `models.Optional[float64]` | Optional | in mW, the current power requested by PD |
| `SystemDesc` | `models.Optional[string]` | Optional | description provided by switch |
| `SystemName` | `models.Optional[string]` | Optional | name of the switch |

## Example (as JSON)

```json
{
  "port_id": "ge-0/0/4",
  "system_desc": "uniper Networks, Inc. ex4300-48t internet router, kernel JUNOS 20.4R3-S7.2, Build date: 2023-04-21 19:47:18 UTC Copyright (c) 1996-2023 Juniper Networks, Inc.",
  "system_name": "Core-AE23",
  "chassis_id": "chassis_id8",
  "lldp_med_supported": false,
  "mgmt_addr": "mgmt_addr6",
  "mgmt_addrs": [
    "mgmt_addrs1",
    "mgmt_addrs2",
    "mgmt_addrs3"
  ],
  "port_desc": "port_desc2"
}
```

