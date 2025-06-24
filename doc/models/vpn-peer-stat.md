
# Vpn Peer Stat

*This model accepts additional fields of type interface{}.*

## Structure

`VpnPeerStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IsActive` | `*bool` | Optional | Redundancy status of the associated interface |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Latency` | `*float64` | Optional | - |
| `Mac` | `*string` | Optional | Router mac address<br><br>**Constraints**: *Minimum Length*: `1` |
| `Mos` | `*float64` | Optional | - |
| `Mtu` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PeerMac` | `*string` | Optional | Peer router mac address<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerPortId` | `*string` | Optional | Peer router device interface<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerRouterName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `PeerSiteId` | `*uuid.UUID` | Optional | - |
| `PortId` | `*string` | Optional | Router device interface<br><br>**Constraints**: *Minimum Length*: `1` |
| `RouterName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Type` | `*string` | Optional | `ipsec`for SRX, `svr` for 128T<br><br>**Constraints**: *Minimum Length*: `1` |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "last_seen": 1470417522.0,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "peer_site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "is_active": false,
  "latency": 109.82,
  "mac": "mac2",
  "mos": 168.08,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

