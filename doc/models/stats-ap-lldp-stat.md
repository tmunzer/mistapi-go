
# Stats Ap Lldp Stat

LLDP Stat (neighbor information, power negotiations)

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApLldpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisId` | `models.Optional[string]` | Optional | - |
| `LldpMedSupported` | `models.Optional[bool]` | Optional | Whether it support LLDP-MED |
| `MgmtAddr` | `models.Optional[string]` | Optional | Switch’s management address (if advertised), can be IPv4, IPv6, or MAC |
| `MgmtAddrs` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `PortDesc` | `models.Optional[string]` | Optional | ge-0/0/4 |
| `PortId` | `models.Optional[string]` | Optional | - |
| `PowerAllocated` | `models.Optional[float64]` | Optional | In mW, provided/allocated by PSE |
| `PowerDraw` | `models.Optional[float64]` | Optional | In mW, total power needed by PD |
| `PowerRequestCount` | `models.Optional[int]` | Optional | Number of negotiations, if it keeps increasing, we don’ t have a stable power |
| `PowerRequested` | `models.Optional[float64]` | Optional | In mW, the current power requested by PD |
| `SystemDesc` | `models.Optional[string]` | Optional | Description provided by switch |
| `SystemName` | `models.Optional[string]` | Optional | Name of the switch |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port_id": "ge-0/0/4",
  "system_desc": "uniper Networks, Inc. ex4300-48t internet router, kernel JUNOS 20.4R3-S7.2, Build date: 2023-04-21 19:47:18 UTC Copyright (c) 1996-2023 Juniper Networks, Inc.",
  "system_name": "Core-AE23",
  "chassis_id": "chassis_id6",
  "lldp_med_supported": false,
  "mgmt_addr": "mgmt_addr4",
  "mgmt_addrs": [
    "mgmt_addrs9",
    "mgmt_addrs0"
  ],
  "port_desc": "port_desc0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

