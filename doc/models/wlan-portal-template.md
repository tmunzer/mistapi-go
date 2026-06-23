
# Wlan Portal Template

Guest portal template payload for a WLAN

*This model accepts additional fields of type interface{}.*

## Structure

`WlanPortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortalTemplate` | [`*models.WlanPortalTemplateSetting`](../../doc/models/wlan-portal-template-setting.md) | Optional | Portal template settings for the WLAN guest portal |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanPortalTemplate := models.WlanPortalTemplate{
        PortalTemplate:       models.ToPointer(models.WlanPortalTemplateSetting{
            AccessCodeAlternateEmail:  models.ToPointer("accessCodeAlternateEmail4"),
            Alignment:                 models.ToPointer(models.PortalTemplateAlignmentEnum_RIGHT),
            Ar:                        models.ToPointer(models.WlanPortalTemplateSettingLocale{
                AuthButtonAmazon:          models.ToPointer("authButtonAmazon6"),
                AuthButtonAzure:           models.ToPointer("authButtonAzure4"),
                AuthButtonEmail:           models.ToPointer("authButtonEmail2"),
                AuthButtonFacebook:        models.ToPointer("authButtonFacebook6"),
                AuthButtonGoogle:          models.ToPointer("authButtonGoogle2"),
            }),
            AuthButtonAmazon:          models.ToPointer("authButtonAmazon4"),
            AuthButtonAzure:           models.ToPointer("authButtonAzure4"),
            PageTitle:                 "pageTitle0",
        }),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

