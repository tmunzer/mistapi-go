
# Ap Port Config

Ethernet port behavior settings for an access point

## Structure

`ApPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | Whether this AP Ethernet port is disabled<br><br>**Default**: `false` |
| `DynamicVlan` | [`*models.ApPortConfigDynamicVlan`](../../doc/models/ap-port-config-dynamic-vlan.md) | Optional | Dynamic VLAN assignment settings for AP port authentication |
| `EnableMacAuth` | `*bool` | Optional | Whether MAC authentication is enabled on this AP port<br><br>**Default**: `false` |
| `Forwarding` | [`*models.ApPortConfigForwardingEnum`](../../doc/models/ap-port-config-forwarding-enum.md) | Optional | enum:<br><br>* `all`: local breakout, All VLANs<br>* `limited`: local breakout, only the VLANs configured in `port_vlan_id` and `vlan_ids`<br>* `mxtunnel`: central breakout to an Org Mist Edge (requires `mxtunnel_id`)<br>* `site_mxedge`: central breakout to a Site Mist Edge (requires `mxtunnel_name`)<br>* `wxtunnel`': central breakout to an Org WxTunnel (requires `wxtunnel_id`)<br><br>**Default**: `"all"` |
| `MacAuthPreferred` | `*bool` | Optional | When `true`, we'll do dot1x then mac_auth. enable this to prefer mac_auth<br><br>**Default**: `false` |
| `MacAuthProtocol` | [`*models.ApPortConfigMacAuthProtocolEnum`](../../doc/models/ap-port-config-mac-auth-protocol-enum.md) | Optional | if `enable_mac_auth`==`true`, allows user to select an authentication protocol. enum: `eap-md5`, `eap-peap`, `pap`<br><br>**Default**: `"pap"` |
| `MistNac` | [`*models.WlanMistNac`](../../doc/models/wlan-mist-nac.md) | Optional | Mist NAC RADIUS settings for a WLAN |
| `MxTunnelId` | `*uuid.UUID` | Optional | If `forwarding`==`mxtunnel`, vlan_ids comes from mxtunnel |
| `MxtunnelName` | `*string` | Optional | If `forwarding`==`site_mxedge`, vlan_ids comes from site_mxedge (`mxtunnels` under site setting) |
| `PortAuth` | [`*models.ApPortConfigPortAuthEnum`](../../doc/models/ap-port-config-port-auth-enum.md) | Optional | When doing port auth. enum: `dot1x`, `none`<br><br>**Default**: `"none"` |
| `PortVlanId` | `*int` | Optional | If `forwarding`==`limited`. VLAN ID allowed on this AP Ethernet port<br><br>**Constraints**: `>= 1`, `<= 4094` |
| `RadiusConfig` | [`*models.RadiusConfig`](../../doc/models/radius-config.md) | Optional | Junos RADIUS authentication and accounting configuration |
| `Radsec` | [`*models.Radsec`](../../doc/models/radsec.md) | Optional | RadSec settings for sending RADIUS traffic over TLS |
| `VlanId` | `*int` | Optional | Optional to specify the VLAN ID for a tunnel if forwarding is for `wxtunnel`, `mxtunnel` or `site_mxedge`.<br><br>* if vlan_id is not specified then it will use first one in vlan_ids[] of the mxtunnel.<br>* if forwarding == site_mxedge, vlan_ids comes from site_mxedge (`mxtunnels` under site setting)<br><br>**Constraints**: `>= 1`, `<= 4094` |
| `VlanIds` | `*string` | Optional | If `forwarding`==`limited`, comma separated list of additional VLAN IDs allowed on this port |
| `WxtunnelId` | `*uuid.UUID` | Optional | If `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session |
| `WxtunnelRemoteId` | `*string` | Optional | If `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    apPortConfig := models.ApPortConfig{
        Disabled:             models.ToPointer(false),
        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
            DefaultVlanId:        models.ToPointer(34),
            Enabled:              models.ToPointer(false),
            Type:                 models.ToPointer(models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
            Vlans:                map[string]string{
                "key0": "vlans1",
            },
        }),
        EnableMacAuth:        models.ToPointer(false),
        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_ALL),
        MacAuthPreferred:     models.ToPointer(false),
        MacAuthProtocol:      models.ToPointer(models.ApPortConfigMacAuthProtocolEnum_PAP),
        MxTunnelId:           models.ToPointer(uuid.MustParse("08cd7499-5841-51c8-e663-fb16b6f3b45e")),
        PortAuth:             models.ToPointer(models.ApPortConfigPortAuthEnum_NONE),
        PortVlanId:           models.ToPointer(1),
        VlanId:               models.ToPointer(9),
        VlanIds:              models.ToPointer("10,20,30"),
        WxtunnelId:           models.ToPointer(uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")),
        WxtunnelRemoteId:     models.ToPointer("wifiguest"),
    }

}
```

