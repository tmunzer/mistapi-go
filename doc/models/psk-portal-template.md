
# Psk Portal Template

Portal UI customization payload

*This model accepts additional fields of type interface{}.*

## Structure

`PskPortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortalTemplate` | [`*models.PskPortalTemplateSetting`](../../doc/models/psk-portal-template-setting.md) | Optional | Custom UI settings for the PSK portal template |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pskPortalTemplate := models.PskPortalTemplate{
        PortalTemplate:       models.ToPointer(models.PskPortalTemplateSetting{
            Alignment:            models.ToPointer(models.PortalTemplateAlignmentEnum_RIGHT),
            Color:                models.ToPointer("color8"),
            Logo:                 models.NewOptional(models.ToPointer("logo6")),
            PoweredBy:            models.ToPointer(false),
            Tos:                  models.ToPointer(false),
        }),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

