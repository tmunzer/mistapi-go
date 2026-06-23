
# Bgp Config

BGP session configuration. BFD is enabled when either bfd_minimum_interval or bfd_multiplier is configured

## Structure

`BgpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKey` | `*string` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan` |
| `BfdMinimumInterval` | `models.Optional[int]` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`, when bfd_multiplier is configured alone. Default:<br><br>* 1000 if `type`==`external`<br>* 350 `type`==`internal`<br><br>**Default**: `350`<br><br>**Constraints**: `>= 1`, `<= 255000` |
| `BfdMultiplier` | `models.Optional[int]` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`, when bfd_minimum_interval_is_configured alone<br><br>**Default**: `3`<br><br>**Constraints**: `>= 1`, `<= 255` |
| `DisableBfd` | `*bool` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. BFD provides faster path failure detection and is enabled by default<br><br>**Default**: `false` |
| `Export` | `*string` | Optional | Routing policy applied to routes exported by this BGP session |
| `ExportPolicy` | `*string` | Optional | Default export policies if no per-neighbor policies defined |
| `ExtendedV4Nexthop` | `*bool` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. By default, either inet/net6 unicast depending on neighbor IP family (v4 or v6). For v6 neighbors, to exchange v4 nexthop, which allows dual-stack support, enable this |
| `GracefulRestartTime` | `*int` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. `0` means disable<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 4095` |
| `HoldTime` | `*int` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. Default is 90.<br><br>**Default**: `90`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `Import` | `*string` | Optional | Routing policy applied to routes imported by this BGP session |
| `ImportPolicy` | `*string` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. Default import policies if no per-neighbor policies defined |
| `LocalAs` | [`*models.BgpLocalAs`](../../doc/models/containers/bgp-local-as.md) | Optional | Required if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. BGP AS, value in range 1-4294967295 |
| `NeighborAs` | [`*models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Neighbors` | [`map[string]models.BgpConfigNeighbors`](../../doc/models/bgp-config-neighbors.md) | Optional | Required if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. If per-neighbor as is desired. Property key is the neighbor address |
| `Networks` | `[]string` | Optional | Optional if `via`==`lan`. List of networks where we expect BGP neighbor to connect to/from |
| `NoPrivateAs` | `*bool` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. If true, we will not advertise private ASNs (AS 64512-65534) to this neighbor<br><br>**Default**: `false` |
| `NoReadvertiseToOverlay` | `*bool` | Optional | Optional if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. By default, we'll re-advertise all learned BGP routers toward overlay<br><br>**Default**: `false` |
| `TunnelName` | `*string` | Optional | Optional if `via`==`tunnel`; tunnel name used for this BGP session |
| `Type` | [`*models.BgpConfigTypeEnum`](../../doc/models/bgp-config-type-enum.md) | Optional | Required if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. enum: `external`, `internal`<br><br>**Constraints**: *Minimum Length*: `1` |
| `Via` | [`models.BgpConfigViaEnum`](../../doc/models/bgp-config-via-enum.md) | Required | enum: `lan`, `tunnel`, `vpn`, `wan`<br><br>**Default**: `"lan"` |
| `VpnName` | `*string` | Optional | Optional if `via`==`vpn`; VPN name used for this BGP session |
| `WanName` | `*string` | Optional | Optional if `via`==`wan`; WAN interface name used for this BGP session |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    bgpConfig := models.BgpConfig{
        AuthKey:                models.ToPointer("auth_key4"),
        BfdMinimumInterval:     models.NewOptional(models.ToPointer(350)),
        BfdMultiplier:          models.NewOptional(models.ToPointer(3)),
        DisableBfd:             models.ToPointer(false),
        Export:                 models.ToPointer("export2"),
        GracefulRestartTime:    models.ToPointer(0),
        HoldTime:               models.ToPointer(90),
        LocalAs:                models.ToPointer(),
        NeighborAs:             models.ToPointer(),
        NoPrivateAs:            models.ToPointer(false),
        NoReadvertiseToOverlay: models.ToPointer(false),
        Via:                    models.BgpConfigViaEnum_LAN,
    }

}
```

