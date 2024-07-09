
# Ap Stats Lldp Stat

LLDP Stat (neighbor information, power negotiations)

## Structure

`ApStatsLldpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisId` | `*string` | Optional | - |
| `LldpMedSupported` | `*bool` | Optional | whether it support LLDP-MED |
| `MgmtAddr` | `*string` | Optional | switch’s management address (if advertised), can be IPv4, IPv6, or MAC |
| `PortDesc` | `*string` | Optional | port description, e.g. “2/20”, “Port 2 on Switch0” |
| `PowerAllocated` | `*float64` | Optional | in mW, provided/allocated by PSE |
| `PowerDraw` | `*float64` | Optional | in mW, total power needed by PD |
| `PowerRequestCount` | `*int` | Optional | number of negotiations, if it keeps increasing, we don’t have a stable power |
| `PowerRequested` | `*float64` | Optional | in mW, the current power requested by PD |
| `SystemDesc` | `*string` | Optional | description provided by switch, e.g. “HP J9729A 2920-48G-POE+ Switch” |
| `SystemName` | `*string` | Optional | name of the switch, e.g. “TC2-OWL-Stack-01” |

## Example (as JSON)

```json
{
  "chassis_id": "chassis_id8",
  "lldp_med_supported": false,
  "mgmt_addr": "mgmt_addr6",
  "port_desc": "port_desc2",
  "power_allocated": 92.78
}
```

