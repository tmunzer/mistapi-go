
# Wxlan Tunnel

WxLAn Tunnel

## Structure

`WxlanTunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Dmvpn` | [`*models.WxlanTunnelDmvpn`](../../doc/models/wxlan-tunnel-dmvpn.md) | Optional | Dynamic Multipoint VPN configurations |
| `ForMgmt` | `*bool` | Optional | determined during creation time and cannot be toggled. A management tunnel cannot be used by wxlan rule or by wlan<br>**Default**: `false` |
| `ForSite` | `*bool` | Optional | - |
| `HelloInterval` | `*int` | Optional | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries.<br>**Default**: `60`<br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `*int` | Optional | **Default**: `7`<br>**Constraints**: `>= 2`, `<= 30` |
| `Hostname` | `*string` | Optional | optional, overwrite the hostname in SCCRQ control message, default is “” or null, %H and %M can be used, which will be replace with corresponding values:<br><br>* %H: name of the ap if provided (and will be stripped so it can be used for hostname) and fallbacks to MAC<br>* %M: MAC (e.g. 5c5b350e0060) |
| `Id` | `*uuid.UUID` | Optional | - |
| `Ipsec` | [`*models.WxlanTunnelIpsec`](../../doc/models/wxlan-tunnel-ipsec.md) | Optional | IPSec-related configurations; requires DMVPN be enabled |
| `IsStatic` | `*bool` | Optional | whether it’s static/unmanaged (i.e. no control session). As the session configurations are not compatible, cannot be toggled.<br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional | - |
| `Mtu` | `*int` | Optional | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 1500` |
| `Name` | `string` | Required | The name of the tunnel |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Peers` | `[]string` | Optional | list of remote peers’ IP or hostname |
| `RouterId` | `*string` | Optional | optional, overwrite the router-id in SCCRQ control message, default is “0” or null, can also be an IPv4 address |
| `Secret` | `*string` | Optional | secret, ‘’ if no auth is used |
| `Sessions` | [`[]models.WxlanTunnelSession`](../../doc/models/wxlan-tunnel-session.md) | Optional | sessions to be established with the tunnel. Has to be >= 1 in order for this tunnel to be useful. For management tunnel, it can only have 1<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `UdpPort` | `*int` | Optional | udp port if `use_udp`==`true` |
| `UseUdp` | `*bool` | Optional | whether to use UDP instead of IP (proto=115, which is default of L2TPv3)<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "for_mgmt": false,
  "hello_interval": 60,
  "hello_retries": 7,
  "is_static": false,
  "mtu": 0,
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "use_udp": false,
  "created_time": 227.84,
  "dmvpn": {
    "enabled": false,
    "holding_time": 188,
    "host_routes": [
      "host_routes6",
      "host_routes7"
    ]
  },
  "for_site": false
}
```
