
# Ap Mesh

Mesh AP settings

## Structure

`ApMesh`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether mesh is enabled on this AP<br>**Default**: `false` |
| `Group` | `models.Optional[int]` | Optional | mesh group, base AP(s) will only allow remote AP(s) in the same mesh group to join, 1-9, optional<br>**Constraints**: `>= 1`, `<= 9` |
| `Role` | [`*models.ApMeshRoleEnum`](../../doc/models/ap-mesh-role-enum.md) | Optional | enum: `base`, `remote` |

## Example (as JSON)

```json
{
  "enabled": false,
  "group": 1,
  "role": "base"
}
```

