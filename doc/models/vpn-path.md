
# Vpn Path

## Structure

`VpnPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BfdProfile` | [`*models.VpnPathBfdProfileEnum`](../../doc/models/vpn-path-bfd-profile-enum.md) | Optional | **Default**: `"broadband"` |
| `Ip` | `*string` | Optional | if different from the wan port |
| `Pod` | `*int` | Optional | **Default**: `1`<br>**Constraints**: `>= 1`, `<= 128` |

## Example (as JSON)

```json
{
  "bfd_profile": "broadband",
  "pod": 2,
  "ip": "ip0"
}
```

