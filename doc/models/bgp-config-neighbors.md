
# Bgp Config Neighbors

*This model accepts additional fields of type interface{}.*

## Structure

`BgpConfigNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | If true, the BGP session to this neighbor will be administratively disabled/shutdown<br>**Default**: `false` |
| `ExportPolicy` | `*string` | Optional | - |
| `HoldTime` | `*int` | Optional | **Default**: `90`<br>**Constraints**: `>= 0`, `<= 65535` |
| `ImportPolicy` | `*string` | Optional | - |
| `MultihopTtl` | `*int` | Optional | Assuming BGP neighbor is directly connected<br>**Constraints**: `>= 0`, `<= 255` |
| `NeighborAs` | [`*models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967295 |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": false,
  "hold_time": 90,
  "neighbor_as": 65000,
  "export_policy": "export_policy4",
  "import_policy": "import_policy8",
  "multihop_ttl": 72,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

