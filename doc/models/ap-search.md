
# Ap Search

## Structure

`ApSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24Bandwidth` | `*string` | Optional | Bandwidth of band_24 |
| `Band24Channel` | `*int` | Optional | Channel of band_24 |
| `Band24Power` | `*int` | Optional | - |
| `Band5Bandwidth` | `*string` | Optional | Bandwidth of band_5 |
| `Band5Channel` | `*int` | Optional | Channel of band_5 |
| `Band5Power` | `*int` | Optional | - |
| `Band6Bandwidth` | `*string` | Optional | - |
| `Band6Channel` | `*int` | Optional | Channel of band_6 |
| `Band6Power` | `*int` | Optional | - |
| `Eth0PortSpeed` | `*int` | Optional | Port speed of eth0 |
| `ExtIp` | `*string` | Optional | - |
| `Hostname` | `[]string` | Optional | Partial / full hostname |
| `InactiveWiredVlans` | `[]int` | Optional | - |
| `Ip` | `*string` | Optional | IP Address |
| `LastHostname` | `*string` | Optional | - |
| `LldpMgmtAddr` | `*string` | Optional | LLDP management ip address |
| `LldpPortDesc` | `*string` | Optional | - |
| `LldpPortId` | `*string` | Optional | LLDP port id |
| `LldpPowerAllocated` | `*int` | Optional | - |
| `LldpPowerDraw` | `*int` | Optional | - |
| `LldpSystemDesc` | `*string` | Optional | LLDP system description |
| `LldpSystemName` | `*string` | Optional | LLDP system name |
| `Mac` | `*string` | Optional | Device model |
| `Model` | `*string` | Optional | - |
| `MxedgeId` | `*string` | Optional | Mist Edge id, if AP is connecting to a Mist Edge |
| `MxedgeIds` | `*string` | Optional | Comma separated list of Mist Edge ids, if AP is connecting to a Mist Edge |
| `MxtunnelStatus` | `string` | Required | MxTunnel status |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PowerConstrained` | `bool` | Required | - |
| `PowerOpmode` | `string` | Required | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Sku` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Uptime` | `*int` | Optional | - |
| `Version` | `*string` | Optional | Version |
| `Wlans` | [`[]models.ApSearchWlan`](../../doc/models/ap-search-wlan.md) | Required | - |

## Example (as JSON)

```json
{
  "mxtunnel_status": "mxtunnel_status4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "power_constrained": false,
  "power_opmode": "power_opmode0",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "wlans": [
    {
      "id": "00001c56-0000-0000-0000-000000000000",
      "ssid": "ssid8"
    }
  ],
  "band_24_bandwidth": "band_24_bandwidth2",
  "band_24_channel": 200,
  "band_24_power": 154,
  "band_5_bandwidth": "band_5_bandwidth0",
  "band_5_channel": 132
}
```

