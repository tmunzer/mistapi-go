
# Switch Network

A network represents a network segment. It can either represent a VLAN (then usually ties to a L3 subnet), optionally associate it with a subnet which can later be used to create addition routes. Used for ports doing `family ethernet-switching`. It can also be a pure L3-subnet that can then be used against a port that with `family inet`.

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | only required for EVPN-VXLAN networks, IPv4 Virtual Gateway |
| `Gateway6` | `*string` | Optional | only required for EVPN-VXLAN networks, IPv6 Virtual Gateway |
| `Isolation` | `*bool` | Optional | whether to stop clients to talk to each other, default is false (when enabled, a unique isolation_vlan_id is required)<br>NOTE: this features requires uplink device to also a be Juniper device and `inter_switch_link` to be set<br>**Default**: `false` |
| `IsolationVlanId` | `*string` | Optional | - |
| `Subnet` | `*string` | Optional | optional for pure switching, required when L3 / routing features are used |
| `Subnet6` | `*string` | Optional | optional for pure switching, required when L3 / routing features are used |
| `VlanId` | [`models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "isolation": false,
  "isolation_vlan_id": "3070",
  "vlan_id": "String9",
  "gateway": "gateway0",
  "gateway6": "gateway66",
  "subnet": "subnet8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

