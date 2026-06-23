
# Search Wxtag Apps Item

Application metadata available for WxTag matching

## Structure

`SearchWxtagAppsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Group` | `string` | Required | Application category group containing this WxTag app |
| `Key` | `string` | Required | Stable application key used in WxTag configuration |
| `Name` | `string` | Required | Display name of the application |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    searchWxtagAppsItem := models.SearchWxtagAppsItem{
        Group:                "Emails",
        Key:                  "gmail",
        Name:                 "Gmail - web/app",
    }

}
```

