
# Vars Annotation

Annotation for a single var, helping identify its purpose and enabling auto-complete/enumeration in UI

## Structure

`VarsAnnotation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Note` | `*string` | Optional | User-provided note to describe what this var was created for |
| `Type` | `*string` | Optional | Used to identify where to enumerate / auto-complete the field from. Default is `generic` (plain string, no special handling).<br><br>enum: `generic`, `mxtunnel_id`<br><br>**Default**: `"generic"` |

## Example (as JSON)

```json
{
  "type": "generic",
  "note": "note4"
}
```

