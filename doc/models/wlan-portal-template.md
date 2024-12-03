
# Wlan Portal Template

*This model accepts additional fields of type interface{}.*

## Structure

`WlanPortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortalTemplate` | [`*models.WlanPortalTemplateSetting`](../../doc/models/wlan-portal-template-setting.md) | Optional | portal template wlan settings |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "portal_template": {
    "accessCodeAlternateEmail": "accessCodeAlternateEmail4",
    "alignment": "right",
    "ar": {
      "authButtonAmazon": "authButtonAmazon6",
      "authButtonAzure": "authButtonAzure4",
      "authButtonEmail": "authButtonEmail2",
      "authButtonFacebook": "authButtonFacebook6",
      "authButtonGoogle": "authButtonGoogle2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "authButtonAmazon": "authButtonAmazon4",
    "authButtonAzure": "authButtonAzure4",
    "pageTitle": "pageTitle0",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

