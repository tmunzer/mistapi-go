
# Bgp Config Neighbors

## Structure

`BgpConfigNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | If true, the BGP session to this neighbor will be administratively disabled/shutdown<br><br>**Default**: `false` |
| `ExportPolicy` | `*string` | Optional | - |
| `HoldTime` | `*int` | Optional | **Default**: `90`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `ImportPolicy` | `*string` | Optional | - |
| `MultihopTtl` | `*int` | Optional | Assuming BGP neighbor is directly connected<br><br>**Constraints**: `>= 0`, `<= 255` |
| `NeighborAs` | [`models.BgpAs`](../../doc/models/containers/bgp-as.md) | Required | BGP AS, value in range 1-4294967294 |

## Example (as JSON)

```json
{
  "disabled": false,
  "hold_time": 90,
  "neighbor_as": 65000,
  "export_policy": "export_policy4",
  "import_policy": "import_policy8",
  "multihop_ttl": 72
}
```

