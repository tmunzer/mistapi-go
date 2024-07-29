
# Psk Portal Template

## Structure

`PskPortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alignment` | [`*models.PskPortalTemplateAlignmentEnum`](../../doc/models/psk-portal-template-alignment-enum.md) | Optional | defines alignment on portal. enum: `center`, `left`, `right`<br>**Default**: `"center"` |
| `Color` | `*string` | Optional | **Default**: `"#1074bc"` |
| `Logo` | `models.Optional[string]` | Optional | custom logo.  default null, uses Juniper Mist Logo |
| `PoweredBy` | `*bool` | Optional | whether to hide "Powered by Juniper Mist" and email footers<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "alignment": "center",
  "color": "#1074bc",
  "poweredBy": false,
  "logo": "logo4"
}
```

