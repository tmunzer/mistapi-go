
# Wxlan Tunnel

WxLAn Tunnel

*This model accepts additional fields of type interface{}.*

## Structure

`WxlanTunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Dmvpn` | [`*models.WxlanTunnelDmvpn`](../../doc/models/wxlan-tunnel-dmvpn.md) | Optional | Dynamic Multipoint VPN configurations |
| `ForMgmt` | `*bool` | Optional | Determined during creation time and cannot be toggled. A management tunnel cannot be used by wxlan rule or by wlan<br>**Default**: `false` |
| `ForSite` | `*bool` | Optional | - |
| `HelloInterval` | `*int` | Optional | In seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries.<br>**Default**: `60`<br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `*int` | Optional | **Default**: `7`<br>**Constraints**: `>= 2`, `<= 30` |
| `Hostname` | `*string` | Optional | Optional, overwrite the hostname in SCCRQ control message, default is  or null, %H and %M can be used, which will be replace with corresponding values:<br><br>* %H: name of the ap if provided (and will be stripped so it can be used for hostname) and fallbacks to MAC<br>* %M: MAC (e.g. 5c5b350e0060) |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Ipsec` | [`*models.WxlanTunnelIpsec`](../../doc/models/wxlan-tunnel-ipsec.md) | Optional | IPSec-related configurations; requires DMVPN be enabled |
| `IsStatic` | `*bool` | Optional | Whether it’s static/unmanaged (i.e. no control session). As the session configurations are not compatible, cannot be toggled.<br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Mtu` | `*int` | Optional | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 1500` |
| `Name` | `string` | Required | The name of the tunnel |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Peers` | `[]string` | Optional | List of remote peers’ IP or hostname |
| `RouterId` | `*string` | Optional | Optional, overwrite the router-id in SCCRQ control message, default is "" or null, can also be an IPv4 address |
| `Secret` | `*string` | Optional | Secret, ‘’ if no auth is used |
| `Sessions` | [`[]models.WxlanTunnelSession`](../../doc/models/wxlan-tunnel-session.md) | Optional | Sessions to be established with the tunnel. Has to be >= 1 in order for this tunnel to be useful. For management tunnel, it can only have 1<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `UdpPort` | `*int` | Optional | UDP port if `use_udp`==`true` |
| `UseUdp` | `*bool` | Optional | Whether to use UDP instead of IP (proto=115, which is default of L2TPv3)<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "for_mgmt": false,
  "hello_interval": 60,
  "hello_retries": 7,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "is_static": false,
  "mtu": 0,
  "name": "name0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "use_udp": false,
  "created_time": 45.2,
  "dmvpn": {
    "enabled": false,
    "holding_time": 188,
    "host_routes": [
      "host_routes6",
      "host_routes7"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "for_site": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

