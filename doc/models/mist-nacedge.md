
# Mist Nacedge

*This model accepts additional fields of type interface{}.*

## Structure

`MistNacedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthTtl` | `*int` | Optional | Cache of last auth result; in seconds<br>**Default**: `604800`<br>**Constraints**: `>= 60`, `<= 2592000` |
| `DefaultDot1xVlan` | `*string` | Optional | Default vlan for all dot1x devices, if different from default_vlan |
| `DefaultVlan` | `*string` | Optional | Default vlan to assign for devices not in the cache |
| `Enabled` | `*bool` | Optional | - |
| `MxedgeHosts` | `[]string` | Optional | List of NAC Edges in this site |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auth_ttl": 604800,
  "default_dot1x_vlan": "20",
  "default_vlan": "test_vlan",
  "mxedge_hosts": [
    "mxedge1.local"
  ],
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

