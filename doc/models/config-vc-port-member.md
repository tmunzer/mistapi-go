
# Config Vc Port Member

## Structure

`ConfigVcPortMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `float64` | Required | - |
| `VcPorts` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "member": 64.66,
  "vc_ports": [
    "vc_ports4",
    "vc_ports5",
    "vc_ports6"
  ]
}
```

