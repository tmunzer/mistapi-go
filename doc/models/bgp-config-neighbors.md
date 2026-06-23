
# Bgp Config Neighbors

Per-neighbor BGP session settings

## Structure

`BgpConfigNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | If true, the BGP session to this neighbor will be administratively disabled/shutdown<br><br>**Default**: `false` |
| `ExportPolicy` | `*string` | Optional | Export policy applied only to this BGP neighbor |
| `HoldTime` | `*int` | Optional | BGP hold time for this neighbor, in seconds<br><br>**Default**: `90`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `ImportPolicy` | `*string` | Optional | Import policy applied only to this BGP neighbor |
| `MultihopTtl` | `*int` | Optional | Assuming BGP neighbor is directly connected<br><br>**Constraints**: `>= 0`, `<= 255` |
| `NeighborAs` | [`models.BgpAs`](../../doc/models/containers/bgp-as.md) | Required | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `TunnelVia` | [`*models.TunnelViaEnum`](../../doc/models/tunnel-via-enum.md) | Optional | If `via`==`tunnel`, specifies which tunnel (primary/secondary) this neighbor is associated with. enum: `primary`, `secondary`<br><br>**Default**: `"primary"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    bgpConfigNeighbors := models.BgpConfigNeighbors{
        Disabled:             models.ToPointer(false),
        ExportPolicy:         models.ToPointer("export_policy4"),
        HoldTime:             models.ToPointer(90),
        ImportPolicy:         models.ToPointer("import_policy8"),
        MultihopTtl:          models.ToPointer(128),
        NeighborAs:           models.BgpAsContainer.FromNumber(65000),
        TunnelVia:            models.ToPointer(models.TunnelViaEnum_PRIMARY),
    }

}
```

