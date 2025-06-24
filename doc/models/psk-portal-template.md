
# Psk Portal Template

*This model accepts additional fields of type interface{}.*

## Structure

`PskPortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alignment` | [`*models.PortalTemplateAlignmentEnum`](../../doc/models/portal-template-alignment-enum.md) | Optional | defines alignment on portal. enum: `center`, `left`, `right`<br><br>**Default**: `"center"` |
| `Color` | `*string` | Optional | **Default**: `"#1074bc"` |
| `Logo` | `models.Optional[string]` | Optional | Custom logo with "data:image/png;base64," format.  default null, uses Juniper Mist Logo |
| `PoweredBy` | `*bool` | Optional | Whether to hide "Powered by Juniper Mist" and email footers<br><br>**Default**: `false` |
| `Tos` | `*bool` | Optional | Whether to show Terms of Service |
| `TosAcceptLabel` | `*string` | Optional | Terms of Service accept button label<br><br>**Default**: `"I accept the Terms of Service"` |
| `TosError` | `*string` | Optional | Terror message for not accepting tos<br><br>**Default**: `"Please review and accept the Terms of Service"` |
| `TosLink` | `*string` | Optional | **Default**: `"Terms of Service"` |
| `TosText` | `*string` | Optional | terms and service text displayed in footer if tos is enabled<br><br>**Default**: `"<< provide your Terms of Service here >>"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "alignment": "center",
  "color": "#1074bc",
  "poweredBy": false,
  "tosAcceptLabel": "I accept the Terms of Service",
  "tosError": "Please review and accept the Terms of Service",
  "tosLink": "Terms of Service",
  "tosText": "<< provide your Terms of Service here >>",
  "logo": "logo8",
  "tos": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

