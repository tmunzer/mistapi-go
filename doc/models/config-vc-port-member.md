
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
  "member": 40.02,
  "vc_ports": [
    "vc_ports0",
    "vc_ports1"
  ]
}
```

