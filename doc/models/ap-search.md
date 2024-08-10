
# Ap Search

## Structure

`ApSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24Bandwith` | `*string` | Optional | Bandwith of band_24 |
| `Band24Channel` | `*int` | Optional | Channel of band_24 |
| `Band24Power` | `*int` | Optional | - |
| `Band5Bandwith` | `*string` | Optional | Bandwith of band_5 |
| `Band5Channel` | `*int` | Optional | Channel of band_5 |
| `Band5Power` | `*int` | Optional | - |
| `Band6Bandwith` | `*string` | Optional | - |
| `Band6Channel` | `*int` | Optional | Channel of band_6 |
| `Band6Power` | `*int` | Optional | - |
| `Eth0PortSpeed` | `*int` | Optional | Port speed of eth0 |
| `ExtIp` | `*string` | Optional | - |
| `Hostname` | `[]string` | Optional | partial / full hostname |
| `Ip` | `*string` | Optional | ip address |
| `LldpMgmtAddr` | `*string` | Optional | LLDP management ip address |
| `LldpPortDesc` | `*string` | Optional | - |
| `LldpPortId` | `*string` | Optional | LLDP port id |
| `LldpPowerAllocated` | `*int` | Optional | - |
| `LldpPowerDraw` | `*int` | Optional | - |
| `LldpSystemDesc` | `*string` | Optional | LLDP system description |
| `LldpSystemName` | `*string` | Optional | LLDP system name |
| `Mac` | `*string` | Optional | device model |
| `Model` | `*string` | Optional | - |
| `MxedgeId` | `*string` | Optional | Mist Edge id, if AP is connecting to a Mist Edge |
| `MxtunnelStatus` | `*string` | Optional | MxTunnel status |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PowerConstrained` | `*bool` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Sku` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `Version` | `*string` | Optional | version |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "band_24_bandwith": "band_24_bandwith8",
  "band_24_channel": 200,
  "band_24_power": 154,
  "band_5_bandwith": "band_5_bandwith6",
  "band_5_channel": 132
}
```

