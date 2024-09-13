
# Mist Nacedge

## Structure

`MistNacedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthTtl` | `*int` | Optional | cache of last auth result; in seconds<br>**Default**: `604800`<br>**Constraints**: `>= 60`, `<= 2592000` |
| `DefaultDot1xVlan` | `*string` | Optional | default vlan for all dot1x devices, if different from default_vlan |
| `DefaultVlan` | `*string` | Optional | Default vlan to assign for devices not in the cache |
| `Enabled` | `*bool` | Optional | - |
| `MxedgeHosts` | `[]string` | Optional | list of NAC Edges in this site |

## Example (as JSON)

```json
{
  "auth_ttl": 604800,
  "default_dot1x_vlan": "20",
  "default_vlan": "test_vlan",
  "mxedge_hosts": [
    "mxedge1.local"
  ],
  "enabled": false
}
```

