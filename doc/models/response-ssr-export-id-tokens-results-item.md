
# Response Ssr Export Id Tokens Results Item

SSR ID token record for a device

## Structure

`ResponseSsrExportIdTokensResultsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | SSR device MAC address for the exported ID token |
| `Token` | `*string` | Optional | ID token to import into Conductor for SSR device onboarding |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSsrExportIdTokensResultsItem := models.ResponseSsrExportIdTokensResultsItem{
        Mac:                  models.ToPointer("mac6"),
        Token:                models.ToPointer("token4"),
    }

}
```

