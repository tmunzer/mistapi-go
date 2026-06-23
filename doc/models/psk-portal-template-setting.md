
# Psk Portal Template Setting

Custom UI settings for the PSK portal template

## Structure

`PskPortalTemplateSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alignment` | [`*models.PortalTemplateAlignmentEnum`](../../doc/models/portal-template-alignment-enum.md) | Optional | defines alignment on portal. enum: `center`, `left`, `right`<br><br>**Default**: `"center"` |
| `Color` | `*string` | Optional | Primary hex color used by the portal template<br><br>**Default**: `"#1074bc"` |
| `Logo` | `models.Optional[string]` | Optional | Custom logo with "data:image/png;base64," format. default null, uses Juniper Mist Logo |
| `PoweredBy` | `*bool` | Optional | Whether to hide "Powered by Juniper Mist" and email footers<br><br>**Default**: `false` |
| `Tos` | `*bool` | Optional | Whether to show Terms of Service |
| `TosAcceptLabel` | `*string` | Optional | Terms of Service accept button label<br><br>**Default**: `"I accept the Terms of Service"` |
| `TosError` | `*string` | Optional | Error message shown when the user has not accepted the Terms of Service<br><br>**Default**: `"Please review and accept the Terms of Service"` |
| `TosLink` | `*string` | Optional | Terms of Service link label displayed in the portal footer<br><br>**Default**: `"Terms of Service"` |
| `TosText` | `*string` | Optional | Terms of Service text displayed in the footer when Terms are enabled<br><br>**Default**: `"<< provide your Terms of Service here >>"` |
| `TosUrl` | `*string` | Optional | Custom URL for the Terms of Service policy |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pskPortalTemplateSetting := models.PskPortalTemplateSetting{
        Alignment:            models.ToPointer(models.PortalTemplateAlignmentEnum_CENTER),
        Color:                models.ToPointer("#1074bc"),
        Logo:                 models.NewOptional(models.ToPointer("logo0")),
        PoweredBy:            models.ToPointer(false),
        Tos:                  models.ToPointer(false),
        TosAcceptLabel:       models.ToPointer("I accept the Terms of Service"),
        TosError:             models.ToPointer("Please review and accept the Terms of Service"),
        TosLink:              models.ToPointer("Terms of Service"),
        TosText:              models.ToPointer("<< provide your Terms of Service here >>"),
        TosUrl:               models.ToPointer("https://company.com/wifi-policy"),
    }

}
```

