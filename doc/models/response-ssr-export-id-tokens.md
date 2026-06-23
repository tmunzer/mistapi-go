
# Response Ssr Export Id Tokens

Response containing SSR ID tokens exported for device onboarding

## Structure

`ResponseSsrExportIdTokens`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.ResponseSsrExportIdTokensResultsItem`](../../doc/models/response-ssr-export-id-tokens-results-item.md) | Optional | Exported SSR ID token records |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSsrExportIdTokens := models.ResponseSsrExportIdTokens{
        Results:              []models.ResponseSsrExportIdTokensResultsItem{
            models.ResponseSsrExportIdTokensResultsItem{
                Mac:                  models.ToPointer("mac0"),
                Token:                models.ToPointer("token0"),
            },
        },
    }

}
```

