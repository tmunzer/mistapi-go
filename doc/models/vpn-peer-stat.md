
# Vpn Peer Stat

## Structure

`VpnPeerStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IsActive` | `*bool` | Optional | Redundancy status of the associated interface |
| `LastSeen` | `*float64` | Optional | - |
| `Latency` | `*float64` | Optional | - |
| `Mac` | `*string` | Optional | router mac address<br>**Constraints**: *Minimum Length*: `1` |
| `Mos` | `*float64` | Optional | - |
| `Mtu` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PeerMac` | `*string` | Optional | peer router mac address<br>**Constraints**: *Minimum Length*: `1` |
| `PeerPortId` | `*string` | Optional | peer router device interface<br>**Constraints**: *Minimum Length*: `1` |
| `PeerRouterName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `PeerSiteId` | `*uuid.UUID` | Optional | - |
| `PortId` | `*string` | Optional | router device interface<br>**Constraints**: *Minimum Length*: `1` |
| `RouterName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Type` | `*string` | Optional | `ipsec`for SRX, `svr` for 128T<br>**Constraints**: *Minimum Length*: `1` |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "peer_site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "is_active": false,
  "last_seen": 131.36,
  "latency": 27.94,
  "mac": "mac0",
  "mos": 249.96
}
```

