
# Bgp Config

BFD is enabled when either bfd_minimum_interval or bfd_multiplier is configured

*This model accepts additional fields of type interface{}.*

## Structure

`BgpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKey` | `*string` | Optional | - |
| `BfdMinimumInterval` | `models.Optional[int]` | Optional | When bfd_multiplier is configured alone. Default:<br><br>* 1000 if `type`==`external`<br>* 350 `type`==`internal`<br>**Default**: `350`<br>**Constraints**: `>= 1`, `<= 255000` |
| `BfdMultiplier` | `models.Optional[int]` | Optional | When bfd_minimum_interval_is_configured alone<br>**Default**: `3`<br>**Constraints**: `>= 1`, `<= 255` |
| `DisableBfd` | `*bool` | Optional | BFD provides faster path failure detection and is enabled by default<br>**Default**: `false` |
| `Export` | `*string` | Optional | - |
| `ExportPolicy` | `*string` | Optional | Default export policies if no per-neighbor policies defined |
| `ExtendedV4Nexthop` | `*bool` | Optional | By default, either inet/net6 unicast depending on neighbor IP family (v4 or v6). For v6 neighbors, to exchange v4 nexthop, which allows dual-stack support, enable this |
| `GracefulRestartTime` | `*int` | Optional | `0` means disable<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 4095` |
| `HoldTime` | `*int` | Optional | **Default**: `90`<br>**Constraints**: `>= 0`, `<= 65535` |
| `Import` | `*string` | Optional | - |
| `ImportPolicy` | `*string` | Optional | Default import policies if no per-neighbor policies defined |
| `LocalAs` | `*int` | Optional | - |
| `NeighborAs` | `*int` | Optional | - |
| `Neighbors` | [`map[string]models.BgpConfigNeighbors`](../../doc/models/bgp-config-neighbors.md) | Optional | If per-neighbor as is desired. Property key is the neighbor address |
| `Networks` | `[]string` | Optional | If `type`!=`external`or `via`==`wan`networks where we expect BGP neighbor to connect to/from |
| `NoReadvertiseToOverlay` | `*bool` | Optional | By default, we'll re-advertise all learned BGP routers toward overlay<br>**Default**: `false` |
| `TunnelName` | `*string` | Optional | If `type`==`tunnel` |
| `Type` | [`*models.BgpConfigTypeEnum`](../../doc/models/bgp-config-type-enum.md) | Optional | enum: `external`, `internal`<br>**Constraints**: *Minimum Length*: `1` |
| `Via` | [`*models.BgpConfigViaEnum`](../../doc/models/bgp-config-via-enum.md) | Optional | network name. enum: `lan`, `tunnel`, `vpn`, `wan`<br>**Default**: `"lan"` |
| `VpnName` | `*string` | Optional | - |
| `WanName` | `*string` | Optional | If `via`==`wan` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bfd_minimum_interval": 350,
  "bfd_multiplier": 3,
  "disable_bfd": false,
  "graceful_restart_time": 0,
  "hold_time": 90,
  "no_readvertise_to_overlay": false,
  "via": "lan",
  "auth_key": "auth_key8",
  "export": "export6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

