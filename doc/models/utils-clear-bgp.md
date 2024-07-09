
# Utils Clear Bgp

## Structure

`UtilsClearBgp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Neighbor` | `string` | Required | neighbor ip-address or 'all' |
| `Type` | [`models.UtilsClearBgpTypeEnum`](../../doc/models/utils-clear-bgp-type-enum.md) | Required | **Default**: `"hard"` |
| `Vrf` | `*string` | Optional | vrf name |

## Example (as JSON)

```json
{
  "neighbor": "neighbor0",
  "type": "hard",
  "vrf": "vrf4"
}
```

