
# Nac Client Coa

## Structure

`NacClientCoa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CoaType` | [`*models.NacCoaTypeEnum`](../../doc/models/nac-coa-type-enum.md) | Optional | CoA type to send. enum: `reauth`, `disconnect`<br><br>**Default**: `"reauth"` |

## Example (as JSON)

```json
{
  "coa_type": "reauth"
}
```

