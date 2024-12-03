
# Config Vc Port Member

*This model accepts additional fields of type interface{}.*

## Structure

`ConfigVcPortMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `float64` | Required | - |
| `VcPorts` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "member": 64.66,
  "vc_ports": [
    "vc_ports4",
    "vc_ports5",
    "vc_ports6"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

