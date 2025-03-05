
# Ap Port Config Forwarding Enum

enum:

* `all`: local breakout, All VLANs
* `limited`: local breakout, only the VLANs configured in `port_vlan_id` and `vlan_ids`
* `mxtunnel`: central breakout to an Org Mist Edge (requires `mxtunnel_id`)
* `site_mxedge`: central breakout to a Site Mist Edge (requires `mxtunnel_name`)
* `wxtunnel`': central breakout to an Org WxTunnel (requires `wxtunnel_id`)

## Enumeration

`ApPortConfigForwardingEnum`

## Fields

| Name |
|  --- |
| `all` |
| `limited` |
| `mxtunnel` |
| `site_mxedge` |
| `wxtunnel` |

## Example

```
all
```

