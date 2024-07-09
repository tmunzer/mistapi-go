
# Switch Network

A network represents a network segment. It can either represent a VLAN (then usually ties to a L3 subnet), optionally associate it with a subnet which can later be used to create addition routes. Used for ports doing `family ethernet-switching`. It can also be a pure L3-subnet that can then be used against a port that with `family inet`.

## Structure

`SwitchNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Isolation` | `*bool` | Optional | whether to stop clients to talk to each other, default is false (when enabled, a unique isolation_vlan_id is required)<br>NOTE: this features requires uplink device to also a be Juniper device and `inter_switch_link` to be set<br>**Default**: `false` |
| `IsolationVlanId` | `*string` | Optional | - |
| `Subnet` | `*string` | Optional | optional for pure switching, required when L3 / routing features are used |
| `VlanId` | `int` | Required | - |

## Example (as JSON)

```json
{
  "isolation": false,
  "isolation_vlan_id": "3070",
  "vlan_id": 84,
  "subnet": "subnet6"
}
```

