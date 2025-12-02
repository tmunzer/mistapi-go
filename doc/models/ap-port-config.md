
# Ap Port Config

## Structure

`ApPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `DynamicVlan` | [`*models.ApPortConfigDynamicVlan`](../../doc/models/ap-port-config-dynamic-vlan.md) | Optional | Optional dynamic vlan |
| `EnableMacAuth` | `*bool` | Optional | **Default**: `false` |
| `Forwarding` | [`*models.ApPortConfigForwardingEnum`](../../doc/models/ap-port-config-forwarding-enum.md) | Optional | enum:<br><br>* `all`: local breakout, All VLANs<br>* `limited`: local breakout, only the VLANs configured in `port_vlan_id` and `vlan_ids`<br>* `mxtunnel`: central breakout to an Org Mist Edge (requires `mxtunnel_id`)<br>* `site_mxedge`: central breakout to a Site Mist Edge (requires `mxtunnel_name`)<br>* `wxtunnel`': central breakout to an Org WxTunnel (requires `wxtunnel_id`)<br><br>**Default**: `"all"` |
| `MacAuthPreferred` | `*bool` | Optional | When `true`, we'll do dot1x then mac_auth. enable this to prefer mac_auth<br><br>**Default**: `false` |
| `MacAuthProtocol` | [`*models.ApPortConfigMacAuthProtocolEnum`](../../doc/models/ap-port-config-mac-auth-protocol-enum.md) | Optional | if `enable_mac_auth`==`true`, allows user to select an authentication protocol. enum: `eap-md5`, `eap-peap`, `pap`<br><br>**Default**: `"pap"` |
| `MistNac` | [`*models.WlanMistNac`](../../doc/models/wlan-mist-nac.md) | Optional | - |
| `MxTunnelId` | `*uuid.UUID` | Optional | If `forwarding`==`mxtunnel`, vlan_ids comes from mxtunnel |
| `MxtunnelName` | `*string` | Optional | If `forwarding`==`site_mxedge`, vlan_ids comes from site_mxedge (`mxtunnels` under site setting) |
| `PortAuth` | [`*models.ApPortConfigPortAuthEnum`](../../doc/models/ap-port-config-port-auth-enum.md) | Optional | When doing port auth. enum: `dot1x`, `none`<br><br>**Default**: `"none"` |
| `PortVlanId` | `*int` | Optional | If `forwarding`==`limited`<br><br>**Constraints**: `>= 1`, `<= 4094` |
| `RadiusConfig` | [`*models.RadiusConfig`](../../doc/models/radius-config.md) | Optional | Junos Radius config |
| `Radsec` | [`*models.Radsec`](../../doc/models/radsec.md) | Optional | RadSec settings |
| `VlanId` | `*int` | Optional | Optional to specify the vlan id for a tunnel if forwarding is for `wxtunnel`, `mxtunnel` or `site_mxedge`.<br><br>* if vlan_id is not specified then it will use first one in vlan_ids[] of the mxtunnel.<br>* if forwarding == site_mxedge, vlan_ids comes from site_mxedge (`mxtunnels` under site setting)<br><br>**Constraints**: `>= 1`, `<= 4094` |
| `VlanIds` | `[]int` | Optional | If `forwarding`==`limited`<br><br>**Constraints**: `>= 1`, `<= 4094` |
| `WxtunnelId` | `*uuid.UUID` | Optional | If `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session |
| `WxtunnelRemoteId` | `*string` | Optional | If `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session |

## Example (as JSON)

```json
{
  "disabled": false,
  "enable_mac_auth": false,
  "forwarding": "all",
  "mac_auth_preferred": false,
  "mac_auth_protocol": "pap",
  "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
  "port_auth": "none",
  "port_vlan_id": 1,
  "vlan_id": 9,
  "vlan_ids": [
    1,
    10,
    50
  ],
  "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
  "wxtunnel_remote_id": "wifiguest",
  "dynamic_vlan": {
    "default_vlan_id": 34,
    "enabled": false,
    "type": "type6",
    "vlans": {
      "key0": "vlans1"
    }
  }
}
```

