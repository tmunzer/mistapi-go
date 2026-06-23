
# Const Language

Language option returned by the constants API

## Structure

`ConstLanguage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | Localized language display name in the API default locale |
| `DisplayNative` | `string` | Required | Language display name in its native locale |
| `Key` | `string` | Required | Locale key used to select this language |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constLanguage := models.ConstLanguage{
        Display:              "English (US)",
        DisplayNative:        "English (US)",
        Key:                  "en-US",
    }

}
```

