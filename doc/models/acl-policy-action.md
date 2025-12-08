
# Acl Policy Action

## Structure

`AclPolicyAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | enum: `allow`, `deny` |
| `DstTag` | `string` | Required | - |

## Example (as JSON)

```json
{
  "dst_tag": "corp",
  "action": "allow"
}
```

