
# Bgp Config

BFD is enabled when either bfd_minimum_interval or bfd_multiplier is configured

## Structure

`BgpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKey` | `*string` | Optional | - |
| `BfdMinimumInterval` | `models.Optional[int]` | Optional | when bfd_multiplier is configured alone<br>default:<br><br>* 1000 if `type`==`external``<br>* 350 `type`==`internal`<br>**Default**: `350`<br>**Constraints**: `>= 1`, `<= 255000` |
| `BfdMultiplier` | `models.Optional[int]` | Optional | when bfd_minimum_interval_is_configured alone<br>**Default**: `3`<br>**Constraints**: `>= 1`, `<= 255` |
| `Communities` | [`[]models.BgpConfigCommunity`](../../doc/models/bgp-config-community.md) | Optional | - |
| `DisableBfd` | `*bool` | Optional | BFD provides faster path failure detection and is enabled by default<br>**Default**: `false` |
| `Export` | `*string` | Optional | - |
| `ExportPolicy` | `*string` | Optional | default export policies if no per-neighbor policies defined |
| `ExtendedV4Nexthop` | `*bool` | Optional | by default, either inet/net6 unicast depending on neighbor IP family (v4 or v6)<br>for v6 neighbors, to exchange v4 nexthop, which allows dual-stack support, enable this |
| `GracefulRestartTime` | `*int` | Optional | `0` means disable<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 4095` |
| `HoldTime` | `*int` | Optional | **Default**: `90`<br>**Constraints**: `>= 0`, `<= 65535` |
| `Import` | `*string` | Optional | - |
| `ImportPolicy` | `*string` | Optional | default import policies if no per-neighbor policies defined |
| `LocalAs` | `*int` | Optional | - |
| `NeighborAs` | `*int` | Optional | - |
| `Neighbors` | [`map[string]models.BgpConfigNeighbors`](../../doc/models/bgp-config-neighbors.md) | Optional | if per-neighbor as is desired. Property key is the neighbor address |
| `Networks` | `[]string` | Optional | if `type`!=`external`or `via`==`wan`networks where we expect BGP neighbor to connect to/from |
| `NoReadvertiseToOverlay` | `*bool` | Optional | by default, we'll re-advertise all learned BGP routers toward overlay<br>**Default**: `false` |
| `Type` | [`*models.BgpConfigTypeEnum`](../../doc/models/bgp-config-type-enum.md) | Optional | **Constraints**: *Minimum Length*: `1` |
| `Via` | [`*models.BgpConfigViaEnum`](../../doc/models/bgp-config-via-enum.md) | Optional | network name<br>**Default**: `"lan"` |
| `VpnName` | `*string` | Optional | - |
| `WanName` | `*string` | Optional | if `via`==`wan` |

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
  "auth_key": "auth_key6",
  "communities": [
    {
      "id": "id8",
      "local_preference": 56,
      "vpn_name": "vpn_name0"
    },
    {
      "id": "id8",
      "local_preference": 56,
      "vpn_name": "vpn_name0"
    }
  ]
}
```

