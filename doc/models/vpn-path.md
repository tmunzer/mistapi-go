
# Vpn Path

## Structure

`VpnPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BfdProfile` | [`*models.VpnPathBfdProfileEnum`](../../doc/models/vpn-path-bfd-profile-enum.md) | Optional | enum: `broadband`, `lte`<br><br>**Default**: `"broadband"` |
| `BfdUseTunnelMode` | `*bool` | Optional | If `type`==`mesh` and for SSR only, whether to use tunnel mode<br><br>**Default**: `false` |
| `Ip` | `*string` | Optional | If different from the wan port |
| `PeerPaths` | [`map[string]models.VpnPathPeerPathsPeer`](../../doc/models/vpn-path-peer-paths-peer.md) | Optional | If `type`==`mesh`, Property key is the Peer Interface name |
| `Pod` | `*int` | Optional | **Default**: `1`<br><br>**Constraints**: `>= 1`, `<= 128` |
| `TrafficShaping` | [`*models.VpnPathTrafficShaping`](../../doc/models/vpn-path-traffic-shaping.md) | Optional | - |

## Example (as JSON)

```json
{
  "bfd_profile": "broadband",
  "bfd_use_tunnel_mode": false,
  "pod": 2,
  "ip": "ip4",
  "peer_paths": {
    "key0": {
      "preference": 144
    }
  }
}
```

