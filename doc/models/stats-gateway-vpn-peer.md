
# Stats Gateway Vpn Peer

## Structure

`StatsGatewayVpnPeer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IsActive` | `*bool` | Optional | Redundancy status of the associated interface |
| `Jitter` | `*float64` | Optional | Jitter in milliseconds<br><br>**Constraints**: `>= 0` |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Latency` | `*float64` | Optional | Latency in milliseconds<br><br>**Constraints**: `>= 0` |
| `Loss` | `*float64` | Optional | Packet loss in percentage<br><br>**Constraints**: `>= 0`, `<= 100` |
| `Mos` | `*float64` | Optional | Mean Opinion Score, a measure of the quality of the VPN link<br><br>**Constraints**: `>= 0`, `<= 5` |
| `Mtu` | `*int` | Optional | - |
| `PeerMac` | `*string` | Optional | Peer router mac address<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerPortId` | `*string` | Optional | Peer router device interface<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerRouterName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `PeerSiteId` | `*uuid.UUID` | Optional | - |
| `PortId` | `*string` | Optional | Router device interface<br><br>**Constraints**: *Minimum Length*: `1` |
| `RouterName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Type` | `*string` | Optional | `ipsec`for SRX, `svr` for 128T<br><br>**Constraints**: *Minimum Length*: `1` |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "last_seen": 1470417522.0,
  "peer_site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "is_active": false,
  "jitter": 234.76,
  "latency": 42.94,
  "loss": 100.0
}
```

