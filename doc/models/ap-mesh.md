
# Ap Mesh

Mesh AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`ApMesh`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bands` | [`[]models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Optional | List of bands that the mesh should apply to. For relay, the first viable one will be picked. For relay, the first viable one will be picked. enum: `24`, `5`, `6` |
| `Enabled` | `*bool` | Optional | Whether mesh is enabled on this AP<br><br>**Default**: `false` |
| `Group` | `models.Optional[int]` | Optional | Mesh group, base AP(s) will only allow remote AP(s) in the same mesh group to join, 1-9, optional<br><br>**Constraints**: `>= 1`, `<= 9` |
| `Role` | [`*models.ApMeshRoleEnum`](../../doc/models/ap-mesh-role-enum.md) | Optional | enum: `base`, `remote` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "group": 1,
  "role": "base",
  "bands": [
    "6"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

