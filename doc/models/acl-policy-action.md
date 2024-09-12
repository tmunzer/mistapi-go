
# Acl Policy Action

## Structure

`AclPolicyAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | enum: `allow`, `deny`<br>**Default**: `"allow"` |
| `DstTag` | `string` | Required | - |

## Example (as JSON)

```json
{
  "action": "allow",
  "dst_tag": "corp"
}
```

