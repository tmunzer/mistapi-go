
# Bgp Config Neighbors

## Structure

`BgpConfigNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | If true, the BGP session to this neighbor will be administratively disabled/shutdown<br>**Default**: `false` |
| `ExportPolicy` | `*string` | Optional | - |
| `HoldTime` | `*int` | Optional | **Default**: `90`<br>**Constraints**: `>= 0`, `<= 65535` |
| `ImportPolicy` | `*string` | Optional | - |
| `MultihopTtl` | `*int` | Optional | assuming BGP neighbor is directly connected<br>**Constraints**: `>= 0`, `<= 255` |
| `NeighborAs` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": false,
  "hold_time": 90,
  "export_policy": "export_policy4",
  "import_policy": "import_policy8",
  "multihop_ttl": 36
}
```

