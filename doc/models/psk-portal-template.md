
# Psk Portal Template

*This model accepts additional fields of type interface{}.*

## Structure

`PskPortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alignment` | [`*models.PortalTemplateAlignmentEnum`](../../doc/models/portal-template-alignment-enum.md) | Optional | defines alignment on portal. enum: `center`, `left`, `right`<br>**Default**: `"center"` |
| `Color` | `*string` | Optional | **Default**: `"#1074bc"` |
| `Logo` | `models.Optional[string]` | Optional | custom logo with “data:image/png;base64,” format.  default null, uses Juniper Mist Logo |
| `PoweredBy` | `*bool` | Optional | whether to hide "Powered by Juniper Mist" and email footers<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "alignment": "center",
  "color": "#1074bc",
  "poweredBy": false,
  "logo": "logo8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

