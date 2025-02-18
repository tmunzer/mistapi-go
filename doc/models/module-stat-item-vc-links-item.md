
# Module Stat Item Vc Links Item

*This model accepts additional fields of type interface{}.*

## Structure

`ModuleStatItemVcLinksItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NeighborModuleIdx` | `*int` | Optional | - |
| `NeighborPortId` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "neighbor_module_idx": 1,
  "neighbor_port_id": "vcp-255/1/0",
  "port_id": "vcp-255/1/0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

