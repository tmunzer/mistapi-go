
# Vpn Path

*This model accepts additional fields of type interface{}.*

## Structure

`VpnPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BfdProfile` | [`*models.VpnPathBfdProfileEnum`](../../doc/models/vpn-path-bfd-profile-enum.md) | Optional | enum: `broadband`, `lte`<br>**Default**: `"broadband"` |
| `Ip` | `*string` | Optional | If different from the wan port |
| `Pod` | `*int` | Optional | **Default**: `1`<br>**Constraints**: `>= 1`, `<= 128` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bfd_profile": "broadband",
  "pod": 2,
  "ip": "ip4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

