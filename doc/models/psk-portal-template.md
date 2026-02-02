
# Psk Portal Template

*This model accepts additional fields of type interface{}.*

## Structure

`PskPortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortalTemplate` | [`*models.PskPortalTemplateSetting`](../../doc/models/psk-portal-template-setting.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "portal_template": {
    "alignment": "right",
    "color": "color8",
    "logo": "logo6",
    "poweredBy": false,
    "tos": false
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

