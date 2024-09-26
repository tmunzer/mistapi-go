
# Optional Stat Vpn Peer

## Structure

`OptionalStatVpnPeer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IsActive` | `*bool` | Optional | Redundancy status of the associated interface |
| `LastSeen` | `*float64` | Optional | - |
| `Latency` | `*float64` | Optional | - |
| `Mos` | `*float64` | Optional | - |
| `Mtu` | `*int` | Optional | - |
| `PeerMac` | `*string` | Optional | peer router mac address<br>**Constraints**: *Minimum Length*: `1` |
| `PeerPortId` | `*string` | Optional | peer router device interface<br>**Constraints**: *Minimum Length*: `1` |
| `PeerRouterName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `PeerSiteId` | `*uuid.UUID` | Optional | - |
| `PortId` | `*string` | Optional | router device interface<br>**Constraints**: *Minimum Length*: `1` |
| `RouterName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Type` | `*string` | Optional | `ipsec`for SRX, `svr` for 128T<br>**Constraints**: *Minimum Length*: `1` |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "peer_site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "is_active": false,
  "last_seen": 64.7,
  "latency": 94.6,
  "mos": 183.3,
  "mtu": 232
}
```
