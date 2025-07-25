
# Switch Ospf Config

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchOspfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Areas` | [`map[string]models.SwitchOspfConfigArea`](../../doc/models/switch-ospf-config-area.md) | Optional | Property key is the area name. Defines the OSPF areas configured on the switch. |
| `Enabled` | `*bool` | Optional | Enable OSPF on the switch<br><br>**Default**: `false` |
| `ReferenceBandwidth` | [`*models.SwitchOspfConfigReferenceBandwidth`](../../doc/models/containers/switch-ospf-config-reference-bandwidth.md) | Optional | Reference bandwidth. Integer(100000) or String (10g) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "reference_bandwidth": 100000,
  "areas": {
    "key0": {
      "no_summary": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "no_summary": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "no_summary": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

