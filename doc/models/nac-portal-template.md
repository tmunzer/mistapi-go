
# Nac Portal Template

Visual template settings for a NAC portal

*This model accepts additional fields of type interface{}.*

## Structure

`NacPortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alignment` | [`*models.PortalTemplateAlignmentEnum`](../../doc/models/portal-template-alignment-enum.md) | Optional | defines alignment on portal. enum: `center`, `left`, `right`<br><br>**Default**: `"center"` |
| `Color` | `*string` | Optional | Primary color used by the NAC portal template<br><br>**Default**: `"#1074bc"` |
| `Logo` | `*string` | Optional | Custom logo custom logo with "data:image/png;base64," format. default null, uses Juniper Mist Logo. |
| `PoweredBy` | `*bool` | Optional | Whether to hide "Powered by Juniper Mist" and email footers<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacPortalTemplate := models.NacPortalTemplate{
        Alignment:            models.ToPointer(models.PortalTemplateAlignmentEnum_CENTER),
        Color:                models.ToPointer("#1074bc"),
        Logo:                 models.ToPointer("logo8"),
        PoweredBy:            models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

