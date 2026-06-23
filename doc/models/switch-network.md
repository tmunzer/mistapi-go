
# Switch Network

A network represents a network segment. It can either represent a VLAN (then usually ties to a L3 subnet), optionally associate it with a subnet which can later be used to create addition routes. Used for ports doing `family ethernet-switching`. It can also be a pure L3-subnet that can then be used against a port that with `family inet`.

## Structure

`SwitchNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | Only required for EVPN-VXLAN networks, IPv4 Virtual Gateway |
| `Gateway6` | `*string` | Optional | Only required for EVPN-VXLAN networks, IPv6 Virtual Gateway |
| `Isolation` | `*bool` | Optional | whether to stop clients to talk to each other, default is false (when enabled, a unique isolation_vlan_id is required). NOTE: this features requires uplink device to also a be Juniper device and `inter_switch_link` to be set. See also `inter_isolation_network_link` and `community_vlan_id` in port_usage<br><br>**Default**: `false` |
| `IsolationVlanId` | `*string` | Optional | Required when `isolation`==`true`. Unique VLAN ID used for client isolation |
| `Subnet` | `*string` | Optional | Optional for pure switching, required when L3 / routing features are used |
| `Subnet6` | `*string` | Optional | Optional for pure switching, required when L3 / routing features are used |
| `VlanId` | [`models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Required | VLAN ID, either numeric or expressed as a template variable string |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchNetwork := models.SwitchNetwork{
        Gateway:              models.ToPointer("gateway4"),
        Gateway6:             models.ToPointer("gateway60"),
        Isolation:            models.ToPointer(false),
        IsolationVlanId:      models.ToPointer("3070"),
        Subnet:               models.ToPointer("subnet2"),
        VlanId:               models.VlanIdWithVariableContainer.FromString("String3"),
    }

}
```

